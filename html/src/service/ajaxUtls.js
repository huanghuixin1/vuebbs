import ajax from "ajax";

/**
 * 发送get请求
 * @param url url地址
 * @param success 成功回调
 * @param err 失败回调
 */
export function sendGet(url, success, err) {
    ajax().get(url).then((response, xhr)=> {
        success && success(response);
    }).catch(function (response, xhr) {
        err && err(response);
    });
}