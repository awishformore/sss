package sss

import (
	"bytes"
	"testing"
)

func TestRandPoly(t *testing.T) {
	b := []byte{1, 2, 3}

	expected := polynomial{10, 1, 2, 3}
	actual, err := randPoly(3, 10, bytes.NewReader(b))
	if err != nil {
		t.Error(err)
	}

	if !equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestRandPolyEOF(t *testing.T) {
	b := []byte{1}

	p, err := randPoly(3, 10, bytes.NewReader(b))
	if p != nil {
		t.Errorf("Expected an error but got %v", p)
	}

	if err == nil {
		t.Error("No error returned")
	}
}

func TestRandPolyEOFFullSize(t *testing.T) {
	b := []byte{1, 2, 0, 0, 0, 0}

	p, err := randPoly(3, 10, bytes.NewReader(b))
	if p != nil {
		t.Errorf("Expected an error but got %v", p)
	}

	if err == nil {
		t.Error("No error returned")
	}
}

func TestRandPolyFullSize(t *testing.T) {
	b := []byte{1, 2, 0, 4}

	expected := polynomial{10, 1, 2, 4}
	actual, err := randPoly(3, 10, bytes.NewReader(b))
	if err != nil {
		t.Error(err)
	}

	if !equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}