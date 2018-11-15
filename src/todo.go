package main

import "time"

type Todo struct {
	// Name      string    `json:"name"`
	// Completed bool      `json:"completed"`
	// Due       time.Time `json:"due"`
	Version  string    `json:"version"`
	Hostname string    `json:"hostname"`
	Time     time.Time `json:"time"`
	Db       string    `json:"db"`
	Apikey   string    `json:"apikey"`
}

type Todos []Todo
