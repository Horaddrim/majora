package majora

import "testing"

func TestRunScript(t *testing.T) {
	script := `
		local json = require("json")
		local fs = require("fs")

		local content = fs.read_file("test/data/mobsf_apk_scan.json")

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

	result := RunScript(script)

	if result == nil {
		t.Errorf("RunScript should not return nil")
	}
}