package main

import (
	"sort"
	"testing"
)

// Helper function to compare slices of strings without considering order
func compareUnorderedSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// Create copies of the slices to avoid modifying the originals
	aCopy := make([]string, len(a))
	bCopy := make([]string, len(b))
	copy(aCopy, a)
	copy(bCopy, b)

	// Sort both slices
	sort.Strings(aCopy)
	sort.Strings(bCopy)

	// Compare the sorted slices
	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}
	return true
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
    <html>
        <body>
            <a href="/path/one">
                <span>Boot.dev</span>
            </a>
            <a href="https://other.com/path/one">
                <span>Boot.dev</span>
            </a>
        </body>
    </html>
    `,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "multiple links with different formats",
			inputURL: "https://example.com",
			inputBody: `
    <html>
        <body>
            <a href="/">Home</a>
            <a href="/about">About</a>
            <a href="https://example.com/contact">Contact</a>
            <a href="//cdn.example.com/image.jpg">Image</a>
            <a href="mailto:info@example.com">Email</a>
        </body>
    </html>
    `,
			expected: []string{
				"https://example.com/",
				"https://example.com/about",
				"https://example.com/contact",
				"https://cdn.example.com/image.jpg",
				"mailto:info@example.com",
			},
		},
		{
			name:     "no links",
			inputURL: "https://example.com",
			inputBody: `
    <html>
        <body>
            <p>This is a paragraph with no links.</p>
        </body>
    </html>
    `,
			expected: []string{},
		},
		{
			name:     "duplicate links",
			inputURL: "https://blog.boot.dev",
			inputBody: `
    <html>
        <body>
            <a href="/path/one">Link 1</a>
            <a href="/path/one">Link 2</a>
            <a href="https://other.com/path/two">Link 3</a>
            <a href="https://other.com/path/two">Link 4</a>
        </body>
    </html>
    `,
			expected: []string{
				"https://blog.boot.dev/path/one",
				"https://blog.boot.dev/path/one",
				"https://other.com/path/two",
				"https://other.com/path/two",
			},
		},
		{
			name:     "nested links",
			inputURL: "https://example.com",
			inputBody: `
    <html>
        <body>
            <div>
                <p>
                    <a href="/nested1">Nested 1</a>
                </p>
            </div>
            <a href="/nested2">
                <span>
                    <a href="/nested3">Nested 3</a>
                </span>
            </a>
        </body>
    </html>
    `,
			expected: []string{
				"https://example.com/nested1",
				"https://example.com/nested2",
				"https://example.com/nested3",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("'%s' FAIL: unexpected error: %v", tc.name, err)
				return
			}
			if !compareUnorderedSlices(actual, tc.expected) {
				t.Errorf("%s FAIL: expected URLs: %v, actual: %v", tc.name, tc.expected, actual)
			}
		})
	}
}
