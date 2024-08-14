package shared

import (
	_ "embed"
)

var (
	//go:embed "styles/host-background-color.css"
	StyleHostBackgroundColor []byte
	//go:embed "styles/host-foreground-color.css"
	StyleHostForegroundColor []byte
	//go:embed "styles/host-row-gap.css"
	StyleHostRowGap []byte
	//go:embed "styles/host-column-gap.css"
	StyleHostColumnGap []byte
)
