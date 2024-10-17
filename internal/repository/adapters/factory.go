package adapters

import (
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository"
	"github.com/JTGlez/GoWeb-IT_V2/internal/repository/adapters/memory"
	"os"
)

func NewRepository() (repository.DataInterface, error) {
	adapter := os.Getenv("DATA_SOURCE")

	if adapter == "" {
		return nil, repository.ErrorMissingAdapter
	}

	switch adapter {
	case "json":
		return memory.NewDatabase(), nil
	default:
		return nil, repository.ErrorUnimplementedAdapter
	}

}
