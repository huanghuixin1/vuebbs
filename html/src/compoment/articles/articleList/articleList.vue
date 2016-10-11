<template>

    <article>
        <v_header></v_header>
        <v_dropload></v_dropload>
        <article class="articleList" ref="articlelist">
            <v_article_item v-for="item in listData" track-by="$index" :item="item"></v_article_item>

            <a class="articleList-loadData" @click="loadBottom()" ref="loadbottom">点击加载</a>
        </article>

        <v_footer></v_footer>
    </article>
</template>
<script>

    import v_header from '../../header_footer/index_header/index_header.vue';
    import v_footer from '../../header_footer/index_footer/index_footer.vue';

    import v_article_item from "../articleItem/articleItem.vue";
    import v_dropload from "../../loading/dropload.vue"; //下拉加载组件
    import store from "./store";
    import {mapGetters, mapActions} from 'vuex';

    let cid;//当前url的模块id

    export default {
        computed: mapGetters({
            listData: 'articleList'
        }),
        components: {
            v_article_item, v_dropload,v_header,v_footer
        },
        created() {
            cid = this.$route.query["cid"] || 0; //获取到cid
            this.$store.dispatch("setArticleList", {cidOrData: cid, minId: 0, maxId: 0});
        },
        mounted() {
            //初始化下拉刷新
            v_dropload.initDropLoad(this.$refs.articlelist, setLoaded => {
                this.$store.dispatch("setArticleList", {
                    cidOrData: cid, minId: this.listData[0].Id, callback: () => {
                        setTimeout(() => {
                            setLoaded();
                        }, 1000);
                    },
                    //数据获取失败
                    errCallback: () => {
                        alert("数据获取失败,请重试");
                        setLoaded();
                    }
                })
            })
        },
        methods: {
            //底部滚动加载
            loadBottom: function () {
                let loadbottom = this.$refs.loadbottom;
                loadbottom.innerHTML = "加载中...";

                this.$store.dispatch("setArticleList", {
                    cidOrData: cid, maxId: this.listData[this.listData.length - 1].Id, callback: () => {
                        setTimeout(() => {
                            loadbottom.innerHTML = "点击加载";
                        }, 1000);
                    },
                    //数据获取失败
                    errCallback: () => {
                        alert("数据获取失败,请重试");
                    }
                });
            }
        },
        store
    }
</script>