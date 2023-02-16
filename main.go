package main

import "xgo/server"

func Login(ctx server.Context) {
	ctx.Response.Write([]byte("login"))
}
func Logout(ctx server.Context) {
	ctx.Response.Write([]byte("logout"))
}

func main() {
	s := server.HttpServer{
		Port: 8083,
		Handler: server.HandlerOnMap{
			HandlerMap: map[string]func(ctx server.Context){},
		},
	}
	s.Route("POST", "/login", Login)
	s.Route("POST", "/logout", Logout)
	s.Start()
}
