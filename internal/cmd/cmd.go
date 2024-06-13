package cmd

import (
	"context"
	"gadmin-backend/internal/consts"
	"gadmin-backend/internal/controller/role"
	"gadmin-backend/internal/controller/system"
	"gadmin-backend/internal/service"
	"github.com/gogf/gf/v2/net/goai"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func:  mainFunc,
	}
)

func mainFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	var (
		s = g.Server()
	)

	oaiSetting()

	// bind
	baseGroup() // base => prefix: /
	roleGroup() // role => prefix: /role

	// TODO: user

	// TODO: menu

	s.Run()
	return nil
}

func oaiSetting() {
	var (
		s   = g.Server()
		oai = s.GetOpenApi()
	)

	s.SetSwaggerUITemplate(consts.SwaggerUITemplate)
	oai.Components.SecuritySchemes = make(goai.SecuritySchemes)
	oai.Components.SecuritySchemes["BearerAuth"] = goai.SecuritySchemeRef{
		Value: &goai.SecurityScheme{
			Type:   "http",
			Scheme: "bearer",
		},
	}
}

func baseGroup() {
	var (
		s = g.Server()
	)

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().CORS)
		group.Bind(
			system.NewV1().Login,
		)

		priv := group.Group("")
		priv.Middleware(service.Middleware().Auth)
		priv.Bind(
			system.NewV1().Logout,
			system.NewV1().UserInfo,
		)
	})
}

func roleGroup() {
	var (
		s = g.Server()
	)

	s.Group("/role", func(group *ghttp.RouterGroup) {
		// default, cors, auth
		group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().CORS, service.Middleware().Auth)

		group.Bind(
			role.NewV1(),
		)
	})
}
