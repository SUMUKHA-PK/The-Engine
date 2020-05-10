package filefinder

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func Find(parentDir, fileName string) (string, error) {
	// command := "find " + parentDir + " -iname *" + fileName + "*"
	cmd := exec.Command("find", parentDir, "-iname", "*"+fileName+"*")
	// cmd := exec.Command("tr", "a-z", "A-Z")
	// cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Returns: %q\n", out.String())

	return "", nil
}
