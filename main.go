package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var GlobalNextNamespaceIDCounter int = 1

func main() {
	fmt.Println("Starting minictr........")
	fmt.Println(`"Available commands for minictr are: 
	NEWNS to create a new PID namespace
	FORK <NS ID> <PROCESS NAME> to create a new Process in a certain namespace
	PS <NS> to list all processes found in a certain namespace
	PS to list all namespaces each with its own processes
"`)
	k := &KernelManager{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter text (Press Ctrl+D or Ctrl+C to exit):")

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.EqualFold(line, "NEWNS"):
			k.createNamespace()
		case strings.Contains(line, "FORK"):
			k.createProcess(line)
		case strings.EqualFold(line, "PS"):
			k.DisplayNamespaces()
		case strings.Contains(line, "PS"):
			k.DisplayProcesses(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
}

type KernelManager struct {
	Namespaces []*Namespace
}

type Namespace struct {
	ID        int
	Processes []*Process
	NextPID   int
}

type Process struct {
	PID    int
	Name   string
	State  string
	Parent *Process
	Child  *Process
}

// We need to declare the functions that will be called based on the user input.
// NEWNS -> Should call: createNamespace() ===== DONE
// FORK NAMESPACEID PROCESSNAME -> should call: createProcess(NAMESPACEID, PROCESSNAME)
// PS NAMESPACE -> should call: DisplayProcesses(NAMESPACEID)

func (k *KernelManager) createNamespace() *Namespace {
	ns := &Namespace{
		ID:      GlobalNextNamespaceIDCounter,
		NextPID: 1,
	}
	GlobalNextNamespaceIDCounter++
	k.Namespaces = append(k.Namespaces, ns)
	fmt.Printf("PID Namespace %d created\n", ns.ID)
	return ns
}

func (k *KernelManager) createProcess(line string) *Process {
	args := strings.Split(line, " ")
	nsID := args[1]
	namespace := k.getNamespace(nsID)
	if namespace == nil {
		namespace = k.createNamespace()
	}
	if len(args) != 3 {
		log.Fatalf("insufficient params to create a process")
		return nil
	}
	process := namespace.createProcess(args[2])
	return process
}

func (k *KernelManager) getNamespace(nsID string) *Namespace {
	id, err := strconv.Atoi(nsID)
	if err != nil {
		log.Fatalf("error retrieving namespace ID %v", err)
	}
	for _, namespace := range k.Namespaces {
		if namespace.ID == id {
			return namespace
		}
	}
	return nil
}

func (n *Namespace) createProcess(pName string) *Process {
	p := &Process{
		Name:  pName,
		PID:   n.NextPID,
		State: "Running",
	}
	n.NextPID++
	fmt.Printf("PID: %d | Process %s at PID Namespace %d created\n", p.PID, p.Name, n.ID)
	n.Processes = append(n.Processes, p)
	return p
}

func (k *KernelManager) DisplayProcesses(line string) {
	args := strings.Split(line, " ")
	if len(args) != 2 {
		log.Fatalf("insufficient params to display all namespace processes. namespace id is required.")
	}
	namespace := k.getNamespace(args[1])
	fmt.Printf("Namespace %d\n", namespace.ID)
	for _, p := range namespace.Processes {
		fmt.Printf("PID: %d       Process: %s       State: %s\n", p.PID, p.Name, p.State)
	}
}

func (k *KernelManager) DisplayNamespaces() {
	for _, n := range k.Namespaces {
		line := fmt.Sprintf("PS %d", n.ID)
		k.DisplayProcesses(line)
	}
}
