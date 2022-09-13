export const timeFilter = (val) => {
    function p(t) {
        return t < 10 ? '0' + t : t;
    }

    var h = Math.floor(val / 60 / 60);
    var m = Math.floor(val / 60 % 60);
    var s = Math.floor(val % 60);
    if (h === 0) {
        return p(m) + ':' + p(s)
    } else {
        return p(h) + ':' + p(m) + ':' + p(s);
    }
    return str
}

export const timeStrToSec = (timeStr) => {
    var hour = timeStr.split(':')[0];
    var min = timeStr.split(':')[1];
    var sec = timeStr.split(':')[2];
    if (hour === '00') {
        return Number(hour * 60) + Number(min);
    } else {
        return Number(hour * 3600) + Number(min * 60) + Number(sec);
    }
}

export function getUTCTime() {
    let d1 = new Date();
    let d2 = new Date(d1.getUTCFullYear(), d1.getUTCMonth(), d1.getUTCDate(), d1.getUTCHours(), d1.getUTCMinutes(), d1.getUTCSeconds());
    return Date.parse(d2) / 1000;
}

