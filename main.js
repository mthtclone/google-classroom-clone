// main.js
import { createApp } from 'https://unpkg.com/vue@3/dist/vue.esm-browser.js';
import Navbar from '/src/components/NavBar.vue';
import AsideBar from '/src/components/Sidebar.vue';
import MainPanel from '/src/components/MainPanel.vue';
import AssignmentPanel from '/src/components/AssignmentPanel.vue'

import NoticeboardComponent from '/src/components/NoticeBoard.vue';

const app = createApp({
    data() {
        return {
            isSidebarOpen: true
        };
    },
    methods: {
        toggleSidebar() {
            this.isSidebarOpen = !this.isSidebarOpen;
            // console.log("Sidebar Triggered.")
        }
    },
    // watch: {
    //     isSidebarOpen(newVal) {
    //         console.log("Sidebar state:", newVal);
    //     }
    // }
});

app.component('navbar-component', Navbar);
app.component('asidebar-component', AsideBar);
app.component('main-panel', MainPanel);

app.component('notice-board', NoticeboardComponent);
app.component('assignment-panel', AssignmentPanel);
app.mount('#app'); 
