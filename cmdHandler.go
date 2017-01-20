package tegola

import (
	"sync"
	"errors"
	"regexp"
)

type CmdHandler func( *Bot, CmdData, Update ) error

type CmdManager struct {
	Routes map[string]CmdHandler
}

func NewCmdManager () *CmdManager {
	var manager = &CmdManager{Routes: map[string]CmdHandler{}}
	manager.Routes["default"] = func (*Bot, CmdData, Update) error {
		return nil
	}
	return manager
}

type CmdData struct {
	Cmd string
	BotName string
	Args string
}

func (c *CmdManager) AddRoute ( route string, cmd CmdHandler, overwrite bool ) error {
	// If route already exists and I don't force overwriting, fire an error
	if _, exists := c.Routes[route]; exists && !overwrite {
		return errors.New("Route " + route + " already exists.")
	}
	c.Routes[route] = cmd
	return nil
}

func (c *CmdManager) DeleteRoute ( route string ) {
	delete(c.Routes, route)
}

func (c *CmdManager) ModifyRoute ( route string, cmd CmdHandler ) error {
	if _, exists := c.Routes[route]; !exists {
		return errors.New("Route " + route +  "does not exists.")
	}
	c.Routes[route] = cmd
	return nil
}

// Echo repeats last user message back to the chat
func (c *CmdManager) CmdRouter (b *Bot, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		u Update
		cmdData CmdData
		route string
		handler CmdHandler
		handled bool
	)

	for u = range b.UpdatesChan {
		handled = false
		// It may be empty
		cmdData = DetectCmd(u.Message.Text)
		// Call the right handler for the command
		for route, handler = range c.Routes {
			if route == cmdData.Cmd {
				handler(b, cmdData, u)
				handled = true
				break
			}
		}
		if !handled {
			c.Routes["default"](b, cmdData, u)
		}
	}
}

func DetectCmd (text string) CmdData {
	var (
		cmdData CmdData
		reg = regexp.MustCompile(CommandRegString)
		match []string
		res map[string]string
	)

	match = reg.FindStringSubmatch(text)
	if len(match) == 0 {
		return cmdData
	}

	res = map[string]string{}
	for i, n := range reg.SubexpNames() {
		if i != 0 { // first name is empty because index 0 is the whole match
			res[n] = match[i]
		}
	}

	cmdData.Cmd = res["command"]
	cmdData.BotName = res["bot"]
	cmdData.Args = res["args"]

	return cmdData
}
