package models

import "time"

type Nodes struct {
	HostName string
	Role string
	JoinTime time.Time
}

type Tokens struct {
	HostName string
	Token string
}

type Labels struct {
	HostName string
	Label string
}

type Actions struct {
	Time time.Time
	Context string
}

type Logs struct {
	Types string
	Time time.Time
	Context string
	Node string
}