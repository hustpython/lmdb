// 导入axios实例
import httpRequest from './request'

// 定义接口的传参


// 获取Video信息
export function GetVideList() {
	return httpRequest({
		url: '/movie',
		method: 'get',
	})
}

export function SyncVideo(param) {
	return httpRequest({
		url: '/filter',
		method: "post",
		data: JSON.stringify(param),
	})
}