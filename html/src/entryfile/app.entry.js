import v_app from "../compoment/index/app.vue";
import VueRouter from "vue-router";
import Vue from "vue";
import {requireCss} from "../common/_require";
import v_articleList from "../views/articles/articleList/articleList.vue";


Vue.use(VueRouter);
var App = Vue.extend(v_app);

const PathPre = "/view";
const routes = [
    {
        path: "/",
        redirect: to=> {
            return PathPre + "/articles/list";
        }
    },
    {
        path: PathPre + "/categories/list", ////板块列表
        component: (resolve)=> {
            requireCss("/views/categories/list/list.css");
            require(["views/categories/list/list.vue"], resolve);
        },
        name: "categories-list"
    },
    {
        path: PathPre + "-userinfo",
        component: (resolve) => {
            requireCss("/compoment/userInfo/index.css");
            require(["../compoment/userInfo/index.vue"], resolve);
        },
        name:"userinfo"
    },
    {
        path: PathPre + "/testpage",
        component: (resolve) => {
            require(["../compoment/testpage/p1.vue"], resolve);
        }
    },
    {
        path: PathPre + "/articles/list",
        component: v_articleList
    },
    {
        path: PathPre + "/articles/detail/:id",
        component: resolve => {
            requireCss("/views/articles/articleDetail/articleDetail.css");

            require(["../views/articles/articleDetail/articleDetail.vue"], resolve);
        },
        name: "article-detail"
    }, {
        path: PathPre + "/login",
        component: (resolve)=> {
            requireCss("/common/login_register.css");
            require(["../views/userInfo/userLogin/userLogin.vue"], resolve);
        }
    }, {
        path: PathPre + "/register",
        component: (resolve)=> {
            requireCss("/common/login_register.css");
            require(["../views/userInfo/userRegister/userRegister.vue"], resolve);
        }
    }];

const router = new VueRouter({
    mode: "history",
    routes // （缩写） 相当于 routes: routes
});

new Vue({
    render: h => h(App),
    router
}).$mount("#app");