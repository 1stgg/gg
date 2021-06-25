
彩色打印go变量


使用方法
```golang
gg.Log(arg,arg2,arg3...)
```

颜色对照表
|变量类型|颜色|示例|
|-|-|-|
|Nil|白色加粗|<span style="color:white;font-weight:bolder;">nil</span>|
|Bool|黄色|<span style="color:yellow;">true</span>|
|Int,Float,Complex|黄色|<span style="color:yellow;">1</span>|
|String|绿色|<span style="color:green;">"string"</span>|
|Func|青色|<span style="color:#3eb1d5;">[Function: main.main.func1]</span>|
|Array,Slice||<span style="color:white;">[ <span style="color:green;">"string"</span>  <span style="color:green;">"string"</span> ] [ <span style="color:yellow;">1</span> <span style="color:yellow;">1</span> ]|
|Map||<span style="color:white;">{ <span style="color:green;">"string"</span>: <span style="color:green;">"string"</span> } { <span style="color:yellow;">1</span>: <span style="color:green;">"string"</span> } { <span style="color:green;">"string"</span>: <span style="color:yellow;">1</span> }</span>|
|Struct||{ title: <span style="color:green;">""标题"</span>, id: <span style="color:yellow;">1</span>, page: { num: <span style="color:yellow;">1</span> } }|
|Interface||按值类型打印|
|Others|白色|<span style="color:white;">0xc000076060</span>|


