// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: c575929542
// Version Date: 2020年 6月17日 星期三 04时40分15秒 UTC

package main

import (
	"flag"

	// This Service
	"github.com/rachihoaoi/daily-exercise.git/truss_sample/internal/hello/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(server.DefaultConfig)
}
