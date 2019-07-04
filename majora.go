package majora

import "github.com/yuin/gopher-lua"

func loadVM(extraInfo map[string]string) (vm *lua.LState) {
	vm = lua.NewState()

	vm.PreloadModule("json", JSONLoader)

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

	for globalVariableName, globalVariableValue := range extraInfo {
		vm.SetGlobal(globalVariableName, lua.LString(globalVariableValue))
	}

	return
}

// ExecScript function runs the given Lua script string
// using a predefined environment and VM.
// WARNING: This function does not handle errors.
// @TODO: Handle custom errors.
func ExecScript(extraVariables map[string]string, script string) (parserJSON []byte) {
	vm := loadVM(extraVariables)
	defer vm.Close()

	if err := vm.DoString(script); err != nil {
    	panic(err)
	}

	parserJSON = []byte(vm.Get(-1).(lua.LString))
	return
}

// ExecFile runs a given Lua script file
// using the same context as the ExecScript function.
// WARNING: This function does not handle errors.
// @TODO: Handle custom errors.
func ExecFile(extraVariables map[string]string, filename string) (parserJSON []byte) {
	vm := loadVM(extraVariables)
	defer vm.Close()

	if err := vm.DoFile(filename); err != nil {
		panic(err)
	}

	parserJSON = []byte(vm.Get(-1).(lua.LString))
	return
}