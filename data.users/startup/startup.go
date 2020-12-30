package startup

import (
	"os"

	di "go.kicksware.com/api/service-common/container"
	"go.kicksware.com/api/service-common/core"

	"github.com/timoth-y/scrapnote-api/data.users/config"
	cnf "github.com/timoth-y/scrapnote-api/data.users/container/config"
)

func InitializeServer() (srv core.Server) {
	config, err := config.ReadServiceConfig(os.Getenv("CONFIG_PATH")); if err != nil {
		return nil
	}
	container := di.NewServiceContainer()
	cnf.ConfigureContainer(container, config)
	container.Resolve(&srv)
	return
}