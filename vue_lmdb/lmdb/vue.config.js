const {defineConfig} = require('@vue/cli-service')
const path = require('path')
module.exports = defineConfig({
    transpileDependencies: true,
    lintOnSave: false,
    chainWebpack: (chain) => {
        const oneofsMap = chain.module.rule('scss').oneOfs.store
        oneofsMap.forEach(item => {
            item
                .use('sass-resources-loader')
                .loader('sass-resources-loader')
                .options({
                    resources: './src/styles/mixin.scss',
                })
        })
    },

    publicPath: process.env.NODE_ENV === 'production' ? './' : '/',

    outputDir: 'dist',

    runtimeCompiler: true, //关键点在这

    // 调整内部的 webpack 配置。

    // 查阅 https://github.com/vuejs/vue-doc-zh-cn/vue-cli/webpack.md


    configureWebpack: () => {
    },

    // 配置 webpack-dev-server 行为。

    devServer: {

        open: process.platform === 'darwin',

        host: '0.0.0.0',

        port: 8080,

        https: false,

        hotOnly: false,

        // 查阅 https://github.com/vuejs/vue-doc-zh-cn/vue-cli/cli-service.md#配置代理

        proxy: null, // string | Object

        before: app => {
        }

    }

})
