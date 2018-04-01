function getDefaultDt() {
    $.get(getUrl(), function (data, status) {
        refreshData(data);
    });
}

//获取当前时间，格式YYYY-MM-DD
function getD() {
    var date = new Date();
    var year = date.getFullYear();
    var month = date.getMonth() + 1;
    var strDate = date.getDate();
    if (month >= 1 && month <= 9) {
        month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
    }
    var currentdate = year + month + strDate;
    return currentdate;
}

function getDt() {
    var date = new Date();
    var year = date.getFullYear();
    var month = date.getMonth() + 1;
    var strDate = date.getDate();
    var m = date.getMinutes();
    var h = date.getHours();
    if (month >= 1 && month <= 9) {
        month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
    }
    if (m >= 0 && m <= 9) {
        m = "0" + m;
    }
    if (h >= 0 && h <= 9) {
        h = "0" + h;
    }
    var currentdate = year + month + strDate + h + m;
    return currentdate;
}