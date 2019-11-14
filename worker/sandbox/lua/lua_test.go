package lua

import (
	"testing"

	"github.com/stretchr/testify/require"
	lua "github.com/yuin/gopher-lua"
)

var lfunc *lua.LFunction

func flatMap(L *lua.LState) int {
	lfunc = L.CheckFunction(-1)
	return 0
}

func lMap(L *lua.LState) int {
	return 0
}

func TestPackFunc(t *testing.T) {
	L := lua.NewState()
	// defer L.Close()
	L.OpenLibs()
	L.SetGlobal("flatMap", L.NewFunction(flatMap))
	L.SetGlobal("map", L.NewFunction(lMap))

	err := L.DoFile("./data/dump.lua")

	require.NoError(t, err)

	L.Close()

	L = lua.NewState()
	L.OpenLibs()

	L.Push(lfunc)

	err = L.PCall(0, 0, nil)

	require.NoError(t, err)

	println(lfunc.String())

	println(L.GetGlobal("say_hi2").String())
}
