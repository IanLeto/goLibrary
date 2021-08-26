package viperDemo

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)

func RunEcho() {
	cmd := exec.Command("echo", "hello world")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	logrus.Info(string(out))
}

