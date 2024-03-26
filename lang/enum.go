package lang

const (
	arrLen   = _end - _start - 1
	langSize = 7
)

var (
	Langs     [langSize]string
	LangCodes [langSize]uint16
	_arr      [langSize + 1][arrLen]string
)

const (
	I18N_ZH_CN = "zh_CN"
	I18N_ZH_AA = "zh-aa"
	I18N_EN_US = "en_US"
	I18N_EN_BB = "en -bb"
	I18N_CC = "cc"
	I18N_BB = "bb"
	I18N_AA = "aa"
)

func init() {
	Langs[0] = I18N_ZH_CN
	Langs[1] = I18N_ZH_AA
	Langs[2] = I18N_EN_US
	Langs[3] = I18N_EN_BB
	Langs[4] = I18N_CC
	Langs[5] = I18N_BB
	Langs[6] = I18N_AA
	for i := uint16(0); i < langSize; i++ {
		LangCodes[i] = i + 1
		_Code_supported[Langs[i]] = i + 1
	}
	SetDefaultLocale(I18N_ZH_CN)
}

const (
	_start Code = 1000 + iota
	COLUMN_1 // 区域 //  // Test Name`~!@#$%^&*()_+-=/?><.,qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM;:'"[]{}|\ //  // 区域 // 区域 // Test Name`~!@#$%^&*()_+-=/?><.,qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM;:'"[]{}|\
	COLUMN_2 // 区域测试 //  // Space //  // 区域测试 // 区域测试 // Space
	COLUMN_3 // 区域1 //  // Space1 //  // 区域2 // 区域3 // Space2
	COLUMN_4 // 区域2 //  // Space2 //  // 区域3 // 区域4 // Space3
	COLUMN_5 // 区域3 //  // Space3 //  // 区域4 // 区域5 // Space4
	COLUMN_6 // 区域4 //  // Space4 //  // 区域5 // 区域6 // Space5
	COLUMN_7 // 区域5 //  // Space5 //  // 区域6 // 区域7 // Space6
	COLUMN_8 // 区域6 //  // Space6 //  // 区域7 // 区域8 // Space7
	COLUMN_9 // 区域7 //  // Space7 //  // 区域8 // 区域9 // Space8
	COLUMN_10 // 区域8 //  // Space8 //  // 区域9 // 区域10 // Space9
	COLUMN_11 // 区域9 //  // Space9 //  // 区域10 // 区域11 // Space10
	SHEET_1 // 区域10 //  // Space10 //  // 区域11 // 区域12 // Space11
	SHEET_2 // 区域11 //  // Space11 //  // 区域12 // 区域13 // Space12
	SHEET_3 // 区域12 //  // Space12 //  // 区域13 // 区域14 // Space13
	SHEET_4 // 区域13 //  // Space13 //  // 区域14 // 区域15 // Space14
	SHEET_5 // 区域14 //  // Space14 //  // 区域15 // 区域16 // Space15
	SHEET_6 // 区域15 //  // Space15 //  // 区域16 // 区域17 // Space16
	SHEET_7 // 区域16 //  // Space16 //  // 区域17 // 区域18 // Space17
	SHEET_8 // 区域17 //  // Space17 //  // 区域18 // 区域19 // Space18
	SHEET_9 // 区域18 //  // Space18 //  // 区域19 // 区域20 // Space19
	SHEET_10 // 区域19 //  // Space19 //  // 区域20 // 区域21 // Space20
	SHEET_11 // 区域20 //  // Space20 //  // 区域21 // 区域22 // Space21
	TOTAL // 总计 //  // Total //  // 总计 // 总计 // Total
	STAT // 统计 //  // Statistics //  // 统计 // 统计 // Statistics
	EXPORT_TIME // 导出时间： //  // Export time: //  // 导出时间： // 导出时间： // Export time:
	ERROR_CODE_10001 // 参数错误 // 参数错误 // The parameter is incorrect // The parameter is incorrect //  //  // 
	ERROR_CODE_10002 // 获取二维码失败，请检查设备身份! // 获取二维码失败，请检查设备身份! // Failed to get the QR code, please check the device identity // Failed to get the QR code, please check the device identity //  //  // 
	ERROR_CODE_10003 // 未知的位号 // 未知的位号 // Unknow tag // Unknow tag //  //  // 
	ERROR_CODE_10004 // 操作异常，请重试 // 操作异常，请重试 // The operation is abnormal, please try again // The operation is abnormal, please try again //  //  // 
	ERROR_CODE_10005 // 数据初始化中，请稍等 // 数据初始化中，请稍等 // Initializing, please wait // Initializing, please wait //  //  // 
	STR_INTERFACE_RESPONSE_146 // 只能上传XML文件 // 只能上传XML文件 // Only XML files can be uploaded! // Only XML files can be uploaded! //  //  // 
	STR_INTERFACE_RESPONSE_147 // 下载（%s）文件失败 // 下载（%s）文件失败 // The download of the file (%s) failed // The download of the file (%s) failed //  //  // 
	STR_INTERFACE_RESPONSE_148 // 文件夹名称不能为空 // 文件夹名称不能为空 // The folder name cannot be empty // The folder name cannot be empty //  //  // 
	STR_INTERFACE_RESPONSE_149 // 删除对象为空 // 删除对象为空 // The object to be deleted is empty // The object to be deleted is empty //  //  // 
	STR_INTERFACE_RESPONSE_150 // 不支持文件夹移动 // 不支持文件夹移动 // Folder movement is not supported // Folder movement is not supported //  //  // 
	STR_INTERFACE_RESPONSE_151 // 文件不存在 // 文件不存在 // The file does not exist // The file does not exist //  //  // 
	STR_INTERFACE_RESPONSE_152 // 查询失败 // 查询失败 // The query failed // The query failed //  //  // 
	STR_INTERFACE_RESPONSE_153 // 存在同名文件夹，修改后重试 // 存在同名文件夹，修改后重试 // Existing folder with the same name, please try again after modification. // Existing folder with the same name, please try again after modification. //  //  // 
	STR_INTERFACE_RESPONSE_154 // 文件夹名称不能超过%d个字符 // 文件夹名称不能超过%d个字符 // The folder name cannot exceed %d characters // The folder name cannot exceed %d characters //  //  // 
	STR_INTERFACE_RESPONSE_155 // 文件夹名称不能出现以下特殊字符：\/:*?"<>| // 文件夹名称不能出现以下特殊字符：\/:*?"<>| // The folder name cannot contain the following special characters: \/:*?"<>| // The folder name cannot contain the following special characters: \/:*?"<>| //  //  // 
	STR_INTERFACE_RESPONSE_156 // 下载失败 // 下载失败 // Download failed // Download failed //  //  // 
	STR_INTERFACE_RESPONSE_157 // 下载成压缩包失败，下载数量少于%d个 // 下载成压缩包失败，下载数量少于%d个 // The download failed as a compressed file, with fewer than %d items downloaded // The download failed as a compressed file, with fewer than %d items downloaded //  //  // 
	STR_INTERFACE_RESPONSE_158 // 导出数据为空 // 导出数据为空 // The export data is empty // The export data is empty //  //  // 
	STR_INTERFACE_RESPONSE_159 // 不支持的后缀类型 // 不支持的后缀类型 // Unsupported suffix types // Unsupported suffix types //  //  // 
	STR_INTERFACE_RESPONSE_165 // 非法文件，请检查后重新导入 // 非法文件，请检查后重新导入 // Invalid file, please check and import again // Invalid file, please check and import again //  //  // 
	STR_INTERFACE_RESPONSE_166 // 上传失败：请检查配置文件 // 上传失败：请检查配置文件 // Upload failed, please check the configuration file // Upload failed, please check the configuration file //  //  // 
	_end
)
