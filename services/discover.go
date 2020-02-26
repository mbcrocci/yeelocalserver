package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/mbcrocci/yeelocalsrv/entities"
)

const (
	message  = "M-SEARCH * HTTP/1.1\r\n HOST:239.255.255.250:1982\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n"
	ssdpAddr = "239.255.255.250:1982"
)

type DiscoverService struct {
	conn  *net.UDPConn
	laddr *net.UDPAddr
	maddr *net.UDPAddr
}

func NewDiscoverService() *DiscoverService {
	return &DiscoverService{}
}

func (ds *DiscoverService) Init() error {
	maddr, err := net.ResolveUDPAddr("udp4", ssdpAddr)
	if err != nil {
		return err
	}

	laddr, err := net.ResolveUDPAddr("udp4", ":0")
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		return err
	}

	ds.maddr = maddr
	ds.laddr = laddr
	ds.conn = conn
	return nil
}

func (ds *DiscoverService) Start(c chan *entities.Light, e chan error) {
	go func() {
		for {
			time.Sleep(5 * time.Second)

			_, err := ds.conn.WriteToUDP([]byte(message), ds.maddr)
			if err != nil {
				e <- err
			}

			ds.conn.SetReadDeadline(time.Now().Add(5 * time.Second))

			buff := make([]byte, 1024)
			n, _, err := ds.conn.ReadFrom(buff)
			if err != nil {
				e <- err
			}

			light, err := ds.ParseLight(string(buff[:n]))
			if err != nil {
				e <- err
				continue
			}
			c <- light
		}

	}()
	fmt.Println("Discover service started!")
}

func (ds *DiscoverService) Shutdown() error {
	return ds.conn.Close()
}

func (ds DiscoverService) ParseLight(ls string) (*entities.Light, error) {
	if len(ls) == 0 {
		return nil, errors.New("Empty device")
	}

	lines := strings.Split(ls, "\r\n")

	jsonStrs := []string{}
	for _, l := range lines {
		if !strings.Contains(l, "HTTP/1.1") &&
			!strings.Contains(l, "Server") &&
			!strings.Contains(l, "Cache") &&
			!strings.Contains(l, "Ext") &&
			!strings.Contains(l, "Date") {

			keyvalues := strings.Split(l, ":")
			if len(keyvalues) > 1 {
				value := strings.Join(keyvalues[1:], "") // in "Location" the value is "yeelight://<ip>", which makes it have more than 1 value

				k, v := strings.TrimSpace(keyvalues[0]), strings.TrimSpace(value)

				jsonStrs = append(jsonStrs, fmt.Sprintf("\"%s\": \"%s\",", k, v))
			}
		}
	}
	jsonStr := strings.Join(jsonStrs, "\n")
	jsonStr = jsonStr[:len(jsonStr)-1]
	jsonStr = fmt.Sprintf("{%s}", jsonStr)

	var light *entities.Light
	err := json.Unmarshal([]byte(jsonStr), &light)
	if err != nil {
		return nil, err
	}

	return light, nil
}
