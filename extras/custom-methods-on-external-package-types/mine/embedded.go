package mine

import (
	"fmt"

	"github.com/jboursiquot/iis-go-foundation/extras/custom-methods-on-external-package-types/theirs"
)

// DoMore is not allowed here because we cannot
// add methods to a type we do not own in this package.
// func (ct theirs.CustomType) DoMore() {

// However, we can create a custom type that embeds
// the external type and attach methods to that.
type myCustomType struct {
	theirs theirs.CustomType
}

func (m *myCustomType) DoMore() {
	m.theirs.Do()
}

func useEmbededType() {
	m := myCustomType{theirs: theirs.CustomType{}}
	fmt.Println(m)
}
