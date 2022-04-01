package route

import (
	"search_server/server/model"
	"search_server/server/services"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	var handles = services.Services{}

	r := e.Group("/migration")

	r.GET("", func(c echo.Context) error {
		go handles.MigrationAll()
		return c.JSON(200, "ok")
	})

	r.GET("/migration-index", func(c echo.Context) error {
		go handles.MigrationProduct()
		return c.JSON(200, "ok")
	})
	r.GET("/index-orders", func(c echo.Context) error {
		go handles.MigrationOrder(true)
		return c.JSON(200, "ok")
	})
	r.GET("/index-users", func(c echo.Context) error {
		go handles.MigrationUser()
		return c.JSON(200, "ok")
	})
	r.GET("/index-keywords", func(c echo.Context) error {
		go handles.MigrationKeyword()
		return c.JSON(200, "ok")
	})

	rSearch := e.Group("/search")
	rSearch.GET("", func(c echo.Context) error {
		var (
			typeSearch = c.QueryParam("search_type")
			q          = &model.CommonQuery{
				Page:    0,
				Limit:   20,
				Keyword: c.QueryParam("keyword"),
			}
			ctx = c.Request().Context()
		)
		switch typeSearch {
		case "product":
			data := handles.SearchProductAppByES(ctx, q)
			return c.JSON(200, echo.Map{
				"data": data,
			})
		case "order":
			data, _ := handles.SearchOrderAdminByES(ctx, q)
			return c.JSON(200, echo.Map{
				"data": data,
			})
		case "user":
			data, _ := handles.SearchUserAdminByES(ctx, q)
			return c.JSON(200, echo.Map{
				"data": data,
			})
		case "keyword":
			data, _ := handles.SearchKeywordAppByES(ctx, q)
			return c.JSON(200, echo.Map{
				"data": data,
			})
		}
		return c.JSON(200, echo.Map{
			"data":    "",
			"message": "type not support",
		})
	})
}
