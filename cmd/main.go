package main

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/ldd27/go-starter-kit/cmd/cmd"
)

func main() {
	now.WeekStartDay = time.Monday
	if err := cmd.RootCMD.Execute(); err != nil {
		panic(err)
	}
}
