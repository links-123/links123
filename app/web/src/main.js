import Vue from 'vue'
import App from './App.vue'
import store from './store'

import VueClipboards from "vue-clipboards";
import VueJSModal from 'vue-js-modal'
import VueRouter from 'vue-router';
import Toasted from 'vue-toasted';
import {routes} from './router/routes';

Vue.use(VueClipboards);
Vue.use(Toasted);
Vue.use(VueJSModal);

// Router
Vue.use(VueRouter);
const router = new VueRouter({
    routes,
    linkActiveClass: 'open active',
    scrollBehavior: () => ({y: 0}),
    mode: 'hash'
});

var truncationFilter = function (text, length, clamp) {
    clamp = clamp || '...';
    var node = document.createElement('div');
    node.innerHTML = text;
    var content = node.textContent;
    return content.length > length ? content.slice(0, length) + clamp : content;
};

Vue.filter('truncate', truncationFilter);

new Vue({
    el: '#app',
    router,
    store,
    render: h => h(App)
});
