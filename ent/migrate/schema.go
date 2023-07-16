// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// VulnInformationsColumns holds the columns for the "vuln_informations" table.
	VulnInformationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString, Default: ""},
		{Name: "description", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "severity", Type: field.TypeString, Default: ""},
		{Name: "cve", Type: field.TypeString, Default: ""},
		{Name: "disclosure", Type: field.TypeString, Default: ""},
		{Name: "solutions", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "references", Type: field.TypeJSON, Nullable: true},
		{Name: "tags", Type: field.TypeJSON, Nullable: true},
		{Name: "github_search", Type: field.TypeJSON, Nullable: true},
		{Name: "from", Type: field.TypeString, Default: ""},
		{Name: "pushed", Type: field.TypeBool, Default: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// VulnInformationsTable holds the schema information for the "vuln_informations" table.
	VulnInformationsTable = &schema.Table{
		Name:       "vuln_informations",
		Columns:    VulnInformationsColumns,
		PrimaryKey: []*schema.Column{VulnInformationsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		VulnInformationsTable,
	}
)

func init() {
}
