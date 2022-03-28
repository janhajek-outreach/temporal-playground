package activity_timeout

import (
	"time"

	"github.com/janhajek-outreach/temporal-playground/solutions"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// ActivityTimeoutWorkflow01
// StartToCloseTimeout is lower than activity duration, you can see overlapping actions because temporal starts a new action immediately after
// the timeout is done
// the workflow will fail because all of these actions will fail with a timeout error
//
// tctl workflow run --taskqueue playground --execution_timeout 60 --workflow_type ActivityTimeoutWorkflow01
func ActivityTimeoutWorkflow01(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           solutions.TaskQueue,
		StartToCloseTimeout: 3 * time.Second, // timeout is lower than activity duration
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 1.0,
			MaximumInterval:    20 * time.Second,
			MaximumAttempts:    5,
		},
	})

	err := workflow.ExecuteActivity(ctx, ActivityTimeoutSleepActivity, ActivityInput{
		Attempts: []Attempt{
			{
				Duration: 5 * time.Second,
			},
			{
				Duration: 5 * time.Second,
			},
			{
				Duration: 5 * time.Second,
			},
			{
				Duration: 5 * time.Second,
			},
			{
				Duration: 5 * time.Second,
			},
		},
	}).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

// ActivityTimeoutWorkflow02
// set activity concurrency
// w := worker.New(c, solutions.TaskQueue, worker.Options{
//   MaxConcurrentActivityExecutionSize: 1,
// })
// and start two workflows, the first one starts to process activity
// the second one will time out after 3 seconds because of ScheduleToStartTimeout
//
// tctl workflow run --taskqueue playground --execution_timeout 60 --workflow_type ActivityTimeoutWorkflow02
func ActivityTimeoutWorkflow02(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              solutions.TaskQueue,
		ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout:    10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 1.0,
			MaximumInterval:    20 * time.Second,
			MaximumAttempts:    5,
		},
	})

	err := workflow.ExecuteActivity(ctx, ActivityTimeoutSleepActivity, ActivityInput{
		Attempts: []Attempt{
			{
				Duration: 8 * time.Second,
			},
		},
	}).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

// ActivityTimeoutWorkflow03
//
// tctl workflow run --taskqueue playground --execution_timeout 60 --workflow_type ActivityTimeoutWorkflow03
func ActivityTimeoutWorkflow03(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              solutions.TaskQueue,
		ScheduleToCloseTimeout: 3 * time.Second,
		StartToCloseTimeout:    6 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 1.0,
			MaximumInterval:    20 * time.Second,
			MaximumAttempts:    5,
		},
	})

	err := workflow.Sleep(ctx, time.Second*5)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, ActivityTimeoutSleepActivity, ActivityInput{}).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

// ActivityTimeoutWorkflow04
// the first 4 attempts run longer than timeout, the last one is shorter and the workflow ends successfully
//
// tctl workflow run --taskqueue playground --execution_timeout 60 --workflow_type ActivityTimeoutWorkflow04
func ActivityTimeoutWorkflow04(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           solutions.TaskQueue,
		StartToCloseTimeout: 3 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 1.0,
			MaximumInterval:    20 * time.Second,
			MaximumAttempts:    5,
		},
	})

	err := workflow.ExecuteActivity(ctx, ActivityTimeoutSleepActivity, ActivityInput{
		Attempts: []Attempt{
			{
				Duration: 5 * time.Second,
			},
			{
				Duration: 5 * time.Second,
			},
			{
				Duration: 5 * time.Second,
			},
			{
				Duration: 5 * time.Second,
			},
			{
				Duration: 2 * time.Second,
			},
		},
	}).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
