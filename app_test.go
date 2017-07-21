package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestTopPasswords(t *testing.T) {
  assert.Equal(t, 100, len(getTopPasswords()), "there should be one hundred")
}
