package main
import "errors"

type Message interface {
	Pri()
}

type Command struct {
	Pri int
	Func func(Message) error
}


type Command interface {
	Exec()
}

func (cmd *Command) Exec() error {
	if cmd.Func == nil {
		return errors.New("command func not defined")
	}
	return cmd.Func()
}

