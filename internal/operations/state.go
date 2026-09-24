package operations

import "fmt"

type StateOpts struct {
	ID string
}

func State(opts *StateOpts) (string, error) {
	fmt.Println("State Function")
	fmt.Printf("State of container %s is running\n", opts.ID)
	return "", nil
}
