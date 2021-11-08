package init

import (
	"github.com/docker-slim/docker-slim/pkg/app/master/commands/backend"
)

func init() {
	backend.RegisterCommand()
}
