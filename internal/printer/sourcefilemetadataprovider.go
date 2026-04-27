package printer

import (
	"github.com/sngl-lang/typescript-go/internal/ast"
	"github.com/sngl-lang/typescript-go/internal/tspath"
)

type SourceFileMetaDataProvider interface {
	GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData
}
