package operations

import "fmt"

type StartOpts struct {
	ID string
}

func Start(opts *StartOpts) error {
	fmt.Println("Start Function")
	fmt.Printf("Starting container %s\n", opts.ID)
	return nil
}
