// Code generated by "dbusutil-gen em -type Config"; DO NOT EDIT.

package dsync

import (
	"pkg.deepin.io/lib/dbusutil"
)

func (v *Config) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "Get",
			Fn:      v.Get,
			OutArgs: []string{"data"},
		},
		{
			Name:   "Set",
			Fn:     v.Set,
			InArgs: []string{"data"},
		},
	}
}