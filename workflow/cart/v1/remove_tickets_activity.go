package workflow

import (
	"context"

	cartv1 "github.com/taltemus/stacked-diff-example/gen/cart/v1"
)

func (w *CartWorkflow) RemoveTicketsActivity(ctx context.Context, input *cartv1.RemoveTicketsInput) (*cartv1.RemoveTicketsOutput, error) {
	// Do some business logic
	return &cartv1.RemoveTicketsOutput{}, nil
}
