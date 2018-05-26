// 内建函数 make 也可以使用 map 关键字来定义 Map:
// 如果不初始化 map，那么就会创建一个 nil map。
// nil map 不能用来存放键值对

/* 1. 声明变量，默认 map 是 nil */
// var map_variable map[key_data_type]value_data_type
/* 创建map */
// countryCapitalMap := map[string]string{"France": "Paris"}
/* 2. 使用 make 函数 */
// map_variable := make(map[key_data_type]value_data_type)

package main

import "fmt"

func main() {
	// var countryCapitalMap map[string]string
	/*创建集合 */
	countryCapitalMap := make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roman"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New delhi"

	/*删除元素*/
	delete(countryCapitalMap, "India")

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "capital is", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["Americ"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(captial) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("hahaaa", captial)
	} else {
		fmt.Println("nil")
	}
}
