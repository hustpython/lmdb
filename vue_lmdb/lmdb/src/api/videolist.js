// 导入axios实例
import httpRequest from './request'

// 获取Video信息
export function GetVideList() {
    return httpRequest({
        url: '/movie',
        method: 'get',
    })
}

export function SyncVideo(param) {
    return httpRequest({
        url: '/movie',
        method: "post",
        data: JSON.stringify(param),
    })
}

export function UpdateVideo(param) {
    return httpRequest({
        url: '/movie',
        method: "PUT",
        data: JSON.stringify(param),
    })
}

export function GetAllTags() {
    return httpRequest({
        url: '/tag/AllTag',
        method: "GET",
    })
}

export function GetMoviesByTag(param) {
    return httpRequest({
        url: '/tag/?Tag=' + param,
        method: "GET",
    })
}

export function GetAllColl() {
    return httpRequest({
        url: '/coll/AllColl',
        method: "GET",
    })
}

export function GetMovieAndColl() {
    return httpRequest({
        url: '/coll/MovieAndColl',
        method: "GET",
    })
}

export function BatchAddColl(param) {
    return httpRequest({
        url: '/coll',
        method: "POST",
        data: JSON.stringify(param),
    })
}

export function GetMoviesByColl(param) {
    return httpRequest({
        url: '/coll/?Coll=' + param,
        method: "GET",
    })
}

export function GetMoviesByRecent() {
    return httpRequest({
        url: '/recent',
        method: "GET",
    })
}

export function DeleteMovieColl(param) {
    return httpRequest({
        url: '/coll/?MId=' + param,
        method: "DELETE",
    })
}

export function GetComment(param) {
    return httpRequest({
        url: '/comment/mid/?MId=' + param,
        method: "GET",
    })
}

export function DeleteComment(param) {
    return httpRequest({
        url: '/comment',
        method: "DELETE",
        data: JSON.stringify(param),
    })
}

export function AddComment(param) {
    return httpRequest({
        url: '/comment',
        method: "POST",
        data: JSON.stringify(param),
    })
}

export function GetCommentByColl(param) {
    return httpRequest({
        url: '/comment/coll/?Coll=' + param,
        method: "GET",
    })
}

export function CutVideoByMId(param) {
    return httpRequest({
        url: '/ffmpeg/cut',
        method: "POST",
        data: JSON.stringify(param),
    })
}