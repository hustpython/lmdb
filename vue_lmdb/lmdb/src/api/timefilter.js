export const timeFilter = (val) => {
    function p(t) {
        return t < 10 ? '0' + t: t;
    }
    var h = Math.floor(val / 60 / 60);
    var m = Math.floor(val  / 60 % 60);
    var s = Math.floor(val  % 60);
    var str = p(h) + ':' + p(m) + ':' + p(s);
    return str
}
