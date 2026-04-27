// Package snglts re-exports the typescript-go internals that SNGL's
// node:// import handler needs. Upstream's API is unstable and lives
// under internal/, so this package is the controlled surface SNGL
// pins against. Type and var aliases only — no logic.
package snglts

import (
	"github.com/sngl-lang/typescript-go/internal/ast"
	"github.com/sngl-lang/typescript-go/internal/core"
	"github.com/sngl-lang/typescript-go/internal/parser"
	"github.com/sngl-lang/typescript-go/internal/scanner"
	"github.com/sngl-lang/typescript-go/internal/tspath"
)

// AST and parser types.
type (
	SourceFile             = ast.SourceFile
	SourceFileParseOptions = ast.SourceFileParseOptions
	Node                   = ast.Node
	NodeList               = ast.NodeList
	Kind                   = ast.Kind
	ModifierFlags          = ast.ModifierFlags
	ScriptKind             = core.ScriptKind
	Path                   = tspath.Path
)

// ParseSourceFile parses TS/JS source text into a SourceFile.
var ParseSourceFile = parser.ParseSourceFile

// ToPath produces the canonical Path used by the parser and module resolver.
var ToPath = tspath.ToPath

// Position helpers — convert Node.Pos()/End() (byte offsets) to 0-based
// line + UTF-16 character. SNGL converts to 1-based line/column.
var (
	GetECMALineAndUTF16CharacterOfPosition = scanner.GetECMALineAndUTF16CharacterOfPosition
	GetECMALineAndByteOffsetOfPosition     = scanner.GetECMALineAndByteOffsetOfPosition
)

// Script kinds.
const (
	ScriptKindUnknown  = core.ScriptKindUnknown
	ScriptKindJS       = core.ScriptKindJS
	ScriptKindJSX      = core.ScriptKindJSX
	ScriptKindTS       = core.ScriptKindTS
	ScriptKindTSX      = core.ScriptKindTSX
	ScriptKindJSON     = core.ScriptKindJSON
	ScriptKindExternal = core.ScriptKindExternal
)

// Node kinds — declarations and type expressions SNGL inspects.
const (
	KindSourceFile           = ast.KindSourceFile
	KindFunctionDeclaration  = ast.KindFunctionDeclaration
	KindClassDeclaration     = ast.KindClassDeclaration
	KindInterfaceDeclaration = ast.KindInterfaceDeclaration
	KindTypeAliasDeclaration = ast.KindTypeAliasDeclaration
	KindEnumDeclaration      = ast.KindEnumDeclaration
	KindEnumMember           = ast.KindEnumMember
	KindVariableStatement    = ast.KindVariableStatement
	KindVariableDeclaration  = ast.KindVariableDeclaration
	KindExportDeclaration    = ast.KindExportDeclaration
	KindExportAssignment     = ast.KindExportAssignment
	KindExportSpecifier      = ast.KindExportSpecifier
	KindModuleDeclaration    = ast.KindModuleDeclaration

	KindParameter         = ast.KindParameter
	KindPropertySignature = ast.KindPropertySignature
	KindMethodSignature   = ast.KindMethodSignature
	KindIndexSignature    = ast.KindIndexSignature
	KindTypeReference     = ast.KindTypeReference
	KindTypeLiteral       = ast.KindTypeLiteral
	KindFunctionType      = ast.KindFunctionType
	KindArrayType         = ast.KindArrayType
	KindTupleType         = ast.KindTupleType
	KindUnionType         = ast.KindUnionType
	KindIntersectionType  = ast.KindIntersectionType
	KindLiteralType       = ast.KindLiteralType
	KindTypeQuery         = ast.KindTypeQuery
	KindParenthesizedType = ast.KindParenthesizedType
	KindOptionalType      = ast.KindOptionalType

	KindStringKeyword    = ast.KindStringKeyword
	KindNumberKeyword    = ast.KindNumberKeyword
	KindBigIntKeyword    = ast.KindBigIntKeyword
	KindBooleanKeyword   = ast.KindBooleanKeyword
	KindVoidKeyword      = ast.KindVoidKeyword
	KindNullKeyword      = ast.KindNullKeyword
	KindUndefinedKeyword = ast.KindUndefinedKeyword
	KindAnyKeyword       = ast.KindAnyKeyword
	KindUnknownKeyword   = ast.KindUnknownKeyword
	KindNeverKeyword     = ast.KindNeverKeyword
	KindObjectKeyword    = ast.KindObjectKeyword

	KindIdentifier = ast.KindIdentifier
)

// Modifier flags.
const (
	ModifierFlagsExport = ast.ModifierFlagsExport
	ModifierFlagsAsync  = ast.ModifierFlagsAsync
	ModifierFlagsNone   = ast.ModifierFlagsNone
)
