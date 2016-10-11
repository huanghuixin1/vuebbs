<!--首页的头部-->
<template>
    <header>
        <a class="header-logo" href="/app.html" title="首页">
            <img src="/imgs/index_logo.jpg"/>
        </a>

        <section class="header-desc">
            <h1>HHX官方论坛 | {{title}}</h1>

            <h2>
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
            totalArticles: "totalArticles"
        }),
        created: function () {
            //先初始化一次信息
            this.$store.dispatch("SetTotalMembers");
            let cid = this.$route.query["cid"] || 0;
            this.$store.dispatch("SetTotalArticles", {cid: cid});

            //初始化标题
//            获取url参数
            let category = getQueryString("cid");
            this.$store.dispatch("SetTitle", cid);

            //路由变化后
            this.$router.afterEach(transition => {
                //对会员数量刷新
                this.$store.dispatch("SetTotalMembers");
                //当前板块的帖子数量刷新

                //获取板块id
                let cid = this.$route.query["cid"] || 0;

                this.$store.dispatch("SetTotalArticles", {cid: cid});
            });
        }
    }
</script>