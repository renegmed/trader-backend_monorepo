package mongo

import (
	"context"
	"fmt"

	"github.com/renegmed/trader-backend_monorepo/internal/domain"
	"github.com/renegmed/trader-backend_monorepo/internal/ports"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type strategiesRepos struct {
	strategiesColl *mongo.Collection
}

func NewStrategiesRepository(strategiesColl *mongo.Collection) ports.StrategiesRepository {
	repo := &strategiesRepos{
		strategiesColl: strategiesColl,
	}

	return repo
}

func (repo *strategiesRepos) Insert(ctx context.Context, strategy *domain.Strategy) (string, error) {
	// translate domain.Strategy into StrategyDTO 
	strategyDTO, err := fromStrategyCoreToDTO(strategy)
	if err != nil {
		return "", err
	}
	result, err := repo.strategiesColl.InsertOne(ctx, strategyDTO)
	if err != nil {
		return "", fmt.Errorf("error on insert strategy: %w", err)
	}

	strategyID, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return "", fmt.Errorf("error getting inserted id from mongo")
	}

	return strategyID.Hex(), nil

}

func (repo *strategiesRepos) GetByID(ctx context.Context, id string) (*domain.Strategy, error) {
	// TODO: Get strategy from MongoDB
	panic("implement me")
}
