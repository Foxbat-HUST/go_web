package gen

import "gorm.io/gorm"

type User struct {

	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`

	//[ 1] name                                           varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Name *string `gorm:"column:name;type:varchar;size:30;"`

	//[ 2] type                                           varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	Type string `gorm:"column:type;type:varchar;size:10;"`

	//[ 3] age                                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Age *uint32 `gorm:"column:age;type:int;"`

	//[ 4] email                                          varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Email *string `gorm:"column:email;type:varchar;size:30;"`

	//[ 5] password                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Password string `gorm:"column:password;type:varchar;size:255;"`
}

var usersTableInfo = &TableInfo{
	Name: "users",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
		},

		{
			Index:              1,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Name",
			GoFieldType:        "*string",
		},

		{
			Index:              2,
			Name:               "type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "Type",
			GoFieldType:        "string",
		},

		{
			Index:              3,
			Name:               "age",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Age",
			GoFieldType:        "*uint32",
		},

		{
			Index:              4,
			Name:               "email",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Email",
			GoFieldType:        "*string",
		},

		{
			Index:              5,
			Name:               "password",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Password",
			GoFieldType:        "string",
		},

		{
			Index:              6,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "*time.Time",
		},

		{
			Index:              7,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "*time.Time",
		},

		{
			Index:              8,
			Name:               "deleted_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "DeletedAt",
			GoFieldType:        "*time.Time",
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *User) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *User) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *User) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *User) TableInfo() *TableInfo {
	return usersTableInfo
}
