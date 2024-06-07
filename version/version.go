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


    __         __                |
   / /__  ____/ /_ ___  __ _  __ | Maintainer : {{.Maintainer}}
  / / _ \/ __/   // -_) _ \ |/ / | Version: {{.Version}}
 /_/\___/\__/_/\_\\__/_//_/___/  | Build: {{.Build}}
                                 | Build Date: {{.Date}} 
                                 |              
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
