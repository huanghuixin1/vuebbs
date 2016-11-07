"use strict";
var gulp = require("gulp"),
    gulp_sourcemaps = require("gulp-sourcemaps"),
    autoprefixer = require("gulp-autoprefixer"),//增加私有变量前缀
    fs = require("fs"),
    webpackStream = require("webpack-stream"),
    webpack = require("webpack"),
    path = require("path"),
    glob = require("glob"),
    md5 = require("gulp-md5-plus"),
    gulpif = require("gulp-if"),//if判断，用来区别生产环境还是开发环境的
    replace = require("gulp-replace"),
    sass = require("gulp-sass"),
    imagemin = require("gulp-imagemin"),
    clean = require("gulp-clean"),//清理文件
    AssetsPlugin = require("assets-webpack-plugin"),
    assetsPluginInstance = new AssetsPlugin({filename: "assets.json", path: path.join(__dirname, "dist")}),
    replaceAssets = require("gulp-replace-assets");

const DEBUG = process.argv[2] !== "build"


gulp.task("test", ()=> {
    //如果是开发模式 则进行持续构建
    console.log(path.join(__dirname, "src/vendor/ajax.js"))
})


gulp.task("dev", ()=> {
    //如果是开发模式 则进行持续构建

    gulp.start("html");
    gulp.start("img");
    gulp.start("css");
    gulp.start("js");
    gulp.start("watch");
})

gulp.task("build", ["clean"], function () {
    //build:img -> js,css -> html

    gulp.start("img");
})

gulp.task("clean", function () {
    return gulp.src(["dist/*", "!.gitignore"])
        .pipe(clean({force: true}));
});

const cssSrc = ["src/**/*.scss"];
const htmlSrc = "src/**/*.html";
// const jsSrc = "["src/**/*.entry.js","src/vendor.js",src/global.js]"
const imgSrc = ["src/imgs/**/*.jpg", "src/imgs/**/*.png", "src/imgs/**/*.gif"];

gulp.task("watch", function () {
    gulp.start("js");
    gulp.watch(imgSrc, ["img"]);
    gulp.watch(cssSrc, ["css"]);
    gulp.watch(htmlSrc, ["html"]);
    console.log("监控文件...");
});

gulp.task("img", DEBUG ? "" : ["js", "html", "css"], ()=> {
    return gulp.src(imgSrc)
        .pipe(gulpif(!DEBUG, imagemin()))
        .pipe(gulp.dest("dist/imgs"))
        .pipe(gulpif(!DEBUG, md5(8, ["./dist/**/*.html", "./dist/**/*.js", "./dist/**/*.css"])))
        .pipe(gulp.dest("dist/imgs"));
});

gulp.task("html", DEBUG ? "" : ["js"], ()=> {
    if (DEBUG) {
        return gulp.src(htmlSrc)
            .pipe(replace(".scss\"", ".css\""))
            .pipe(gulp.dest("dist"));
    } else {
        // replaceAssets
        fs.readFile(path.join(__dirname, "dist", "assets.json"), (err, data) => {
            if (err) throw err;
            var resuJson = JSON.parse(data.toString());
            var jsAssets = {}; //最后需要替换的字符串
            for (var key in resuJson) {
                if (resuJson[key].hasOwnProperty("js")){
                    let temp = resuJson[key].js;
                    jsAssets[key + ".js"] = temp.substr(1,temp.length - 1)
                }

                if (resuJson[key].hasOwnProperty("css")) {
                    jsAssets[key + ".scss"] = resuJson[key].css;

                }
            }

            return gulp.src(htmlSrc)
                .pipe(replaceAssets(jsAssets))
                .pipe(gulp.dest("dist"));

        });
    }
    // .pipe(replace("[{jsPath}]/","../../public/"))
});


//sass编译为css
gulp.task("css", DEBUG ? "" : ["html"], ()=> {
    return gulp.src(cssSrc)
        .pipe(gulpif(DEBUG, gulp_sourcemaps.init())) //如果debug则生成sourcemap
        .pipe(sass({outputStyle: DEBUG ? "expand" : "compressed"}))
        .pipe(autoprefixer({browserslist: ["iOS9", "Android >= 4.4"]}))
        .pipe(gulpif(DEBUG, gulp_sourcemaps.write("dist"))) //如果debug则生成sourcemap
        .pipe(gulp.dest("dist"))
        .pipe(gulpif(!DEBUG, md5(8, ["./dist/**/*.html", "./dist/**/*.js"])))

        .pipe(gulp.dest("dist"));
});

gulp.task("js", ()=> {
    function entries(globPath) {
        var files = glob.sync(globPath);
        var entries = {}, entry, dirname, basename;

        for (var i = 0; i < files.length; i++) {
            entry = files[i];
            // if (files[i].toString().match(/js\/store/i) || files[i].toString().match(/js\/lib/i) || files[i].toString().match(/js\/compoment/i) )
            //     continue;
            //过滤一下非入口文件
            if (!files[i].toString().match(/.+\.entry\.js/i)) {
                continue;
            }
            dirname = path.dirname(entry);
            basename = path.basename(entry, ".js");

            //干掉前面的assets
            entries[path.join(dirname, basename).replace("src/", "").replace("src\\", "").replace("\\", "/")] = "./" + entry;
            // entries[path.join(dirname, basename)] = "./" + entry;
        }

        entries["common/global"] = "./src/common/global.js";
        entries["common/vendor"] = ["vue", "vuex", "./src/common/vendor.js"];

        console.log(entries);
        return entries;
    }

    return gulp.src("")
        .pipe(webpackStream({
            watch: DEBUG,
            entry: entries("src/**/*.entry.js"),// entries("js/**/*.js"),
            output: {
                publicPath: '/',
                filename: "[name]" + (DEBUG ? "" : "-[chunkhash:8]") + ".js"
            },
            module: {
                loaders: [
                    {
                        test: /\.js$/,
                        loader: "babel!eslint",
                        exclude: /node_modules/
                    },

                    {
                        test: /\.vue$/,
                        loader: "vue",
                        exclude: /node_modules/
                    }
                ]
            },
            eslint: {
                configFile: path.join(__dirname, ".eslintrc.js"),
                ignorePath: path.join(__dirname, ".eslintignore")
            },
            resolve: {
                root: path.resolve('./src'),
                extensions: ["", ".js"],
                alias: {
                    "ajax": path.join(__dirname, "src/vendor/ajax.js")
                }
            },
            babel: {
                presets: ["es2015"]
            },
            plugins: [
                // new ExtractTextWebpackPlugin("[name]" + (DEBUG ? "" : "-[contenthash:8]") + ".css"),
                new webpack.optimize.CommonsChunkPlugin("common/vendor", "common/vendor" + ( DEBUG ? "" : "-[chunkhash:8]") + ".js"),
                // new ExtractTextPlugin("style" + (DEBUG ? "" : "-[contenthash]") + ".css"),
                new webpack.DefinePlugin({}),
                DEBUG ? function () {
                } : new webpack.optimize.UglifyJsPlugin({
                    compress: {
                        warnings: false
                    }
                }),
                new webpack.optimize.OccurenceOrderPlugin(),

                new webpack.DefinePlugin({
                    'process.env': {
                        NODE_ENV: DEBUG ? '"debug"' : '"production"'
                    }
                }),
                //生成资源文件的键值对
                assetsPluginInstance
            ]
        }))
        // .pipe(replace("{[imgPath]}", ""))
        .pipe(gulp.dest("dist"));
});