---
title: 默认模块
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# 默认模块

Base URLs:

# Authentication

- HTTP Authentication, scheme: bearer

# Default

## GET 通过id查询单词

GET /api/v1/words/{wordId}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|wordId|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "data": {
    "id": 12,
    "img": "2e5188d21742dd5408b10552f51528e9_189487_1732240746.jpg",
    "word": "awkward",
    "phonetic": "/ˈɔːkwərd/",
    "mean": "adj.尴尬的；笨拙的；不雅观的；棘手的；不方便的",
    "sound": "awkward.mp3",
    "created_at": "2025-10-15T22:02:30+08:00"
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 查询用户生词

GET /api/v1/user-new-words

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|nick|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 21,
        "word_id": 10,
        "nick": "柚子皮",
        "is_mastered": false,
        "img": "78e1cc3da722b51b21203c8969321d3c.jpg",
        "word": "at this rate",
        "phonetic": "/æt ðɪs reɪt/",
        "mean": "n.照这样下去",
        "sound": "at this rate.mp3",
        "created_at": "2025-12-13T16:08:12+08:00"
      },
      {
        "id": 19,
        "word_id": 11,
        "nick": "柚子皮",
        "is_mastered": true,
        "img": "0e8ef21e3ff40ada50158e9ba4c2700b_185138_1571630239.jpeg",
        "word": "attractive",
        "phonetic": "/əˈtræktɪv/",
        "mean": "adj.（事物）有吸引力的，诱人的；有魅力的，性感的；引起兴趣的",
        "sound": "attractive.mp3",
        "created_at": "2025-12-13T15:49:07+08:00"
      },
      {
        "id": 10,
        "word_id": 13,
        "nick": "柚子皮",
        "is_mastered": false,
        "img": "5ef73b3793a90788f352a20c9e0d1ba3_81099_1538624010.jpg",
        "word": "balance",
        "phonetic": "/ˈbæləns/",
        "mean": "n.平衡，平衡能力；天平；  v.平衡；保持平衡",
        "sound": "balance.mp3",
        "created_at": "2025-12-13T15:27:39+08:00"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 32,
      "total": 3,
      "total_pages": 1
    }
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 单词列表

GET /api/v1/words

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 1,
        "img": "fbd3d74be33a645b04e9b2136aeaf151_118640_1718087848.jpg",
        "word": "ache",
        "phonetic": "/eɪk/",
        "mean": "n.疼痛；  v.觉得疼痛",
        "sound": "ache.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 2,
        "img": "ba231fcdd3650a847c268c8e44c29c65_145026_1595471950.jpeg",
        "word": "adult",
        "phonetic": "/əˈdʌlt/",
        "mean": "n.成年人，成年动物；  adj.成年的；成年人的；（内容）色情的",
        "sound": "adult.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 3,
        "img": "61ff9cdeb0e93da7c9408f222ee53360_151049_1538648807.jpg",
        "word": "aircraft",
        "phonetic": "/ˈerkræft/",
        "mean": "n.飞机，航空器",
        "sound": "aircraft.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 4,
        "img": "dbec85012975f16afbbcec3b39a0f5e3_27345_1534129612.jpg",
        "word": "alcohol",
        "phonetic": "/ˈælkəhɔːl/",
        "mean": "n.酒；酒精，乙醇",
        "sound": "alcohol.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 5,
        "img": "feb29b8ed86bc498b84ce24a06046a8e_78348_1565343414.jpg",
        "word": "allowance",
        "phonetic": "/əˈlaʊəns/",
        "mean": "n.零用钱",
        "sound": "allowance.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 6,
        "img": "20120429_04_43_54_887-gigapixel-scale.jpg",
        "word": "analysis",
        "phonetic": "/əˈnæləsɪs/",
        "mean": "n.分析报告",
        "sound": "analysis.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 7,
        "img": "194bd1e9d5520073d0d8590679fae7b9_199324_1564738253.jpeg",
        "word": "appeal",
        "phonetic": "/əˈpiːl/",
        "mean": "v.请求；上诉；有吸引力；  n.呼吁；请求；吸引力； 筹款努力",
        "sound": "appeal.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 8,
        "img": "cc1f3ee890569084b1efeb8b4e2ac5c0_97405_1584613498.jpeg",
        "word": "assist",
        "phonetic": "/əˈsɪst/",
        "mean": "v.帮助；  n.帮助",
        "sound": "assist.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 9,
        "img": "ddd6cc71bb8876ef3b58482a99d50804_169119_1565696258.jpg",
        "word": "at the same time",
        "phonetic": "/ət ðə seɪm taɪm/",
        "mean": "phr.同时；然而",
        "sound": "at the same time.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 10,
        "img": "78e1cc3da722b51b21203c8969321d3c.jpg",
        "word": "at this rate",
        "phonetic": "/æt ðɪs reɪt/",
        "mean": "n.照这样下去",
        "sound": "at this rate.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 11,
        "img": "0e8ef21e3ff40ada50158e9ba4c2700b_185138_1571630239.jpeg",
        "word": "attractive",
        "phonetic": "/əˈtræktɪv/",
        "mean": "adj.（事物）有吸引力的，诱人的；有魅力的，性感的；引起兴趣的",
        "sound": "attractive.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 12,
        "img": "2e5188d21742dd5408b10552f51528e9_189487_1732240746.jpg",
        "word": "awkward",
        "phonetic": "/ˈɔːkwərd/",
        "mean": "adj.尴尬的；笨拙的；不雅观的；棘手的；不方便的",
        "sound": "awkward.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 13,
        "img": "5ef73b3793a90788f352a20c9e0d1ba3_81099_1538624010.jpg",
        "word": "balance",
        "phonetic": "/ˈbæləns/",
        "mean": "n.平衡，平衡能力；天平；  v.平衡；保持平衡",
        "sound": "balance.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 14,
        "img": "081f47cf56367074ea7e9574aa6aecd7.jpeg",
        "word": "be helpful to",
        "phonetic": "/bi helpfl tuː/",
        "mean": "phr.有助于...",
        "sound": "be helpful to.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 15,
        "img": "fbd01b065db074aa8dd738c0956b17f5_97714_1604029987.jpg",
        "word": "beyond",
        "phonetic": "/bɪˈjɑːnd/",
        "mean": "prep.超出；  adv.在更远处，在另一边",
        "sound": "beyond.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 16,
        "img": "02af981cb698f16be86c1d33952961d3.jpg",
        "word": "bring to",
        "phonetic": "/brɪŋ tu/",
        "mean": "phr.使恢复知觉",
        "sound": "bring to.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 17,
        "img": "capture_9756_20231106_0d0a9a22.jpg",
        "word": "capture",
        "phonetic": "/ˈkæptʃər/",
        "mean": "v.捕获，抓住；留存；  n.战利品",
        "sound": "capture.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 18,
        "img": "3_0_20160720071835_740_c-gigapixel-scale.jpg",
        "word": "chart",
        "phonetic": "/tʃɑːrt/",
        "mean": "n.图表，表格；  v.制成图表",
        "sound": "chart.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 19,
        "img": "7e182082307dc2d66003613104529b44_97847_1600243414.jpg",
        "word": "check in",
        "phonetic": "/tʃek ɪn/",
        "mean": "phr.登记（在旅馆、机场等）",
        "sound": "check in.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 20,
        "img": "6563661a468c6a9bed1498ae8c9e2698_29991_1584955979.jpeg",
        "word": "clear away",
        "phonetic": "/klɪr əˈweɪ/",
        "mean": "phr.把…收走",
        "sound": "clear away.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 21,
        "img": "519e510bb27e1be2a7c4fea1901d9634_58379_1559730507.jpg",
        "word": "compensate",
        "phonetic": "/ˈkɑːmpenseɪt/",
        "mean": "v.赔偿",
        "sound": "compensate.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 22,
        "img": "aa3618119db26ae12ffd276523fc57f4.jpg",
        "word": "complain about",
        "phonetic": "/kəmˈplein əˈbaʊt/",
        "mean": "phr.抱怨",
        "sound": "complain about.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 23,
        "img": "i_9_7117_0_2_20150808154112-gigapixel-scale.jpg",
        "word": "completely",
        "phonetic": "/kəmˈpliːtli/",
        "mean": "adv.彻底地",
        "sound": "completely.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 24,
        "img": "1976f0d81c338ccb41a0bc758e5eb43f_149347_1561953069.jpg",
        "word": "comprise",
        "phonetic": "/kəmˈpraɪz/",
        "mean": "v.构成，由……组成；包含，包括",
        "sound": "comprise.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 25,
        "img": "i_1_10952_0_4_150925180737-gigapixel-scale.jpg",
        "word": "conclusion",
        "phonetic": "/kənˈkluːʒn/",
        "mean": "n.结论，推论",
        "sound": "conclusion.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 26,
        "img": "i_24_7956_0_6_151123182919-gigapixel-scale.jpg",
        "word": "connect",
        "phonetic": "/kəˈnekt/",
        "mean": "v.连接，关联；与……有同感；把……联系起来",
        "sound": "connect.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 27,
        "img": "76d2d6159165bdc68f13225913fa53e8_207790_1749720161.jpg",
        "word": "consideration",
        "phonetic": "/ kənˌsɪdəˈreɪʃ(ə)n /",
        "mean": "n.考虑，思考；体贴，体谅，顾及",
        "sound": "consideration.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 28,
        "img": "557a572d5e8ebaf607682bb1938dd288_197385_1586599882.jpeg",
        "word": "constraint",
        "phonetic": "/kənˈstreɪnt/",
        "mean": "n.限制；约束，束缚",
        "sound": "constraint.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 29,
        "img": "957939ffb255381263b4c2f6f79c720c.jpg",
        "word": "convince of",
        "phonetic": "/kənˈvɪns ʌv/",
        "mean": "phr.使信服",
        "sound": "convince of.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 30,
        "img": "3c33bac518f30b3a2eead70863f66062.jpeg",
        "word": "crazy about",
        "phonetic": "/ˈkrezi əˈbaʊt/",
        "mean": "phr.痴迷于…，迷恋…",
        "sound": "crazy about.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 31,
        "img": "b05b41f4261bd0079e3d7d58f0c3311c_56439_1565684272.jpg",
        "word": "crowd",
        "phonetic": "/kraʊd/",
        "mean": "n.人群，观众；民众；朋友；  v.聚集在……周围；涌入（脑海）；挤满",
        "sound": "crowd.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      },
      {
        "id": 32,
        "img": "i_13_4579_0_4_151221175044-gigapixel-scale.jpg",
        "word": "cure",
        "phonetic": "/kjʊr/",
        "mean": "v.治愈，治疗",
        "sound": "cure.mp3",
        "created_at": "2025-10-15T22:02:30+08:00"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 32,
      "total": 4056,
      "total_pages": 127
    }
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 登录

POST /api/v1/user/login

> Body 请求参数

```json
{
  "nick": "柚子皮",
  "password": "123456"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "data": {
    "email": "",
    "id": 6,
    "nick": "柚子皮",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJuaWNrIjoi5p-a5a2Q55quIiwiaXNzIjoieXpwIiwiZXhwIjoxNzk3MTUyMzQwLCJpYXQiOjE3NjU2MTYzNDB9.FbujT8l4WMGh0JTQ04WuYab0UCofOQytyL87bz8bLkQ"
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 注册

POST /api/v1/user/register

> Body 请求参数

```json
{
  "nick": "柚子皮",
  "password": "123456"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 400,
  "data": {
    "msg": "昵称已存在"
  },
  "msg": "参数错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 已登录查询生词

GET /api/v1/new-word

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 21,
        "word_id": 10,
        "nick": "柚子皮",
        "is_mastered": false,
        "img": "78e1cc3da722b51b21203c8969321d3c.jpg",
        "word": "at this rate",
        "phonetic": "/æt ðɪs reɪt/",
        "mean": "n.照这样下去",
        "sound": "at this rate.mp3",
        "created_at": "2025-12-13T16:08:12+08:00"
      },
      {
        "id": 19,
        "word_id": 11,
        "nick": "柚子皮",
        "is_mastered": true,
        "img": "0e8ef21e3ff40ada50158e9ba4c2700b_185138_1571630239.jpeg",
        "word": "attractive",
        "phonetic": "/əˈtræktɪv/",
        "mean": "adj.（事物）有吸引力的，诱人的；有魅力的，性感的；引起兴趣的",
        "sound": "attractive.mp3",
        "created_at": "2025-12-13T15:49:07+08:00"
      },
      {
        "id": 10,
        "word_id": 13,
        "nick": "柚子皮",
        "is_mastered": false,
        "img": "5ef73b3793a90788f352a20c9e0d1ba3_81099_1538624010.jpg",
        "word": "balance",
        "phonetic": "/ˈbæləns/",
        "mean": "n.平衡，平衡能力；天平；  v.平衡；保持平衡",
        "sound": "balance.mp3",
        "created_at": "2025-12-13T15:27:39+08:00"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 32,
      "total": 3,
      "total_pages": 1
    }
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 新建生词

POST /api/v1/new-word

> Body 请求参数

```json
{
  "word_id": 2
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## DELETE 删除生词

DELETE /api/v1/new-word/{wordId}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|wordId|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## PUT 掌握生词

PUT /api/v1/new-word/{wordId}/master

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|wordId|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "data": {
    "msg": "标记为已掌握"
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## PUT 取消掌握生词

PUT /api/v1/new-word/{wordId}/unmaster

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|wordId|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "data": {
    "msg": "标记为未掌握"
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 数据模型

