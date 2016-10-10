# 项目说明
//jy0322007:1

## GO
服务器端开发接口环境,数据库使用 **maraidb** ,缓存使用 **redis**


## node.js
```
对应 gulp和 webpack 作为前端代码的构建
```

### 编译命令：

#####开发:持续构建 开发环境中使用
```
npm run dev 
```

##### 部署:删除dist下所有文件，重新生成并添加文件的MD5
```
npm run build
```

### 静态文件目录

```
- html  --- 静态文件
  - gulpfile.js gulp的脚本 
  - src  源代码目录
  - dist 将src编译后的文件目录
  - imgs 存放图片
```

#### 关于开发的约定
```
1.js的入口文件都以 .entry.js来结尾  如 index.entry.js
2.scss文件,非入口需要以"_"为文件开头, 如 _header.scss
```

## 开发模式说明
总的来说,是以spa的方式来进行开发,除了首页内容,其他的组件都以异步的形式进行加载
```
1.使用vue的组件化进行开发,每个组件的样式写到单独的.scss文件中,并且要在同一目录,参考 /html/compoment/header_footer/index_footer
2.所有的html入口都放到src的根目录,在入口的基础上,使用vue-router对url进行匹配并加载对应的页面
注: 组件的scss文件是单独写的,所以要在require组件之前,把样式require进来.具体的操作参考 /html/compoment/app.vue
3.因为构建的原因,所以引入的时候是 xxx.css 而不要引入成xxx.scss 否则无法替换引用,所以,一定要删除.vue文件中的style节点!
```

###用到的相关库

> 前端框架
* [Vue1 官方网站](http://cn.vuejs.org/)
* [vue-router 文档](http://vuejs.github.io/vue-router/zh-cn/index.html)
* [vuex 文档](http://vuex.vuejs.org/zh-cn/index.html)
* [ajax库](https://github.com/fdaciuk/ajax)
* [hackernews 的 Vue 实现实例项目](https://github.com/vuejs/vue-hackernews)


> 服务器端
* [orm:GORM](https://github.com/jinzhu/gorm)
* [web框架:baa](https://github.com/go-baa/baa)

###