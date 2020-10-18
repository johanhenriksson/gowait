package gowait

import (
	"bufio"
	"log"
	"os"
	"time"
)

var RealStdout = os.Stdout

type Dict map[string]interface{}
type Result interface{}
type TaskFunc func(*Task) (Result, error)
type TaskMap map[string]TaskFunc

func Main(tasks TaskMap) {
	// load task definition from env
	// load cluster definition from env
	// connect upstream
	// hook stdout
	// execute task
	// send result

	// unpack task definition from environment
	taskdef := &Taskdef{}
	if err := unpackstruct(envTaskdef, taskdef); err != nil {
		log.Fatal("err unpacking taskdef:", err)
	}

	// connect upstream as soon as possible.
	// any error that occurs before the upstream connection
	// is established is effectively lost.
	client := Connect(taskdef.Upstream)

	// find task function by name
	taskfunc, exists := tasks[taskdef.Name]
	if !exists {
		log.Fatal("No such task:", taskdef.Name)
	}

	// set as virtual for now
	// this avoids task lost errors while in test mode
	taskdef.Meta["virtual"] = true

	// initialize task object
	// this will have more stuff added to it later
	task := Task{
		Taskdef: taskdef,
	}

	// capture stdout and send logs upstream
	r, w, _ := os.Pipe()
	os.Stdout = w
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			client.SendLog(taskdef.ID, "stdout", scanner.Text())
			RealStdout.WriteString(scanner.Text())
			RealStdout.WriteString("\n")
		}
	}()

	// send task initialization message
	client.SendInit(taskdef)

	// execute task function
	result, err := taskfunc(&task)

	// restore stdout
	os.Stdout = RealStdout

	// send return value or error upstream
	if err != nil {
		client.SendError(taskdef.ID, err.Error())
	} else {
		client.SendReturn(taskdef.ID, result)
	}

	// wait a bit for everything to finish
	// todo: replace with proper synchronization
	time.Sleep(100 * time.Millisecond)
	client.Close()
}
