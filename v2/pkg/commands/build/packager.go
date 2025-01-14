package build

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/wailsapp/wails/v2/internal/fs"
)

// PackageProject packages the application
func packageProject(options *Options, platform string) error {

	var err error
	switch platform {
	case "darwin", "windows":
		err = packageApplication(options)
	default:
		err = fmt.Errorf("packing not supported for %s yet", platform)
	}

	if err != nil {
		return err
	}

	return nil
}

// cleanBuildDirectory will remove an existing build directory and recreate it
func cleanBuildDirectory(options *Options) error {

	buildDirectory := options.BuildDirectory

	// Clear out old builds
	if fs.DirExists(buildDirectory) {
		err := os.RemoveAll(buildDirectory)
		if err != nil {
			return err
		}
	}

	// Create clean directory
	err := os.MkdirAll(buildDirectory, 0700)
	if err != nil {
		return err
	}

	return nil
}

// Gets (and creates) the build base directory
func getBuildBaseDirectory(options *Options) (string, error) {
	buildDirectory := filepath.Join(options.ProjectData.Path, "build")
	if !fs.DirExists(buildDirectory) {
		err := os.MkdirAll(buildDirectory, 0700)
		if err != nil {
			return "", err
		}
	}
	return buildDirectory, nil
}

// Gets the platform dependent package assets directory
func getPackageAssetsDirectory() string {
	return fs.RelativePath("internal/packager", runtime.GOOS)
}
