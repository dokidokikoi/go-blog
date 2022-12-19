package meta

import (
	"fmt"
	"testing"
)

func TestValidateListOption(t *testing.T) {
	listOption := &ListOption{Order: "name asc", Page: 1, PageSize: 1}
	errs := listOption.Validate()
	for _, err := range errs {
		fmt.Printf("%s\n", err)
	}

}
