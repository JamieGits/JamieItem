package main

import "fmt"

func main()  {
	var countryCapitalMap  map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["甘肃"] = "兰州"
	countryCapitalMap["江西"] = "南昌"
	countryCapitalMap["湖北"] = "武汉"
	countryCapitalMap["海南"] = "海口"
	keys := make([]string, 0, len(countryCapitalMap))
	fmt.Println("原始map")
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
		keys = append(keys, country)
		fmt.Printf("keys: %v\n",keys)
	}

	/* 查看元素在集合中是否存在 */

	capital, ok := countryCapitalMap["湖南"]

	if(ok){
		fmt.Println("Capital of 湖南 is",capital)
	}else{
		fmt.Println("Capital of 湖南 is not present")
	}
   //删除元素
   delete(countryCapitalMap,"海南")
   fmt.Println("Entry for 海南 is deleted")
   fmt.Println("删除后的map元素")
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
   fmt.Println(uint64(15))
}
