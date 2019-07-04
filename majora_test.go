package majora

import "testing"

func TestExecScript(t *testing.T) {
	script := `
		local json = require("json")
		local fs = require("fs")

		local content = fs.read_file(datapath)

		local json_parsed = json.decode(content)

		local new_permissions = {}
		for permission,permissionData in pairs(json_parsed["permissions"]) do
			local new_permission = {}
			new_permission.title = permission

			for key, value in pairs(permissionData) do
				new_permission[key] = value
			end

			table.insert(new_permissions, new_permission)
		end

		json_parsed["permissions"] = new_permissions

		return json.encode(json_parsed)
	`

	globalVariables := make(map[string]string)

	globalVariables["datapath"] = "test/data/mobsf_apk_scan.json"

	result := ExecScript(globalVariables, script)

	if result == nil {
		t.Errorf("ExecScript should not return nil.")
	}
}

func TestExecFile(t *testing.T) {
	globalVariables := make(map[string]string)

	globalVariables["datapath"] = "test/data/mobsf_apk_scan.json"

	result := ExecFile(globalVariables, "test/mobsf.lua")

	if result == nil {
		t.Errorf("ExecFile should not return nil.")
	}
}