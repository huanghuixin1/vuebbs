<template>
    <div class="user_login">
        <div class="logo_dz">
            <img alt="logo" src="http://att.discuz.net/data/attachment/common/26/common_36_i94wrTB4.jpg">
            <h2>hhxblogs</h2>
        </div>
        <div class="login_set">
            <form id="loginBox" class="login_set">
                <p>
                    <input type="text" name="username" id="username" placeholder="请输入帐户" value="">
                </p>
                <p>
                    <input type="password" name="password" id="password" placeholder="请输入密码">
                </p>
                <p id="qainput" style="display:none">
                    <input type="text" name="answer" id="answer" placeholder="请输入答案">
                </p>
                <div class="log_bar"><input id="loginBtn" type="button" value="登录" class="lb_lq_btn"></div>
                <div id="toQuickLogin" class="to_choose">

                    <router-link to="/register" class="lr" id="registerBtn">注册账户</router-link>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
    export default{
        data(){
            return {
                message: 'Helo, Vue.js 2.0'
            }
        }
    }
</script>
