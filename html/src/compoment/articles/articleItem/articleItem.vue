<template>
    <section class="articleItem">

        <!--发帖人信息-->
        <v_authorinfo :authorinfo="item"></v_authorinfo>

        <!--帖子摘要内容-->
        <router-link :to="{ path: '/articles/detail/' + item.Id}" class="articleItem-content">
            <!--标题-->
            <h2>
                {{item.Title}}
            </h2>

            <!--摘要-->
            <p>
                {{item.Summary}}
            </p>

            <!--回帖数量-->
            <div>
                <img src="/imgs/icon_replay.png"/>
                <i>999</i>
            </div>
        </router-link>
    </section>
</template>
<script>
    import v_authorinfo from "../../userInfo/authorInfo/authorInfo.vue";

    export default{
        props: ["item"],
        components: {
            v_authorinfo
        }
    }
</script>
