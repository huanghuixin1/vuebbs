<!--帖子详情页面-->
<template>
    <article class="articleDetail-container">
        <v_header></v_header>

        <section class="articleDetail">
            <!--发布者相关-->
            <div class="author-container">
                <v_authorinfo :authorinfo="article"></v_authorinfo>
                <span>楼主 1L</span>
            </div>

            <h1>{{ article.Title }}</h1>

            <div v-html="article.AContent">

            </div>
        </section>
    </article>
</template>
<script>
    import v_header from '../../header_footer/index_header/index_header.vue';
    import v_authorinfo from "../../userInfo/authorInfo/authorInfo.vue";

    import * as articleService from "../../../service/articlesService";

    import store from "./store";

    let id;
    export default{
        data(){
            return {
                article: {}
            }
        },
        components: {
            v_header, v_authorinfo
        },
        beforeCreate: function () {
            //获取到cid
            id = this.$route.params.id || 0;

            this.$store.dispatch("SetIsShowHeaderDetail", {isShow: false});//设置不显示详细信息
            this.$store.dispatch("SetTitle", {title: "帖子详情"});

            let _this = this;
            //初始化信息
            articleService.getDetail({
                id, success: (ret)=> {
                    _this.article = ret.Data;
                }
            });
        },
        store
    }
</script>