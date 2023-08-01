package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	scriptTestString  = "{ \"Wires\", iconicData._Wires },"
	dataTestStringRaw = "   00000000 00000000 000    00   000000000000000000000000  0000  000000000000000000000000 0000000000000000000000000000000 00000000000000000000000000000000 0000 000000000000000000000000000 00  0000000000000000000000000000    0000000000000000000000000000000 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 000000000000000000000000000000 0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 000000000000000000000000000000 0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 000000000000000000000000000000  000000000000000000000000000000    00000000 00000000 00000000   "
	dataTestString    = "public static string[] _Wires = {\"" + dataTestStringRaw + "\", \"Somewhere on the module\"};"
)

func TestScriptRegexp(t *testing.T) {
	assert.Equal(t, []string{scriptTestString, "Wires", "Wires"}, reScript.FindStringSubmatch(scriptTestString))
	assert.Nil(t, reScript.FindStringSubmatch(" invalid string "))
}

func TestDataRegexp(t *testing.T) {
	assert.Equal(t, []string{dataTestString, "Wires", dataTestStringRaw, "\"Somewhere on the module\""}, reData.FindStringSubmatch(dataTestString))
	assert.Nil(t, reData.FindStringSubmatch(" invalid string "))
}