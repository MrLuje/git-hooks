// +build !windows

package main

import "os"

func isExecutable(info os.FileInfo, folder string) bool {
	return info.Mode()&0111 != 0
}
