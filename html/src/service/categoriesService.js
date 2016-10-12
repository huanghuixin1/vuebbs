import config from "../compoment/utls/config";
import ajax from "ajax";

//url前缀
const urlPre = config.host + "/categoriesApi";

/**
 * 根据cid获取板块名称
 * @param cid
 */
export function getNameById({cid, success, err}) {
    let url = urlPre + "/getNameById?cid=" + cid;
    ajax().get(url).then((response, xhr)=> {
        success && success(response);
    }).catch(function (response, xhr) {
        err && err(response);
    });
}

//获取板块列表和帖子数量
export function getCategoryListName({success, err}) {
    let url = urlPre + "/getListName";
    ajax().get(url).then((rep)=> {
        success && success(rep);
    }).catch((rep)=> {
        err && err(response);
    });
}