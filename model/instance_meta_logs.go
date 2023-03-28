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


Table: instance_meta_logs
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] instance_uuid                                  TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] instance_host_name                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] instance_node_name                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] os_details                                     TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] public_ip                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] memory_limit                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 7] cpu_limit                                      INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 8] storage_limit                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 9] disable_request                                BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[10] disable_commitment_piece_generation            BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[11] disable_storage_deal                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[12] disable_online_deals                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[13] disable_offline_deals                          BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[14] number_of_cpus                                 INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[15] storage_in_bytes                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[16] system_memory                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[17] heap_memory                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[18] heap_in_use                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[19] stack_in_use                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[20] instance_start                                 TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[21] bytes_per_cpu                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[22] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[23] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[24] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[25] system_instance_meta_id                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[26] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[27] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 13,    "instance_uuid": "QZqJQFhMxvKdcoPYRhoYIGKXY",    "instance_host_name": "NkObRXNTrSgVTmrnWfaqQtSMG",    "instance_node_name": "uUtfSFLYSNkXgRVjSHnQbEUPM",    "os_details": "IOJXCpfFNFyWWXWGtwIamAglk",    "public_ip": "GJlJfHlSXyKHEaurudKHiDnHf",    "memory_limit": 16,    "cpu_limit": 6,    "storage_limit": 58,    "disable_request": true,    "disable_commitment_piece_generation": true,    "disable_storage_deal": false,    "disable_online_deals": true,    "disable_offline_deals": true,    "number_of_cpus": 60,    "storage_in_bytes": 55,    "system_memory": 96,    "heap_memory": 50,    "heap_in_use": 76,    "stack_in_use": 52,    "instance_start": "2097-04-10T11:10:18.072518734-04:00",    "bytes_per_cpu": 62,    "node_info": "wOyeKPZPqxkMJQmhilThcIMoZ",    "requester_info": "njINpEXWJnnqYmgtpxmHybHPq",    "delta_node_uuid": "yhiYTereSvRGaSamamuAPWUPV",    "system_instance_meta_id": 5,    "created_at": "2147-06-07T15:53:28.278550966-04:00",    "updated_at": "2070-01-08T01:13:47.373733577-05:00"}



*/

// InstanceMetaLogs struct is a row record of the instance_meta_logs table in the estuary database
type InstanceMetaLogs struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] instance_uuid                                  TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	InstanceUUID sql.NullString `gorm:"column:instance_uuid;type:TEXT;"`
	//[ 2] instance_host_name                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	InstanceHostName sql.NullString `gorm:"column:instance_host_name;type:TEXT;"`
	//[ 3] instance_node_name                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	InstanceNodeName sql.NullString `gorm:"column:instance_node_name;type:TEXT;"`
	//[ 4] os_details                                     TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	OsDetails sql.NullString `gorm:"column:os_details;type:TEXT;"`
	//[ 5] public_ip                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	PublicIP sql.NullString `gorm:"column:public_ip;type:TEXT;"`
	//[ 6] memory_limit                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	MemoryLimit sql.NullInt64 `gorm:"column:memory_limit;type:INT8;"`
	//[ 7] cpu_limit                                      INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	CPULimit sql.NullInt64 `gorm:"column:cpu_limit;type:INT8;"`
	//[ 8] storage_limit                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	StorageLimit sql.NullInt64 `gorm:"column:storage_limit;type:INT8;"`
	//[ 9] disable_request                                BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableRequest sql.NullBool `gorm:"column:disable_request;type:BOOL;"`
	//[10] disable_commitment_piece_generation            BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableCommitmentPieceGeneration sql.NullBool `gorm:"column:disable_commitment_piece_generation;type:BOOL;"`
	//[11] disable_storage_deal                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableStorageDeal sql.NullBool `gorm:"column:disable_storage_deal;type:BOOL;"`
	//[12] disable_online_deals                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableOnlineDeals sql.NullBool `gorm:"column:disable_online_deals;type:BOOL;"`
	//[13] disable_offline_deals                          BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableOfflineDeals sql.NullBool `gorm:"column:disable_offline_deals;type:BOOL;"`
	//[14] number_of_cpus                                 INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	NumberOfCpus sql.NullInt64 `gorm:"column:number_of_cpus;type:INT8;"`
	//[15] storage_in_bytes                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	StorageInBytes sql.NullInt64 `gorm:"column:storage_in_bytes;type:INT8;"`
	//[16] system_memory                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemMemory sql.NullInt64 `gorm:"column:system_memory;type:INT8;"`
	//[17] heap_memory                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	HeapMemory sql.NullInt64 `gorm:"column:heap_memory;type:INT8;"`
	//[18] heap_in_use                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	HeapInUse sql.NullInt64 `gorm:"column:heap_in_use;type:INT8;"`
	//[19] stack_in_use                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	StackInUse sql.NullInt64 `gorm:"column:stack_in_use;type:INT8;"`
	//[20] instance_start                                 TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	InstanceStart time.Time `gorm:"column:instance_start;type:TIMESTAMPTZ;"`
	//[21] bytes_per_cpu                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	BytesPerCPU sql.NullInt64 `gorm:"column:bytes_per_cpu;type:INT8;"`
	//[22] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	NodeInfo sql.NullString `gorm:"column:node_info;type:TEXT;"`
	//[23] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequesterInfo sql.NullString `gorm:"column:requester_info;type:TEXT;"`
	//[24] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DeltaNodeUUID sql.NullString `gorm:"column:delta_node_uuid;type:TEXT;"`
	//[25] system_instance_meta_id                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemInstanceMetaID sql.NullInt64 `gorm:"column:system_instance_meta_id;type:INT8;"`
	//[26] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[27] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
}

var instance_meta_logsTableInfo = &TableInfo{
	Name: "instance_meta_logs",
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
			Name:               "instance_uuid",
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
			GoFieldName:        "InstanceUUID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "instance_uuid",
			ProtobufFieldName:  "instance_uuid",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "instance_host_name",
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
			GoFieldName:        "InstanceHostName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "instance_host_name",
			ProtobufFieldName:  "instance_host_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "instance_node_name",
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
			GoFieldName:        "InstanceNodeName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "instance_node_name",
			ProtobufFieldName:  "instance_node_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "os_details",
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
			GoFieldName:        "OsDetails",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "os_details",
			ProtobufFieldName:  "os_details",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "public_ip",
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
			GoFieldName:        "PublicIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "public_ip",
			ProtobufFieldName:  "public_ip",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "memory_limit",
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
			GoFieldName:        "MemoryLimit",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "memory_limit",
			ProtobufFieldName:  "memory_limit",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "cpu_limit",
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
			GoFieldName:        "CPULimit",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "cpu_limit",
			ProtobufFieldName:  "cpu_limit",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "storage_limit",
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
			GoFieldName:        "StorageLimit",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "storage_limit",
			ProtobufFieldName:  "storage_limit",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "disable_request",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "DisableRequest",
			GoFieldType:        "sql.NullBool",
			JSONFieldName:      "disable_request",
			ProtobufFieldName:  "disable_request",
			ProtobufType:       "bool",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "disable_commitment_piece_generation",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "DisableCommitmentPieceGeneration",
			GoFieldType:        "sql.NullBool",
			JSONFieldName:      "disable_commitment_piece_generation",
			ProtobufFieldName:  "disable_commitment_piece_generation",
			ProtobufType:       "bool",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "disable_storage_deal",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "DisableStorageDeal",
			GoFieldType:        "sql.NullBool",
			JSONFieldName:      "disable_storage_deal",
			ProtobufFieldName:  "disable_storage_deal",
			ProtobufType:       "bool",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "disable_online_deals",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "DisableOnlineDeals",
			GoFieldType:        "sql.NullBool",
			JSONFieldName:      "disable_online_deals",
			ProtobufFieldName:  "disable_online_deals",
			ProtobufType:       "bool",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "disable_offline_deals",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "DisableOfflineDeals",
			GoFieldType:        "sql.NullBool",
			JSONFieldName:      "disable_offline_deals",
			ProtobufFieldName:  "disable_offline_deals",
			ProtobufType:       "bool",
			ProtobufPos:        14,
		},

		{
			Index:              14,
			Name:               "number_of_cpus",
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
			GoFieldName:        "NumberOfCpus",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "number_of_cpus",
			ProtobufFieldName:  "number_of_cpus",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		{
			Index:              15,
			Name:               "storage_in_bytes",
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
			GoFieldName:        "StorageInBytes",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "storage_in_bytes",
			ProtobufFieldName:  "storage_in_bytes",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		{
			Index:              16,
			Name:               "system_memory",
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
			GoFieldName:        "SystemMemory",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "system_memory",
			ProtobufFieldName:  "system_memory",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		{
			Index:              17,
			Name:               "heap_memory",
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
			GoFieldName:        "HeapMemory",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "heap_memory",
			ProtobufFieldName:  "heap_memory",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		{
			Index:              18,
			Name:               "heap_in_use",
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
			GoFieldName:        "HeapInUse",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "heap_in_use",
			ProtobufFieldName:  "heap_in_use",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		{
			Index:              19,
			Name:               "stack_in_use",
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
			GoFieldName:        "StackInUse",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "stack_in_use",
			ProtobufFieldName:  "stack_in_use",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		{
			Index:              20,
			Name:               "instance_start",
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
			GoFieldName:        "InstanceStart",
			GoFieldType:        "time.Time",
			JSONFieldName:      "instance_start",
			ProtobufFieldName:  "instance_start",
			ProtobufType:       "uint64",
			ProtobufPos:        21,
		},

		{
			Index:              21,
			Name:               "bytes_per_cpu",
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
			GoFieldName:        "BytesPerCPU",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bytes_per_cpu",
			ProtobufFieldName:  "bytes_per_cpu",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		{
			Index:              22,
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_info",
			ProtobufFieldName:  "node_info",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		{
			Index:              23,
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "requester_info",
			ProtobufFieldName:  "requester_info",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		{
			Index:              24,
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
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "delta_node_uuid",
			ProtobufFieldName:  "delta_node_uuid",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		{
			Index:              25,
			Name:               "system_instance_meta_id",
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
			GoFieldName:        "SystemInstanceMetaID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "system_instance_meta_id",
			ProtobufFieldName:  "system_instance_meta_id",
			ProtobufType:       "int32",
			ProtobufPos:        26,
		},

		{
			Index:              26,
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        27,
		},

		{
			Index:              27,
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
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        28,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *InstanceMetaLogs) TableName() string {
	return "instance_meta_logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *InstanceMetaLogs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *InstanceMetaLogs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *InstanceMetaLogs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *InstanceMetaLogs) TableInfo() *TableInfo {
	return instance_meta_logsTableInfo
}
