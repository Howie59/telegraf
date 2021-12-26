//go:build goplugin
// +build goplugin

package goplugin

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"plugin"
	"strings"
)

// loadExternalPlugins 从指定路径下共享库中加载额外的插件(ilogtail中也有类似的机制)
func LoadExternalPlugins(rootDir string) error {
	return filepath.Walk(rootDir, func(pth string, info os.FileInfo, err error) error {
		// Stop if there was an error.
		if err != nil {
			return err
		}

		// Ignore directories.
		if info.IsDir() {
			return nil
		}

		// Ignore files that aren't shared libraries.
		ext := strings.ToLower(path.Ext(pth))
		if ext != ".so" && ext != ".dll" {
			return nil
		}

		// Load plugin.
		_, err = plugin.Open(pth)
		if err != nil {
			return fmt.Errorf("error loading %s: %s", pth, err)
		}

		return nil
	})
}
