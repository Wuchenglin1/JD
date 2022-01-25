# API



# 常规说明

如无特殊说明，则返回一个一下格式的json

```json
{
    status: true, // true：成功， false：失败
    data: "" // 提示信息
}
```



# 用户属性

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



# 用户注册

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

|              | 说明                                      | 必须 |
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

# 用户登录

#### `/register/normal`

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

# 用户获取通过获取token

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

| status  | data              |                                   |
| ------- | ----------------- | --------------------------------- |
| `false` | `parseTokenError` | `refreshToken`无效                |
| `false` | `errToken`        | `refreshToken`错误                |
| `false` | `expiredToken`    | `refreshToken`已过期              |
| `true`  | `""`              | `refreshToken`正常，分发新的token |





# 商品详情页

| 请求参数        | 说明                                             | 必选 |
| --------------- | ------------------------------------------------ | ---- |
| `type`          | 商品类别                                         | 是   |
| `name`          | 商品名称                                         | 是   |
| `token`         | 用户的token                                      | 是   |
| `cover`         | 商品封面（二进制文件流）这里是单文件             | 是   |
| `describePhoto` | 商品展示的图片（二进制文件流）这里可以是多个文件 | 是   |
| `describeVideo` | 商品展示的视频（二进制文件流）这里可以是多个文件 | 是   |
| `detailPhoto`   | 商品的介绍（二进制文件流）这里可以是多个文件     | 是   |

`女士衬衫` `0520101`

| 请求参数         | 说明                                                         | 必选 |
| ---------------- | ------------------------------------------------------------ | ---- |
| `price`          | 价格                                                         | 是   |
| `brand`          | 品牌                                                         | 否   |
| `womenClothing`  | 女装                                                         | 否   |
| `size`           | 尺码                                                         | 否   |
| `color`          | 颜色                                                         | 否   |
| `version`        | 版型                                                         | 否   |
| `length`         | 衣长                                                         | 否   |
| `sleeveLength`   | 袖长                                                         | 否   |
| `suitableAge`    | 适用年龄(18-24周岁请求参数为`1`,25-29周岁返回`2`,30-34周岁返回`3`,35-39周岁返回`4`,40-49周岁返回`5`) | 否   |
| `getModel`       | 领型                                                         | 否   |
| `style`          | 风格                                                         | 否   |
| `material`       | 材质                                                         | 否   |
| `pattern`        | 图案                                                         | 否   |
| `wearingWay`     | 穿着方式                                                     | 否   |
| `popularElement` | 流行元素                                                     | 否   |
| `sleeveType`     | 袖型                                                         | 否   |
| `clothesPlacket` | 衣门襟                                                       | 否   |
| `marketTime`     | 上市时间                                                     | 否   |
| `fabric`         | 面料                                                         | 否   |
| `other`          | 其他分类                                                     | 否   |

`牛仔长裤` `0520201`

| 请求参数         | 说明     | 必选 |
| ---------------- | -------- | ---- |
| `brand`          | 品牌     | 否   |
| `size`           | 尺码     | 否   |
| `color`          | 颜色     | 否   |
| `waistType`      | 腰型     | 否   |
| `height`         | 裤长     | 否   |
| `pants`          | 裤型     | 否   |
| `thick`          | 厚度     | 否   |
| `stretch`        | 弹力     | 否   |
| `material`       | 材质     | 否   |
| `suitableAge`    | 适用年龄 | 否   |
| `markeTime`      | 上市时间 | 否   |
| `popularElement` | 流行元素 | 否   |
| `fabric`         | 面料     | 否   |
| `frontPants`     | 裤门襟   | 否   |


| 返回参数 | 说明   |
| -------- | ------ |
| `status` | 状态码 |
| `data`   | 信息   |

| `status` | `data`                   | 说明                           |
| -------- | ------------------------ | ------------------------------ |
| `false`  | `expiredToken`           | `token`过期                    |
| `false`  | `parseTokenError`        | `token`错误                    |
| `false`  | `errToken`               | `token`无效                    |
| `false`  | `类型不正确`             | `type`不能为空或不存在         |
| `false`  | `商品名称不能为空`       | `name`为空                     |
| `false`  | `商品名称太长啦`         | `name`大于30字节               |
| `false`  | `属性名称太长啦`         | 女士衬衫某一属性长度大于30字节 |
| `false`  | `价格填写不正确`         | `price`为空或为负数            |
| `false`  | `请正确填写适用年龄`     | `suitableAge`填写不规范        |
| `false`  | `cover上传失败`          | 服务器错误                     |
| `false`  | `封面文件不能为空的啦！` | `cover`大小为0                 |
| `false`  | `封面文件太大的啦`       | `cover`大小大于10mb            |
| `false`  | `商品展示图不能为空呀！` | `describePhoto`为空            |
| `false`  | `商品展示图太大啦！`     | `describePhoto`大于30mb        |
| `false`  | `商品展示视频不能为空`   | `describeVideo`为空            |
| `false`  | `商品展示视频太大啦`     | `describeVideo`大于1个g        |
| `false`  | `商品介绍不能为空`       | `detailPhoto`为空              |
| `false`  | `商品介绍图片太大`       | `detailPhoto`大于10mb          |

