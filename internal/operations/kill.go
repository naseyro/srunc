package operations

import "fmt"

type KillOpts struct {
	ID     string
	Signal string
}

func Kill(opts *KillOpts) error {
	fmt.Println("Kill Function")
	fmt.Printf("Sending Signal %s to container %s\n", opts.Signal, opts.ID)
	return nil
}
