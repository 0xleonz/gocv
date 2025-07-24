package assets

import "embed"

//go:embed ../../assets/config.yml
//go:embed ../../assets/templates/*
var Files embed.FS
