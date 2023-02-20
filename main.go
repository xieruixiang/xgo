package main

import (
	"fmt"
	"time"
	"xgo/server"
)

func Login(ctx server.Context) {
	ctx.Response.Write([]byte("login"))
}

func Logout(ctx server.Context) {
	ctx.Response.Write([]byte("logout"))
}

func LogoutAbc(ctx server.Context) {
	ctx.Response.Write([]byte("logout-abc"))
}

func LogoutAbcGet(ctx server.Context) {
	ctx.Response.Write([]byte("logout-abc-get"))
}

func main() {
	s := server.NewServerHttp(8083, func(next server.Filter) server.Filter {
		return func(ctx server.Context) {
			start := time.Now().Nanosecond()
			next(ctx)
			end := time.Now().Nanosecond()
			fmt.Printf("耗时纳秒:%d \n", end-start)
		}
	})
	s.Route("POST", "/login", Login)
	s.Route("POST", "/logout", Logout)
	s.Route("POST", "/logout/gg/abc", LogoutAbc)
	s.Route("GET", "/logout/gg/abc", LogoutAbcGet)
	s.Start()
}
