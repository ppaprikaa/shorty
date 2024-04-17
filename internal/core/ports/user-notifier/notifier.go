package usernotifier

import "context"

type Notifier interface {
	NotifyActivation(context.Context, *ActivationArgs) error
}

type ActivationArgs struct {
	Destination string
	Code        string
}
