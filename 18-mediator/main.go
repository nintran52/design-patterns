package main

import "fmt"

// Step 1: Define interface
type User interface {
	Send(message string)
	Receive(message string)
	GetName() string
}

type Mediator interface {
	Send(message string, user User)
	Register(user User)
}

// Step 2: Implement concrete User
type ChatUser struct {
	name     string
	mediator Mediator
}

func NewChatUser(name string, mediator Mediator) *ChatUser {
	return &ChatUser{name: name, mediator: mediator}
}

func (u *ChatUser) Send(message string) {
	fmt.Printf("%s sends: %s\n", u.name, message)
	u.mediator.Send(message, u)
}

func (u *ChatUser) Receive(message string) {
	fmt.Printf("%s received: %s\n", u.name, message)
}

func (u *ChatUser) GetName() string {
	return u.name
}

// Step 3: Implement concrete Mediator
type ChatRoom struct {
	users []User
}

func (c *ChatRoom) Register(user User) {
	c.users = append(c.users, user)
}

func (c *ChatRoom) Send(message string, sender User) {
	for _, u := range c.users {
		if u != sender {
			u.Receive(fmt.Sprintf("[%s]: %s", sender.GetName(), message))
		}
	}
}

// Step 4: Client code
func main() {
	room := &ChatRoom{}

	quan := NewChatUser("Quan", room)
	tai := NewChatUser("Tai", room)
	hieu := NewChatUser("Hieu", room)

	room.Register(quan)
	room.Register(tai)
	room.Register(hieu)

	quan.Send("Hi everyone 👋")
	tai.Send("Hey Quan!")
	hieu.Send("Nice to meet you both 😊")
}
