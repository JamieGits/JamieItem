package main

 {
	me{
	name
  }
}
{
	"me":{
		"name":"lihao"
}
}

{
	me{
		name,
		profilePicture(size:300){
			width,
			heigth,
			url
   }
  }
}
{
 "me":{
 	"name":"lihao",
 	"profilePicture":{
 		"width":"300",
 		"heigth":"300",
 		"url":"https://cdn/300.jpg"
}
}
}
//嵌套查询
{
	me{
		name,
		friends{
			name
}
}
}
{
	"me"{
		"name":"lihao",
		"friends":[
		 {
		 	"name":"zhangsan"
		 },
		 {
		 	"name":"lisi"
		 	}
		]
}
}
//friends就是朋友的列表，返回来就是一个数组，里面是name的属性，比如不但要朋友列表的用户名，还需要朋友参加活动的名字
{
	me{
	name,
	friends{
		name,
		events{
		name
 }
}
 }
}


{
	"me":{
		"name":"lihao",
		"friends":[
			"name":"zhangsan",
			"events":[
			{
			"name":"啤酒节"
			},
			{
			"name":"万圣节"
			}
			],
			"name":"zhangsan",
			"events":[
			{
			"name":"圣诞节"
			},
			{
			"name":"春节"
			}
			]
]
}
}


//查询里也可以做一些限制，比如排序的标准，或者限制返回来的结果数量
{
	me{
		name,
		friends(orderby: IMPORTANCE, first: 1){
			name,
			events(first: 1){
				name
}
}
}
}

func main()  {
	
}



