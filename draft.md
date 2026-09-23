This is a draft program in Go to write all ideas.

type Namespace struct {
	ID int
	NextPID int
	Processes []Process
}

type Process struct {
	PID int
	Name string
	ExecutablePath string
	State string
	Parent Process
	Child Process
}

We need to implement the following commands:
1. NEWNS -> Create a new Namespace with its unique ID and its own version of NextPID and its own PID table
2. FORK -> Create a new process into a specific namespace and assign the Next available PID to this process dynamically.
3. WAIT -> Wait is called by the Parent process when its child process is going to die.
4. EXIT -> Delete a process from a specific namespace. We don't need to delete the process itself from the Namespace's processes, all we need is to change the State to be "exited"
5. PS -> Read all available processes in a specific namespace. (Print and loop over all namespace.Processes)


Should we use os.Process? or create our own process struct?

We need to parse input lines based on each function. or actually not, we have a formula
that always the first parameter is the namespace id and the second one is the process name for create process
