<template>
    <div class="user_login">
        <div class="logo_dz">
            <img alt="logo" src="http://att.discuz.net/data/attachment/common/26/common_36_i94wrTB4.jpg">
            <h2>Discuz! 官方站</h2>
        </div>
        <div class="login_set">
            <form id="loginBox" class="login_set">
                <p>
                    <input type="text" name="username" id="username" placeholder="请输入帐户" value="">
                </p>
                <p>
                    <input type="password" name="password" id="password" placeholder="请输入密码">
                </p>
                <p>
                    <input type="text" id="questionclick" placeholder="安全提问(未设置请忽略)" value="">
                    <input type="hidden" name="questionid" id="questionid" value="0">
                </p>
                <p id="qainput" style="display:none">
                    <input type="text" name="answer" id="answer" placeholder="请输入答案">
                </p>
                <div class="log_bar"><input id="loginBtn" type="button" value="登录" class="lb_lq_btn"></div>
                <div id="toQuickLogin" class="to_choose">
                    <a class="fr" id="registerBtn">注册账户</a>
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
