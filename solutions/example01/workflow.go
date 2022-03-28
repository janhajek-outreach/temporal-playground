package example01

import (
	"go.temporal.io/sdk/workflow"
)

type Input struct {
	A int
	B int
}

type Output struct {
	Result int
}

/**
tctl workflow run --taskqueue playground --execution_timeout 60 --workflow_type Workflow1
*/
func Workflow1(ctx workflow.Context, input Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 01")

	return Output{input.A + input.B}, nil
}
