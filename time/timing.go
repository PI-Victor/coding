package main

import (
	"fmt"
	"time"
)

type stepInfo struct {
	stepName  string
	timeStart time.Time
	timeStop  time.Time
}

func main() {
	steps := []stepInfo{}
	steps = append(steps,
		stepInfo{
			stepName:  "step1",
			timeStart: time.Now(),
		})

	steps = append(steps,
		stepInfo{
			stepName:  "step2",
			timeStart: time.Now(),
		})

	steps = append(steps,
		stepInfo{
			stepName:  "step3",
			timeStart: time.Now(),
		})

	wantedStep := "step2"
	for i, step := range steps {
		if step.stepName == wantedStep {
			steps[i] = stepInfo{}
		}
		fmt.Println(steps)
	}
	fmt.Printf("test\n")
	for _, step := range steps {
		if step.stepName == "" {
			fmt.Println("You should watch your step!")
		}
		fmt.Println(step.stepName)
	}
}
