// Dynamic Loading Modules

// Views
const Dashboard = resolve => {
    require.ensure(['../layouts/Dashboard.vue'],
        () => {
            resolve(require('../layouts/Dashboard.vue'));
        });
};

export const routes = [
    {
        path: '',
        name: 'home',
        components: {
            default: Dashboard
        }
    },
    {
        path: '/dashboard',
        component: {
            render(c) {
                return c('router-view')
            }
        },
        children: [{
            path: "",
            component: Dashboard
        }]
    }
];