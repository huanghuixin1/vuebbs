<template>
   <span>
       <h2 class="userInfo-title">
            用户信息页面
        </h2>
        <v_userlist></v_userlist>

        <v_usertest></v_usertest>
   </span>
</template>
<script>
    import v_userlist from "./userList/listInfo.vue"

    import store from "./store"
    export default{
        components: {
            v_userlist,
            v_usertest
        },
        store
    }
</script>
