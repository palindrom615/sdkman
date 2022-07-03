package util

import "runtime"

func Platform() string {
	platform := runtime.GOOS
	isx64 := runtime.GOARCH == "amd64"
	isx86 := runtime.GOARCH == "386" || runtime.GOARCH == "amd64p32"
	isArm := runtime.GOARCH == "arm"
	isArm64 := runtime.GOARCH == "arm64"

	if platform == "windows" {
		platform = "MSYS_NT-10.0"
		if isx86 {
			platform = "MINGW32_NT-6.2"
		}
	} else {
		if isx86 {
			platform += "x32"
		}
		if isx64 {
			platform += "x64"
		}
		if isArm {
			platform += "arm32"
		}
		if isArm64 {
			platform += "arm64"
		}
	}

	return platform
}
