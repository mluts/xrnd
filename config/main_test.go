package config

import (
	"testing"
)

var configExample = struct {
	path         string
	layoutsCount int
	layoutString map[LayoutName]string
}{
	"testdata/config.yml",
	1,
	map[LayoutName]string{
		"s": "--output HDMI1 --auto --output eDP1 --left-of HDMI1 --primary",
	},
}

func TestConfig(t *testing.T) {
	cfg, err := Read(configExample.path)

	if err != nil {
		t.Errorf("Expected %v to be nil", err)
	}

	if len(cfg.Layouts) != configExample.layoutsCount {
		t.Errorf("Expected layouts count to be %d, but have %d",
			configExample.layoutsCount,
			len(cfg.Layouts))
	}

	for layoutName, layoutOptions := range configExample.layoutString {
		layout := cfg.Layouts[layoutName]

		if layout.String() != layoutOptions {
			t.Errorf("Expected layout to be\n%q,\n but was\n%q",
				layoutOptions,
				layout.String())
		}
	}
}
