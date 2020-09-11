package app

import (
	"context"

	pb "github.com/voidwell/ps2census.rpc"
	"google.golang.org/grpc"
)

type PS2CensusClient struct {
	client pb.PS2CensusClient
}

func NewPS2CensusClient(ps2censusEndpoint string) (*PS2CensusClient, error) {
	conn, err := grpc.Dial(ps2censusEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &PS2CensusClient{
		client: pb.NewPS2CensusClient(conn),
	}, nil
}

func (c *PS2CensusClient) GetCharacter(characterID string) (*pb.CharacterResult, error) {
	character, err := c.client.GetCharacter(context.Background(), &pb.CharacterQuery{CharacterID: characterID})
	if err != nil {
		return nil, err
	}

	return character, nil
}

func (c *PS2CensusClient) GetOutfit(outfitID string) (*pb.OutfitResult, error) {
	outfit, err := c.client.GetOutfit(context.Background(), &pb.OutfitQuery{OutfitID: outfitID})
	if err != nil {
		return nil, err
	}

	return outfit, nil
}
