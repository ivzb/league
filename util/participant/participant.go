package participant

import (
	"strings"

	"league/util/file"
)

const newLine = "\n"

func New(file file.File, path string) ([]string, error) {
	bytes, err := file.Read(path)

	if err != nil {
		return nil, err
	}

	participants := []string{}

	for _, line := range strings.Split(strings.TrimSpace(string(bytes)), newLine) {
		words := strings.Split(line, " ")
		participants = append(participants, words[0])
	}

	return participants, nil
}
