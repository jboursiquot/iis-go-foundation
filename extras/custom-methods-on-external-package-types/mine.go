package mine

import (
	"fmt"

	"github.com/jboursiquot/iis-go-foundation/extras/custom-methods-on-external-package-types/theirs"
)

func do() {
	ci := theirs.CustomType{}
	fmt.Println(ci)
}

// DoMore is not allowed here because we cannot
// add methods to a type we do not own in this package.
// func (ct theirs.CustomType) DoMore() {}

// However, we can create a local type based on the
// external type and attach methods to it.
// type localCustomType theirs.CustomType
// func (ctw *localCustomType) DoMore() {}
