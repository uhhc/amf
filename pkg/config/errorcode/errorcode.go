package errorcode

// Code define
const (
	OK int32 = 200

	// system error
	SystemError  int32 = 1000
	DbError      int32 = 1001
	NoPermission int32 = 1002

	// business error
	RequestParamError int32 = 2001
	CreateDataError   int32 = 2002
	UpdateDataError   int32 = 2003
	DeleteDataError   int32 = 2004
	GetDataError      int32 = 2005
	ResourceNotFound  int32 = 2006
)

// CodeMap is a mapping for code and error info
var CodeMap = map[int32]map[string]string{
	OK: {
		"cn": "成功",
		"en": "Success",
	},
	SystemError: {
		"cn": "系统错误",
		"en": "System error",
	},
	DbError: {
		"cn": "数据库错误",
		"en": "DB error",
	},
	NoPermission: {
		"cn": "您没有权限进行该操作",
		"en": "No permission",
	},
	RequestParamError: {
		"cn": "请求参数错误",
		"en": "Request params error",
	},
	CreateDataError: {
		"cn": "创建失败",
		"en": "Create data failure",
	},
	UpdateDataError: {
		"cn": "更新失败",
		"en": "Update data failure",
	},
	DeleteDataError: {
		"cn": "删除失败",
		"en": "Delete data failure",
	},
	GetDataError: {
		"cn": "获取信息失败",
		"en": "Get data failure",
	},
	ResourceNotFound: {
		"cn": "您要找的资源不存在",
		"en": "The resource does not exist",
	},
}
