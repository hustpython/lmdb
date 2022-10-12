const {defineConfig} = require('@vue/cli-service')
module.exports = {
    publicPath: './',

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
}
