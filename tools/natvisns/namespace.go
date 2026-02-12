// Copyright (C) 2026 The Qt Company Ltd.
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// AddNamespace takes clean natvis content (no ##NAMESPACE## markers) and a
// namespace string, and returns the content with all Qt type names prefixed
// by the namespace. It discovers type names from Type/@Name and
// AlternativeType/@Name attributes, plus <!-- qt-type: TypeName --> annotations.
func AddNamespace(content, namespace string) string {
	return addNamespace(content, namespace, nil)
}

// addNamespace is the core implementation. extraTypes allows callers (tests)
// to provide additional type names beyond what's discovered from the content.
func addNamespace(content, namespace string, extraTypes []string) string {
	if namespace == "" {
		return content
	}

	names := collectTypeNames(content)

	// Merge extra types
	seen := make(map[string]bool, len(names))
	for n := range toSet(names) {
		seen[n] = true
	}
	for _, n := range extraTypes {
		if !seen[n] {
			names = append(names, n)
			seen[n] = true
		}
	}

	if len(names) == 0 {
		return content
	}

	// Sort longest-first to prevent partial replacements
	// (e.g. "QLinkedListNode" before "QLinkedList")
	sort.Slice(names, func(i, j int) bool {
		return len(names[i]) > len(names[j])
	})

	// Extract XML comments to protect them from false-positive replacements.
	// (e.g. "The Qt Company" in copyright headers must not be modified.)
	// The annotations were already read by collectTypeNames above.
	var comments []string
	commentRe := regexp.MustCompile(`(?s)<!--.*?-->`)
	content = commentRe.ReplaceAllStringFunc(content, func(match string) string {
		idx := len(comments)
		comments = append(comments, match)
		return fmt.Sprintf("\x00COMMENT_%d\x00", idx)
	})

	for _, name := range names {
		// Match the type name when NOT preceded by :: or a word character,
		// and followed by a word boundary. Go's RE2 doesn't support
		// lookbehinds, so we capture the preceding character (or start)
		// and restore it in the replacement.
		re := regexp.MustCompile(`(?m)(^|[^:\w])` + regexp.QuoteMeta(name) + `\b`)
		content = re.ReplaceAllStringFunc(content, func(match string) string {
			prefix := match[:len(match)-len(name)]
			return prefix + namespace + "::" + name
		})
	}

	// Restore XML comments
	for i, c := range comments {
		content = strings.Replace(content, fmt.Sprintf("\x00COMMENT_%d\x00", i), c, 1)
	}

	return content
}

// collectTypeNames extracts Qt type base names from:
// 1. Type Name="..." and AlternativeType Name="..." attributes
// 2. <!-- qt-type: TypeName --> annotations
func collectTypeNames(content string) []string {
	seen := make(map[string]bool)

	// Match Type Name="SomeName" or AlternativeType Name="SomeName"
	// The name may include template parameters — in natvis XML these are
	// encoded as &lt;...&gt; rather than raw <...>.
	attrRe := regexp.MustCompile(`(?:Type|AlternativeType)\s+Name\s*=\s*"([^"]+)"`)
	for _, match := range attrRe.FindAllStringSubmatch(content, -1) {
		name := match[1]
		// Strip template arguments: both raw < and XML-encoded &lt;
		if idx := strings.Index(name, "<"); idx >= 0 {
			name = name[:idx]
		} else if idx := strings.Index(name, "&lt;"); idx >= 0 {
			name = name[:idx]
		}
		// Extract the base name (first segment before ::)
		// e.g. QHashPrivate::Node -> QHashPrivate
		if idx := strings.Index(name, "::"); idx >= 0 {
			name = name[:idx]
		}
		name = strings.TrimSpace(name)
		if name != "" && !seen[name] {
			seen[name] = true
		}
	}

	// Match <!-- qt-type: TypeName --> annotations
	annotRe := regexp.MustCompile(`<!--\s*qt-type:\s*(\S+)\s*-->`)
	for _, match := range annotRe.FindAllStringSubmatch(content, -1) {
		name := strings.TrimSpace(match[1])
		if name != "" && !seen[name] {
			seen[name] = true
		}
	}

	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	return names
}

func toSet(names []string) map[string]bool {
	s := make(map[string]bool, len(names))
	for _, n := range names {
		s[n] = true
	}
	return s
}
