import config from "../compoment/utls/config";
import * as ajaxUtls from "./ajaxUtls";

//url前缀
const urlPre = config.host + "/categoriesApi";

/**
 * 根据cid获取板块名称
 */
export function getNameById({cid, success, err}) {
    let url = urlPre + "/getNameById?cid=" + cid;
    ajaxUtls.sendGet(url, success, err);
}

//获取板块列表和帖子数量
export function getCategoryListName({success, err}) {
    let url = urlPre + "/getListName";
    ajaxUtls.sendGet(url, success, err);
}