package version

import (
	"fmt"
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

var url = "https://github.com/zaac04/locksmith"

var Art = `
 _             _                _       _        |
| |           | |              (_)  _  | |       | Maintainer : {{.Maintainer}}
| | ___   ____| |  _  ___ ____  _ _| |_| |__     | Version: {{.Version}}
| |/ _ \ / ___) |_/ )/___)    \| (_   _)  _ \    | Build: {{.Build}}
| | |_| ( (___|  _ (|___ | | | | | | |_| | | |   | Build Date: {{.Date}} 
 \_)___/ \____)_| \_|___/|_|_|_|_|  \__)_| |_|   |
                                                 
`

func PrintVersion() string {
	Maintainer = fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, Maintainer)
	build := BuildMeta{
		Version:    Version,
		Maintainer: Maintainer,
		Build:      BuildNo,
		Date:       Date,
	}
	tmp, _ := template.New("art").Parse(Art)
	var result strings.Builder
	tmp.Execute(&result, build)
	return result.String()
}
