<!--作者信息 放在帖子列表和帖子详情的顶部位置 -->
<template>
    <router-link :to="{ path: '/userinfos/userdetail/' + authorinfo.UserId}" class="authorInfo">
        <!--左边头像-->
        <img v-bind:src="headerImg"/>

        <!--右边作者介绍-->
        <address>
            {{authorinfo.NickName}} <i>lv12</i>

            <p>{{authorinfo.CreateTime | tamp2span}}</p>
        </address>
    </router-link>
</template>

<script>
    import _ from "../../utls/filters";
    import config from "../../utls/config"

    export default{
        computed: {
            headerImg(){ //处理没有上传头像的默认头像问题
                return this.authorinfo.HeaderImg || config.defaltHeaderImg;
            }
        },
        props: ["authorinfo"]
    }
</script>