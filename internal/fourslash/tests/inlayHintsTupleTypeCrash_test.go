package fourslash_test

import (
	"testing"

	"github.com/sngl-lang/typescript-go/internal/core"
	"github.com/sngl-lang/typescript-go/internal/fourslash"
	"github.com/sngl-lang/typescript-go/internal/ls/lsutil"
	"github.com/sngl-lang/typescript-go/internal/testutil"
)

func TestInlayHintsTupleTypeCrash(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function iterateTuples(tuples: [string][]): void {
  tuples.forEach((l) => {})
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayFunctionParameterTypeHints: core.TSTrue,
		},
	})
}
