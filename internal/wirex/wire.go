//go:build wireinject
// +build wireinject

package wirex

// The build tag makes sure the stub is not built in the final build.

import (
	"context"

	"github.com/google/wire"

	"github.com/fsyabc/test/internal/mods"
	"github.com/fsyabc/test/pkg/util"
)

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	wire.Build(
		InitCacher,
		InitDB,
		InitAuth,
		wire.NewSet(wire.Struct(new(util.Trans), "*")),
		wire.NewSet(wire.Struct(new(Injector), "*")),
		mods.Set,
	) // end
	return new(Injector), nil, nil
}
