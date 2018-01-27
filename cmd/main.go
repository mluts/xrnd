package main

import (
	"fmt"
	"github.com/mluts/xrnd/config"
	"os"
	"os/exec"
	"strings"
)

var configPath = "${HOME}/.config/xrnd/config.yml"

func init() {
	configPath = os.ExpandEnv(configPath)
}

func edit(path string) {
	editor := os.ExpandEnv("${EDITOR}")
	if len(editor) == 0 {
		editor = "vi"
	}

	fmt.Println("Running", editor, path)
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func xrandr(optionsStr string) {
	options := strings.Split(optionsStr, " ")
	fmt.Printf("Running: xrandr %v", options)

	cmd := exec.Command("xrandr", options...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func abort(f string, msg ...interface{}) {
	fmt.Printf(f, msg...)
	os.Exit(1)
}

func layoutName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}

	return ""
}

func main() {
	layout := layoutName()

	cfg, err := config.Read(configPath)

	if err != nil {
		abort("Can't read config %s: %v\n", configPath, err)
	}

	if len(layout) == 0 {
		fmt.Println("Layouts:")
		for name := range cfg.Layouts {
			fmt.Println("-", name)
		}
	} else if layout == "edit" {
		edit(configPath)
		os.Exit(0)
	} else {
		layoutConfig, ok := cfg.Layouts[config.LayoutName(layout)]
		if !ok {
			abort("Layout %s doesn't exist\n", layout)
		}
		xrandr(layoutConfig.String())
	}
}
