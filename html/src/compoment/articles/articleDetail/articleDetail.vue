<!--帖子详情页面-->
<template>
    <article class="articleDetail">
        <v_header></v_header>
    </article>
</template>
<script>
    import v_header from '../../header_footer/index_header/index_header.vue';

    import store from "./store";

    export default{
        components: {
            v_header
        },
        beforeCreate: function () {
            this.$store.dispatch("SetIsShowHeaderDetail", {isShow: false});//设置不显示详细信息
            this.$store.dispatch("SetTitle", {title: "帖子详情"});
            //初始化信息
            //获取板块列表 (因为这里的数据很简单  所以直接获取即可)
//            categoriesApi.getCategoryListName({success: data=>{
//                this.listCategories = data.Data;
//            }});

        },
        store
    }
</script>