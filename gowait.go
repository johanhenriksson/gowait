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

func Execute(taskfunc TaskFunc) {
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

	// set as virtual for now
	taskdef.Meta["virtual"] = true

	// initialize task object
	// this will have more stuff added to it later
	task := Task{
		Taskdef: taskdef,
	}

	// connect upstream and send init message
	client := Connect(taskdef.Upstream)
	client.SendInit(taskdef)

	// stdout -> log pump
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

	// execute task code
	result, err := taskfunc(&task)

	// re-attach stdout
	os.Stdout = RealStdout

	// send return value or error
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
