package boot

import (
	"fmt"
	"net/http"

	"github.com/muh-hizbe/cashier-api/internal/config"
	"github.com/muh-hizbe/cashier-api/internal/route"
)

func Init() {
	var w http.ResponseWriter
	var r *http.Request
	route.Init(w, r)
	fmt.Println("Server running di localhost:", config.Port)
	http.ListenAndServe(config.Port, nil)
}
