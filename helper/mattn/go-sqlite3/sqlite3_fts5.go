// Copyright (C) 2014 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//go:build fts5
// +build fts5

package sqlite3

/*
#cgo CFLAGS: -DSQLITE_ENABLE_FTS5
#cgo LDFLAGS: -lm
*/
import "C"
