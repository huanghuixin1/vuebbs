<template>
    <div id="app">

        <!-- 路由外链 -->
        <article class="main">
            <transition name="fade" mode="out-in">
                <router-view></router-view>
            </transition>
        </article>
    </div>
</template>

<script>
    export default{
        methods:{
            test:()=>{
                
            }
        }
    }
</script>