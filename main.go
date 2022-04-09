package main

import (
	gohttp "net/http"
	"time"

	"github.com/himynamej/api-test.git/repository/mocked"
	"github.com/himynamej/api-test.git/service/accounts"
	"github.com/himynamej/api-test.git/transport/http"
	"github.com/pingcap/errors"
)

func main() {
	//ctx := context.Background()
	userRepo := mocked.NewUserRepository()
	tokenRepo := mocked.NewTokenRepository()
	srv := accounts.New(userRepo, tokenRepo, "ali", "123", "ali/123", accounts.SetPasswordCost(5))
	h := http.NewHandler(srv, time.Local)
	err := gohttp.ListenAndServe(":8080", h)
	if err != nil {
		panic(errors.Wrap(err, "error on listen and serve http"))
	}
}
