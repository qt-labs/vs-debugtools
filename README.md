# vs-debugtools

Natvis debugger visualizers for Qt types, plus tooling to add C++ namespace
prefixes when Qt is built inside a custom namespace.

## Natvis files

The `natvis/` directory contains natvis visualizer definitions for Visual Studio
and VS Code:

- `natvis/qt5.natvis` -- Qt 5 types
- `natvis/qt6.natvis` -- Qt 6 types

These files can be used directly by dropping them into your debugger's
visualizer directory.

### Qt namespace support

When Qt is built with a custom namespace (e.g. `MyNamespace`), all type names
in the natvis files need to be prefixed accordingly (e.g. `QString` becomes
`MyNamespace::QString`). The `natvisns` tool handles this automatically.

Type names are discovered from `Type/@Name` and `AlternativeType/@Name`
attributes in the natvis XML. Types that only appear in expressions (not in
type declarations) are annotated with `<!-- qt-type: TypeName -->` comments
so the tool can find them too.

## Building natvisns

Requires [Go](https://go.dev/) 1.22 or later.

```
go build ./tools/natvisns
```

This produces a `natvisns` binary (or `natvisns.exe` on Windows) in the
current directory.

### Usage

```
natvisns <input.natvis> <namespace> <output.natvis>
```

Example:

```
natvisns natvis/qt6.natvis MyNamespace qt6_namespaced.natvis
```

### Running tests

```
go test ./tools/natvisns
```

## Test application

The `test-app/` directory contains a CMake-based C++ application that
exercises various Qt types, useful for manually testing the natvis
visualizers under a debugger. Build it with CMake:

```
cmake -B build -S .
cmake --build build
```

Requires a Qt 5 or Qt 6 installation discoverable by CMake.

## License

MIT -- see `LICENSES/MIT`.
