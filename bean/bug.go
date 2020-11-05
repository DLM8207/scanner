package bean

/**
 *bug type
 */
type Bug struct {
	/* bug id */
	id string
	/* 发布日期 */
	publicDate string
	/* 级别 */
	level string
	/* 影响产品 */
	product string
	/* 版本，可能多个版本 */
	version []string
	/* bug描述 */
	desc string
	/* bug类型，eg:通用bug */
	bugType string
	/* 相关链接 */
	link string
	/* 解决方法 */
	solution string
	/* 补丁 */
	patch Patch
}
