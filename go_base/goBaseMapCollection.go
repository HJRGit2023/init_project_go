package main

import "fmt"

func main() {
	var sitemap map[string]string /* 创建集合 */
	sitemap = make(map[string]string) /* 初始化集合 */
	/* map 插入 key - value 对,各个公司网站 */
	sitemap["Baidu"] = "www.baidu.com" /* 添加元素 */
	sitemap["Google"] = "www.google.com"
	sitemap["Taobao"] = "www.taobao.com"
	sitemap["Sina"] = "www.sina.com.cn"
	sitemap["Sohu"] = "www.sohu.com"
	sitemap["Tencent"] = "www.tencent.com"
	sitemap["Wiki"] = "www.wikipedia.org"
	fmt.Println(sitemap) /* 输出集合 */
	/*使用键输出website的值值 */ 
	for key := range sitemap {
		fmt.Printf("公司名称：%s 网站地址：%s\n",key, sitemap[key])
	}
	/*查看元素在集合中是否存在 */
	name, ok := sitemap["QQ"] /*如果确定是真实的,则存在,否则不存在 */
	if ok {
		fmt.Printf("公司名称：%s 真实存在\n", name)
	} else {
		fmt.Printf("公司名称：%s 不存在\n", "QQ")
	}

	/* 创建map */
    countryCapitalMap :=map[string]string{"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New delhi"}
    fmt.Println("原始地图")
	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "：首都是", countryCapitalMap[country])
	}
	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("删除元素", "France")

	fmt.Println("删除元素后地图")
	/*打印地图*/
	for country :=range countryCapitalMap {
		fmt.Println(country,"：首都是", countryCapitalMap [ country ])
	}
}