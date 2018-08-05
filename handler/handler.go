package handler

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/nongdenchet/identicon/endpoint"
	"github.com/nongdenchet/identicon/helpers"
	"github.com/nongdenchet/identicon/middleware"
	"github.com/nongdenchet/identicon/repository"
	"github.com/nongdenchet/identicon/service"
)

func NewHandler() *httptransport.Server {
	repo := repository.IdenticonRepoImpl{}

	var s endpoint.IdenticontService
	{
		s = service.IdenticonServiceImpl{Repo: repo}
		s = middleware.NewLoggingMiddleware(s)
		s = middleware.NewIntrumentationMiddleware(s)
	}

	return httptransport.NewServer(
		endpoint.MakeGenerateEndpoint(s),
		helpers.DecodeGenerateRequest,
		helpers.EncodeResponse,
	)
}
