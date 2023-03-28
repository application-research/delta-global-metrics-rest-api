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
{    "id": 93,    "instanceUuid": "BMIMASFbwVoZulNBUFGZhfSVW",    "instanceHostName": "CxuBGAQEDwdvgxuJhqFibVCfh",    "instanceNodeName": "IZJXagGdYFLncCRDLYwhmfkZS",    "osDetails": "nnstxhjfCoqFGQIwDGWKcfbsV",    "publicIp": "BycLLOtyYWbmWyCoxUhEubLWN",    "memoryLimit": 62,    "cpuLimit": 10,    "storageLimit": 26,    "disableRequest": true,    "disableCommitmentPieceGeneration": true,    "disableStorageDeal": true,    "disableOnlineDeals": false,    "disableOfflineDeals": false,    "numberOfCpus": 78,    "storageInBytes": 41,    "systemMemory": 30,    "heapMemory": 10,    "heapInUse": 97,    "stackInUse": 65,    "instanceStart": "2253-01-06T10:26:22.503703655-05:00",    "bytesPerCpu": 75,    "nodeInfo": "shManTAukFURoYpTdAIePHgKx",    "requesterInfo": "WKdYjoCMUgHxZWUuYUJLMmuJt",    "deltaNodeUuid": "ioDKKXECtpdyPvcQOFnHjpXai",    "systemInstanceMetaId": 64,    "createdAt": "2153-05-14T11:36:44.108263386-04:00",    "updatedAt": "2109-10-16T05:28:35.049000199-04:00"}



*/

// InstanceMetaLogs struct is a row record of the instance_meta_logs table in the estuary database
type InstanceMetaLogs struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;"`
	//[ 1] instance_uuid                                  TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	InstanceUUID null.String `gorm:"column:instance_uuid;type:TEXT;"`
	//[ 2] instance_host_name                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	InstanceHostName null.String `gorm:"column:instance_host_name;type:TEXT;"`
	//[ 3] instance_node_name                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	InstanceNodeName null.String `gorm:"column:instance_node_name;type:TEXT;"`
	//[ 4] os_details                                     TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	OsDetails null.String `gorm:"column:os_details;type:TEXT;"`
	//[ 5] public_ip                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	PublicIP null.String `gorm:"column:public_ip;type:TEXT;"`
	//[ 6] memory_limit                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	MemoryLimit null.Int `gorm:"column:memory_limit;type:INT8;"`
	//[ 7] cpu_limit                                      INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	CPULimit null.Int `gorm:"column:cpu_limit;type:INT8;"`
	//[ 8] storage_limit                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	StorageLimit null.Int `gorm:"column:storage_limit;type:INT8;"`
	//[ 9] disable_request                                BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableRequest null.Int `gorm:"column:disable_request;type:BOOL;"`
	//[10] disable_commitment_piece_generation            BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableCommitmentPieceGeneration null.Int `gorm:"column:disable_commitment_piece_generation;type:BOOL;"`
	//[11] disable_storage_deal                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableStorageDeal null.Int `gorm:"column:disable_storage_deal;type:BOOL;"`
	//[12] disable_online_deals                           BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableOnlineDeals null.Int `gorm:"column:disable_online_deals;type:BOOL;"`
	//[13] disable_offline_deals                          BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	DisableOfflineDeals null.Int `gorm:"column:disable_offline_deals;type:BOOL;"`
	//[14] number_of_cpus                                 INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	NumberOfCpus null.Int `gorm:"column:number_of_cpus;type:INT8;"`
	//[15] storage_in_bytes                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	StorageInBytes null.Int `gorm:"column:storage_in_bytes;type:INT8;"`
	//[16] system_memory                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemMemory null.Int `gorm:"column:system_memory;type:INT8;"`
	//[17] heap_memory                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	HeapMemory null.Int `gorm:"column:heap_memory;type:INT8;"`
	//[18] heap_in_use                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	HeapInUse null.Int `gorm:"column:heap_in_use;type:INT8;"`
	//[19] stack_in_use                                   INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	StackInUse null.Int `gorm:"column:stack_in_use;type:INT8;"`
	//[20] instance_start                                 TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	InstanceStart null.Time `gorm:"column:instance_start;type:TIMESTAMPTZ;"`
	//[21] bytes_per_cpu                                  INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	BytesPerCPU null.Int `gorm:"column:bytes_per_cpu;type:INT8;"`
	//[22] node_info                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	NodeInfo null.String `gorm:"column:node_info;type:TEXT;"`
	//[23] requester_info                                 TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	RequesterInfo null.String `gorm:"column:requester_info;type:TEXT;"`
	//[24] delta_node_uuid                                TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DeltaNodeUUID null.String `gorm:"column:delta_node_uuid;type:TEXT;"`
	//[25] system_instance_meta_id                        INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	SystemInstanceMetaID null.Int `gorm:"column:system_instance_meta_id;type:INT8;"`
	//[26] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:TIMESTAMPTZ;"`
	//[27] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;"`
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
			GoFieldType:        "null.String",
			JSONFieldName:      "instanceUuid",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "instanceHostName",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "instanceNodeName",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "osDetails",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "publicIp",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "memoryLimit",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "cpuLimit",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "storageLimit",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "disableRequest",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "disableCommitmentPieceGeneration",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "disableStorageDeal",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "disableOnlineDeals",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "disableOfflineDeals",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "numberOfCpus",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "storageInBytes",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "systemMemory",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "heapMemory",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "heapInUse",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "stackInUse",
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
			GoFieldType:        "null.Time",
			JSONFieldName:      "instanceStart",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "bytesPerCpu",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "nodeInfo",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "requesterInfo",
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
			GoFieldType:        "null.String",
			JSONFieldName:      "deltaNodeUuid",
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
			GoFieldType:        "null.Int",
			JSONFieldName:      "systemInstanceMetaId",
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
			GoFieldType:        "null.Time",
			JSONFieldName:      "createdAt",
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
			GoFieldType:        "null.Time",
			JSONFieldName:      "updatedAt",
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
