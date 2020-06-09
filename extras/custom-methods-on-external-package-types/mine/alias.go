package mine

import (
	"fmt"

	"github.com/jboursiquot/iis-go-foundation/extras/custom-methods-on-external-package-types/theirs"
)

func useAliasType() {
	ct := theirs.CustomType{}
	fmt.Println(ct)
}

// DoMore is not allowed here because we cannot
// add methods to a type we do not own in this package.
// func (ct theirs.CustomType) DoMore() {}

// However, we can create a local alias based on the
// external type and attach methods to it.
// type aliasCustomType theirs.CustomType
// func (a *aliasCustomType) DoMore() {}
