### Hexlet tests and linter status:
[![Actions Status](https://github.com/kicha3007-boop/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/kicha3007-boop/go-project-242/actions)
[![CI](https://github.com/kicha3007-boop/go-project-242/actions/workflows/ci.yml/badge.svg)](https://github.com/kicha3007-boop/go-project-242/actions/workflows/ci.yml)

# hexlet-path-size

Консольная утилита и библиотека на Go, которая считает размер файла или директории — упрощённый аналог `du`.

## Установка

Нужен Go 1.25 или новее.

```bash
git clone https://github.com/kicha3007-boop/go-project-242.git
cd go-project-242
make build            # бинарник в bin/hexlet-path-size
```

## Использование

```bash
./bin/hexlet-path-size -h
NAME:
   hexlet-path-size - print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)

USAGE:
   hexlet-path-size [global options] <path>

GLOBAL OPTIONS:
   --recursive, -r  recursive size of directories
   --human, -H      human-readable sizes (auto-select unit)
   --all, -a        include hidden files and directories
   --help, -h       show help
```

| Флаг | Что делает |
|---|---|
| `--recursive`, `-r` | учитывает вложенные директории; без флага считаются только файлы первого уровня |
| `--human`, `-H` | выводит размер в подходящей единице: `B`, `KB`, `MB`, `GB`, `TB`, `PB`, `EB` (основание 1024) |
| `--all`, `-a` | учитывает скрытые файлы и директории (имя начинается с точки) |

Вывод — `<размер>\t<путь>`:

```bash
./bin/hexlet-path-size output.dat
25165824B	output.dat

./bin/hexlet-path-size --human output.dat
24.0MB	output.dat

./bin/hexlet-path-size -H -a -r project/
31.0MB	project/
```

## Как библиотека

```go
import "code"

size, err := code.GetPathSize("project", true, true, false) // recursive, human, all
// size == "24.0MB"
```

## Разработка

```bash
make test       # тесты
make lint       # golangci-lint (конфиг .golangci.yml)
make lint-fix   # автоисправление
```

Проверка Hexlet локально (Docker): образ `hexletprojects/go_path_size_project_ru` содержит тесты автопроверки,
код кладётся в `/project/code`:

```bash
docker run --rm -v "$PWD:/src:ro" --entrypoint sh mirror.gcr.io/hexletprojects/go_path_size_project_ru:latest \
  -c 'rm -rf /project/code && cp -r /src /project/code && cd /project && make setup && make test && make lint'
```

Демо записывается так: `asciinema rec demo.cast`, затем сборка и вызовы из раздела «Использование».
