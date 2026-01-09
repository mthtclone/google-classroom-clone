<!-- components/Sidebar.vue -->
<template>
  <transition name="slide">
    <aside 
        :style="{ display: open ? 'block' : 'none' }"    
        class="sidebar"  
        v-if="open"
      >
        <nav class="menu">
            <ul>
                <li> 
                  <a href="/">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-house-icon lucide-house"><path d="M15 21v-8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v8"/><path d="M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/></svg>
                    Home 
                  </a>
                </li>
                <li>
                  <a href="/academic-calender.html">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-calendar-icon lucide-calendar"><path d="M8 2v4"/><path d="M16 2v4"/><rect width="18" height="18" x="3" y="4" rx="2"/><path d="M3 10h18"/></svg> 
                    Calender 
                  </a>
                </li>
            </ul>

            <ul>
               <li class="enrolled">
                <div class="left-group">
                  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-icon customizable lucide-graduation-cap-icon lucide-graduation-cap lucide-icon customizable"><path d="M21.42 10.922a1 1 0 0 0-.019-1.838L12.83 5.18a2 2 0 0 0-1.66 0L2.6 9.08a1 1 0 0 0 0 1.832l8.57 3.908a2 2 0 0 0 1.66 0z"></path><path d="M22 10v6"></path><path d="M6 12.5V16a6 3 0 0 0 12 0v-3.5"></path></svg>
                  Enrolled 
                 </div>
                 <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-icon customizable lucide-chevron-down-icon lucide-chevron-down lucide-icon customizable"><path d="m6 9 6 6 6-6"></path></svg>
               </li>
            </ul>
            <ul>
                <li 
                    v-for="(course, index) in enrolled"
                    :key="index"
                    class="enrolled-item"
                >   
                    {{ course }}
                </li>
            </ul>

            <ul>
                <li>
                  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-archive-icon lucide-archive"><rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8"/><path d="M10 12h4"/></svg> 
                  Archived Classes 
                </li>
                <li>
                  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-settings-icon lucide-settings"><path d="M9.671 4.136a2.34 2.34 0 0 1 4.659 0 2.34 2.34 0 0 0 3.319 1.915 2.34 2.34 0 0 1 2.33 4.033 2.34 2.34 0 0 0 0 3.831 2.34 2.34 0 0 1-2.33 4.033 2.34 2.34 0 0 0-3.319 1.915 2.34 2.34 0 0 1-4.659 0 2.34 2.34 0 0 0-3.32-1.915 2.34 2.34 0 0 1-2.33-4.033 2.34 2.34 0 0 0 0-3.831A2.34 2.34 0 0 1 6.35 6.051a2.34 2.34 0 0 0 3.319-1.915"/><circle cx="12" cy="12" r="3"/></svg> 
                  Settings  
                </li>
            </ul>
        </nav>
    </aside>
  </transition>
</template>

<script>
export default {
    name: "Sidebar",

    props: {
        open: {
          type: Boolean,
          default: true
        },
        enrolled: {
            type: Array,
            default: () => []
        }
    }
};
</script>

<style>
@layer sidebar {
  .sidebar {
    width: 250px;
    height: 100vh;
    background: var(--primary-color);
    padding: 20px;
    box-sizing: border-box;
    overflow-y: auto;
  }

  .menu,
  ul {
    display: flex;
    flex-direction: column;
    list-style: none;
    padding: 2px;
    margin: 0 0 0px 0;
  }

  .menu li {
    font-size: 0.9rem;
    gap: 15px;
  }

  .section-title {
    margin-top: 20px;
    margin-bottom: 10px;
    font-weight: 700;
    font-size: 0.9rem;
    font-family: var(--secondary-font);
    color: #555;
  }

  .section-title > li {
    text-overflow: ellipsis;
    overflow: hidden;
  }

  li {
    display: flex; 
    align-items: center;
    gap: 8px;
    padding: 10px;
    border-radius: 6px;
    cursor: pointer;
    font-family: var(--secondary-font);
  }

  li > svg {
    flex-shrink: 0;
    width: 19px;
    height: 19px;
  }

  li > a {
    display: flex;
    align-items: center;
    gap: 8px;
    text-decoration: none;
    color: inherit;
    width: 100%;
  }

  li:hover {
    background: #e0e0e0;
    padding: 10px;
  }

  li:active {
    background-color: #a8a8a8;
  }

  li.enrolled {
    justify-content: space-between;
  }

  li.enrolled > .left-group {
    display: flex;               
    align-items: center;        
    gap: 6px;                    
  }

  .enrolled-item {
    font-weight: 600;
  }

  .slide-enter-from, .slide-leave-to {
    transform: translateX(-100%);
    opacity: 0;
  }

  .slide-enter-to, .slide-leave-from {
    transform: translateX(0);
    opacity: 1;
  }

  .slide-enter-active, .slide-leave-active {
    transition: transform 0.3s ease, opacity 0.3s ease;
  }
}
</style>