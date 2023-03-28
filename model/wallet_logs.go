package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: wallet_logs
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] uu_id                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] addr                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] owner                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] key_type                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] private_key                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 7] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 9] system_wallet_id                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[10] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[11] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[12] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 80,    "uuId": "peFbMgHivRACqmNCvSorweBqP",    "addr": "DZHnFIoBEmueGkCKsbhBavqtY",    "owner": "EiInDJEyANGkwlltXGTjhuYGx",    "keyType": "vDZKOPdjBbTrnnYuaGJHQaYtJ",    "privateKey": "wUkalhcPceXJobZGwJmepOywB",    "nodeInfo": "vgQknMVpqfrrAcQbfEKeknEyr",    "requesterInfo": "oByfGPwsEqniJLssjQCfDFPAQ",    "requestingApiKey": "MRemQfcdJkewUiFHrosfOFoKp",    "systemWalletId": 70,    "createdAt": "2067-08-04T12:59:08.474556194-04:00",    "updatedAt": "2094-09-21T03:52:50.420485709-04:00",    "deltaNodeUuid": "BJLhHabuFNncXQTqTrxbJQkLE"}



*/

// WalletLogs struct is a row record of the wallet_logs table in the estuary database
type WalletLogs struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] uu_id                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	UuID null.String `gorm:"column:uu_id;type:TEXT;"`
	//[ 2] addr                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Addr null.String `gorm:"column:addr;type:TEXT;"`
	//[ 3] owner                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Owner null.String `gorm:"column:owner;type:TEXT;"`
	//[ 4] key_type                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	KeyType null.String `gorm:"column:key_type;type:TEXT;"`
	//[ 5] private_key                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	PrivateKey null.String `gorm:"column:private_key;type:TEXT;"`
	//[ 6] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	NodeInfo null.String `gorm:"column:node_info;type:TEXT;"`
	//[ 7] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequesterInfo null.String `gorm:"column:requester_info;type:TEXT;"`
	//[ 8] requesting_api_key                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequestingAPIKey null.String `gorm:"column:requesting_api_key;type:TEXT;"`
	//[ 9] system_wallet_id                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemWalletID null.Int `gorm:"column:system_wallet_id;type:INT8;"`
	//[10] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[11] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
	//[12] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DeltaNodeUUID null.String `gorm:"column:delta_node_uuid;type:TEXT;"`
}

var wallet_logsTableInfo = &TableInfo{
	Name: "wallet_logs",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "uu_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "UuID",
			GoFieldType:        "null.String",
			JSONFieldName:      "uuId",
			ProtobufFieldName:  "uu_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "addr",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Addr",
			GoFieldType:        "null.String",
			JSONFieldName:      "addr",
			ProtobufFieldName:  "addr",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "owner",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Owner",
			GoFieldType:        "null.String",
			JSONFieldName:      "owner",
			ProtobufFieldName:  "owner",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "key_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "KeyType",
			GoFieldType:        "null.String",
			JSONFieldName:      "keyType",
			ProtobufFieldName:  "key_type",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "private_key",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "PrivateKey",
			GoFieldType:        "null.String",
			JSONFieldName:      "privateKey",
			ProtobufFieldName:  "private_key",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "node_info",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "NodeInfo",
			GoFieldType:        "null.String",
			JSONFieldName:      "nodeInfo",
			ProtobufFieldName:  "node_info",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "requester_info",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "RequesterInfo",
			GoFieldType:        "null.String",
			JSONFieldName:      "requesterInfo",
			ProtobufFieldName:  "requester_info",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "requesting_api_key",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "RequestingAPIKey",
			GoFieldType:        "null.String",
			JSONFieldName:      "requestingApiKey",
			ProtobufFieldName:  "requesting_api_key",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "system_wallet_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "SystemWalletID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "systemWalletId",
			ProtobufFieldName:  "system_wallet_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "createdAt",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "updatedAt",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "delta_node_uuid",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "DeltaNodeUUID",
			GoFieldType:        "null.String",
			JSONFieldName:      "deltaNodeUuid",
			ProtobufFieldName:  "delta_node_uuid",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WalletLogs) TableName() string {
	return "wallet_logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WalletLogs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WalletLogs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WalletLogs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WalletLogs) TableInfo() *TableInfo {
	return wallet_logsTableInfo
}
