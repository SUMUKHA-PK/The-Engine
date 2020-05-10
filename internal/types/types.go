package types

import "time"

// Command is the structure of any command running at any point of time
type Command struct {
	Timestamp   time.Time
	CommandName string
	// some sort of context is needed in order to cancel the commands when I want to
}

// RunningQueue holds all the commans that are currently running
var RunningQueue chan Command

// LogQueue holds all the logs
var LogQueue chan Command

// Runnable is the struct that describes a command to be run.
type Runnable struct {
	Category string // TV Show, Movie
	Name     string // Name of the playable
	Season   string
	Episode  string
}
