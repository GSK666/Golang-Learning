# 切片   
slice是一个比数组更加强大的序列接口   
slice的类型仅由所包含的元素决定，需要使用内建的方法make     
make([]类型，大小)       
len返回slice的长度   
append方法添加元素   
copy复制slice    
slice通过slice[low:high]语法进行切片操作     
*******
# 关联数组   
map是go内置关联数组类型     
使用make创建空map，make([key-type]val-type)    
使用典型的make[key] = val语法进行键值设置     
println输出所有的键值    
name[key]获取一个键的值    
len返回键值对的数目     
delete从一个map中移除键值对     
可以直接申明和初始化一个新的map    
**********
# range迭代   
range可以迭代各种各样的数据结构    
range在数组和slice中都提供每个项的索引和值，忽略使用空值定义符_     
range在map中迭代键值对    
range在字符串中地阿呆unicode编码，第一个是rune的起始字节位置，第二个是rune的值     
*******
