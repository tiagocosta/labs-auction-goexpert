package auction

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) UpdateAuctionStatus(ctx context.Context, auction *auction_entity.Auction) *internal_error.InternalError {
	filter := bson.D{primitive.E{Key: "_id", Value: auction.Id}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "status", Value: auction.Status},
	},
	}}

	_, err := ar.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return internal_error.NewInternalServerError("error trying to update auction status")
	}

	return nil
}
