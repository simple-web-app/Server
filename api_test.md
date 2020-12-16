

# API Test

## Hint

初始用户为user1*~*user10

初始密码为pass1*~*pass10

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

