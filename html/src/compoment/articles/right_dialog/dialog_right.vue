<template>
    <article class="articlelist-dialogright-container" v-bind:class="{ 'articlelist-dialogright-container-show': isShowDialog }">
        <article class="articlelist-dialogright">
            <h3>话题筛选</h3>

            <ul class="dialogright-types">
                <li class="dialogright-types-selected">全部 <i>[ok]</i></li>
                <li>公告 <i>[ok]</i></li>
                <li>待解决 <i>[ok]</i></li>
                <li>已解决 <i>[ok]</i></li>
            </ul>
        </article>
    </article>
</template>

<script lang="babel">
    import {mapGetters, mapActions} from 'vuex';

    export default{
        computed: mapGetters({
            isShowDialog : "isShowRightDialog"
        })
    }
</script>