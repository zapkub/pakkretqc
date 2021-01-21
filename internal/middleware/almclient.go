package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zapkub/pakkretqc/internal/conf"
	"github.com/zapkub/pakkretqc/internal/perrors"
	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

func MustGetALMClient(ctx context.Context) *almsdk.Client {
	var err error
	token, ok := GetSessionToken(ctx)
	if !ok {
		panic(fmt.Errorf("cannot get AML client token notfound: %w", perrors.Unauthenticated))
	}

	var almclient = almsdk.New(&almsdk.ClientOptions{
		Endpoint: conf.ALMEndpoint(),
	})
	err = almclient.Authenticate(ctx, token)
	if err != nil {
		panic(fmt.Errorf("cannot connect to alm %w", err))
	}
	return almclient
}

func ALMClient(n http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		n.ServeHTTP(rw, r)
	})
}
