package app

import (
	"fmt"

	"github.com/Kortivex/connected_roots/internal/connected_roots/config"
)

const (
	banner = `
 _____                             _           _  ______            _
/  __ \                           | |         | | | ___ \          | |
| /  \/ ___  _ __  _ __   ___  ___| |_ ___  __| | | |_/ /___   ___ | |_ ___
| |    / _ \| '_ \| '_ \ / _ \/ __| __/ _ \/ _| | |    // _ \ / _ \| __/ __|
| \__/\ (_) | | | | | | |  __/ (__| ||  __/ (_| | | |\ \ (_) | (_) | |_\__ \
 \____/\___/|_| |_|_| |_|\___|\___|\__\___|\__,_| \_| \_\___/ \___/ \__|___/

%s:
	* HTTP server will be running at [http://%s:%d]
	* Frontend server will be running at [http://%s:%d]
https://github.com/Kortivex/connected_roots
_____________________________________________________________________________________________

`
)

// GenerateBanner generates the final string to display the final banner to print.
func GenerateBanner(conf *config.Config) string {
	return fmt.Sprintf(banner, conf.App.Name, "localhost", conf.API.Port, "localhost", conf.Frontend.Port)
}
