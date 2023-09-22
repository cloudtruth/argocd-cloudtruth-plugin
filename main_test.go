package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMergeArgoEnv(t *testing.T) {
	pluginConfig := make(map[string]string)
	defaultValue := "default_value"

	// Test case 0: No environment variables set
	result := mergeArgoEnv("MY_PLUGIN_VAR", defaultValue, pluginConfig)

	if result != defaultValue {
		t.Errorf("Expected %s, but got %s", defaultValue, result)
	}

	// Test case 1: System environment variable set
	os.Setenv("MY_PLUGIN_VAR", "system_env_value")

	result = mergeArgoEnv("MY_PLUGIN_VAR", defaultValue, pluginConfig)
	expected := "system_env_value"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test case 2: Plugin configuration parameter set
	// os.Unsetenv("MY_PLUGIN_VAR")
	pluginConfig["my-plugin-var"] = "config_param_value"

	result = mergeArgoEnv("MY_PLUGIN_VAR", defaultValue, pluginConfig)
	expected = "config_param_value"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test case 3: ARGOCD_ENV environment variable set
	// delete(pluginConfig, "my-plugin-var")
	os.Setenv("ARGOCD_ENV_MY_PLUGIN_VAR", "argocd_env_value")

	result = mergeArgoEnv("MY_PLUGIN_VAR", defaultValue, pluginConfig)
	expected = "argocd_env_value"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test case 4: PARAM environment variable set
	// os.Unsetenv("ARGOCD_ENV_MY_PLUGIN_VAR")
	os.Setenv("PARAM_MY_PLUGIN_VAR", "param_value")

	result = mergeArgoEnv("MY_PLUGIN_VAR", defaultValue, pluginConfig)
	expected = "param_value"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

}

func TestReadPluginConfig(t *testing.T) {
	// Create a temporary plugin config file
	file, err := os.CreateTemp("", "plugin-config.yaml")
	if err != nil {
		t.Fatalf("Failed to create temporary plugin config file: %s", err)
	}
	defer os.Remove(file.Name())

	// Write test data to the plugin config file
	configData := `
my-plugin-var: value1
another-plugin-var: value2
`
	err = os.WriteFile(file.Name(), []byte(configData), 0644)
	if err != nil {
		t.Fatalf("Failed to write data to the plugin config file: %s", err)
	}

	// Test reading the plugin config file
	result := readPluginConfig(file.Name())

	// Verify the expected data is returned
	expected := map[string]string{
		"my-plugin-var":      "value1",
		"another-plugin-var": "value2",
	}
	if len(result) != len(expected) {
		t.Errorf("Expected %d entries in the plugin config map, but got %d", len(expected), len(result))
	}

	for key, value := range expected {
		if result[key] != value {
			t.Errorf("Expected value %s for key %s, but got %s", value, key, result[key])
		}
	}

	// Test reading a non-existent plugin config file
	result = readPluginConfig("non-existent-file.yaml")

	// Verify an empty map is returned
	if len(result) != 0 {
		t.Errorf("Expected an empty map, but got %v", result)
	}

	// Write bad data to the plugin config file
	err = os.WriteFile(file.Name(), []byte("&"), 0644)
	if err != nil {
		t.Fatalf("Failed to write data to the plugin config file: %s", err)
	}

	// Test reading a bad plugin config file
	result = readPluginConfig(file.Name())

	// Verify an empty map is returned
	if len(result) != 0 {
		t.Errorf("Expected an empty map, but got %v", result)
	}

}

// Define a helper function to create a temporary test file with the given contents
func createTestFile(contents string) (string, error) {
	file, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(contents)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func TestFileReplace(t *testing.T) {
	// Test case 1: file with no pattern match
	fileContents := "This is a test file."
	filePath, err := createTestFile(fileContents)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	parameters := map[string]Parameter{"name": {value: "John"}}
	templates := map[string]Template{"name": {value: func() string { return "Mary" }}}
	ctproject := &CTProject{nil, "env", "envuuid", "proj", "projuuid", "", parameters, templates}

	replacedContents, err := fileReplace(filePath, "{%s}", ctproject)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if replacedContents != fileContents {
		t.Errorf("Expected replaced contents to be '%s', but got '%s'", fileContents, replacedContents)
	}

	// Test case 2: file with a single pattern match
	fileContents = "This is a test file. My name is {name}."
	filePath, err = createTestFile(fileContents)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	replacedContents, err = fileReplace(filePath, "{%s}", ctproject)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	expectedContents := strings.Replace(fileContents, "{name}", "John", -1)
	if replacedContents != expectedContents {
		t.Errorf("Expected replaced contents to be '%s', but got '%s'", expectedContents, replacedContents)
	}

	// Test case 3: file with multiple pattern matches
	fileContents = "This is a test file. My name is {name}. Nice to meet you, {name}."
	filePath, err = createTestFile(fileContents)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	replacedContents, err = fileReplace(filePath, "{%s}", ctproject)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	expectedContents = strings.Replace(fileContents, "{name}", "John", -1)
	if replacedContents != expectedContents {
		t.Errorf("Expected replaced contents to be '%s', but got '%s'", expectedContents, replacedContents)
	}

	// Test case 4: non-existing file
	filePath = "nonexistent.txt"
	replacedContents, err = fileReplace(filePath, "{%s}", ctproject)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if replacedContents != "" {
		t.Errorf("Expected replaced contents to be empty, but got '%s'", replacedContents)
	}
	expectedError := fmt.Sprintf("open %s: no such file or directory", filePath)
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}

	// Test case 5: empty pattern
	fileContents = "This is a test file."
	filePath, err = createTestFile(fileContents)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	replacedContents, err = fileReplace(filePath, "", ctproject)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if replacedContents != fileContents {
		t.Errorf("Expected replaced contents to be '%s', but got '%s'", fileContents, replacedContents)
	}

	// Test case 6: file with a single template pattern match
	fileContents = "This is a test file. My name is {templates.name}."
	filePath, err = createTestFile(fileContents)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	replacedContents, err = fileReplace(filePath, "{%s}", ctproject)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	expectedContents = strings.Replace(fileContents, "{templates.name}", "Mary", -1)
	if replacedContents != expectedContents {
		t.Errorf("Expected replaced contents to be '%s', but got '%s'", expectedContents, replacedContents)
	}

}
