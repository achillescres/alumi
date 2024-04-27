// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MatchesColumns holds the columns for the "matches" table.
	MatchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "accepted", "declined"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "mentor_id", Type: field.TypeInt},
		{Name: "menti_id", Type: field.TypeInt},
	}
	// MatchesTable holds the schema information for the "matches" table.
	MatchesTable = &schema.Table{
		Name:       "matches",
		Columns:    MatchesColumns,
		PrimaryKey: []*schema.Column{MatchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "matches_mentors_mentor",
				Columns:    []*schema.Column{MatchesColumns[3]},
				RefColumns: []*schema.Column{MentorsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "matches_mentis_menti",
				Columns:    []*schema.Column{MatchesColumns[4]},
				RefColumns: []*schema.Column{MentisColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "match_mentor_id_menti_id",
				Unique:  true,
				Columns: []*schema.Column{MatchesColumns[3], MatchesColumns[4]},
			},
		},
	}
	// MentisColumns holds the columns for the "mentis" table.
	MentisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// MentisTable holds the schema information for the "mentis" table.
	MentisTable = &schema.Table{
		Name:       "mentis",
		Columns:    MentisColumns,
		PrimaryKey: []*schema.Column{MentisColumns[0]},
	}
	// MentorsColumns holds the columns for the "mentors" table.
	MentorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "search_for", Type: field.TypeString},
	}
	// MentorsTable holds the schema information for the "mentors" table.
	MentorsTable = &schema.Table{
		Name:       "mentors",
		Columns:    MentorsColumns,
		PrimaryKey: []*schema.Column{MentorsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString},
		{Name: "message_from", Type: field.TypeInt},
		{Name: "message_match", Type: field.TypeInt},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_users_from",
				Columns:    []*schema.Column{MessagesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "messages_matches_match",
				Columns:    []*schema.Column{MessagesColumns[4]},
				RefColumns: []*schema.Column{MatchesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "message_created_at_message_match_message_from",
				Unique:  false,
				Columns: []*schema.Column{MessagesColumns[1], MessagesColumns[4], MessagesColumns[3]},
			},
		},
	}
	// RealExperiencesColumns holds the columns for the "real_experiences" table.
	RealExperiencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "where", Type: field.TypeString},
		{Name: "as", Type: field.TypeString},
		{Name: "how", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "when_started", Type: field.TypeTime, Nullable: true},
		{Name: "when_ended", Type: field.TypeTime, Nullable: true},
		{Name: "user_real_experiences", Type: field.TypeInt, Nullable: true},
	}
	// RealExperiencesTable holds the schema information for the "real_experiences" table.
	RealExperiencesTable = &schema.Table{
		Name:       "real_experiences",
		Columns:    RealExperiencesColumns,
		PrimaryKey: []*schema.Column{RealExperiencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "real_experiences_users_real_experiences",
				Columns:    []*schema.Column{RealExperiencesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "login", Type: field.TypeString},
		{Name: "hashed_password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "bio", Type: field.TypeString},
		{Name: "education_info", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "telegram", Type: field.TypeString, Nullable: true},
		{Name: "other_contacts", Type: field.TypeString, Nullable: true},
		{Name: "skills", Type: field.TypeJSON},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"menti", "mentor"}},
		{Name: "user_menti", Type: field.TypeInt, Nullable: true},
		{Name: "user_mentor", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_mentis_menti",
				Columns:    []*schema.Column{UsersColumns[11]},
				RefColumns: []*schema.Column{MentisColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_mentors_mentor",
				Columns:    []*schema.Column{UsersColumns[12]},
				RefColumns: []*schema.Column{MentorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MatchesTable,
		MentisTable,
		MentorsTable,
		MessagesTable,
		RealExperiencesTable,
		UsersTable,
	}
)

func init() {
	MatchesTable.ForeignKeys[0].RefTable = MentorsTable
	MatchesTable.ForeignKeys[1].RefTable = MentisTable
	MessagesTable.ForeignKeys[0].RefTable = UsersTable
	MessagesTable.ForeignKeys[1].RefTable = MatchesTable
	RealExperiencesTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = MentisTable
	UsersTable.ForeignKeys[1].RefTable = MentorsTable
}
