package lib

import (
	"fmt"
	"os/exec"
)

func RunCommand(command []string) error {
	out, err := exec.Command(command[0], command[1:]...).Output()
	fmt.Printf("%v", string(out))
	if err != nil {
		fmt.Println(err)
	}
	return err
}
