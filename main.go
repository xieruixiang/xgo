package main

import "xgo/server"

func Login(ctx server.Context) {
	ctx.Response.Write([]byte("login"))
}
func Logout(ctx server.Context) {
	ctx.Response.Write([]byte("logout"))
}

func main() {
	s := server.NewServerHttp(8083)
	s.Route("POST", "/login", Login)
	s.Route("POST", "/logout", Logout)
	s.Start()
}
