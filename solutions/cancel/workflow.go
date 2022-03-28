package cancel

import (
	"errors"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

// CancelWorkflow is a Workflow Definition that shows how it can be canceled.
//
// tctl workflow run --taskqueue playground --execution_timeout 60 --workflow_type CancelWorkflow  --workflow_id cancel_test
// tctl workflow cancel --workflow_id cancel_test
func CancelWorkflow(ctx workflow.Context) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Minute,
		HeartbeatTimeout:    5 * time.Second,
		WaitForCancellation: true,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	logger.Info("cancel workflow started")
	var a *Activities // Used to call Activities by function pointer
	defer func() {

		if !errors.Is(ctx.Err(), workflow.ErrCanceled) {
			return
		}

		logger.Info("canceled xxxxxxxxxxxxxxxxxxxxx")

		// When the Workflow is canceled, it has to get a new disconnected context to execute any Activities
		//newCtx, _ := workflow.NewDisconnectedContext(ctx)
		//err := workflow.ExecuteActivity(newCtx, a.CleanupActivity).Get(ctx, nil)
		//if err != nil {
		//	logger.Error("CleanupActivity failed", "Error", err)
		//}
	}()

	var result string
	err := workflow.ExecuteActivity(ctx, a.ActivityToBeCanceled).Get(ctx, &result)
	logger.Info(fmt.Sprintf("ActivityToBeCanceled returns %v, %v", result, err))

	err = workflow.ExecuteActivity(ctx, a.ActivityToBeSkipped).Get(ctx, nil)
	logger.Error("Error from ActivityToBeSkipped", "Error", err)

	logger.Info("Workflow Execution complete.")

	return nil
}
