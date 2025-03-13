package workflow

import (
	cartv1 "github.com/taltemus/stacked-diff-example/gen/cart/v1"
	"go.temporal.io/sdk/workflow"
)

type CartWorkflowReq struct {
	Request  *cartv1.CartRequest
	Response *cartv1.CartResponse
}

// CartWorkflowImpl implements the shopping cart workflow.
type CartWorkflow struct {
	*CartWorkflowReq
	cancelled bool
}

// Run executes the workflow.
func (w *CartWorkflow) Run(ctx workflow.Context) {
	panic("not implemented")
}
