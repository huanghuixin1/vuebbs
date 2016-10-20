<template>
    <article>
        <v_header></v_header>

        <article class="categories-list">
            <router-link :to="{ path: '/articles/list?cid=' + item.Id }" v-for="item in listCategories">
                <h2>{{item.CName}}</h2>
                <i>主题: {{item.ArticleCount}}</i>
            </router-link>
        </article>

        <v_footer></v_footer>
    </article>

</template>

<script lang="babel">
    import v_header from "../../../compoment/header_footer/index_header/index_header.vue";
    import v_footer from "../../../compoment/header_footer/footer_onlyback/footer_onlyback.vue";

    import * as categoriesApi from "../../../service/categoriesService";

    import store from "./store";

    export default{
        components: {
            v_header,v_footer
        },
        data(){
            return {
                listCategories:[]
            }
        },
        beforeCreate: function () {
            this.$store.dispatch("SetIsShowHeaderDetail", {isShow: false});//设置不显示详细信息

            //初始化信息
            //获取板块列表 (因为这里的数据很简单  所以直接获取即可)
            categoriesApi.getCategoryListName({success: data=>{
                this.listCategories = data.Data;
            }});

        },
        created: function () {
            this.$store.dispatch("SetTitle", {title: "板块列表"});
        },
        store
    }
</script>