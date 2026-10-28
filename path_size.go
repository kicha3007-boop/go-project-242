// Package code считает размер файла или директории — упрощённый аналог утилиты du.
package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize возвращает размер файла или директории path строкой: "123B", а с human — "1.2MB".
// recursive — учитывать вложенные директории, all — учитывать скрытые файлы и директории.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := pathSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}

// pathSize считает размер одного элемента: у файла — его размер, у директории — сумму содержимого.
func pathSize(path string, recursive, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("stat %s: %w", path, err)
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	return dirSize(path, recursive, all)
}

// dirSize суммирует файлы директории; вложенные директории учитываются только при recursive.
func dirSize(dir string, recursive, all bool) (int64, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, fmt.Errorf("read dir %s: %w", dir, err)
	}

	var total int64

	for _, entry := range entries {
		if !all && isHidden(entry.Name()) {
			continue
		}

		if entry.IsDir() && !recursive {
			continue
		}

		size, err := pathSize(filepath.Join(dir, entry.Name()), recursive, all)
		if err != nil {
			return 0, err
		}

		total += size
	}

	return total, nil
}

// isHidden — скрытым считается элемент, чьё имя начинается с точки.
func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}
