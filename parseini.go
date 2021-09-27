package parseini

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func IniFile(filepath string) (map[string]interface{}, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data := make(map[string]interface{})
	mapKey := ""

	for {
		text, err := reader.ReadSlice('\n')
		if err == io.EOF {
			break
		}

		tracer := strings.Trim(string(text), " \t\n\r\x0B")
		traLen := len(tracer)
		if traLen == 0 || tracer[0] == ';' || tracer[0] == '#' {
			continue
		}

		traLen = strings.Index(tracer, ";")
		if traLen != -1 {
			tracer = tracer[0:traLen]
		}

		traLen = strings.Index(tracer, "#")
		if traLen != -1 {
			tracer = tracer[0:traLen]
		}

		traLen = len(tracer)

		if string(tracer[0]) == "[" && string(tracer[traLen-1]) == "]" {
			mapKey = tracer[1 : traLen-1]
			continue
		}
		traLen = strings.Index(tracer, "=")
		key := strings.Trim(tracer[0:traLen], " \t")
		val := strings.Trim(tracer[traLen+1:], " '\"`")
		if mapKey == "" {
			data[key] = val
		} else if tmp, ok := data[mapKey]; !ok {
			data[mapKey] = map[string]string{
				key: val,
			}
		} else {
			tmp.(map[string]string)[key] = val
			data[mapKey] = tmp
		}
	}
	return data, nil
}
