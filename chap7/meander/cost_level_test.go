package meander_test

import (
	"testing"

	"../meander"
	"github.com/cheekybits/is"
)

func TestCostValue(t *testing.T) {
	is := is.New(t)
	is.Equal(int(meander.Cost1), 1)
	is.Equal(int(meander.Cost2), 2)
	is.Equal(int(meander.Cost3), 3)
	is.Equal(int(meander.Cost4), 4)
	is.Equal(int(meander.Cost5), 5)
}

func TestCostString(t *testing.T) {
	is := is.New(t)
	is.Equal(int(meander.Cost1.String()), "$")
	is.Equal(int(meander.Cost2.String()), "$$")
	is.Equal(int(meander.Cost3.String()), "$$$")
	is.Equal(int(meander.Cost4.String()), "$$$$")
	is.Equal(int(meander.Cost5.String()), "$$$$$")
}
