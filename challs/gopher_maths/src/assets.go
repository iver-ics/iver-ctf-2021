package main

import _ "embed"

//go:embed assets/banner.txt
var assetsBanner string

//go:embed assets/ohno.txt
var assetsOhNo string

//go:embed assets/rfc1436.txt
var assetsRFC1436 []byte
