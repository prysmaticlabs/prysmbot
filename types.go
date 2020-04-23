package main

type botCommandGroup struct {
	name string
	displayName string
	shorthand string
	helpText string
	flags []*botCommand
}

type botCommand struct {
	group string
	command string
	shorthand string
	helpText string
	responseText string
}