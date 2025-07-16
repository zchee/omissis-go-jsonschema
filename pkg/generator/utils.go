package generator

import (
	"slices"

	"github.com/atombender/go-jsonschema/pkg/codegen"
	"github.com/atombender/go-jsonschema/pkg/schemas"
)

const additionalProperties = "AdditionalProperties"

func sortedKeys[T any](props map[string]T, sortField bool) []string {
	names := make([]string, 0, len(props))
	for name := range props {
		names = append(names, name)
	}

	if sortField {
		slices.Sort(names)
	}

	return names
}

func sortDefinitionsByName(defs schemas.Definitions, sortField bool) []string {
	names := make([]string, 0, len(defs))

	for name := range defs {
		names = append(names, name)
	}

	if sortField {
		slices.Sort(names)
	}

	return names
}

func isNamedType(t codegen.Type) bool {
	switch x := t.(type) {
	case *codegen.NamedType:
		return true

	case *codegen.PointerType:
		if _, ok := x.Type.(*codegen.NamedType); ok {
			return true
		}
	}

	return false
}

func isMapType(t codegen.Type) bool {
	_, isMapType := t.(*codegen.MapType)

	return isMapType
}
