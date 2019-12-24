# 流量统计

## 1.统计js代码

```javascript
$(document).ready(function () {

    /**
     * 上报用户信息，以及访问数据上报到打点服务器.
     */
    $.get(
        'http://mytongji.lo/dot.gif',
        {
            "time":gettime(),
            "ip":getip(),
            "url":geturl(),
            "refer":getrefer(),
            "ua" :getuser_agent(),
        }
    );

})
function gettime() {
    var nowDate = new Date();
    return nowDate.toLocaleString();
}

function geturl() {
    return window.location.href;
}

function getip() {
    return returnCitySN["cip"] + ',' + returnCitySN["cname"];
}

function getrefer() {
    return document.referrer;
}

function getcookie() {
    return document.cookie;
}

function getuser_agent() {
    return navigator.userAgent;
}
```

## 2.利用nginx的 ngx_http_empty_gif_module模块实现打点服务器

[ngx_http_empty_gif_module说明](http://nginx.org/en/docs/http/ngx_http_empty_gif_module.html)



nginx 配置：

```ng
server {
    listen 80;
    server_name mytongji.lo;
    root /Users/linzl/webroot/tongji/;
    index index.php index.html index.htm;
	
	access_log  /data/log/tongji_access.log  main;
    location = /_.gif {
         empty_gif;
         error_page 405 =200 $request_uri; #get或post请求一个静态文件会返回405，需要把他的状态码设置为200
    }
}
```



