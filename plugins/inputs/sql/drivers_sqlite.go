// +build linux,freebsd
// +build !mips !mips64

package sql

import (
	_ "modernc.org/sqlite"
)
