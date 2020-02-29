package entities

// Command represent an yeelight command object.
type Command struct {
	ID     int           `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

func NewCommand(id int, method string, params []interface{}) *Command {
	return &Command{
		ID:     id,
		Method: method,
		Params: params,
	}
}
