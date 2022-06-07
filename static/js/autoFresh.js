// 定时刷新页面
var url = window.location.href;
setInterval(function() {
    window.location.href = url;
}, 3000)