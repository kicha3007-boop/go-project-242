package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// makeTree создаёт во временной директории файлы с заданным содержимым и возвращает её путь.
func makeTree(t *testing.T, files map[string]string) string {
	t.Helper()

	root := t.TempDir()

	for name, content := range files {
		path := filepath.Join(root, filepath.FromSlash(name))
		require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
		require.NoError(t, os.WriteFile(path, []byte(content), 0o600))
	}

	return root
}

func TestGetPathSize(t *testing.T) {
	root := makeTree(t, map[string]string{
		"file.txt":            "file\n",
		"dir/a.txt":           "a\n",
		"dir/.hidden.txt":     "hidden\n",
		"dir/nested/deep.txt": "deep\n",
		"dir/.git/config":     "config\n",
		"empty/.keep":         "",
	})
	require.NoError(t, os.Remove(filepath.Join(root, "empty", ".keep")))

	tests := []struct {
		name      string
		path      string
		recursive bool
		all       bool
		expected  string
	}{
		{name: "file", path: "file.txt", expected: "5B"},
		{name: "directory first level only", path: "dir", expected: "2B"},
		{name: "recursive", path: "dir", recursive: true, expected: "7B"},
		{name: "hidden files on first level", path: "dir", all: true, expected: "9B"},
		{name: "recursive with hidden", path: "dir", recursive: true, all: true, expected: "21B"},
		{name: "empty directory", path: "empty", recursive: true, all: true, expected: "0B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPathSize(filepath.Join(root, tt.path), tt.recursive, false, tt.all)
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}

func TestGetPathSizeHuman(t *testing.T) {
	root := makeTree(t, map[string]string{"large.txt": string(make([]byte, 2048))})

	got, err := GetPathSize(filepath.Join(root, "large.txt"), false, true, false)
	require.NoError(t, err)
	require.Equal(t, "2.0KB", got)
}

func TestGetPathSizeMissingPath(t *testing.T) {
	_, err := GetPathSize(filepath.Join(t.TempDir(), "missing"), false, false, false)
	require.ErrorContains(t, err, "missing")
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		size     int64
		human    bool
		expected string
	}{
		{size: 123, expected: "123B"},
		{size: 25165824, expected: "25165824B"},
		{size: 123, human: true, expected: "123B"},
		{size: 1023, human: true, expected: "1023B"},
		{size: 1024, human: true, expected: "1.0KB"},
		{size: 1234567, human: true, expected: "1.2MB"},
		{size: 25165824, human: true, expected: "24.0MB"},
		{size: 1 << 40, human: true, expected: "1.0TB"},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, formatSize(tt.size, tt.human), "size %d, human %v", tt.size, tt.human)
	}
}
