package main

import (
	"encoding/json"
	"os"
	"reflect"

	"github.com/alecthomas/jsonschema"
	"github.com/kubernetes-sigs/prometheus-adapter/pkg/config"
	pmodel "github.com/prometheus/common/model"
)


func schemaOverride(t reflect.Type) *jsonschema.Type {
	durationType := reflect.TypeOf((*pmodel.Duration)(nil)).Elem()
	if t == durationType {
		return &jsonschema.Type{
			Type: "string",
		}
	}
	return nil
}

func main() {

	t := reflect.TypeOf((*config.MetricsDiscoveryConfig)(nil)).Elem()
	r := &jsonschema.Reflector{
		//TypeNamer:           schemaTypeName,
		YAMLEmbeddedStructs: true,
		TypeMapper:          schemaOverride,
		//AdditionalFields:    schemaAddFields,
	}
	schema := r.ReflectFromType(t)
	err := json.NewEncoder(os.Stdout).Encode(schema)
	if err != nil {
		os.Exit(1)
	}

}
