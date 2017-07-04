package cmdr

import (
	"fmt"
	"log"
	"os/exec"

	shellwords "github.com/mattn/go-shellwords"
)

// Dispatcher is dispatcher
type Dispatcher struct {
	commands map[string]string
}

// NewDispatcher create a new dispatcher
func NewDispatcher(cmap map[string]string) *Dispatcher {
	return &Dispatcher{
		commands: cmap,
	}
}

// Do dispatch command and return capture output
func (d *Dispatcher) Do(c string) string {
	args, err := shellwords.Parse(c)
	mcmd := d.commands[args[0]]
	fmt.Printf("=> Found predefined command -> %s\n", mcmd)
	params := args[1:]
	parami := make([]interface{}, len(params))
	for i, v := range params {
		parami[i] = v
	}
	newcmd := fmt.Sprintf(mcmd, parami...)
	newargs, err := shellwords.Parse(newcmd)
	fmt.Printf("==> Executing -> %v\n", newcmd)
	out, err := exec.Command(newargs[0], newargs[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
