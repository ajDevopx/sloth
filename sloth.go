package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// task template
type task struct {
	TaskID   string
	TaskMsg  string
	TaskTime string
}

// get message from the main and
// convert it into slice to string
func getMessage(msg []string) (stringMsg string) {
	stringMsg = strings.Join(msg, "")
	return stringMsg
}

// get unix timestamp for an event
// converted into string
func getTimestamp() string {
	unixTime := time.Now().Unix()
	return strconv.FormatInt(unixTime, 10)
}

// construct hash value from the message and
// the timestamp of the instance we are created
func constructHashAllother(getMsg []string) (hash, strMsg, strTimestamp string) {
	// get msg and the timestamp
	strMsg = getMessage(getMsg)
	strTimestamp = getTimestamp()
	hashVictim := strMsg + strTimestamp
	hasher := sha1.New()
	hasher.Write([]byte(hashVictim))
	finalhash := hasher.Sum(nil)
	hash = hex.EncodeToString(finalhash)
	return
}

// create new task
func newTask(rawMsg []string) *task {
	p := new(task)
	hash, msg, timestamp := constructHashAllother(rawMsg)
	p.TaskID = hash
	p.TaskMsg = msg
	p.TaskTime = timestamp
	return p
}

// add the task to list file
func addMsg(message []string) {
	taskPtr := newTask(message)
	bs, err := json.Marshal(taskPtr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

func main() {

	// subcommands
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)

	// add subcommand pointer.
	addCommandPtr := addCommand.String("m", "", "To specify the task message")

	// To verify the subcommand is provided.
	if len(os.Args) < 2 {
		fmt.Println("Two more subcommand is required.")
		os.Exit(1)
	}

	// Parse the flag for appropriate flagSet
	// Subcommand parsing.
	switch os.Args[1] {
	case "add":
		addCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check the flagset for subcommand is parsed
	// Check the false for the subcommand value
	if addCommand.Parsed() {
		if *addCommandPtr == "" {
			addCommand.PrintDefaults()
			os.Exit(1)
		}
	}
	//command line message
	message := os.Args[3:]
	// get the message from the prompt
	addMsg(message)

}
