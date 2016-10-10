"use strict";
import articleList from "./modules";
import Vuex from "vuex";
import Vue from "vue";

import * as actions from "./action";
import * as getters from "./getters";

Vue.use(Vuex);

export default new Vuex.Store({
    actions,
    getters,
    modules: {
        articleList
    }
});