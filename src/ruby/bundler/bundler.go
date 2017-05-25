package bundler

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Stager interface {
	BuildDir() string
	LinkDirectoryInDepDir(destDir, depSubDir string) error
}
type Manifest interface {
	InstallOnlyVersion(depName string, installDir string) error
}
type Command interface {
	Execute(dir string, stdout io.Writer, stderr io.Writer, program string, args ...string) error
}
type Logger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	Protip(tip string, help_url string)
}

type Bundler struct {
	Stager   Stager
	Manifest Manifest
	Log      Logger
	Command  Command
}

func (b *Bundler) Install() error {
	if err := s.Manifest.InstallOnlyVersion("bundler", filepath.Join(c.DepDir, "bundler")); err != nil {
		return err
	}

	gempath := new(bytes.Buffer)
	if err := s.Command(c.Stager.BuildDir(), gempath, io.Discard, "gem", "env", "gempath"); err != nil {
		return err
	}
	os.Setenv("GEM_PATH", filepath.Join(c.DepDir, "bundler")+":"+strings.Trim(gempath.String()))

	// FIXME: Set GEM_PATH in env dir

	return nil
}

func (b *Bundler) RubyVersion() error {
	buffer := new(bytes.Buffer)
	if err := s.Command(c.Stager.BuildDir(), buffer, io.Discard, "bundle", "platform", "--ruby"); err != nil {
		return err
	}

	return nil
}
