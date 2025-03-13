package cartactor

import (
	"context"
	"log"

	cartv1 "github.com/taltemus/stacked-diff-example/gen/cart/v1"
)

type CartActor struct {
	cartActorCleint cartv1.CartActorClient
}

func NewCartActor(cartActorCleint cartv1.CartActorClient) *CartActor {
	return &CartActor{
		cartActorCleint: cartActorCleint,
	}
}

func (c *CartActor) StartWorkflow(ctx context.Context, tickets string) (string, error) {
	var cartID string
	// A bunch of business logic
	run, err := c.cartActorCleint.StartCartAsync(ctx, &cartv1.CartRequest{
		CartId:    cartID,
		TicketIds: tickets,
	})
	if err != nil {
		return "", err
	}
	log.Printf("Started cart workflow with run ID: %s", run.RunID())
	return cartID, nil
}
