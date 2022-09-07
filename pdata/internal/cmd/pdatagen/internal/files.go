// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal // import "go.opentelemetry.io/collector/pdata/internal/cmd/pdatagen/internal"

import (
	"os"
	"path/filepath"
	"strings"
)

const header = `// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".`

// AllFiles is a list of all files that needs to be generated.
var AllFiles = []*File{
	commonFile,
	metricsFile,
	resourceFile,
	immutableSliceFile,
	traceFile,
	logFile,
}

// File represents the struct for one generated file.
type File struct {
	Name        string
	PackageName string
	imports     []string
	testImports []string
	// Can be any of sliceOfPtrs, sliceOfValues, messageValueStruct, or messagePtrStruct
	structs []baseStruct
}

// GenerateFile generates the configured data structures for this File.
func (f *File) GenerateFile() error {
	fp, err := os.Create(filepath.Join(".", "pdata", f.PackageName, generateFileName(f.Name)))
	if err != nil {
		return err
	}

	var sb strings.Builder
	generateHeader(&sb, f.PackageName)

	// Add imports
	sb.WriteString("import (" + newLine)
	for _, imp := range f.imports {
		if imp != "" {
			sb.WriteString("\t" + imp + newLine)
		} else {
			sb.WriteString(newLine)
		}
	}
	sb.WriteString(")")

	// Write all structs
	for _, s := range f.structs {
		sb.WriteString(newLine + newLine)
		s.generateStruct(&sb)
	}
	sb.WriteString(newLine)

	_, err = fp.WriteString(sb.String())
	if err != nil {
		return err
	}
	return fp.Close()
}

// GenerateTestFile generates tests for the configured data structures for this File.
func (f *File) GenerateTestFile() error {
	fp, err := os.Create(filepath.Join(".", "pdata", f.PackageName, generateTestFileName(f.Name)))
	if err != nil {
		return err
	}

	var sb strings.Builder
	generateHeader(&sb, f.PackageName)

	// Add imports
	sb.WriteString("import (" + newLine)
	for _, imp := range f.testImports {
		if imp != "" {
			sb.WriteString("\t" + imp + newLine)
		} else {
			sb.WriteString(newLine)
		}
	}
	sb.WriteString(")")

	// Write all tests
	for _, s := range f.structs {
		sb.WriteString(newLine + newLine)
		s.generateTests(&sb)
	}

	_, err = fp.WriteString(sb.String())
	if err != nil {
		return err
	}
	return fp.Close()
}

// GenerateInternalFile generates the internal pdata structures for this File.
func (f *File) GenerateInternalFile() error {
	fp, err := os.Create(filepath.Join(".", "pdata", "internal", generateInternalFileName(f.Name)))
	if err != nil {
		return err
	}

	var sb strings.Builder
	generateHeader(&sb, "internal")

	// Add imports
	sb.WriteString("import (" + newLine)
	for _, imp := range f.imports {
		if imp != "" {
			sb.WriteString("\t" + imp + newLine)
		} else {
			sb.WriteString(newLine)
		}
	}
	sb.WriteString(")")

	// Write all types and funcs
	for _, s := range f.structs {
		s.generateInternal(&sb)
	}
	sb.WriteString(newLine)

	// Write all tests generate value
	for _, s := range f.structs {
		sb.WriteString(newLine + newLine)
		s.generateTestValueHelpers(&sb)
	}
	sb.WriteString(newLine)

	_, err = fp.WriteString(sb.String())
	if err != nil {
		return err
	}
	return fp.Close()
}

func generateFileName(fileName string) string {
	return "generated_" + fileName + ".go"
}

func generateInternalFileName(fileName string) string {
	return "generated_wrapper_" + fileName + ".go"
}

func generateTestFileName(fileName string) string {
	return "generated_" + fileName + "_test.go"
}

func generateHeader(sb *strings.Builder, packageName string) {
	sb.WriteString(header)
	sb.WriteString(newLine + newLine)
	sb.WriteString("package " + packageName)
	sb.WriteString(newLine + newLine)
}