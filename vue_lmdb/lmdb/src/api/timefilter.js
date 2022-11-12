export const timeFilter = (val) => {
    function p(t) {
        return t < 10 ? '0' + t : t;
    }

    var h = Math.floor(val / 60 / 60);
    var m = Math.floor(val / 60 % 60);
    var s = Math.floor(val % 60);
    if (s < 0) {
        return "00:00:00"
    }
    return p(h) + ':' + p(m) + ':' + p(s);
}

export const timeStrToSec = (timeStr) => {
    var hour = timeStr.split(':')[0];
    var min = timeStr.split(':')[1];
    var sec = timeStr.split(':')[2];
    return Number(hour * 3600) + Number(min * 60) + Number(sec);
}

export function getUTCTime() {
    let d1 = new Date();
    let d2 = new Date(d1.getUTCFullYear(), d1.getUTCMonth(), d1.getUTCDate(), d1.getUTCHours(), d1.getUTCMinutes(), d1.getUTCSeconds());
    return Date.parse(d2) / 1000;
}

export function getCurrentTime() {
    let date = new Date();
    let year = date.getFullYear(); //获取当前年份
    let mon = date.getMonth() + 1; //获取当前月份
    let da = date.getDate(); //获取当前日
    let h = date.getHours(); //获取小时
    let m = date.getMinutes(); //获取分钟
    let s = date.getSeconds(); //获取秒
    return `${year} ${mon} ${da} ${h}:${m}:${s}`;
}

export function percent2Point(percent) {
    let point = percent.replace("%", "");
    point = point / 100;
    return point;
}