package assetsPkg

import "embed"

//go:embed "*"
var EmbeddedAssets embed.FS
