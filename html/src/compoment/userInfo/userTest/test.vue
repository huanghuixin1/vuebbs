<template>
    <div>
        <button @click="setUserCount(1)">测试添加用户数量</button>
    </div>
</template>
<script>
    import {setUserCount} from "../userList/action"

    export default{
        vuex: {
            actions:{
                setUserCount
            }
        }
    }
</script>
