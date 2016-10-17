import config from "../compoment/utls/config";
import * as ajaxUtls from "./ajaxUtls";

//url前缀
const urlPre = config.host + "/userinfosApi";

//获取当前用户信息
export function getCurrentUser({success}) {
    let url = urlPre + "/getCurrent";
    ajaxUtls.sendGet(url, success);
}

//获取会员总数
export function getCount(callback, errCallback) {
    let url = urlPre + "/count";
    ajaxUtls.sendGet(url, callback, errCallback);
}

/**
 * 用户注册
 * @param user 用户信息对象
 * @param success 成功回调
 * @param err 失败回调
 */
export function registerUser({user, success, err}) {
    let url = urlPre + "/regist";
    ajaxUtls.sendPost({url, data: user, success, err});
}
/**
 * 用户登录
 * @param user 用户信息对象
 * @param success 成功回调
 * @param err 失败回调
 */

export function loginUser({user, success, err}) {
    let url = urlPre + "/login";
    ajaxUtls.sendPost({url, data: user, success, err});
}
