package middleware

import (
	"myapp/data"

	"github.com/dendosan/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
