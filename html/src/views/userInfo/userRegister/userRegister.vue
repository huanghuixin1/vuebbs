<template>
    <div class="user_login">
        <div class="logo_dz">
            <img alt="logo" src="http://att.discuz.net/data/attachment/common/26/common_36_i94wrTB4.jpg">
            <h2>Discuz! 官方站</h2>
        </div>
        <div class="login_set">
            <form id="loginBox" class="login_set" @submit.prevent="registSubmit">
                <p>
                    <input type="text" v-model="user.LoginName" placeholder="请输入帐户">
                </p>
                <p>
                    <input type="password" v-model="user.LoginPwd"  placeholder="请输入密码">
                </p>
                <p>
                    <input type="password"  v-model="user.regCfnPwd" placeholder="请确认密码">
                </p>
                <p>
                    <input type="text" v-model="user.NickName"  placeholder="请输入昵称">
                </p>
                <p>
                    <input type="text"  v-model="user.EmailAddr"  placeholder="请输入邮箱">
                </p>
                <div class="log_bar">

                    <input id="registerBtn" type="submit" value="注册" class="lb_lq_btn"></div>
                <div id="toQuickLogin" class="to_choose">
                    <router-link to="/login" class="lr" id="loginBtn">登录</router-link>
                </div>
                <!--<ul><li v-for="err in errors" v-text="err"></li></ul>-->
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
                    EmailAddr: '',
                    regCfnPwd:''
                }
            }
        },
        methods: {
            registSubmit() {
                var user = this.user;
                userinfosService.registerUser({
                    user, success: (ret)=> {
                        if(ret.Status=='2'){
                            alert("参数错误");
                        }else if(ret.Status=='100001'){
                            alert("用户名已存在");
                        }else{
                            window.history.go(-1);
                        }
                    }
                })
            }
        }
    }
</script>
