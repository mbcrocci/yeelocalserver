# yeelocalserver
An http server that runs a service to discover yeelights in your local network

### Install and Run
```bash
go get github.com/mbcrocci/yeelocalserver

yeelocalserver
```

### Consume api
#### Lights
```http
GET http://localhost:3000/lights HTTP/1.1
```

#### Command
```http
POST :3000/lights/<light id>/command HTTP/1.1
"Content-type": "application/json"
{
  "id": 1,
  "method": "set_power",
  "params": ["on", "smooth", 0]
}
```
