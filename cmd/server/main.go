package main

import (
	"github.com/tqlong1609/go_backend_ecommerce/internal/routers"
)

func main() {
	r := routers.Init()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") - change port: r.Run(":8002")
}
