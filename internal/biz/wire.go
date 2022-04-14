package biz

import (
	"usm/internal/biz/usecase/account"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	account.NewUsecase,
)
