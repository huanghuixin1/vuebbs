"use strict";
import articleList from "./modules";
import index_header from "../../header_footer/index_header/modules";
import Vuex from "vuex";
import Vue from "vue";

import * as actionsList from "./action";
import * as gettersList from "./getters";
import * as actionHeader from "../../header_footer/index_header/action";
import * as gettersHeader from "../../header_footer/index_header/getters";

Vue.use(Vuex);

let actions = Object.assign({}, actionsList, actionHeader);
let getters = Object.assign({}, gettersList, gettersHeader);

export default new Vuex.Store({
    actions,
    getters,
    modules: {
        articleList,
        index_header
    }
});