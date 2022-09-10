export const timeFilter = (val) => {
    function p(t) {
        return t < 10 ? '0' + t : t;
    }
    var h = Math.floor(val / 60 / 60);
    var m = Math.floor(val / 60 % 60);
    var s = Math.floor(val % 60);
    var str = p(h) + ':' + p(m) + ':' + p(s);
    return str
}

export function getUTCTime() {
    let d1 = new Date();
    let d2 = new Date(d1.getUTCFullYear(), d1.getUTCMonth(), d1.getUTCDate(), d1.getUTCHours(), d1.getUTCMinutes(), d1.getUTCSeconds());
    return Date.parse(d2)/1000;
}

