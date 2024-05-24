package version

import (
	"html/template"
	"strings"
)

type BuildMeta struct {
	Version    string
	Maintainer string
	Build      string
	Date       string
}

var (
	Version    = ""
	Maintainer = ""
	BuildNo    = ""
	Date       = ""
)

var Art = `
 ______                 ___  _  _     | 
|___  /                / _ \| || |    | Maintainer : {{.Maintainer}}
   / / __    __ _  ___| | | | || |_   |	Version: {{.Version}}
  / / / _ \ / _  |/  _| | | |__   _|  |	Build: {{.Build}}
 / /_| (_| | (_| | (__| |_| |  | |    | Build Date: {{.Date}} 
/_____\__,_|\__,_|\___|\___/   |_|    |
`

func PrintVersion() string {

	build := BuildMeta{
		Version:    Version,
		Maintainer: Maintainer,
		Build:      BuildNo,
		Date:       Date,
	}
	tmpl, _ := template.New("art").Parse(Art)
	var result strings.Builder
	tmpl.Execute(&result, build)
	return result.String()
}
