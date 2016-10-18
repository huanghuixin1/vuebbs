"use strict";
import index_header from "../../../compoment/header_footer/index_header/modules";
import Vuex from "vuex";
import Vue from "vue";

import * as actionHeader from "../../../compoment/header_footer/index_header/action";
import * as gettersHeader from "../../../compoment/header_footer/index_header/getters";

Vue.use(Vuex);

let actions = Object.assign({}, actionHeader);
let getters = Object.assign({}, gettersHeader);

export default new Vuex.Store({
    actions,
    getters,
    modules: {
        index_header
    }
});