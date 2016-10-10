<template>
    <section class="articleItem">

        <!--发帖人信息-->

        <a class="articleItem-author" v-bind:href="'/userinfos/userdetail/' + item.UserId">
            <!--左边头像-->
            <img src="http://uc.discuz.net/avatar.php?uid=2810018&size=small"/>


            <!--右边作者介绍-->
            <address>
                {{item.NickName}} <i>lv1</i>

                <p>{{item.CreateTime | tamp2span}}</p>
            </address>
        </a>

        <!--帖子摘要内容-->
        <a class="articleItem-content">
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
        </a>
    </section>
</template>
<script>
    import _ from "../../utls/filters";
    export default{
        props: ["item"]
    }
</script>
