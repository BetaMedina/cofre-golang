package factories

import (
	auth_middleware "secrets-golang/internal/auth/middlewares"
	config "secrets-golang/internal/infra"
)

var AuthorizedMiddlewareFactory auth_middleware.AuthorizedMiddleware

func init() {
	AuthorizedMiddlewareFactory = auth_middleware.NewAuthorizedMiddleware(config.NewTokenInfra())
}
