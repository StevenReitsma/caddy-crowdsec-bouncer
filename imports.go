package crowdsec

import (
	// Import the default CrowdSec modules. Primary reason this
	// file exists is to satisfy the Caddy documentation and download
	// pages to list the modules correctly.
	_ "github.com/StevenReitsma/caddy-crowdsec-bouncer/appsec"
	_ "github.com/StevenReitsma/caddy-crowdsec-bouncer/crowdsec"
	_ "github.com/StevenReitsma/caddy-crowdsec-bouncer/http"
	_ "github.com/StevenReitsma/caddy-crowdsec-bouncer/layer4"
)
