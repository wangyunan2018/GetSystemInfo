package tools

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFile(fileName string) ([]string, error) {
	// 读取文件
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	buf := bufio.NewReader(f)
	var result []string
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				return result, nil
			}

			return nil, err
		}

		result = append(result, line)
	}

	return result, nil
}
