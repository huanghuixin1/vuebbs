"use strict";
import index_header from "../../../compoment/header_footer/index_header/modules";
import Vuex from "vuex";
import Vue from "vue";

import * as actionHeader from "../../../compoment/header_footer/index_header/action";


Vue.use(Vuex);

let actions = Object.assign({}, actionHeader);

export default new Vuex.Store({
    actions,
    getters:{},
    modules: {
        index_header
    }
});