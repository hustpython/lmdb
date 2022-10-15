# LMDB
[![Build Status](https://travis-ci.com/PowerDos/Mall-Vue.svg?branch=master)](https://travis-ci.com/PowerDos/Mall-Vue)
[![VueJS Version](https://img.shields.io/badge/VueJS-v3.2.13-green.svg?style=flat-square)](https://vuejs.org/)
[![Go Version](https://img.shields.io/badge/Go-1.19-green.svg?style=flat-square)](https://vuejs.org/)
[![JavaScript Style Guide](https://img.shields.io/badge/code_style-standard-brightgreen.svg)](https://standardjs.com)
[![Read the Docs (version)](https://img.shields.io/readthedocs/pip/stable.svg)](https://github.com)
[![npm version](https://img.shields.io/badge/npm-v8.13.2-brightgreen.svg)](https://standardjs.com)
[![author](https://img.shields.io/badge/author-mxq-brightgreen.svg)](https://standardjs.com)



> 这是一个基于VUE + Go做的一个本地视频播放器项目，欢迎fork或star。

预览地址(无后端数据) : https://hustpython.github.io/lmdb 

## 使用指南

### 本地开发
``` bash
# 安装依赖
npm install

# 开发模式
npm run dev

# 打包
npm run build

# 启动后端
bee run
```
### 直接使用

- 解压 lmdb\releases.zip

- 进入releases文件夹

- 进入 js 目录，将js文件中 "192.168.31.54:9090" 

  换成 自己的ip+port(主要是为了移动端能够访问本地PC后台)

- 返回 双击 run.bat(本命令是打开index.html文件)，如果需要监听端口方便移动端访问需要在index.html命令

  打开cmd，输入 http-server 回车



## 功能清单
- [x] [本地/移动硬盘视频快速搜索与过滤](#本地视频快速搜索与过滤)
- [x] [数据持久化](#数据持久化)
- [x] [批量创建合集](#批量创建合集)
- [x] [快捷键](#快捷键)

### 本地视频快速搜索与过滤

![](/readmeimg/1.png)

### 数据持久化

![](./readmeimg./2.png)

### 批量创建合集

![](./readmeimg./3.png)

### 快捷键

| 键名    | 功能                  |
| ------- | --------------------- |
| Space   | 播放/暂停             |
| Ctrl    | 打开评论/关闭评论窗口 |
| Enter   | 确认评论              |
| Left    | 左快进6s              |
| Right   | 右快进6s              |
| Insert  | 截图                  |
| Alt + A | 打开/关闭评论列表     |

