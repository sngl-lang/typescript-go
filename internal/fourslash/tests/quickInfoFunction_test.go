package fourslash_test

import (
	"testing"

	"github.com/sngl-lang/typescript-go/internal/fourslash"
	"github.com/sngl-lang/typescript-go/internal/testutil"
)

func TestQuickInfoFunction(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/**/function foo() { return "hi"; }`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "", "function foo(): string", "")
}
