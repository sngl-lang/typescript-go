package fourslash_test

import (
	"testing"

	"github.com/sngl-lang/typescript-go/internal/fourslash"
	"github.com/sngl-lang/typescript-go/internal/testutil"
)

func TestDocumentHighlightNestedRequireDestructureNoCrash1(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: /bar.js
const { a: { b } } = require('./foo');
/**/b;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, "")
}
