package workflow

import (
	"fmt"
	"log"
	"strings"
	"time"

	cartv1 "github.com/taltemus/stacked-diff-example/gen/cart/v1"
	"go.temporal.io/sdk/workflow"
)

type CartWorkflowReq struct {
	Request *cartv1.CartRequest
}

// CartWorkflowImpl implements the shopping cart workflow.
type CartWorkflow struct {
	*CartWorkflowReq
	cancelled bool
}

// Run executes the workflow.
func (w *CartWorkflow) Run(ctx workflow.Context) (*cartv1.CartResponse, error) {
	log.Printf("Workflow started for cart: %s", w.Request.CartId)

	// Split comma-separated ticket IDs.
	ticketIDs := strings.Split(w.Request.TicketIds, ",")

	// Call ReserveTickets (placeholder).
	reserveStatus, _ := cartv1.ReserveTicketsActivity(ctx, &cartv1.ReserveTicketsInput{
		CartId:    w.Request.CartId,
		TicketIds: ticketIDs,
	})

	selector := workflow.NewSelector(ctx)
	selector.AddReceive(cartv1.CancelCartSignalSignalChannel, func(c workflow.ReceiveChannel, more bool) {
		log.Println("Cancellation signal received.")
		w.cancelled = true
	})
	selector.AddFuture(workflow.NewTimer(ctx, 5*time.Minute), func(f workflow.Future) {
		log.Println("Timeout reached.")
		w.cancelled = true
	})
	for selector.HasPending() {
		selector.Select(ctx)
		if w.cancelled {
			break
		}
	}

	var removeStatus string
	if w.cancelled {
		resp, _ := cartv1.RemoveTicketsActivity(ctx, &cartv1.RemoveTicketsInput{
			CartId:    w.Request.CartId,
			TicketIds: ticketIDs,
		})
		removeStatus = resp.Status
	} else {
		cartv1.CheckoutCartActivity(ctx, &cartv1.CheckoutCartInput{})
	}

	return &cartv1.CartResponse{
		Confirmation: fmt.Sprintf("Cart %s: Reserve [%s] -> Remove [%s]",
			w.Request.CartId, reserveStatus, removeStatus),
	}, nil
}
