package activity_timeout

import (
	"context"
	"errors"
	"time"

	"go.temporal.io/sdk/activity"
)

type ActivityInput struct {
	Attempts []Attempt
}

type Attempt struct {
	Duration  time.Duration
	ResultErr bool
}

func ActivityTimeoutSleepActivity(ctx context.Context, input ActivityInput) error {

	duration := time.Duration(0)
	var err bool

	activityInfo := activity.GetInfo(ctx)
	if len(input.Attempts) >= int(activityInfo.Attempt) {
		duration = input.Attempts[int(activityInfo.Attempt)-1].Duration
		err = input.Attempts[int(activityInfo.Attempt)-1].ResultErr
	}

	activity.GetLogger(ctx).Info("staaaaaaaaaaaaaart")

	if duration > 0 {
		time.Sleep(duration)
	}

	activity.GetLogger(ctx).Info("eeeeeeeeeeeeeeeeeeend")

	if err {
		return errors.New("activity error")
	}
	return nil
}
