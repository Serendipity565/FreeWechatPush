# FreeWechatPush

使用免费的微信测试公众号，做微信通知推送

---

原理：申请免费的测试版公众号，可获得免费的模板推送功能。

## 🔎效果图

![img1](https://github.com/Serendipity565/FreeWechatPush/blob/main/img/img.jpg)

## 🚀快速开始

### 准备阶段

1. 在[微信官方测试号](https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login)中注册账号，直接使用微信扫码即可

2. 将代码框架中的 `users_sample.yaml` 改成 `users.yaml`,将 `admin_sample.yaml` 改成 `admin.yaml`

### 配置代码

注册完账号进入界面之后，需要获取四个值，分别是：appID，appsecret，openId 以及 template_id。

首先是 appID 和 appsecret ，如图，然后我们需要在代码的 `admin.yaml` 中对应天上这两个值

![img2](https://github.com/Serendipity565/FreeWechatPush/blob/main/img/img2.png)

![img5](https://github.com/Serendipity565/FreeWechatPush/blob/main/img/img5.png)

接下来是 openId ，也就是为微信号（不是日常使用的那个），想让谁收消息，就让谁用微信扫下面的二维码关注自己创建的测试号，然后对应的用户就会出现在用户列表里，会有一个昵称和微信号，这里的微信号才是我们需要用到的，也就是 openId

![img3](https://github.com/Serendipity565/FreeWechatPush/blob/main/img/img3.png)

接着我们进入 `users.yaml` 中加入用户的姓名，生日月份、日期以及对应的 openID

![img6](https://github.com/Serendipity565/FreeWechatPush/blob/main/img/img6.png)

**说明：**

1. user里的name字段不重要，可以随便写
2. 这里的月份要用英文

然后新增测试模板获得 template_id（模板ID）

![img4](https://github.com/Serendipity565/FreeWechatPush/blob/main/img/img4.png)

这是我的配置模板，经供参考，可以根据需求自己定制

```text
🌏今天：{{date.DATA}}
🏠地区：{{region.DATA}}
☀️天气：{{weather.DATA}}
🌙气温：{{temp.DATA}}℃
🏳️风向：{{wind_dir.DATA}}
✨风力：{{wind_str.DATA}}
🎂距离你的生日：{{birthday.DATA}}
😊对你说的话：{{note.DATA}}{{note1.DATA}}{{note2.DATA}}{{note3.DATA}}
```

![img7](https://github.com/Serendipity565/FreeWechatPush/blob/main/img/img7.png)

接着在 `data\weather.go` 中将 `WeatherTempleId` 改成自己的 template_id 即可

**说明：**

1. 关于城市ID，如果 `data\weather.go` 中没有自己想要的城市，只需要前往[天气网](https://www.weather.com.cn/)中搜索对应的城市，将相应的城市ID填入即可

![img7](https://github.com/Serendipity565/FreeWechatPush/blob/main/img/img8.png)
![img7](https://github.com/Serendipity565/FreeWechatPush/blob/main/img/img9.png)
