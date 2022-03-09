package lib

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/withbioma/interview-backend/lib/errs/systemerr"
	"gorm.io/datatypes"
)

// ParseJSON parses json to a variable.
func ParseJSON(j datatypes.JSON, target interface{}) error {
	tValue := reflect.ValueOf(target)
	if tValue.Kind() != reflect.Ptr {
		return systemerr.New("target variable is not a pointer")
	}
	if err := json.Unmarshal(j, target); err != nil {
		return systemerr.Newf("There was an error parsing JSON %v", j)
	}
	return nil
}

// ParseUint parses string to uint.
func ParseUint(str string) (uint, error) {
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	u := uint(u64)
	return u, nil
}
