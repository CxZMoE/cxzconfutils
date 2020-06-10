package cxzconfutils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ParsedMap a map which is used to store keys and values
type ParsedMap map[string]string

// ReadLines read all lines of the config file,and return a string array contains them.
func ReadLines(path string) ([]string, error) {
	var strArray []string
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = strings.TrimSuffix(str, "\n")

		strArray = append(strArray, str)
	}

	return strArray, nil
}

// ParseLines returns a map with key of string and value of string.
func ParseLines(lines []string) ParsedMap {
	var pm ParsedMap = make(map[string]string, 1)
	for _, v := range lines {
		// To skip the first line of config
		if !strings.Contains(v, "=") {
			continue
		}
		//fmt.Println(v)
		temp := strings.Split(v, "=")
		pm[temp[0]] = strings.TrimSuffix(temp[1], "\n")
	}
	return pm
}

// GetValue returns a value of the map by passing a key to this function.
func GetValue(m ParsedMap, key string) string {
	return m[key]
}

// CheckKey returns true if the key has a value
func CheckKey(m ParsedMap, key string) bool {
	//fmt.Println("value", m[key])

	// Check the key itself exists
	if _, ok := m[key]; ok {
		// Check whether the value is null.
		if m[key] != "" && m[key] != "\"\"" {
			return true
		}
		return false
	}
	return false
}

// CheckFileKey read from file and returns true if the key has a value
func CheckFileKey(path, key string) bool {
	lines, err := ReadLines(path)
	if err != nil {
		return false
	}
	return CheckKey(ParseLines(lines), key)
}

// GetValueFromFile read from file and returns a value of the map by passing a key to this function.
func GetValueFromFile(path string, key string) string {
	lines, err := ReadLines(path)
	if err != nil {
		return ""
	}
	return ParseLines(lines)[key]
}

// ConfIsNotExist check if the config file is exist
func ConfIsNotExist(path string) bool {
	f, err := os.Open(path)

	if err != nil {
		return true
	}
	_, err = f.Stat()
	return os.IsNotExist(err)
}
