package mongo

import (
	"fmt"

	"github.com/renegmed/trader-backend_monorepo/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type StrategyDTO struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
}

func fromStrategyCoreToDTO(input *domain.Strategy) (*StrategyDTO, error) {
	if input == nil {
		return nil, fmt.Errorf("invalid input strategy")
	}
	
	id, err := bson.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid strategy id: %w", err)
	}

	result := &StrategyDTO{  // this will directly put to heap, which save compiler process
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}

	return result, nil
}

func fromStrategyDTOToCore(input StrategyDTO) (domain.Strategy, error) {

	result := domain.Strategy{
		ID:          input.ID.Hex(),
		Name:        input.Name,
		Description: input.Description,
	}

	return result, nil
}
