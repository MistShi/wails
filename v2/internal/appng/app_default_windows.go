//go:build !dev && !production && !bindings && windows

package appng

import (
	"os/exec"

	"github.com/leaanthony/winc/w32"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// App defines a Wails application structure
type App struct{}

func (a *App) Run() error {
	return nil
}

// CreateApp creates the app!
func CreateApp(_ *options.App) (*App, error) {
	result := w32.MessageBox(0,
		`Wails applications will not build without the correct build tags. 
Please use "wails build" or press "OK" to open the documentation on how to use "go build"`,
		"Error",
		w32.MB_ICONERROR|w32.MB_OKCANCEL)
	if result == 1 {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://wails.io").Start()
	}
	return nil, nil
}