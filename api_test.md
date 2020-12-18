

# API Test

## Hint

初始用户为user1\~user10

初始密码为pass1\~pass10

创建评论时需要token对应的用户名作为Arthur，否则无法创建

## The way to use api

### GET

#### 	**Sign in**

```
http://localhost:8080/blog/user/signin
```

Get时需要提供username和password

举例

```
http://localhost:8080/blog/user/signin?username=user1&password=pass1
```

返回值

```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDgwMDk0NzIsImlhdCI6MTYwODAwNTg3Mn0.7kLZW1xH7cdsQTQvM16aC_oSBxiczIQLZLXOqeIm5_c"
}
```

#### 	Get Articles

```
http://localhost:8080/blog/user/articles
```

Get时需要提供page作为参数

举例

```
http://localhost:8080/blog/user/articles?page=1
```

  返回值

```
{
  "PageCount": 2,
  "Articles": [
    {
      "id": 1,
      "name": "QAQ"
    },
    {
      "id": 2,
      "name": "QAQ"
    }
  ]
}
```

#### 	Get Article By Id

```
http://localhost:8080/blog/user/article/{id}
```

举例

```
http://localhost:8080/blog/user/article/1
```

返回值

```
{
  "id": 1,
  "name": "QAQ",
  "tags": [
    {
      "name": "CS"
    },
    {
      "name": "SC"
    }
  ],
  "date": "2019",
  "content": "qwq"
}
```

#### 	Delete Article

```
http://localhost:8080/blog/user/deleteArticle/{id}
```

举例

```
http://localhost:8080/blog/user/deleteArticle/1
```

返回值

```
"Success"
```



#### 	Get Comments

```
http://localhost:8080/blog/user/article/{id}/comments
```

Get时需要提供page作为参数

举例

```
http://localhost:8080/blog/user/article/2/comments?page=1
```

返回值

（其中date为time，content和author为string， articleId为int32）

```
{
  "PageCount": 2,
  "contents": [
    {
      "date": "2020.12.14 20:12:05",
      "content": "qqq",
      "author": "quq",
      "articleId": 2
    },
    {
      "date": "2020.12.14 20:12:04",
      "content": "qqqa",
      "author": "qwe",
      "articleId": 2
    }
  ]
}
```

#### GetTagById

```
http://localhost:8080/blog/user/tag/{id}
```

使用时需要指定tag的ID来找到对应的tag

举例：

```
http://localhost:8080/blog/user/tag/1
```

结果：

![](./fig/getTagById.PNG)



#### GetTags

```
http://localhost:8080/blog/user/getTags
```

使用时不需要任何参数，会直接输出所有tags


结果：

![](./fig/addTag.PNG)



### POST

#### 	Create Comment

```
http://localhost:8080/blog/user/article/{id}/comments
```

举例

```
http://localhost:8080/blog/user/article/1/comments
```

需要headers

Authorization : token（token为登录时返回的token）

需要Body

{"content":"contents"} 属性为string

返回值

```
{
      "date": "2020-12-15 13:02:41",
      "content": "new content3",
      "author": "user1",
      "articleId": 2
    }
```


#### Add Article

```
http://localhost:8080/blog/user/addArticle
```

需要Body内容，格式如下：
![](./fig/addArticle.PNG)

举例：
```
{
"name":"new article",
"tags":[{"name":"test"}],
"content":"user add article test"
}
```

结果：
```
{
  "id": 6,
  "name": "new article",
  "tags": [
    {
      "name": "test"
    }
  ],
  "date": "2020-12-16 22:27:15",
  "content": "user add article test"
}
```

然后查看articles可以看到：

```
{
 http://localhost:8080/blog/user/articles?page=1
```
![](./fig/articles.PNG)

显然，这时已经成功添加文件了。



#### AddTag

```
http://localhost:8080/blog/user/addTag
```

功能为添加Tag,使用时需要Body参数，如下

```
{"name":"tag name"}
```

举例：

```
{"name":"tag2"}
```

然后运行

```
http://localhost:8080/blog/user/getTags
```
可以看到如下结果：

![](./fig/addTag.PNG)

显然，此时已经成功添加了一个新的tag.