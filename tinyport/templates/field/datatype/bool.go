package datatype

import (
	"fmt"

	"github.com/notional-labs/tinyport/tinyport/pkg/multiformatname"
)

var (
	// DataBool bool data type definition
	DataBool = DataType{
		DataType:          func(string) string { return "bool" },
		DefaultTestValue:  "false",
		ValueLoop:         "false",
		ValueIndex:        "false",
		ValueInvalidIndex: "false",
		ProtoType: func(_, name string, index int) string {
			return fmt.Sprintf("bool %s = %d", name, index)
		},
		GenesisArgs: func(name multiformatname.Name, value int) string {
			return fmt.Sprintf("%s: %t,\n", name.UpperCamel, value%2 == 0)
		},
		CLIArgs: func(name multiformatname.Name, _, prefix string, argIndex int) string {
			return fmt.Sprintf(`%s%s, err := cast.ToBoolE(args[%d])
            		if err != nil {
                		return err
            		}`,
				prefix, name.UpperCamel, argIndex)
		},
		ToBytes: func(name string) string {
			return fmt.Sprintf(`%[1]vBytes := []byte{0}
					if %[1]v {
						%[1]vBytes = []byte{1}
					}`, name)
		},
		ToString: func(name string) string {
			return fmt.Sprintf("strconv.FormatBool(%s)", name)
		},
		GoCLIImports: []GoImport{{Name: "github.com/spf13/cast"}},
	}
)
