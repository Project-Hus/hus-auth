// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// HusSessionsColumns holds the columns for the "hus_sessions" table.
	HusSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "iat", Type: field.TypeTime},
		{Name: "exp", Type: field.TypeBool, Default: false},
		{Name: "uid", Type: field.TypeUUID},
	}
	// HusSessionsTable holds the schema information for the "hus_sessions" table.
	HusSessionsTable = &schema.Table{
		Name:       "hus_sessions",
		Columns:    HusSessionsColumns,
		PrimaryKey: []*schema.Column{HusSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hus_sessions_users_hus_sessions",
				Columns:    []*schema.Column{HusSessionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ServicesColumns holds the columns for the "services" table.
	ServicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "protocol", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ServicesTable holds the schema information for the "services" table.
	ServicesTable = &schema.Table{
		Name:       "services",
		Columns:    ServicesColumns,
		PrimaryKey: []*schema.Column{ServicesColumns[0]},
	}
	// SubdomainsColumns holds the columns for the "subdomains" table.
	SubdomainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "subdomain", Type: field.TypeString},
		{Name: "role", Type: field.TypeString},
		{Name: "service_id", Type: field.TypeInt},
	}
	// SubdomainsTable holds the schema information for the "subdomains" table.
	SubdomainsTable = &schema.Table{
		Name:       "subdomains",
		Columns:    SubdomainsColumns,
		PrimaryKey: []*schema.Column{SubdomainsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subdomains_services_subdomains",
				Columns:    []*schema.Column{SubdomainsColumns[3]},
				RefColumns: []*schema.Column{ServicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "provider", Type: field.TypeEnum, Enums: []string{"hus", "google"}},
		{Name: "google_sub", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "email_verified", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString},
		{Name: "given_name", Type: field.TypeString},
		{Name: "family_name", Type: field.TypeString},
		{Name: "birthdate", Type: field.TypeTime, Nullable: true},
		{Name: "profile_picture_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		HusSessionsTable,
		ServicesTable,
		SubdomainsTable,
		UsersTable,
	}
)

func init() {
	HusSessionsTable.ForeignKeys[0].RefTable = UsersTable
	SubdomainsTable.ForeignKeys[0].RefTable = ServicesTable
}
