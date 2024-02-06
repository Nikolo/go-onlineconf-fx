package onlineconf_fx

import (
	"context"

	"github.com/Nikolo/go-onlineconf/pkg/onlineconf"
	"github.com/Nikolo/go-onlineconf/pkg/onlineconfInterface"
	"go.uber.org/fx"
)

func Module(name string) fx.Option {
	return fx.Module(
		"Onlineconf",
		fx.Invoke(
			func(ls fx.Lifecycle, c onlineconfInterface.Instance) {
				ls.Append(
					fx.Hook{
						OnStart: func(ctx context.Context) error {
							return c.StartWatcher(ctx)
						},
						OnStop: func(_ context.Context) error { return c.StopWatcher() },
					},
				)
			},
		),

		fx.Provide(func(options []onlineconfInterface.Option) onlineconfInterface.Instance {
			return onlineconf.Create(options...)
		}),
	)
}
