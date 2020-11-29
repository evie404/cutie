package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable_Filename(t *testing.T) {
	tests := []struct {
		name string
		t    *Table
		want string
	}{
		{
			"uses FilenameOverride when present",
			&Table{
				TableName:        "table_ones",
				FilenameOverride: "table_ones_override",
			},
			"table_ones_override",
		},
		{
			"uses TableName when FilenameOverride is not available",
			&Table{
				TableName: "table_ones",
			},
			"table_ones",
		},
		{
			"uses TableName when FilenameOverride is blank",
			&Table{
				TableName:        "table_ones",
				FilenameOverride: "",
			},
			"table_ones",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.Filename()
			assert.Equal(t, tt.want, got)
		})
	}
}
