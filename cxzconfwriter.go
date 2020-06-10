package cxzconfutils

import (
	"bufio"
	"io/ioutil"
	"os"
)

//WriteConfFile write config map to config file
func WriteConfFile(path string, confMap map[string]string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	// Backup config file
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	ioutil.WriteFile(path+".last", data, 0755)

	// Rewrite new file
	writer := bufio.NewWriter(f)
	for k, v := range confMap {
		writer.WriteString(k + "=" + v + "\n")
	}
	writer.Flush()
	return nil
}
