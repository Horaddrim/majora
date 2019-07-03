package majora

import (
	luaJSON "layeh.com/gopher-json"
	"github.com/yuin/gopher-lua"
)

func loadVM(datapath string) (vm *lua.LState) {
	vm = lua.NewState()

	vm.PreloadModule("json", luaJSON.Loader)

	fileModule := `
		local fs = {}

		function fs.read_file(filename) 
			local file = io.open(filename, "r")

			local content = file:read("*all")
			file:close()

			return content
		end

		return fs
	`

	mod, err := vm.LoadString(fileModule)

	if err != nil {
		panic(err)
	}

	loadedIndex := vm.GetField(vm.GetField(vm.Get(lua.EnvironIndex), "package"), "preload")

	vm.SetField(loadedIndex, "fs", mod)
	vm.SetGlobal("datapath", lua.LString(datapath))
	return
}

// RunScript function runs the given Lua script string
// using a predefined environment and VM.
// WARNING: This function does not handle errors.
// @TODO: Handle custom errors.
func RunScript(script, datapath string) (parserJSON []byte) {
	vm := loadVM(datapath)
	defer vm.Close()

	if err := vm.DoString(script); err != nil {
    	panic(err)
	}

	parserJSON = []byte(vm.Get(-1).(lua.LString))
	return
}

// RunFile runs a given Lua script file
// using the same context as the RunScript function.
// WARNING: This function does not handle errors.
// @TODO: Handle custom errors.
func RunFile(filename, datapath string) (parserJSON []byte) {
	vm := loadVM(datapath)
	defer vm.Close()

	if err := vm.DoFile(filename); err != nil {
		panic(err)
	}

	parserJSON = []byte(vm.Get(-1).(lua.LString))
	return
}