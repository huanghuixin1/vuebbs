<template>
    <div class="user_login">
        <div class="logo_dz">
            <img alt="logo" src="http://att.discuz.net/data/attachment/common/26/common_36_i94wrTB4.jpg">
            <h2>hhxblogs</h2>
        </div>
        <div class="login_set">
            <form id="loginBox" class="login_set" @submit.prevent="loginSubmit">
                <p>
                    <input type="text" v-model="user.LoginName" placeholder="请输入帐户" value="">
                </p>
                <p>
                    <input type="password" v-model="user.LoginPwd" placeholder="请输入密码">
                </p>
                <p id="qainput" style="display:none">
                    <input type="text" name="answer" id="answer" placeholder="请输入答案">
                </p>
                <div class="log_bar"><input id="loginBtn" type="submit" value="登录" class="lb_lq_btn"></div>
                <div id="toQuickLogin" class="to_choose">

                    <router-link to="/register" class="lr" id="registerBtn">注册账户</router-link>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
    import * as userinfosService from "../../../service/userinfosService";

    export default{
        data(){
            return {
                user: {
                    LoginName: '',
                    LoginPwd: ''
                }
            }
        },
        methods: {
            loginSubmit() {
                var user = this.user;
                userinfosService.loginUser({
                    user, success: (ret)=> {
                        if(ret.Status=='100002'){
                            alert('用户名不存在');
                        }else if(ret.Status=='0'){
                            alert('注册成功了,走一波');
                            transition.redirect('/articles/list')
                        }else{
                            alert('what?出错了')
                        }
                    }
                })
            }
        }
    }
</script>
