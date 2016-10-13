<template>
    <div id="app">

        <!-- 路由外链 -->
        <article class="main">
            <transition name="fade" mode="out-in">
                <router-view></router-view>

                <!--<div class="loading-index">加载中...</div>-->
            </transition>
        </article>
    </div>
</template>

<script>
    export default{}
</script>