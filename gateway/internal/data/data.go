package data

import (
	"gateway/internal/conf"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// TODO warpped database client
}

// NewData .
func NewData(c *conf.Data) (*Data, error) {
	return &Data{}, nil
}
