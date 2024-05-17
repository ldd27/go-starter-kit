package version

import (
	"fmt"
	"runtime"
)

var (
	version     = ""
	gitRevision = ""
	gitBranch   = ""
	buildTime   = ""
	goVersion   = runtime.Version()
	oSArch      = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

func Print() {
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Git revision: %s\n", gitRevision)
	fmt.Printf("Git branch: %s\n", gitBranch)
	fmt.Printf("Go Version: %s\n", goVersion)
	fmt.Printf("OS/Arch: %s\n", oSArch)
	fmt.Printf("Build Time: %s\n", buildTime)
}
