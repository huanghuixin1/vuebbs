<template>
    <div class="user_login">
        <div class="logo_dz">
            <img alt="logo" src="http://att.discuz.net/data/attachment/common/26/common_36_i94wrTB4.jpg">
            <h2>Discuz! 官方站</h2>
        </div>
        <div class="login_set">
            <form id="loginBox" class="login_set" @submit.prevent="handleSubmit">
                <p>
                    <input type="text" v-bind="user.LoginName" v-model="LoginName" placeholder="请输入帐户">
                </p>
                <p>
                    <input type="password" v-bind="user.LoginPwd" v-model="LoginPwd"  placeholder="请输入密码">
                </p>
                <p>
                    <input type="password" name="regCfnPwd" v-model="regCfnPwd" placeholder="请确认密码">
                </p>
                <p>
                    <input type="text" v-bind="user.NickName" v-model="NickName"  placeholder="请输入昵称">
                </p>
                <p>
                    <input type="text" v-bind="user.EmailAddr" v-model="EmailAddr"  placeholder="请输入邮箱">
                </p>
                <div class="log_bar">

                    <input id="registerBtn" type="submit" value="注册" class="lb_lq_btn"></div>
                <div id="toQuickLogin" class="to_choose">
                    <router-link to="/login" class="lr" id="loginBtn">登录</router-link>
                </div>
                <ul><li v-for="err in errors" v-text="err"></li></ul>
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
                    LoginPwd: '',
                    NickName: '',
                    EmailAddr: ''
                },
                LoginName: '',
                LoginPwd: '',
                NickName: '',
                EmailAddr: '',
                regCfnPwd:''
            }
        },
        vuerify: {
            LoginName:'required',
            LoginPwd: {
                test: /\w{4,}/,
                message: '至少四位字符'
            },
            regCfnPwd:{
                test (val) {
                    return val === this.LoginPwd
                },
                message: '密码输入不一致'
            }
        },
        computed: {
            errors () {
                return this.$vuerify.$errors
            }
        },
        methods: {
            handleSubmit() {
                if (this.$vuerify.check()) {
                    var registerUser = this.user;
                    userinfosService.registerUser({
                        registerUser, success: (ret)=> {
                            alert(JSON.stringify(ret));
                        }
                    })
                }
            }
        }
    }
</script>
