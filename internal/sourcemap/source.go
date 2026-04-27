package sourcemap

import "github.com/sngl-lang/typescript-go/internal/core"

type Source interface {
	Text() string
	FileName() string
	ECMALineMap() []core.TextPos
}
