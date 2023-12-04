package consts

var e = make(map[int32]string)

const (
	CODE_OK    = iota //成功
	CODE_PARAM        //参数异常
	CODE_FAIL
	CODE_SYSTEM

	CODE_ORG_EXIST
	CODE_ORG_SON_EXIST
	CODE_ACCOUNT_EXIST
	CODE_VERIFICATION
	CODE_PASSWORD_ACCOUNT
	CODE_MULTIPLE_ACCOUNT
	CODE_PERMISSIION
	CODE_H5_TOKEN_INVALID
	CODE_SAMPLE_NOT_EXIST

	CODE_CONNECT_FAILED
	CODE_CONNECT_SUCCESS
	CODE_USER_PWD_NOT_FOUND
	CODE_STORE_FILE_ERROR
	CODE_EXCEL_FORMART_ERROR
	CODE_XLSX_FILE_NAME_ERROR
	CODE_USER_IS_DISABLE
	CODE_DATE_IS_NOT_SUNDAY
	CODE_USER_IS_NOT_CONFIG
	CODE_USER_AUTH_ERROR
	CODE_USER_AUTH_ERROR_COUNT
)

func init() {
	e[CODE_OK] = "操作成功"
	e[CODE_PARAM] = "参数异常"
	e[CODE_FAIL] = "操作失败"
	e[CODE_SYSTEM] = "服务异常"
	e[CODE_ORG_EXIST] = "机构已存在"
	e[CODE_ORG_SON_EXIST] = "存在未删除的下级机构"
	e[CODE_ACCOUNT_EXIST] = "账号已存在"
	e[CODE_VERIFICATION] = "验证码错误"
	e[CODE_PASSWORD_ACCOUNT] = "账号或密码错误"
	e[CODE_MULTIPLE_ACCOUNT] = "此账号在多个企业存在, 请使用账号格式: 账号@企业代号"
	e[CODE_PERMISSIION] = "您的权限不足，请联系系统管理员"
	e[CODE_H5_TOKEN_INVALID] = "token失效"
	e[CODE_SAMPLE_NOT_EXIST] = "样本不存在"
	e[CODE_CONNECT_FAILED] = "连接失败"
	e[CODE_CONNECT_SUCCESS] = "连接成功"
	e[CODE_USER_PWD_NOT_FOUND] = "用户或密码不正确"
	e[CODE_STORE_FILE_ERROR] = "门店运营概况表格名字不正确，正确格式: 门店运营概况周报20220401.xlsx / 门店运营概况日报20220401.xlsx"
	e[CODE_EXCEL_FORMART_ERROR] = "仅支持XLSX、XLS、CSV文件格式"
	e[CODE_XLSX_FILE_NAME_ERROR] = "文件名未按照标准格式编写"
	e[CODE_USER_IS_DISABLE] = "用户已被禁用, 请联系管理员"
	e[CODE_DATE_IS_NOT_SUNDAY] = "文件名称中的日期不是星期日"
	e[CODE_USER_IS_NOT_CONFIG] = "用户未配置科室或角色, 请联系管理员"
	e[CODE_USER_AUTH_ERROR] = "非法访问"
	e[CODE_USER_AUTH_ERROR_COUNT] = "您的账号已被锁定，请五分钟后重试"
}

func GetMessage(code int32) string {
	msg, ok := e[code]
	if !ok {
		return "服务异常,稍后重试."
	}
	return msg
}
