// Copyright (C) 2026 The Qt Company Ltd.
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestAddNamespace_Qt5 verifies that AddNamespace on a clean (marker-free)
// qt5.natvis produces the same result as a simple marker replacement, modulo
// intentional improvements (bare refs the markers missed).
//
// The test fixtures in testdata/ are snapshots of the natvis files before
// markers were removed. If the production natvis files gain new types, the
// fixtures should be regenerated from the marker-era versions to stay useful
// as ground truth.
func TestAddNamespace_Qt5(t *testing.T) {
	testAddNamespace(t, "qt5_with_markers.natvis",
		// Types that only appear in expressions (no Type/@Name declaration).
		// In production, the natvis files carry <!-- qt-type: ... --> annotations.
		// Here we supply them as extra types for the test.
		[]string{"QMetaType", "Qt", "QLinkedListNode", "QHashDummyValue"})
}

// TestAddNamespace_Qt6 verifies the same for qt6.natvis.
func TestAddNamespace_Qt6(t *testing.T) {
	testAddNamespace(t, "qt6_with_markers.natvis",
		[]string{"QMetaType", "Qt", "QtPrivate", "QMetaTypeInterface",
			"QAbstractSocket", "QAtomicInt", "QHashDummyValue"})
}

func testAddNamespace(t *testing.T, fixture string, extraTypes []string) {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("testdata", fixture))
	if err != nil {
		t.Fatalf("reading fixture %s: %v", fixture, err)
	}
	withMarkers := string(data)

	const ns = "TestNamespace"

	// Expected: what the old marker-replacement approach produces
	expected := strings.ReplaceAll(withMarkers, "##NAMESPACE##", ns)

	// Clean input: remove markers to simulate the marker-free natvis files
	clean := strings.ReplaceAll(withMarkers, "##NAMESPACE##::", "")
	clean = strings.ReplaceAll(clean, "##NAMESPACE##", "")

	// Actual: what addNamespace produces from clean input, with the extra
	// types that production files would provide via annotations.
	actual := addNamespace(clean, ns, extraTypes)

	if actual == expected {
		return // Perfect match
	}

	// Collect and classify differences
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(actual, "\n")

	maxLines := len(expectedLines)
	if len(actualLines) > maxLines {
		maxLines = len(actualLines)
	}

	var gaps []string         // actual is missing namespace that expected has
	var improvements []string // actual adds namespace that expected lacks

	for i := 0; i < maxLines; i++ {
		var eLine, aLine string
		if i < len(expectedLines) {
			eLine = expectedLines[i]
		}
		if i < len(actualLines) {
			aLine = actualLines[i]
		}
		if eLine == aLine {
			continue
		}

		diff := formatDiff(i+1, eLine, aLine)

		// Classify: count namespace references in each line
		eCount := strings.Count(eLine, ns+"::")
		aCount := strings.Count(aLine, ns+"::")
		if eCount > aCount {
			gaps = append(gaps, diff)
		} else {
			improvements = append(improvements, diff)
		}
	}

	if len(improvements) > 0 {
		t.Logf("AddNamespace produces %d intentional improvements "+
			"(bare refs now correctly namespaced):\n\n%s",
			len(improvements), strings.Join(improvements, "\n\n"))
	}

	if len(gaps) > 0 {
		t.Errorf("AddNamespace has %d gaps (missing namespace refs). "+
			"Add <!-- qt-type: ... --> annotations for these types:\n\n%s",
			len(gaps), strings.Join(gaps, "\n\n"))
	}
}

func formatDiff(line int, expected, actual string) string {
	return fmt.Sprintf("--- line %d ---\nexpected: %s\nactual:   %s",
		line, strings.TrimSpace(expected), strings.TrimSpace(actual))
}
