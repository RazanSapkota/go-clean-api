package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewPostRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(NewUserRoutes),
)
