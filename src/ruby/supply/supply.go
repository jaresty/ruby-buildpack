package supply

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Stager interface {
	LinkDirectoryInDepDir(destDir, depSubDir string) error
}
type Manifest interface {
	// DefaultVersion(depName string) (libbuildpack.Dependency, error)
	// InstallDependency(dep libbuildpack.Dependency, outputDir string) error
	InstallOnlyVersion(depName string, installDir string) error
	// AllDependencyVersions(string) []string
}
type Command interface {
	Execute(dir string, stdout io.Writer, stderr io.Writer, program string, args ...string) error
}
type Logger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	Protip(tip string, help_url string)
}

type Supply struct {
	Stager   Stager
	Manifest Manifest
	Log      Logger
	Command  Command
}

func (s *Supply) InstallBundler() error {
	if err := s.Manifest.InstallOnlyVersion("bundler", filepath.Join(c.DepDir, "bundler")); err != nil {
		return err
	}

	gempath := new(bytes.Buffer)
	if err := s.Command(c.Stager.BuildDir(), gempath, io.Discard, "gem", "env", "gempath"); err != nil {
		return err
	}
	os.Setenv("GEM_PATH", filepath.Join(c.DepDir, "bundler")+":"+strings.Trim(gempath.String()))

	buffer := new(bytes.Buffer)
	if err := s.Command(c.Stager.BuildDir(), buffer, io.Discard, "bundle", "platform", "--ruby"); err != nil {
		return err
	}

	return nil
}

func (c *Supply) InstallBins() error {
	// c.Log.Info("engines.node (package.json):  " + c.Engines.Node) // TODO "unspecified" see https://github.com/dgodd/nodejs-buildpack/blob/master/bin/compile#L99
	// c.Log.Info("engines.npm (package.json):   " + c.Engines.Npm)  // TODO "unspecified (use default)" see https://github.com/dgodd/nodejs-buildpack/blob/master/bin/compile#L100
	// c.warnNodeEngine()

	// if err := os.MkdirAll(filepath.Join(c.DepDir, "bin"), 0755); err != nil {
	// 	return err
	// }
	// os.Setenv("PATH", filepath.Join(c.DepDir, "bin")+":"+os.Getenv("PATH"))

	// if err := c.InstallNodejs(); err != nil {
	// 	c.Log.Error("Failed to install node")
	// 	return err
	// }

	// if err := c.InstallNpm(); err != nil {
	// 	c.Log.Error("Failed to install npm: %v", err)
	// 	return err
	// }

	// if c.isYarn() {
	// 	if err := c.InstallYarn(); err != nil {
	// 		c.Log.Error("Failed to install yarn: %v", err)
	// 		return err
	// 	}
	// }

	return nil
}

// func (c *Supply) InstallNodejs() error {
// 	version := c.Engines.Node
// 	if version == "" {
// 		dep, err := c.Manifest.DefaultVersion("node")
// 		if err != nil {
// 			return err
// 		}
// 		version = dep.Version
// 	} else {
// 		versionConstraint := version
// 		versions := c.Manifest.AllDependencyVersions("node")
// 		if matchingVersion, err := libbuildpack.FindMatchingVersion(versionConstraint, versions); err == nil {
// 			version = matchingVersion
// 		}
// 	}

// 	dep := libbuildpack.Dependency{Name: "node", Version: version}
// 	if err := c.Manifest.InstallDependency(dep, c.DepDir); err != nil {
// 		return err
// 	}
// 	if err := rename(c.DepDir, "node-v*", "node"); err != nil {
// 		return err
// 	}
// 	if err := os.Symlink(filepath.Join("..", "node", "bin", "npm"), filepath.Join(c.DepDir, "bin", "npm")); err != nil {
// 		return err
// 	}
// 	return os.Symlink(filepath.Join("..", "node", "bin", "node"), filepath.Join(c.DepDir, "bin", "node"))
// }
