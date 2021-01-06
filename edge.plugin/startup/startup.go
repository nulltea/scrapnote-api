package startup

import (
	"os"

	di "github.com/timoth-y/scrapnote-api/lib.common/container"
	"github.com/timoth-y/scrapnote-api/lib.common/core"

	"github.com/timoth-y/scrapnote-api/edge.plugin/config"
	cnf "github.com/timoth-y/scrapnote-api/edge.plugin/container/config"
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

