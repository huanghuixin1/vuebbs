import Vue from "vue";
import * as timeUtls  from "./timeUtls";


(function init() {
    Vue.filter("tamp2span", function (timetamp) {
        return timeUtls.getTimespanDesc(timetamp);
    });
}());