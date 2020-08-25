# onedirve-sdk-golang

## 快速入门

### 申请应用

在调用 Onedrive API 时需要使用到 4 个参数

- client_id: 客户端ID
- client_secret: 客户端密码
- redirect_uri: 重定向地址
- scope: API权限

1. 访问并登录MicroSoft Azure：<https://portal.azure.com/#blade/Microsoft_AAD_RegisteredApps/ApplicationsListBlade>
2. 点击新注册，并填写注册信息，其中的重定向 URI 作为 redirect_uri 参数

    ![README-2020-08-23-04-35-26](http://img.ssrcoder.com/images/picgo/README-2020-08-23-04-35-26.png)

3. 主页上的 应用程序(客户端) ID 作为 client_id

    ![README-2020-08-23-04-39-25](http://img.ssrcoder.com/images/picgo/README-2020-08-23-04-39-25.png)

4. 在 证书和密码 页面新建客户端密码，填写相关信息，生成的值作为 client_secret
   
   ![README-2020-08-23-04-40-36](http://img.ssrcoder.com/images/picgo/README-2020-08-23-04-40-36.png)

   ![README-2020-08-23-04-41-01](http://img.ssrcoder.com/images/picgo/README-2020-08-23-04-41-01.png)

5. 在 API 权限页面选择需要的权限，添加权限，Microsoft Graph，委托的权限，勾选 Files.ReadWrite.All，这将作为 scope 传递

    ![README-2020-08-23-04-46-15](http://img.ssrcoder.com/images/picgo/README-2020-08-23-04-46-15.png)

### 访问用户信息

参考 examples/login/main.go

获取用户信息只需要使用 User.Read 权限，所以我们在scopes里只添加了 User.Read

## 计划API调用方法

1. client := NewClient()
2. fs := NewFileSystem(client)
3. fs.Cache(true)
4. fs.SetExpire(0,0,0,0,10,0) // 年,月,日,时,分,秒
5. root := fs.Root()
6. dir := fs.OpenDir("/foo")
7. fs.Refresh()
8. children := dir.Children()
9. file := children[0]
10. file.IsFile()
11. file.IsDir()
12. file.Write()
13. file.WriteAppend().WriteAppend()
14. file2 := root.Upload("")
15. dir.Copy("/")
