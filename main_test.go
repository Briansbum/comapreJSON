package main

import (
	"errors"
	"testing"

	"github.com/nsf/jsondiff"
	"github.com/stretchr/testify/assert"
)

func testCompareJSON(t *testing.T) {
	testFixtures := []struct {
		json1  []byte
		json2  []byte
		result jsondiff.Difference
		err    error
	}{
		{[]byte{}, []byte{}, jsondiff.Difference(5), errors.New("json documents are different, BothArgsAreInvalidJson")},
		{[]byte(`{"foo": "bar"}`), []byte{}, jsondiff.Difference(4), errors.New("json documents are different, SecondArgIsInvalidJson")},
		{[]byte{}, []byte(`{"foo": "bar"}`), jsondiff.Difference(3), errors.New("json documents are different, FirstArgIsInvalidJson")},
		{[]byte(`{"foo": "bar"}`), []byte(`{"foo": "bar"}`), jsondiff.Difference(0), nil},
		{
			[]byte(`{"foo": "bar"}`),
			[]byte(`{"foo": "bar", "far": {"faz": "baz"}}`),
			jsondiff.Difference(2),
			errors.New("json documents are different, NoMatch"),
		},
		{
			[]byte(`{"foo": "bar", "far": {"faz": "baz"}, "coo": "cat"}`),
			[]byte(`{"foo": "bar", "far": {"faz": "baz"}}`),
			jsondiff.Difference(1),
			nil,
		},
	}

	for _, fixture := range testFixtures {
		result, err := CompareJSON(fixture.json1, fixture.json2)

		if assert.NoError(t, err) {
			assert.Equal(t, result, fixture.result)
		}
	}
}
