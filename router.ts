import { createRouter, createWebHistory } from 'vue-router';
import MainPanel from './components/MainPanee.vue';
import ClassPage from './components/ClassPage.vue';

const routes = [
    {
        path: '/',
        name: 'home',
        component: MainPanel
    },
    {
        path: '/class/:id',
        name: 'class',
        component: ClassPage,
        props: true
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;