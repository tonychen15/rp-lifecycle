package example

import (
	"net/http"

	"github.com/rp-lifecycle/framework"
)

func main() {
	server := &http.Server{
		Handler: framework.NewMicroService(),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
