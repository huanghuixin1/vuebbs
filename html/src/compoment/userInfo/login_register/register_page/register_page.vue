<template>
    <div class="user_login">
        <div class="logo_dz">
            <img alt="logo" src="http://att.discuz.net/data/attachment/common/26/common_36_i94wrTB4.jpg">
            <h2>Discuz! 官方站</h2>
        </div>
        <div class="login_set">
            <form id="loginBox" class="login_set">
                <p>
                    <input type="text" name="LoginName" id="LoginName" placeholder="请输入帐户" value="">
                </p>
                <p>
                    <input type="password" name="LoginPwd" id="LoginPwd" placeholder="请输入密码">
                </p>
                <p>
                    <input type="password" name="LoginConfirmationPwd" id="LoginConfirmationPwd" placeholder="请确认密码">
                </p>
                <p>
                    <input type="password" name="NickName" id="NickName" placeholder="请输入昵称">
                </p>
                <p>
                    <input type="password" name="EmailAddr" id="EmailAddr" placeholder="请输入邮箱">
                </p>
                <div class="log_bar"><input id="registerBtn" type="button" value="登录" class="lb_lq_btn"></div>
                <div id="toQuickLogin" class="to_choose">
                    <a class="fr" id="loginBtn">登录</a>
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
