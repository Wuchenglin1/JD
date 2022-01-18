# API



# 常规说明

如无特殊说明，则返回一个一下格式的json

```json
{
    status: true, // true：成功， false：失败
    data: "" // 提示信息
}
```



# 用户注册

`说明：用户先输入手机号，发送验证码验证手机后，再填写账号信息`

![image-20220118183346565](C:\Users\13366\AppData\Roaming\Typora\typora-user-images\image-20220118183346565.png)

`流程如上图`



##### `/user/register/phone` `POST`

- `application/x-www-form-urlencoded`
- 通过手机注册

`调用接口发送手机验证码，并检查用户名及其密码的规范性，符合规范之后再发送验证码` `<-看你是否拥有手机`

| 请求参数   | 说明        | 必须 |
| ---------- | ----------- | ---- |
| userName   | 用户名/账号 | 是   |
| password   | 密码        | 是   |
| phone      | 用户手机号  | 是   |
| verifyCode | 验证码      | 是   |

| 返回参数 | 说明     |
| -------- | -------- |
| status   | 状态码   |
| data     | 返回信息 |

| status  | data                   | 说明                       |
| ------- | ---------------------- | -------------------------- |
| `false` | `用户名不能为空`       | `userName`为空             |
| `false` | `用户名太长了`         | `userName`长度超过20个字节 |
| `false` | `密码不能小于6个字符`  | `password`长度少于6个字节  |
| `false` | `密码不能大于16个字符` | `password`长度超过16个字节 |
| `false` | `手机号不能为空`       | `phone`为空                |
| `false` | `该手机号填写错误`     | `phone`长度不为11位        |
| `false` | `手机号已被注册`       | `phone`已存在              |
| `false` | `请输入验证码`         | `verifyCode`为空           |
| `false` | `未发送验证码`         | `verifyCode`无对应验证码   |
| `false` | `验证码错误`           | `verifyCode`与验证码不符   |
| `true`  | `注册成功！`           | 参数合法                   |

