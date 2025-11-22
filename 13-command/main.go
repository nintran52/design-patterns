package main

// Step 1: Command interface
type Command interface {
	Execute()
	Undo()
}

// Step 2: Receiver (Devices)
type Light struct {
	location string
}

func (l *Light) On() {
	println(l.location + " light is ON")
}

func (l *Light) Off() {
	println(l.location + " light is OFF")
}

// Step 3: Concrete Commands
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}

// Step 4: Invoker (Remote Control)
type RemoteControl struct {
	onCommand   Command
	offCommand  Command
	lastCommand Command
}

func (rc *RemoteControl) SetCommands(on, off Command) {
	rc.onCommand = on
	rc.offCommand = off

}

func (rc *RemoteControl) PressOn() {
	rc.onCommand.Execute()
	rc.lastCommand = rc.onCommand
}

func (rc *RemoteControl) PressOff() {
	rc.offCommand.Execute()
	rc.lastCommand = rc.offCommand
}

func (rc *RemoteControl) PressUndo() {
	if rc.lastCommand != nil {
		rc.lastCommand.Undo()
	}
}

// Step 5: Usage Example
func main() {
	livingRoomLight := &Light{location: "Living Room"}

	lightOn := &LightOnCommand{light: livingRoomLight}
	lightOff := &LightOffCommand{light: livingRoomLight}

	remote := &RemoteControl{}
	remote.SetCommands(lightOn, lightOff)

	remote.PressOn()
	remote.PressOff()
	remote.PressUndo()
}
