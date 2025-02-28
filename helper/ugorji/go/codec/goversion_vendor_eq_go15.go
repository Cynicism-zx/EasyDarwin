// Copyright (c) 2012-2018 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

//go:build go1.5 && !go1.6
// +build go1.5,!go1.6

package codec

import "os"

var genCheckVendor = os.Getenv("GO15VENDOREXPERIMENT") == "1"
