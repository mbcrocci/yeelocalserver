package entities

// Command represent an yeelight command object.
// The params are kept as an array of string for now to ease development,
// they will be change to either []interface or Command will be an interface.
type Command struct {
	ID int `json:"id"`
	Method string `json:"method"`
	Params []string `json:"params"`
}

func NewCommand(id int, method string, params []string) *Command {
	return &Command{
		ID: id,
		Method: method,
		Params: params,
	}
}