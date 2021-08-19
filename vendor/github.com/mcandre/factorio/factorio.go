package factorio

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

// Version is semver.
const Version = "0.0.1"

// ArtifactToplevelDir names the output location for binaries.
const ArtifactToplevelDir = "bin"

// Banner names the default output subdirectory.
const Banner = "artifact-dev"

// PlatformBlocklist excludes some fragile targets.
var PlatformBlocklist = regexp.MustCompile(`(android\/.*)|(ios\/.*)`)

// BannerParameter controls the environment variable
// for overriding the default Banner subdirectory.
//
// Example configuration: FACTORIO_BANNER=hello-0.0.1
const BannerParameter = "FACTORIO_BANNER"

// PlatformBlocklistParameter controls the environment variable
// for overriding the default PlatformBlocklist.
//
// Example configuration: FACTORIO_PLATFORM_BLOCKLIST=//
const PlatformBlocklistParameter = "FACTORIO_PLATFORM_BLOCKLIST"

// Platform models a basic targetable execution configuration.
type Platform struct {
	// Os denotes a high-level environment.
	//
	// Example: "linux"
	Os string

	// Arch denotes a low-level environment.
	//
	// Example: "amd64"
	Arch string
}

// String renders a platform.
func (o Platform) String() string {
	return fmt.Sprintf("%s/%s", o.Os, o.Arch)
}

func Platforms() ([]Platform, error) {
	var platforms []Platform

	var distOut bytes.Buffer

	cmd := exec.Command("go")
	cmd.Args = []string{"go", "tool", "dist", "list"}
	cmd.Stderr = os.Stderr
	cmd.Stdout = bufio.NewWriter(&distOut)

	if err := cmd.Run(); err != nil {
		return platforms, err
	}

	scanner := bufio.NewScanner(&distOut)
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "/")

		if len(parts) < 2 {
			return platforms, fmt.Errorf("cannot parse platform metadata: %v", line)
		}

		platforms = append(platforms, Platform{Os: parts[0], Arch: parts[1]})
	}

	return platforms, nil
}

// Build generates binaries for the given platform.
func Build(platform Platform, artifactToplevelDir string, args []string) error {
	artifactDir := path.Join(artifactToplevelDir, platform.Os, platform.Arch)

	log.Printf("building %s\n", artifactDir)

	if err := os.MkdirAll(artifactDir, 0755); err != nil {
		return err
	}

	allPackagesPath := fmt.Sprintf(".%c...", os.PathSeparator)

	cmd := exec.Command("go")
	cmd.Args = []string{"go", "build", "-o", artifactDir}
	cmd.Args = append(cmd.Args, args...)
	cmd.Args = append(cmd.Args, allPackagesPath)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%s", platform.Os))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOARCH=%s", platform.Arch))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Port generates a multitude of binaries.
func Port(args []string) error {
	banner, ok := os.LookupEnv(BannerParameter)

	if !ok {
		banner = Banner
	}

	artifactToplevelDir := path.Join(ArtifactToplevelDir, banner)

	platformBlocklist := PlatformBlocklist

	platformBlocklistPattern, ok := os.LookupEnv(PlatformBlocklistParameter)

	if ok {
		pb, err := regexp.Compile(platformBlocklistPattern)

		if err != nil {
			panic(err)
		}

		platformBlocklist = pb
	}

	platforms, err := Platforms()

	if err != nil {
		return err
	}

	for _, platform := range platforms {
		if platformBlocklist.MatchString(platform.String()) {
			log.Printf("skipping %s", platform)
			continue
		}

		if err := Build(platform, artifactToplevelDir, args); err != nil {
			return err
		}
	}

	return nil
}
