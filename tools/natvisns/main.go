// Copyright (C) 2026 The Qt Company Ltd.
// SPDX-License-Identifier: MIT

// natvisns adds a C++ namespace prefix to all Qt type references in a natvis
// file. It is used by the Qt VS Tools build system and VS extension to produce
// namespace-aware natvis files from the clean, marker-free originals.
//
// Usage:
//     natvisns <input.natvis> <namespace> <output.natvis>
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: natvisns <input.natvis> <namespace> <output.natvis>\n")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	namespace := os.Args[2]
	outputPath := os.Args[3]

	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", inputPath, err)
		os.Exit(1)
	}

	result := AddNamespace(string(data), namespace)

	if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", outputPath, err)
		os.Exit(1)
	}
}
