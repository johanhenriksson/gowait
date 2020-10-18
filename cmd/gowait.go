package main

import (
	"fmt"
	"github.com/johanhenriksson/gowait"
	"time"
)

type Input struct {
	Hello int
}

type Result struct {
	Key string
}

func main() {
	gowait.Execute(func(task *gowait.Task) (gowait.Result, error) {
		// this is how to unpack custom inputs:
		inputs := Input{}
		if err := task.Input(&inputs); err != nil {
			return nil, err
		}

		// do something useful
		fmt.Println("input hello:", inputs.Hello)
		for i := 0; i < 3; i++ {
			fmt.Println("hello team", i)
			time.Sleep(time.Second)
		}

		// return a result!
		result := Result{
			Key: "this is the epic result of this massive computation",
		}
		return result, nil
	})
}
