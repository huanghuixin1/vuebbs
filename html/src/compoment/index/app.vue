<template>
    <div id="app">
        <v_header></v_header>
        <!-- 路由外链 -->
        <article class="main">
            <router-view></router-view>
        </article>

        <v_footer></v_footer>
    </div>
</template>

<script>
    import v_header from '../header_footer/index_header/index_header.vue';
    import v_footer from '../header_footer/index_footer/index_footer.vue';

    import * as  headerActions from "../header_footer/index_header/action"
    import store from "./store";
    export default{
        components: {
            v_header,
            v_footer
        },
        methods: {
            changTitle(title){
                headerActions.SetTitle(this.$store, title);
            }
        },
        store
    }
</script>