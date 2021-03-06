package values

import (
	"encoding/json"
	"github.com/echocat/caretakerd/errors"
	"strings"
)

// # Description
//
// This tells caretakerd what to do if a process ends.
type RestartType int

const (
	// @id never
	// Never restart the process.
	Never RestartType = 0
	// @id onFailures
	// Only restart the process on failures.
	OnFailures RestartType = 1
	// @id always
	// Always restart the process. This means on success and on failures.
	Always RestartType = 2
)
var AllRestartTypes = []RestartType{
	Never,
	OnFailures,
	Always,
}

func (i RestartType) String() string {
	result, err := i.CheckedString()
	if err != nil {
		panic(err)
	}
	return result
}

func (i RestartType) CheckedString() (string, error) {
	switch i {
	case Never:
		return "never", nil
	case OnFailures:
		return "onFailures", nil
	case Always:
		return "always", nil
	}
	return "", errors.New("Illegal restart type: %v", i)
}

func (i *RestartType) Set(value string) error {
	lowerValue := strings.ToLower(value)
	for _, candidate := range AllRestartTypes {
		if strings.ToLower(candidate.String()) == lowerValue {
			(*i) = candidate
			return nil
		}
	}
	return errors.New("Illegal restart type: %s", value)
}

func (i RestartType) MarshalYAML() (interface{}, error) {
	return i.CheckedString()
}

func (i *RestartType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}
	return i.Set(value)
}

func (i RestartType) MarshalJSON() ([]byte, error) {
	s, err := i.CheckedString()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(s)
}

func (i *RestartType) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	return i.Set(value)
}

func (i RestartType) OnSuccess() bool {
	return i == Always
}

func (i RestartType) OnFailures() bool {
	return i == OnFailures || i == Always
}

func (i RestartType) Validate() error {
	_, err := i.CheckedString()
	return err
}
