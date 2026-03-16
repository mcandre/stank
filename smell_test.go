package stank_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mcandre/stank"
)

func TestSmellJSONCodecIdempotent(t *testing.T) {
	var smell stank.Smell
	smellJSON, err := json.Marshal(smell)

	if err != nil {
		t.Error(err)
	}

	var smell2 stank.Smell
	err = json.Unmarshal(smellJSON, &smell2)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(smell2, smell) {
		t.Errorf("expected smell %v to equal %v", smell2, smell)
	}
}
