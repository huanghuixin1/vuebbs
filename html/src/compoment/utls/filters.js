import Vue from "vue";
import * as timeUtls  from "./timeUtls";


(function init() {
    Vue.filter("tamp2span", function (timetamp) {
        return timeUtls.getTimespanDesc(timetamp);
    });

    //处理默认头像的问题
    Vue.filter("headerimg", function (header) {
        if(!header){
            return "http://pic.cnblogs.com/face/758427/20161010214740.png";
        }else{
            return header;
        }
    });
}());