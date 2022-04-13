package service

import (
	"usm/internal/service/account"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(account.NewService)
