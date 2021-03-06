# API



# 常规说明

如无特殊说明，则返回一个一下格式的json

```json
{
    "status": true, // true：成功， false：失败
    "data": "" // 提示信息
}
```

# 用户相关

## 用户属性

`用户的属性json格式如下`

```json
{
    "uid":"",//每个用户唯一的uid
    "name":"",//昵称
    "password":"",//密码
    "headPic":"",//用户的头像
    "phone":"",//手机号
    "email":"",//邮箱号
    "money":"",//钱
    "isBan":"",//是否被封禁
    "admin":"",//是否是管理员
}
```



## 用户注册

#### **`/verify/sms/register`** `POST`

`注册先检测手机号是否被注册过和是否能够正常通信,并发送验证码`

| 请求参数 | 说明   | 必须 |
| -------- | ------ | ---- |
| `phone`  | 手机号 | 是   |

| 返回参数 | 说明     |
| -------- | -------- |
| `status` | 状态码   |
| `data`   | 返回信息 |

| 返回参数 | status   | data             | 说明                       |
| -------- | -------- | ---------------- | -------------------------- |
|          | ` false` | `手机号不可为空` | `phone` 为空               |
|          | `false`  | `手机号不合法`   | `phone` 不满足11位长度要求 |
|          | `false`  | `手机号已被注册` | `phone` 已被使用           |
|          | `true`   | `""`             | 发送验证码成功             |

#### `/check/sms/register` `POST`

`检查注册时发送手机短信验证码是否正确`

| 请求参数     | 说明   | 必须 |
| ------------ | ------ | ---- |
| `phone`      | 手机号 | 是   |
| `verifyCode` | 验证码 | 是   |

| 返回参数 | status  | data               | 说明                        |
| -------- | ------- | ------------------ | --------------------------- |
|          | `false` | `"验证码不能为空"` | `verifyCode`为空            |
|          | `false` | `"验证码错误"`     | `phone`与`verifyCode`不符合 |
|          | `true`  | `""`               | 参数正确                    |



#### `/user/register/email` `POST`

- `application/x-www-form-urlencoded`
- 通过`email`注册

`检查所有参数是否合法` 

| 请求参数     | 说明                                      | 必须 |
| ------------ | ----------------------------------------- | ---- |
| `phone`      | 这里的`phone`参数是上面已经验证过的手机号 | 是   |
| `userName`   | 用户名/账号                               | 是   |
| `password`   | 密码                                      | 是   |
| `email`      | 用户邮箱                                  | 是   |
| `verifyCode` | 邮箱验证码                                | 是   |

| 返回参数 | status  | data                             | 说明                       |
| -------- | ------- | -------------------------------- | -------------------------- |
|          | `false` | `"用户名不能为空"`               | `userName`为空             |
|          | `false` | `"用户名太长了"`                 | `userName`长度超过20个字节 |
|          | `false` | `"该用户名已被使用，请更换名称"` | `userName`已存在           |
|          | `false` | `"密码不能小于6个字符"`          | `password`长度少于6个字节  |
|          | `false` | `"密码不能大于16个字符"`         | `password`长度超过16个字节 |
|          | `false` | `"邮箱不能为空"`                 | `email`为空                |
|          | `false` | `"邮箱格式错误"`                 | `email`格式错误            |
|          | `false` | `"邮箱已被注册"`                 | `email`已被使用            |
|          | `false` | `"请输入验证码"`                 | `verifyCode`为空           |
|          | `false` | `"未发送验证码"`                 | `email`无对应`verifyCode`  |
|          | `false` | `"验证码错误"`                   | `verifyCode`与验证码不符   |
|          | `true`  | `"注册成功！"`                   | 参数合法                   |

**`/verify/email/register`** `POST`

`发送注册时的邮箱验证码`

| 请求参数 | 说明     | 必选 |
| -------- | -------- | ---- |
| `email`  | 邮箱地址 | 是   |

| 返回参数 | status  | data             | 说明            |
| -------- | ------- | ---------------- | --------------- |
|          | `false` | `"邮箱不能为空"` | `email`为空     |
|          | `false` | `"邮箱格式错误"` | `email`格式错误 |
|          | `false` | `"邮箱已被注册"` | `email`已存在   |

## 用户登录

#### `/register/normal` `POST`

`JD的登录很简单，就直接是账号-密码式登录，所以只需要一一检对密码和邮箱或者账号是否正确就行`

| 请求参数   | 说明                 | 必选 |
| ---------- | -------------------- | ---- |
| `account`  | 邮箱/用户名/登录手机 | 是   |
| `password` | 密码                 | 是   |

| 返回参数       | 说明         |
| -------------- | ------------ |
| `status`       | 状态码       |
| `data`         | 信息         |
| `token`        | 用户token    |
| `refreshToken` | refreshToken |

| status  | data                             | 说明                          |
| ------- | -------------------------------- | ----------------------------- |
| `false` | `"请输入账户名和密码"`           | `account`和`password`同时为空 |
| `false` | `"请输入账户名"`                 | `account`为空                 |
| `false` | `"请输入密码"`                   | `password`为空                |
| `false` | `"不存在此账号"`                 | `account`未注册过             |
| `false` | `您的账号存在风险，请联系客服`   | `account`已被封号             |
| `false` | `账号不存在`                     | `account`未被注册             |
| `false` | `账户名与密码不匹配，请重新输入` | `account`和`password`对不上   |
| `true`  | `"登录成功！"`                   | 参数合法                      |

## 金钱系统

### 用户充值

#### `/user/rechargeBalance` `POST`

| 请求参数 | 说明           | 必选 |
| -------- | -------------- | ---- |
| `token`  | 用户的`token`  | 是   |
| `gold`   | 充值的金钱数量 | 是   |

| 返回参数 | status  | data              | 说明        |
| -------- | ------- | ----------------- | ----------- |
|          | `false` | `expiredToken`    | `token`过期 |
|          | `false` | `parseTokenError` | `token`错误 |
|          | `false` | `errToken`        | `token`无效 |
|          | `false` | `充值余额不正确`  | 数字不正确  |
|          | `false` | `用户不存在`      | 用户不存在  |
|          | `true`  | `充值$money$成功` | 充值成功    |

`$money$`表示一个int参数

### 用户查询余额

#### `/user/checkBalance` `GET`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |

| 返回参数 | status  | data               | 说明        |
| -------- | ------- | ------------------ | ----------- |
|          | `false` | `expiredToken`     | `token`过期 |
|          | `false` | `parseTokenError`  | `token`错误 |
|          | `false` | `errToken`         | `token`无效 |
|          | `false` | `用户不存在`       | 用户不存在  |
|          | `true`  | 如下格式json字符串 | 参数合法    |

```json
{
    "status":true,
    "data":0
}
```



## 用户获取通过获取token

#### `/token/get` `POST`

`通过refreshToken获取token` 

| 请求参数       | 说明                 | 必选 |
| -------------- | -------------------- | ---- |
| `refreshToken` | 用户的`refreshToken` | 是   |

| 返回参数 | 说明            |
| -------- | --------------- |
| `status` | 状态码          |
| `data`   | 说明            |
| `token`  | 用户的新`token` |

| status  | data              |                                             |
| ------- | ----------------- | ------------------------------------------- |
| `false` | `parseTokenError` | `refreshToken`无效                          |
| `false` | `errToken`        | `refreshToken`错误                          |
| `false` | `expiredToken`    | `refreshToken`已过期                        |
| `true`  | `""`              | `refreshToken`正常，分发新的token(如下格式) |

```json
{
    "data": "",
    "token": "",//这里放的是token
    "status": true
}
```



# 商品相关

## 插入商品

### 创建一个商品的基本信息

#### `/goods/create` `POST`



| 请求参数        | 说明                                             | 必选 |
| --------------- | ------------------------------------------------ | ---- |
| `type`          | 商品类别                                         | 是   |
| `name`          | 商品名称                                         | 是   |
| `token`         | 用户的token                                      | 是   |
| `price`         | 商品的价格                                       | 是   |
| `inventory`     | 商品库存数量                                     | 是   |
| `cover`         | 商品封面（二进制文件流）这里是单文件             | 是   |
| `describePhoto` | 商品展示的图片（二进制文件流）这里可以是多个文件 | 是   |
| `describeVideo` | 商品展示的视频（二进制文件流）这里可以是多个文件 | 否   |
| `detailPhoto`   | 商品的介绍（二进制文件流）这里可以是多个文件     | 是   |

| 返回参数 | 说明      |
| -------- | --------- |
| `gid`    | 货物的gid |
| `status` | 状态码    |
| `data`   | 信息      |

| `status` | `data`                   | 说明                    |
| -------- | ------------------------ | ----------------------- |
| `false`  | `expiredToken`           | `token`过期             |
| `false`  | `parseTokenError`        | `token`错误             |
| `false`  | `errToken`               | `token`无效             |
| `false`  | `类型不正确`             | `type`不能为空或不存在  |
| `false`  | `商品名称不能为空`       | `name`为空              |
| `false`  | `商品名称太长啦`         | `name`大于30字节        |
| `false`  | `cover上传失败`          | 服务器错误              |
| `false`  | `图片太大`               | `colorPhoto`大于5mb     |
| `false`  | `封面文件不能为空的啦！` | `cover`大小为0          |
| `false`  | `封面文件太大的啦`       | `cover`大小大于10mb     |
| `false`  | `商品展示图不能为空呀！` | `describePhoto`为空     |
| `false`  | `商品展示图太大啦！`     | `describePhoto`大于30mb |
| `false`  | `商品展示视频不能为空`   | `describeVideo`为空     |
| `false`  | `商品展示视频太大啦`     | `describeVideo`大于1个g |
| `false`  | `商品介绍不能为空`       | `detailPhoto`为空       |
| `true`   | `""`                     | 参数合法                |

### 为一个商品插入尺寸表

#### `/goods/create/size` `POST`

| 请求参数 | 说明        | 必选   |
| -------- | ----------- | ------ |
| `token`  | 用户的token | **是** |
| `size`   | 尺寸表      | **是** |
| `gid`    | 商品的`gid` | **是** |

说明：商品的`size`以这样的格式：

如有S,M,L,XL,XXL的尺码:

`S;M;L;XL;XXL;`	每个尺码中间用`;`间隔



| `status` | `data`            | 说明               |
| -------- | ----------------- | ------------------ |
| `false`  | `expiredToken`    | `token`过期        |
| `false`  | `parseTokenError` | `token`错误        |
| `false`  | `errToken`        | `token`无效        |
| `false`  | `商品不存在`      | `fGid`不存在或错误 |
| `true`   | `""`              | 参数合法           |



### 为一个商品插入颜色表

#### `/goods/photo/color` `POST`

| 请求参数     | 说明                   | 必选 |
| ------------ | ---------------------- | ---- |
| `token`      | 用户的token            | 是   |
| `gid`        | 商品的`fGid`           | 是   |
| `color`      | 颜色描述               | 是   |
| `colorPhoto` | 颜色图片(二进制文件流) | 是   |

| 返回参数 | 说明   |
| -------- | ------ |
| `status` | 状态码 |
| `data`   | 信息   |

| `status` | `data`            | 说明                  |
| -------- | ----------------- | --------------------- |
| `false`  | `expiredToken`    | `token`过期           |
| `false`  | `parseTokenError` | `token`错误           |
| `false`  | `errToken`        | `token`无效           |
| `false`  | `商品不存在`      | `fGid`不存在或错误    |
| `false`  | `描述太长`        | `color`描述大于15字节 |
| `false`  | `图片不能为空`    | `colorPhoto`为空      |
| `false`  | `图片太大`        | `colorPhoto`大于5mb   |
| `true`   | `""`              | 参数合法              |



### 为一个商品的基本信息插入介绍

插入一条女士衬衫的介绍

#### `/goods/blouse` `POST`

`女士衬衫` `0520101`

| 请求参数         | 说明                                                         | 必选   |
| ---------------- | ------------------------------------------------------------ | ------ |
| `gid`            | 商品的`gid`                                                  | **是** |
| `brand`          | 品牌                                                         | 否     |
| `womenClothing`  | 女装                                                         | 否     |
| `version`        | 版型                                                         | 否     |
| `length`         | 衣长                                                         | 否     |
| `sleeveLength`   | 袖长                                                         | 否     |
| `suitableAge`    | 适用年龄(18-24周岁请求参数为`1`,25-29周岁返回`2`,30-34周岁返回`3`,35-39周岁返回`4`,40-49周岁返回`5`) | 否     |
| `getModel`       | 领型                                                         | 否     |
| `style`          | 风格                                                         | 否     |
| `material`       | 材质                                                         | 否     |
| `pattern`        | 图案                                                         | 否     |
| `wearingWay`     | 穿着方式                                                     | 否     |
| `popularElement` | 流行元素                                                     | 否     |
| `sleeveType`     | 袖型                                                         | 否     |
| `clothesPlacket` | 衣门襟                                                       | 否     |
| `marketTime`     | 上市时间                                                     | 否     |
| `fabric`         | 面料                                                         | 否     |
| `other`          | 其他分类                                                     | 否     |

`牛仔长裤` `0520201`

| 请求参数         | 说明     | 必选   |
| ---------------- | -------- | ------ |
| `brand`          | 品牌     | 否     |
| `size`           | 尺码     | **是** |
| `waistType`      | 腰型     | 否     |
| `height`         | 裤长     | 否     |
| `pants`          | 裤型     | 否     |
| `thick`          | 厚度     | 否     |
| `stretch`        | 弹力     | 否     |
| `material`       | 材质     | 否     |
| `suitableAge`    | 适用年龄 | 否     |
| `markeTime`      | 上市时间 | 否     |
| `popularElement` | 流行元素 | 否     |
| `fabric`         | 面料     | 否     |
| `frontPants`     | 裤门襟   | 否     |


| 返回参数 | 说明   |
| -------- | ------ |
| `status` | 状态码 |
| `data`   | 信息   |

| `status` | `data`               | 说明                           |
| -------- | -------------------- | ------------------------------ |
| `false`  | `expiredToken`       | `token`过期                    |
| `false`  | `parseTokenError`    | `token`错误                    |
| `false`  | `errToken`           | `token`无效                    |
| `false`  | `属性名称太长啦`     | 女士衬衫某一属性长度大于30字节 |
| `false`  | `请正确填写适用年龄` | `suitableAge`填写不规范        |
| `false`  | `商品错误！`         | `fGid`不存在或错误             |
| `true`   | `""`                 | 参数合法                       |

## 获取商品信息

### 浏览所有商品

#### `/goods/browse` `GET`

| 请求参数      | 说明                       |
| ------------- | -------------------------- |
| `arrangement` | 排序方式，具体如下文字所示 |

综合排序↓`0`，↑`1` ，

销量排序↓`2`，↑`3`，

评论数排序↓`4`，↑`5`，

新品排序↓`6`,↑`7`，

价格排序↓`8`，↑`9`





| `status` | `data`       | 说明                   |
| -------- | ------------ | ---------------------- |
| `false`  | `服务器错误` | 服务器某一项出现了问题 |
| `true`   | 如下图所示   | 参数正常               |

返回如下格式的json

```json
{
    "status":"true",
    "data":{
        "gId":"",//商品的id
        "cover":"",//商品的封面，这里是一串url
        "price":"",//价格
        "name":"",//商品名称
        "commentAccount":"",//评论数量
        "ownerName":"",//商家名称
    }
},{
    ...
}
```





### 获取商品的信息

#### `/goods/getInfo` `GET`

获取商品的基本信息

| 请求参数 | 说明        | 必选 |
| -------- | ----------- | ---- |
| `gid`    | 商品的`gid` | 是   |

| 返回参数  | 说明                 |
| --------- | -------------------- |
| `status`  | 状态码               |
| `data`    | 信息                 |
| goodsInfo | 一个如下的json字符串 |

```json
{
    "status":true,
    "data":[
        {
        "type":0,//商品类别
        "name":"",//商品名称
        "gid":0,//商品gid
        "price":0,//价格
        "ownerUid":0,//商家uid
        "ownerName":"",//店铺商家名字
        "saleTime":"",//上架时间
        "volume":0,//成交量
        "favorableRating":0,//好评度
        "cover":"",//封面，这里是一串url
    },
    "",//describePhoto 图片（商品展示的图片url）
    "",//detailPhoto 图片（商品详情图片的url）
    ]
}
```



| status  | data         | 说明            |
| ------- | ------------ | --------------- |
| `false` | `商品不存在` | 商品`gid`不存在 |
| `true`  | `""`         | 商品存在        |



### 获取商品的size

#### `/goods/getSize` `Get`

| 请求参数 | 说明      | 必选 |
| -------- | --------- | ---- |
| `gid`    | 商品`gid` | 是   |



| 返回参数 | 说明                       |
| -------- | -------------------------- |
| `status` | 状态码                     |
| `data`   | 信息                       |
| `size`   | 返回一个以下json结构的size |

```json
//假如有"S","M","L","XL","XXL"的尺码，返回如下所示格式的json字符串
{
    "size": "S"
}{
    "size": "M"
}{
    "size": "L"
}{
    "size": "XL"
}{
    "size": "XXL"
}
```



| status  | data           | 必选        |
| ------- | -------------- | ----------- |
| `false` | `该商品不存在` | `gid`不存在 |
| `true`  | `""`           | `size`存在  |

### 获取商品的颜色

#### `/goods/getColor` `GET`

| 请求参数 | 说明      | 必选 |
| -------- | --------- | ---- |
| `gid`    | 商品`gid` | 是   |



| 返回参数 | 说明                         |
| -------- | ---------------------------- |
| `status` | 状态码                       |
| `data`   | 信息                         |
| `color`  | 返回一个以下json格式的字符串 |
| `url`    | 返回一个以下json格式的字符串 |

```json
{
    "color": "绿色",
    "url": ""
}{
    "color": "红色",
    "url": ""
}{
    ...
}
```



| status  | data         | 说明        |
| ------- | ------------ | ----------- |
| `false` | `商品不存在` | `gid`不存在 |
| `true`  | `""`         | `color`存在 |

## 查询商品

### 按类别查询商品

#### `/goods/browse/type` `GET`

商品的类别详情查询channal文档

| 请求参数 | 说明       | 必选 |
| -------- | ---------- | ---- |
| `type`   | 商品的类别 | 是   |

| 返回参数 | 说明                     |
| -------- | ------------------------ |
| `status` | 状态码                   |
| `data`   | 返回信息                 |
| goods    | 如下格式商品的json字符串 |

```json
{
    "gId":"",
    "cover":"",//封面，这里是一串url
    "price":0,
    "name":"",//商品名称
    "commentAccount":0,//商品评论数量
    "ownerName":"",//商家名称
}{
    ...
}
```



| status  | data         | 说明             |
| ------- | ------------ | ---------------- |
| `false` | `类型错误`   | `type`参数有误   |
| `false` | `类型不存在` | 请求的type不存在 |

### 按关键词，在全部商品里面搜索商品

#### `/goods/browse/all` `POST`

| 请求参数   | 说明   | 必选 |
| ---------- | ------ | ---- |
| `keyWords` | 关键词 | 是   |

| 返回参数    | 说明                 |
| ----------- | -------------------- |
| `status`    | 状态码               |
| `data`      | 信息                 |
| `goodsInfo` | 如下格式的json字符串 |



```json
{
    "status":true,
    "data":{
    "gId":0,//商品id
    "cover":"",//封面url
    "price":0,//价格
    "name":"",//商品名称
    "commentAccount":0,//评论数量
    "ownerName":"",//商家名字
    }
}{
 	...
}
```



| status  | data               | 说明                                   |
| ------- | ------------------ | -------------------------------------- |
| `false` | `没有找到相关商品` | 没有商品拥有`keyWords`                 |
| `true`  | `""`               | 找到相关商品，返回如上格式的json字符串 |



## 将一件物品加入购物车

#### `/goods/add/shoppingCart` `POST`

| 请求参数  | 说明        | 必选 |
| --------- | ----------- | ---- |
| `token`   | 用户的token | 是   |
| `gid`     | 商品的`gid` | 是   |
| `account` | 商品的数量  | 是   |
| `color`   | 商品的颜色  | 否   |
| `size`    | 商品的尺寸  | 否   |
| `style`   | 商品的款式  | 否   |



| 返回参数 | status  | data              | 说明                  |
| -------- | ------- | ----------------- | --------------------- |
|          | `false` | `expiredToken`    | `token`过期           |
|          | `false` | `parseTokenError` | `token`错误           |
|          | `false` | `errToken`        | `token`无效           |
|          | `false` | `商品不存在`      | `gid`错误或商品不存在 |
|          | `false` | `商品数量有误`    | `account`错误         |
|          | `false` | `商品id有误`      | `gid`有误             |
|          | `true`  | `""`              | 参数合法，加入成功    |

## 查看购物车

#### `/user/shoppingCart` `GET`



| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |

| 返回参数 | status  | data                   | 说明           |
| -------- | ------- | ---------------------- | -------------- |
|          | `false` | `expiredToken`         | `token`过期    |
|          | `false` | `parseTokenError`      | `token`错误    |
|          | `false` | `errToken`             | `token`无效    |
|          | `false` | `购物车还是空空如也呀` | 购物车没有东西 |
|          | `true`  | `""`                   | 参数合法       |

若`status`为`true`,返回以下格式的json字符串

如果在`color`或`size`或`style`某一栏没有选择，就直接留空 

```json
{
    "data": {
        "uId": 0,
        "gid": 0,
        "goodsName": "",
        "cover":"",//商品的封面，一串url
        "color": "",
        "size": "",
        "style": "",
        "price": 0,
        "account": 0,
        "totalPrice": 0
    },
    "status": true
}{
   ...
}{
    "totalPrice": 1600
}
```

## 将某物品从购物车中移除

#### `/goods/delete/shoppingCart` `DELETE`

| 请求参数 | 说明                | 必选 |
| -------- | ------------------- | ---- |
| `token`  | 用户的`token`       | 是   |
| `gid`    | 购物车中物品的`gid` | 是   |

| status  | data                 | 说明                    |
| ------- | -------------------- | ----------------------- |
| `false` | `expiredToken`       | `token`过期             |
| `false` | `parseTokenError`    | `token`错误             |
| `false` | `errToken`           | `token`无效             |
| `false` | `gid不正确`          | `gid`格式错误           |
| `false` | `您还没有关注该商品` | 用户还没有关注`gid`商品 |
| `true`  | `删除成功`           | 删除成功                |





## 关注商品

#### `/goods/focus` `POST`

通过插入商品的uid来关注商品

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |
| `gid`    | 商品的`gid`   | 是   |



| status  | data                                 | 说明                       |
| ------- | ------------------------------------ | -------------------------- |
| `false` | `expiredToken`                       | `token`过期                |
| `false` | `parseTokenError`                    | `token`错误                |
| `false` | `errToken`                           | `token`无效                |
| `false` | `gid不正确`                          | `gid`格式错误              |
| `false` | `您已经关注过该商品啦！请勿重复关注` | 该用户已经关注过这个商品了 |
| `true`  | `关注成功`                           | 合法                       |



## 获取关注商品的信息

#### `/goods/getFocus` `GET`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |

| 返回参数 | 说明   |
| -------- | ------ |
| `status` | 状态码 |
| `data`   | 信息   |

| status  | data                         | 说明                 |
| ------- | ---------------------------- | -------------------- |
| `false` | `expiredToken`               | `token`过期          |
| `false` | `parseTokenError`            | `token`错误          |
| `false` | `errToken`                   | `token`无效          |
| `false` | `您还没有关注商品喔`         | 用户的喜爱的商品为空 |
| `true`  | 返回下面格式的一段json字符串 |                      |

```json
{
    "status":true,
    "data":{
        "gId":0,//商品的gid
        "name":"",//商品的名称
        "price":0,//商品的价格
        "cover":"",//商品封面的url
        "commentAccount":0,//商品的评论数量
        "favorableRating":0,//商品的好评率
    }
}
```

## 删除关注的商品

#### `/goods/delete/focus` `DELETE`

| 请求参数 | 说明                      | 必选 |
| -------- | ------------------------- | ---- |
| `token`  | 用户的`token`             | 是   |
| `gid`    | 用户想删除关注商品的`gid` | 是   |



| status  | data                 | 说明                      |
| ------- | -------------------- | ------------------------- |
| `false` | `expiredToken`       | `token`过期               |
| `false` | `parseTokenError`    | `token`错误               |
| `false` | `errToken`           | `token`无效               |
| `false` | `商品id有误`         | `gid`有误                 |
| `false` | `您还没有关注该商品` | 用户没有关注为`gid`的商品 |
| `true`  | `""`                 | 删除成功                  |



## 结算购物车

### 新增一个收货人信息

#### `/order/createConsigneeInfo` `POST`

| 请求参数         | 说明                   | 必选 |
| ---------------- | ---------------------- | ---- |
| `token`          | 用户的`token`          | 是   |
| `name`           | 收件人的名字           | 是   |
| `province`       | 所在省(直辖市和市相同) | 是   |
| `city`           | 所在市                 | 是   |
| `area`           | 所在区                 | 是   |
| `town`           | 所在镇                 | 是   |
| `detailAddress`  | 详细地址               | 是   |
| `phone`          | 收件人电话             | 是   |
| `fixedTelephone` | 固定电话               | 否   |
| `email`          | 邮箱地址               | 否   |
| `addressAlias`   | 地址别名               | 否   |

| 返回参数 | status  | data              | 说明               |
| -------- | ------- | ----------------- | ------------------ |
|          | `false` | `expiredToken`    | `token`过期        |
|          | `false` | `parseTokenError` | `token`错误        |
|          | `false` | `errToken`        | `token`无效        |
|          | `false` | `参数不能为空`    | 必选参数某一项为空 |
|          | `true`  | `添加成功`        | 添加成功           |



### 获取收货人信息

#### `/order/getConsigneeInfo` `GET`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |

| 返回参数 | status  | data                         | 说明                     |
| -------- | ------- | ---------------------------- | ------------------------ |
|          | `false` | `expiredToken`               | `token`过期              |
|          | `false` | `parseTokenError`            | `token`错误              |
|          | `false` | `errToken`                   | `token`无效              |
|          | `false` | `您还没有添加过收件人信息`   | 没有找到创建过收件人信息 |
|          | 无      | 返回一个以下格式的json字符串 |                          |

```json
{
   		 	"uid":0,//创建该收货人的uid
			"cid":0,//收货人的编号
			"name":"",
			"address": "",//直接是地址 省+""+市+""+区+""+镇
			"detailAddress":  "",
			"phone":          "",
			"fixedTelephone": "",
			"email":          "",
			"addressAlias":   "",
}{
    ...
}
```



### 删除一个收货人信息

#### `/order/DeleteConsigneeInfo` `DELETE`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |
| `cid`    | 收货人的编号  | 是   |

| 返回参数 | status  | data              | 说明         |
| -------- | ------- | ----------------- | ------------ |
|          | `false` | `expiredToken`    | `token`过期  |
|          | `false` | `parseTokenError` | `token`错误  |
|          | `false` | `errToken`        | `token`无效  |
|          | `false` | `该收货人不存在`  | 收货人不存在 |
|          | `true`  | `删除成功`        | 删除成功     |



### 创建订单

#### `/goods/settlement` `POST`

| 请求参数     | 说明                                      | 必选 |
| ------------ | ----------------------------------------- | ---- |
| `token`      | 用户的`token`                             | 是   |
| `settlement` | 结算商品(格式如下)                        | 是   |
| `address`    | 用户的住址                                | 是   |
| `phone`      | 联系人电话                                | 是   |
| `name`       | 联系人名字                                | 是   |
| `payWay`     | 支付方式，目前只有在线支付，只发送数字`1` | 是   |

```json
{
  "token":"",
  "settlement":{
      "gid":0,
      "account":0,
      "color":"",
      "size":"",
  },
  "address":"",
  "phone":"",
  "name":"",
  "payWay":""
}
```

| 返回参数 | status  | data                 | 说明                     |
| -------- | ------- | -------------------- | ------------------------ |
|          | `false` | `expiredToken`       | `token`过期              |
|          | `false` | `parseTokenError`    | `token`错误              |
|          | `false` | `errToken`           | `token`无效              |
|          | `false` | `您还没有选择商品`   | 购物车没有东西           |
|          | `false` | `地址不能为空`       | `address`为空            |
|          | `false` | `电话不能为空`       | `phone`为空              |
|          | `false` | `联系人名字不能为空` | `name`为空               |
|          | `false` | `支付方式不正确`     | `payWay`不正确           |
|          | `false` | `商品格式不正确`     | `settlement`不正确       |
|          | `false` | `商品的库存不足`     |                          |
|          | `true`  | `""`                 | 返回一个如下格式的订单号 |

```json
{
    "status":true,
    "data":"",//这里是订单号
}
```



### 支付订单

#### `/order/pay` `POST`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |
| `order`  | 订单号        | 是   |

| 返回参数 | status  | data                  | 说明          |
| -------- | ------- | --------------------- | ------------- |
|          | `false` | `expiredToken`        | `token`过期   |
|          | `false` | `parseTokenError`     | `token`错误   |
|          | `false` | `errToken`            | `token`无效   |
|          | `false` | `订单号不能为空`      | `order`为空   |
|          | `false` | `该订单不存在`        | `order`不存在 |
|          | `false` | `您的余额不足,请充值` | 用户余额不足  |
|          | `true`  | `"支付成功"`          | 支付成功      |



### 取消订单

#### `/order/cancel` `PUT`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |
| `order`  | 订单号        | 是   |

| 返回参数 | status  | data              | 说明          |
| -------- | ------- | ----------------- | ------------- |
|          | `false` | `expiredToken`    | `token`过期   |
|          | `false` | `parseTokenError` | `token`错误   |
|          | `false` | `errToken`        | `token`无效   |
|          | `false` | `订单号不能为空`  | `order`为空   |
|          | `false` | `该订单不存在`    | `order`不存在 |
|          | `true`  | `取消成功`        | 参数合法      |

### 删除订单

#### `/order/delete` `DELETE`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |
| `order`  | 订单号        | 是   |

| 返回参数 | status  | data              | 说明         |
| -------- | ------- | ----------------- | ------------ |
|          | `false` | `expiredToken`    | `token`过期  |
|          | `false` | `parseTokenError` | `token`错误  |
|          | `false` | `errToken`        | `token`无效  |
|          | `false` | `该订单不存在`    | 订单不存在   |
|          | `true`  | `删除成功`        | 订单删除成功 |



### 确认收货

#### `/order/confirm` `PUT`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |
| `order`  | 订单号        | 是   |

| 返回参数 | status  | data              | 说明        |
| -------- | ------- | ----------------- | ----------- |
|          | `false` | `expiredToken`    | `token`过期 |
|          | `false` | `parseTokenError` | `token`错误 |
|          | `false` | `errToken`        | `token`无效 |
|          | `false` | `该订单不存在`    | 订单不存在  |
|          | `true`  | `""`              |             |



## 查看订单

### 查看所有订单

#### `/order/checkAll` `GET`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |

| 返回参数 | status  | data                     | 说明        |
| -------- | ------- | ------------------------ | ----------- |
|          | `false` | `expiredToken`           | `token`过期 |
|          | `false` | `parseTokenError`        | `token`错误 |
|          | `false` | `errToken`               | `token`无效 |
|          | 无      | 返回以下格式的json字符串 | 参数合法    |



```json
{
    "uid": 0,//用户的id
    "orderNumber": "",//订单号
    "consignee": "",//收货人
    "address": "",//地址
    "phone": "",//收货人手机号
    "payWay": "",//支付方式
    "totalPrice": 0,//总价
    "status": "",//订单状态
    "time": "2022-02-09T00:36:49.5132377+08:00",//创建订单的时间
    "settlement": [
        {
            "gid": 0,//商品的gid
            "name": "",//商品名称
            "cover": "",//商品封面的url
            "price": 0,//商品的价格
            "account": 0,//商品的数量
            "color": "",//商品的颜色
            "size": ""//商品的尺寸
        },
        {
            ...
        }
    ]
}{
 			...   
}
```



### 查看指定订单

#### `/order/checkSpecified` `GET`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |
| `order`  | 订单号        | 是   |

| 返回参数                             | 说明   |
| ------------------------------------ | ------ |
| `status`                             | 状态码 |
| `data`                               | 信息   |
| 如果无status，返回以下格式的订单信息 |        |

| status  | data                 | 说明           |
| ------- | -------------------- | -------------- |
| `false` | `expiredToken`       | `token`过期    |
| `false` | `parseTokenError`    | `token`错误    |
| `false` | `errToken`           | `token`无效    |
| `false` | `订单号不能为空`     | 订单号格式错误 |
| `false` | `该订单不存在`       | 订单不存在     |
| 无      | 以下json格式的字符串 |                |

```json
{
    "uid": 0,
    "orderNumber": "",
    "consignee": "",
    "address": "",
    "phone": "",
    "payWay": "",
    "totalPrice": 0,
    "status": "",
    "time": "2022-02-09T01:05:39.6901257+08:00",
    "settlement": [
        {
            "gid": 0,
            "name": "",
            "cover": "",
            "price": 0,
            "account": 0,
            "color": "",
            "size": ""
        },
        {
				...
        }
    ]
}
```



### 按不同的状态查看订单

#### `/order/checkByStatus` `GET`

| 请求参数 | 说明                                                         | 必选 |
| -------- | ------------------------------------------------------------ | ---- |
| `token`  | 用户的`token`                                                | 是   |
| `status` | 数字，`1`查看"待付款",`2`查看"待收货",`3`查看"已完成",`4`查看"已取消" | 是   |

| 返回参数 | status  | data              | 说明               |
| -------- | ------- | ----------------- | ------------------ |
|          | `false` | `expiredToken`    | `token`过期        |
|          | `false` | `parseTokenError` | `token`错误        |
|          | `false` | `errToken`        | `token`无效        |
|          | `false` | `status错误`      | 发送的`status`有误 |
|          | `false` | `该订单不存在`    | 订单不存在         |
|          |         |                   |                    |

```json
{
    "uid": 0,//用户的id
    "orderNumber": "",//订单号
    "consignee": "",//收货人
    "address": "",//地址
    "phone": "",//收货人手机号
    "payWay": "",//支付方式
    "totalPrice": 0,//总价
    "status": "",//订单状态
    "time": "2022-02-09T00:36:49.5132377+08:00",//创建订单的时间
    "settlement": [
        {
            "gid": 0,//商品的gid
            "name": "",//商品名称
            "cover": "",//商品封面的url
            "price": 0,//商品的价格
            "account": 0,//商品的数量
            "color": "",//商品的颜色
            "size": ""//商品的尺寸
        },
        {
            ...
        }
    ]
}{
 			...   
}
```



## 用户评论

### 评论某个商品

#### `/comment/add` `POST`

| 请求参数      | 说明                              | 必选 |
| ------------- | --------------------------------- | ---- |
| `token`       | 用户的`token`                     | 是   |
| `gid`         | 要评论商品的`gid`                 | 是   |
| `comment`     | 评价晒单(限制500字)               | 是   |
| `grade`       | 商品评分`1`、`2`、`3`、`4`、`5`分 | 是   |
| `isAnonymous` | 是否匿名:`1`为匿名,`2`为不匿名    | 是   |
| `video`       | 评论的视频(限制只能上传一个视频)  | 否   |
| `photo`       | 评论的图片(可以上传多个图片)      | 否   |

| status  | data              | 说明              |
| ------- | ----------------- | ----------------- |
| `false` | `expiredToken`    | `token`过期       |
| `false` | `parseTokenError` | `token`错误       |
| `false` | `errToken`        | `token`无效       |
| `false` | `gid有误`         | `gid`有误         |
| `false` | `是否匿名`        | `isAnonymous`有误 |
| `false` | `评分有误`        | `grade`有误       |
| `false` | `上传的视频有误`  | `video`文件有误   |
| `false` | `视频文件太大`    | `video`文件太大   |
| `false` | `图片太大啦`      | `photo`文件太大   |
| `true`  | `""`              |                   |



### 查询商品下的评论

#### `/comment/viewComment` `GET`

| 请求参数 | 说明        | 必选 |
| -------- | ----------- | ---- |
| `gid`    | 商品的`gid` | 是   |

| status  | data               | 说明 |
| ------- | ------------------ | ---- |
| `false` | `gid错误`          |      |
| `false` | `没有查询到评论`   |      |
| `true`  | 返回以下格式的data |      |

```json
{
    "status":true,
    "data":{
        "commentId":0,
        "gid":0,
        "uid":0,
        "comment":"",//评论的内容
        "grade":0,//评分
        "isAnonymous":false,//是否匿名
        "time":时间的默认格式,
        "name":"",
        "photoUrl":[
            "0":"",
            "1":""//对应的url
        ],
        "videoUrl":url,
    },{
    ...
}
```



### 回复某条评论

#### `/comment/reply` `POST`

| 请求参数    | 说明              | 必选 |
| ----------- | ----------------- | ---- |
| `token`     | 用户的`token`     | 是   |
| `gid`       | 商品的`gid`       | 是   |
| `commentId` | 评论的`commentId` | 是   |
| `comment`   | 回复的内容        | 是   |

`这里不带文件`

| status  | data              | 说明                    |
| ------- | ----------------- | ----------------------- |
| `false` | `expiredToken`    | `token`过期             |
| `false` | `parseTokenError` | `token`错误             |
| `false` | `errToken`        | `token`无效             |
| `false` | `评论id有误`      | `commentId`有误         |
| `false` | `评论不能为空`    | `comment`为空           |
| `false` | `gid有误`         | `gid`有误               |
| `false` | `没有找到该评论`  | `commentId`的评论不存在 |
| `true`  | `""`              |                         |



### 查看某条评论下的回复

#### `/comment/viewComment/specific` `GET`

| 请求参数    | 说明              | 必选 |
| ----------- | ----------------- | ---- |
| `commentId` | 评论的`commentId` | 是   |

| status  | data                     | 说明              |
| ------- | ------------------------ | ----------------- |
| `false` | `commentId有误`          | `commentId`有误   |
| `false` | `没有找到该评论`         | `commentId`不存在 |
| `false` | `没有多余的评论啦`       | 该评论没有回复    |
| `true`  | 返回以下格式的json字符串 |                   |

```json
{
    "status":true,
    "data":{
        "commentId":0,
        "sComment":[
           "0":{
            "commentId":0,
            "uid":0,
            "comment":"",
            "grade":0,
            "isAnonymous":false,
            "time":"时间的标准格式",
            },{
    ...
}
        ]
    }
}
```



## 创建一个商品(另辟蹊径)

#### `/goods/createGoods` `POST`

| 请求参数        | 说明                                    | 必选 |
| --------------- | --------------------------------------- | ---- |
| `token`         | 用户的token                             | 是   |
| `type`          | 商品类别                                | 是   |
| `name`          | 商品名称                                | 是   |
| `price`         | 商品的价格                              | 是   |
| `inventory`     | 商品库存数量                            | 是   |
| `cover`         | 商品封面（URL）这里是单图片             | 是   |
| `describePhoto` | 商品展示的图片（URL）这里可以是多个图片 | 是   |
| `detailPhoto`   | 商品的介绍（URL）这里可以是多个图片     | 是   |

| 返回参数 | status  | data                     | 说明            |
| -------- | ------- | ------------------------ | --------------- |
|          | `false` | `expiredToken`           | `token`过期     |
|          | `false` | `parseTokenError`        | `token`错误     |
|          | `false` | `errToken`               | `token`无效     |
|          | `false` | `type有误`               | `type`有误      |
|          | `false` | `价格有误`               | `price`有误     |
|          | `false` | `商品数量有误`           | `inventory`有误 |
|          | `true`  | 返回以下格式的json字符串 |                 |

```json
{
    "status":true,
    "data":0,//商品的gid
}
```

## 店铺相关

### 查看店铺的商品

#### `/store/getGoods` `GET`

| 请求参数 | 说明                        | 必选 |
| -------- | --------------------------- | ---- |
| `token`  | 用户的`token`               | 是   |
| `way`    | 查看方式，`1`按时间查看,`2` | 否   |

综合排序↓`0`，↑`1` ，

销量排序↓`2`，↑`3`，

评论数排序↓`4`，↑`5`，

新品排序↓`6`,↑`7`，

价格排序↓`8`，↑`9`

| 返回参数 | 说明     |
| -------- | -------- |
| `status` | 状态码   |
| `data`   | 返回信息 |

| status  | data                 | 说明           |
| ------- | -------------------- | -------------- |
| `false` | `expiredToken`       | `token`过期    |
| `false` | `parseTokenError`    | `token`错误    |
| `false` | `errToken`           | `token`无效    |
| `false` | `您筛选的方式有误`   | `way`错误      |
| `false` | `没有查询到相关商品` | 没有查询到商品 |
| `true`  | 下列各式的json字符串 |                |

```json
{
    "status":true,
    "data":[{
        "gId":0,
        "cover":"",
        "price":0,
        "name":"",
        "commentAccount":0,
        "ownerName":"",
    },{
        ...
    }]
}
```

### 发布(修改)店铺的公告

#### `/store/postAnnouncement` `PUT`

| 请求参数       | 说明          | 必选 |
| -------------- | ------------- | ---- |
| `token`        | 用户的`token` | 是   |
| `announcement` | 店铺的公告    | 是   |

返回参数

| status  | data              | 说明                           |
| ------- | ----------------- | ------------------------------ |
| `false` | `expiredToken`    | `token`过期                    |
| `false` | `parseTokenError` | `token`错误                    |
| `false` | `errToken`        | `token`无效                    |
| `false` | `公告不能为空`    | `announcement`为空             |
| `false` | `公告太长了`      | `announcement`长度大于45个字符 |
| `true`  | `""`              | 发布成功                       |



### 获取公告

#### `/store/getAnnouncement` `GET`

| 请求参数 | 说明          | 必选 |
| -------- | ------------- | ---- |
| `token`  | 用户的`token` | 是   |

| status  | data              | 说明                       |
| ------- | ----------------- | -------------------------- |
| `false` | `expiredToken`    | `token`过期                |
| `false` | `parseTokenError` | `token`错误                |
| `false` | `errToken`        | `token`无效                |
| `true`  |                   | 返回以下格式的`json`字符串 |

```json
{
	"status":true,
    "data":""
}
```

