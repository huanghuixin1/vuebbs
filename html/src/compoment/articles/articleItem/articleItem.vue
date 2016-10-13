<template>
    <section class="articleItem">

        <!--发帖人信息-->
        <router-link :to="{ path: '/userinfos/userdetail/' + item.UserId}" class="articleItem-author">
            <!--左边头像-->
            <img src="http://uc.discuz.net/avatar.php?uid=2810018&size=small"/>


            <!--右边作者介绍-->
            <address>
                {{item.NickName}} <i>lv1</i>

                <p>{{item.CreateTime | tamp2span}}</p>
            </address>
        </router-link>

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
    import _ from "../../utls/filters";
    export default{
        props: ["item"]
    }
</script>
