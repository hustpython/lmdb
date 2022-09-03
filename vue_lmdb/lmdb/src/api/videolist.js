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