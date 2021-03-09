package primitive

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// DoPrimitive executes primitive command
func DoPrimitive(input, output string, shapes int, mode int) (string, error) {
	args := fmt.Sprintf("-i %s -o %s -n %d -m %d", input, output, shapes, mode)
	cmd := exec.Command("primitive", strings.Fields(args)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
