package operations

import "fmt"

type DeleteOpts struct {
	ID string
}

func Delete(opts *DeleteOpts) error {
	fmt.Println("Delete Function")
	fmt.Printf("Deleting container %s\n", opts.ID)
	return nil
}
