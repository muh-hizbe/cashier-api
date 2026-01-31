package boot

import (
	"fmt"
	"net/http"

	"github.com/muh-hizbe/cashier-api/internal/config"
	"github.com/muh-hizbe/cashier-api/internal/database"
	"github.com/muh-hizbe/cashier-api/internal/route"
)

func Init() {
	cfg := config.LoadConfig()
	pool := database.ConnectPostgresPool(cfg.Database.URL)
	route.Init(pool)
	defer pool.Close()
	addr := "0.0.0.0:" + config.GetAppConfig().Port
	fmt.Println("Server running di localhost:", addr)
	http.ListenAndServe(addr, nil)
}
