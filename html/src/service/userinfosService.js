import config from "../compoment/utls/config";
import ajax from "ajax";

//url前缀
const urlPre = config.host + "/userinfosApi";

//获取会员总数
export function getCount(callback,errCallback) {
    let url = urlPre + "/count";
    ajax().get(url).then((response, xhr)=> {
        callback && callback(response);
    }).catch(function (response, xhr) {
        errCallback && errCallback(response);
    });
}