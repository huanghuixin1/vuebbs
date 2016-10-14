//公共的utls方法
import * as userinfoService from "../../service/userinfosService";

const CurrentUserKey = "user_currentuser"; //用户缓存的key
const SessionIdKey = "vuebbs_sessionid"; //cookie中sessionId的key


export function setCurrentUser(user) {
    sessionStorage.setItem(CurrentUserKey, JSON.stringify(user));
}

export function getCurrentUser({ success }) {
    let strUser = sessionStorage.getItem(CurrentUserKey);

    //如果获取不到
    if (strUser === null || strUser === "") {
        //如果cookie里有sessionId的值
        if (document.cookie[SessionIdKey]) {
            // 拿到这个sessionId 去服务器获取一次用户信息
            userinfoService.getCurrentUser({
                success: function(data) {
                    if (data.Status === 0) {
                        this.setCurrentUser(data.Data);
                        success(data.Data);
                    } else {
                        success(null);
                    }
                }
            });
        } else {
            success(null);
        }
    }

    success(JSON.parse(strUser));
}
