<!--首页的头部-->
<template>
    <header>
        <a class="header-logo" href="/app.html" title="首页">
            <img src="/imgs/index_logo.jpg"/>
        </a>

        <section class="header-desc">
            <h1>HHX官方论坛 | {{title}}</h1>

            <h2 v-show="isShowDetail">
                <i>{{totalArticles}} 帖子</i>
                <i>{{totalMembers}} 会员</i>
            </h2>
        </section>
    </header>
</template>
<script>
    import {mapGetters, mapActions} from 'vuex';
    import {getQueryString} from "../../utls/urlUtls";

    export default {
        computed: mapGetters({
            title: 'title',
            totalMembers: "totalMembers",
            totalArticles: "totalArticles",
            isShowDetail: "isShowHeaderDetail" //是否显示头部的详情
        }),

        created: function () {
            //如果显示详细信息 才初始化数据
            if(this.isShowDetail){
                //先初始化一次信息
                this.$store.dispatch("SetTotalMembers");
                let cid = this.$route.query["cid"] || 0;
                this.$store.dispatch("SetTotalArticles", {cid: cid});

                //初始化标题
                let category = getQueryString("cid");
                this.$store.dispatch("SetTitle", {cid});
            }
        }
    }
</script>