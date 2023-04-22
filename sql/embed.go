package sql

import "embed"

//go:embed 23042201_create_table_user.up.sql
//go:embed 23042203_update_table_user.up.sql
var Fs embed.FS
