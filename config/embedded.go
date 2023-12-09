package config

import _ "embed"

//go:embed embedded/BUILDTOKEN.key
var OFFICIAL_BUILD_TOKEN string

//go:embed embedded/official_priv.key
var OFFICIAL_TLS_PRIV_KEY string

//go:embed embedded/official_servpub.crt
var OFFICIAL_TLS_PUB_KEY string
