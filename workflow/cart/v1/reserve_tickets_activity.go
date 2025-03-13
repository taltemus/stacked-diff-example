package workflow

import (
	"context"

	cartv1 "github.com/taltemus/stacked-diff-example/gen/cart/v1"
)

func (w *CartWorkflow) ReserveTicketsActivity(ctx context.Context, input *cartv1.ReserveTicketsInput) (*cartv1.ReserveTicketsOutput, error) {
	panic("not implemented")
}
