package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/janhajek-outreach/temporal-playground/solutions"
	"github.com/janhajek-outreach/temporal-playground/solutions/activity_timeout"
	"github.com/janhajek-outreach/temporal-playground/solutions/cancel"
	"github.com/janhajek-outreach/temporal-playground/solutions/example01"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create client
	c, err := client.NewClient(client.Options{
		HostPort:  solutions.Address,
		Namespace: solutions.Namespace,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// Create worker
	w := worker.New(c, solutions.TaskQueue, worker.Options{
		MaxConcurrentActivityExecutionSize: 1,
	})

	// Register workflows and activities

	w.RegisterWorkflow(example01.Workflow1)

	w.RegisterWorkflow(activity_timeout.ActivityTimeoutWorkflow01)
	w.RegisterWorkflow(activity_timeout.ActivityTimeoutWorkflow02)
	w.RegisterWorkflow(activity_timeout.ActivityTimeoutWorkflow03)
	w.RegisterWorkflow(activity_timeout.ActivityTimeoutWorkflow04)
	w.RegisterActivity(activity_timeout.ActivityTimeoutSleepActivity)

	w.RegisterWorkflow(cancel.CancelWorkflow)
	w.RegisterActivity(&cancel.Activities{})

	// Start worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
