package cmd

import (
	"fmt"
	"log"
)

func ExecBash() {
	cmd1 := "sudo apt install lm-sensors -y || yum install -y lm_sensors"
	_, _, err := ShellOut(cmd1)
	if err != nil {
		log.Printf("命令执行失败，请检查输入是否有误。%v", err)
	}

	// PrintOut()
	// fmt.Println(out)
	// fmt.Println(errout)

	cmd2 := "sensors | grep 'Core'"
	out2, errout2, err := ShellOut(cmd2)
	if err != nil {
		log.Printf("命令执行失败，请检查输入是否有误。%v", err)
	}
	// PrintOut()
	fmt.Println(out2)
	fmt.Println(errout2)
}
