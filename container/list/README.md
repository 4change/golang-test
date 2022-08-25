# 注意事项
源码版本: go1.11.5
# 双向链表的相关概念
双向链表，包含一个头结点
双向链表支持的操作
- Init(): 清空获取初始化双向链表
- insert(): 在元素at后面插入元素e，扩展双向链表的长度，然后返回被插入的元素e
- PushBack(): 从链表的尾部插入元素，并返回已插入的元素
- PushFront(): 从链表的头部插入元素, 并返回已插入的元素
# 参考代码
[list.go](https://gitee.com/fochange/test-go/blob/master/container/list/list.go)