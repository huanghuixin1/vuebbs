<template>
    <footer class="bottom-container bottomBar">
        <div class="bottomIcon">
            <!--左箭头 返回-->
            <a @click="back()" class="backBtn">
                <i class="icon iconfont iconAnswer">&#xe628;</i>
            </a>

            <!--板块-->
            <router-link :to="{ path: '/categories/list'}" class="moduleIcon module">
                <i class="icon iconfont">&#xe605;</i>模块
            </router-link>
            <!--发帖-->
            <a href="javascript:;" class="module postIcon" @click="publish()">
                <i class="icon iconfont">&#xe615;</i>发帖
            </a>
            <!--更多-->
            <a href="javascript:void(0);" id="info_center" class="moreC db">
                <i class="icon iconfont circle">&#xe633;</i>
            </a>
        </div>
    </footer>
</template>
<script>
    export default{
        methods: {
            publish: function () {
                alert("这个点不得, 点旁边的模块");
                return false;
            },
            back(){
                alert("这个也点不得,点旁边的模块");
                return false;
            }
        }
    }
</script>
