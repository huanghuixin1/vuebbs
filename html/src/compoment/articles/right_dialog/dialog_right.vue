<template>
    <article class="articlelist-dialogright-container"
             v-bind:class="{ 'articlelist-dialogright-container-show': isShowDialog }">
        <article class="articlelist-dialogright">
            <h3>话题筛选</h3>

            <ul class="dialogright-types">
                <li class="dialogright-types-selected">全部 <i>[ok]</i></li>
                <li @click="setType(1)">公告 <i>[ok]</i></li>
                <li @click="setType(2)">待解决 <i>[ok]</i></li>
                <li @click="setType(3)">已解决 <i>[ok]</i></li>
            </ul>
        </article>
    </article>
</template>

<script lang="babel">
    import  * as urlUtls from "../../utls/urlUtls";

    export default{
        computed: {
            isShowDialog(){
                return this.$store.state.articleList.showRightDialog
            }
        },
        methods: {
            setType(type){
                urlUtls.setQueryString("type", type)
            }
        }
    }
</script>