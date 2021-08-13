package mageextras

import (
	"fmt"
	"os"
	"os/exec"
)

// GoVetShadow runs go vet against all Go packages in a project,
// with variable shadow checking enabled.
//
// Depends on golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
func GoVetShadow(args ...string) error {
	shadowPath, err := exec.LookPath("shadow")

	if err != nil {
		return err
	}

	return GoVet(fmt.Sprintf("-vettool=%s", shadowPath))
}

// GoVet runs go vet against all Go packages in a project.
func GoVet(args ...string) error {
	cmdName := "go"

	cmdParameters := []string{cmdName}
	cmdParameters = append(cmdParameters, "vet")
	cmdParameters = append(cmdParameters, args...)
	cmdParameters = append(cmdParameters, AllPackagesPath)

	cmd := exec.Command(cmdName)
	cmd.Args = cmdParameters
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
