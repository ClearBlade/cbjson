package cbjson

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

//  Files used for testing
const (
	testFile    = "good.conf"
	error25File = "errorLine25.conf"
	anotherFile = "example.conf"
)

func Test_PreprocessFile(t *testing.T) {

	theMap, _, err := GetJSONFile(testFile)
	if err != nil {
		t.Errorf("unmarshal failed: %s", err.Error())
	} else {
		validateContents(t, theMap)
	}
}

func Test_PreprocessAnotherFile(t *testing.T) {

	theMap, _, err := GetJSONFile(anotherFile)
	if err != nil {
		t.Errorf("unmarshal failed: %s", err.Error())
	} else {
		validateContents(t, theMap)
	}
}

func Test_ErrorLine25(t *testing.T) {
	_, _, err := GetJSONFile(error25File)
	if err == nil {
		t.Errorf("Got success but expected syntax error")
	} else if !strings.Contains(err.Error(), "line 25") {
		t.Errorf("Got error on wrong line")
	}
}

func preprocessOneFile(filename string) ([]byte, error) {
	byts, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open test file: %s", err.Error)
	}
	p := NewPreprocessor(byts)
	buf, err := p.preprocess()
	if err != nil {
		return nil, fmt.Errorf("Couldn't preprocess test file: %s", err.Error())
	}
	return buf.Bytes(), nil
}

func validateContents(t *testing.T, stuff map[string]interface{}) (bool, error) {
	scoreFunc, ok := stuff["ScoringFunction"]
	if !ok {
		t.Errorf("missing key ScoringFunction")
	}
	if _, ok = scoreFunc.(string); !ok {
		t.Errorf("Scoring function must be a string")
	}

	nodes, ok := stuff["Nodes"]
	if !ok {
		t.Errorf("missing key Nodes")
	}
	if _, ok = nodes.(map[string]interface{}); !ok {
		t.Errorf("Scoring function must be a map[string]interface{}")
	}

	return true, nil
}

func mustHaveAndBeOfType(fieldName string, fieldType interface{}) error {
	return nil
}
