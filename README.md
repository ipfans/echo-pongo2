echo-pongo2
======

Middleware echo-pongo2 is a [pongo2](https://github.com/flosch/pongo2) template engine support for [echo](https://github.com/labstack/echo/).

### Installation

	go get github.com/ipfans/echo-pongo2

## Example

```go
package main

import (
	"github.com/ipfans/echo-pongo2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	serv := echo.New()
	serv.Use(middleware.Logger())
	serv.Use(middleware.Recover())
	serv.Use(pongo2.Pongo2())
	serv.Get("/", func(ctx *echo.Context) error {
		// ctx.Set("ContentType") = "text/html"
		// ctx.Set("encoding") = "UTF-8"
		ctx.Set("template", "index.html")
		ctx.Set("data", map[string]interface{}{
			"user": "ipfans",
		})
		return nil
	})

	serv.Run("127.0.0.1:8080")
}
```

## License

This project is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text.
