<template>
    <div class="main-panel">
      <h2 class="panel-heading">Classroom</h2>
  
      <div class="cla-list">
        <div 
            class="cla-card" 
            v-for="(course, index) in class" 
            :key="index"
        >
            <div 
                class="cla-image"
                :style="{ 
                    backgroundColor: course.color, 
                    boxShadow: 'inset 0 0 100px rgba(0, 0, 0, 0.2)'
                    }"
            >
                <div class="cla-metadata">
                    <a 
                        class="cla-name" 
                        @click.prevent="goToClass(course)" 
                        :data-fullname="course.name"
                        @mouseenter="startToolTipTimer(index, 'name')"
                        @mouseleave="cancelToolTipTimer(index, 'name')"
                        :class="{ 'show-tooltip': tooltipVisible[index].name }"
                    >
                        {{ course.name }}
                    </a>
                    <p class="cla-section">{{ course.section }}</p>
                    <p class="cla-teacher">{{ course.teacher }}</p>
                </div>
                <button 
                    class="cla-button"
                    @click.stop="openCategoryModal(index)"
                    @mouseenter="startToolTipTimer(index, 'metadata')"
                    @mouseleave="cancelToolTipTimer(index, 'metadata')"
                    :class="{ 'show-tooltip': tooltipVisible[index].metadata }"
                    :data-fullname="'View Metadata'"
                    > <!-- Change to View Metadata-->
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-icon customizable lucide-info-icon lucide-info lucide-icon customizable"><circle cx="12" cy="12" r="10"></circle><path d="M12 16v-4"></path><path d="M12 8h.01"></path></svg>
                </button>
            </div>
            <span class="cla-category" v-if="course.category">[{{ course.category }}]</span>
            
            <div class="cla-extra-info">
                <p v-if="course.noticeCount > 0">
                    <span class="notice-count">{{ course.noticeCount }}</span>
                    <span class="notice-text"> announcement{{ course.noticeCount > 1 ? 's' : '' }}</span>
                </p>
                <p v-else>No announcement.</p>

                <p v-if="course.assignmentCount > 0">
                    <span class="notice-count">{{ course.assignmentCount }}</span>
                    <span class="notice-text"> pending assignment{{ course.assignmentCount > 1 ? 's' : '' }}</span>
                </p>
                <p v-else>No pending assignments.</p>
            </div>

            <div class="cla-footer">
                <button 
                    class="icon-btn" 
                    @click.stop="toggleNotifications(index)"
                    @mouseenter="startToolTipTimer(index, 'notifications')"
                    @mouseleave="cancelToolTipTimer(index, 'notifications')"
                    :class="{ 'show-tooltip': tooltipVisible[index].notifications }"
                    :data-fullname="'Toggle Notifications'"
                >
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-bell-icon lucide-bell"><path d="M10.268 21a2 2 0 0 0 3.464 0"/><path d="M3.262 15.326A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673C19.41 13.956 18 12.499 18 8A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.738 7.326"/></svg>
                </button>

                <button  
                    class="icon-btn" 
                    @click.stop="openCategoryModal(index)"
                    @mouseenter="startToolTipTimer(index, 'category')"
                    @mouseleave="cancelToolTipTimer(index, 'category')"
                    :class="{ 'show-tooltip': tooltipVisible[index].category }"
                    :data-fullname="'Categorize Class'"
                >
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-tag-icon lucide-tag"><path d="M12.586 2.586A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8.704 8.704a2.426 2.426 0 0 0 3.42 0l6.58-6.58a2.426 2.426 0 0 0 0-3.42z"/><circle cx="7.5" cy="7.5" r=".5" fill="currentColor"/></svg>
                </button>

                <button 
                    class="icon-btn" 
                    v-if="course.category" 
                    @click.stop="pinClass(index)"
                >
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-pin-icon lucide-pin"><path d="M12 17v5"/><path d="M9 10.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24V16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-.76a2 2 0 0 0-1.11-1.79l-1.78-.9A2 2 0 0 1 15 10.76V7a1 1 0 0 1 1-1 2 2 0 0 0 0-4H8a2 2 0 0 0 0 4 1 1 0 0 1 1 1z"/></svg>
                </button>
            </div>
        </div>
      </div>
    </div>
</template>

<script>
import classData from "/data/metadata.json"
import details from "/data/slug.json"
import assignmentData from '/data/assignment.json'

export default {
  name: "MainPanel",
  data() {
    return {
    tooltipVisible: [],
    tooltipTimers: [],
    class: [],
    };
  },
  created() {
    this.class = classData.map((course, index) => {
        const storedColor = localStorage.getItem(`course-color-${index}`);
        const color = storedColor || this.getRandomColor();
        if (!storedColor) 
            localStorage.setItem(`course-color-${index}`, color);

        const slug = course.name.toLowerCase().replace(/\s+/g, '-').replace(/[^a-z0-9\-]/g, '');
        const noticeCount = details[slug] ? details[slug].length : 0;

        const assignmentCount = assignmentData.assignments.filter(a => a.courseSlug === slug && !a.completed).length;

        return {
            ...course,
            color,
            slug,
            noticeCount,
            assignmentCount
        };
    });

    this.tooltipVisible = this.class.map(() => ({
        name: false,
        notifications: false,
        category: false,
        metadata: false
    }));
    this.tooltipTimers = this.class.map(() => ({
        name: null,
        notifications: null,
        category: null,
        metadata: null
    }));
  },
  methods: {
    startToolTipTimer(index, type) {
        clearTimeout(this.tooltipTimers[index][type]);
        this.tooltipTimers[index][type] = setTimeout(() => {
            this.tooltipVisible[index][type] = true;
        }, 800);
    },
    cancelToolTipTimer(index, type) {
        clearTimeout(this.tooltipTimers[index][type]);
        this.tooltipVisible[index][type] = false;
    },
    openCategoryModal(index) {
      // Placeholder: we’ll add functionality later
      console.log("Categorize class:", this.class[index].name);
    },
    getRandomColor() {
      const r = Math.floor(Math.random() * 127 + 127);
      const g = Math.floor(Math.random() * 127 + 127);
      const b = Math.floor(Math.random() * 127 + 127);
      return `rgb(${r}, ${g}, ${b})`;
    },
    goToClass(course) {
        const classPage = course.name.toLowerCase().replace(/\s+/g, '-').replace(/[^a-z0-9\-]/g, '');
        window.location.href = `src/c/${classPage}.html`;
    },
    toggleNotifications(index) {
        console.log("Toggled notifications for:", this.class[index].name);
    },
    pinClass(index) {
        console.log("Pinned class:", this.class[index].name);
    }
  }
};
</script>

<style>
@layer panel {
    .main-panel {
        flex: 1;
        padding: 20px;
        overflow-y: auto;
        font-family: var(--secondary-font);
        transition: flex-basis 0.4s ease-in-out;
    }

    .panel-heading {
        padding: 0.4rem;
    }

    .cla-list {
        display: flex;
        flex-wrap: wrap;
        gap: 20px;
    }

    .cla-card {
        width: 320px;
        height: 320px;
        border-radius: 10px;
        border: var(--borderline);
        overflow: hidden;
        text-align: center;
        background-color: #fff;
        display: flex;
        flex-direction: column;
        align-items: center;
        position: relative;
        overflow: visible;
    }

    .cla-card:hover {
        box-shadow: 0 2px 6px rgba(0,0,0,0.15);
    }

    .cla-metadata {
        display: flex;
        flex-direction: column;
        gap: 2px;
        max-width: 240px;
        overflow-y: auto;
    }

    .cla-metadata > * {
        white-space: nowrap;       
        overflow: hidden;           
        text-overflow: ellipsis;    
    }

    .cla-name {
        font-size: 1.3rem;
        font-weight: 600;
        color: #fff;
        margin: 0;
        text-decoration: none;
        cursor: pointer;
    }

    .cla-name::after {
        content: attr(data-fullname);
        position: absolute;
        left: 12%;
        bottom: 80%;
        background-color: #333333e1;
        color: #fff;
        padding: 4px 8px;
        border-radius: 4px;
        font-size: 15px;
        white-space: nowrap;
        opacity: 0;
        pointer-events: none;
        transition: opacity 0.2s;
        z-index: 999;
    }

    .cla-name.show-tooltip::after {
        opacity: 1;
    }

    .cla-section {
        text-align: left;
        font-size: 0.9rem;
        font-weight: 500;
        opacity: 1;
    }

    .cla-teacher {
        font-size: 0.85rem;
        opacity: 1;    
        font-weight: 500; 
        margin-top: auto; 
        text-align: left;
    }

    .cla-name:is(:hover, :active, :focus-within) {
        text-decoration: underline;
        text-decoration-thickness: 2px;
    }

    .cla-image {
        position: relative;
        width: 100%;
        height: 110px;
        padding: 15px;
        box-sizing: border-box;
        border-radius: 10px 10px 0 0;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: flex-start; 
        color: #fff;
        box-shadow: var(--inset-shadow-1);
        overflow: visible;
    }

    .cla-button {
        position: absolute;
        top: 8px;
        right: 8px;
        padding: 6px 10px;
        border: none;
        cursor: pointer;
        background-color: transparent;
        object-fit: fill;
        transition: background-color 0.2s ease, 
                box-shadow 0.2s ease,
                transform 0.2s ease;
    }

    .cla-button:hover {
        transform: scale(1.04);
    }

    .cla-button::after {
        content: attr(data-fullname);
        position: absolute;
        bottom: 100%;
        right: 50%;
        transform: translateX(50%);
        background-color: #333333e1;
        color: #fff;
        padding: 4px 8px;
        border-radius: 4px;
        font-size: 12px;
        white-space: nowrap;
        opacity: 0;
        pointer-events: none;
        transition: opacity 0.2s;
        z-index: 999;
    }

    .cla-button.show-tooltip::after {
        opacity: 1;
    }

    .cla-extra-info {
        width: 90%;
        margin: 10px auto 0 auto;
        padding: 10px;
        border-radius: 6px;
        font-size: 0.85rem;
        text-align: left;
        max-height: 80px;
        overflow-y: auto;
    }

    .notice-count {
        font-weight: 700;
        font-size: 1.1rem;
        color: #111; 
    }

    .notice-text {
        font-weight: 500;
        font-size: 0.95rem;
        color: #555; 
    }

    .cla-footer {
        margin-top: auto;
        width: 90%;
        display: flex;
        justify-content: flex-start;
        overflow: visible;
        border-top: var(--borderline);
        gap: 10px;
        padding: 10px 14px 14px 14px;
    }

    .icon-btn {
        width: 32px;
        height: 32px;
        border: none;
        cursor: pointer;
        border-radius: 50%;
        background-color: #ffffff;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 1rem;
        transition: background-color 0.2s ease,
                    transform 0.15s ease,
                    box-shadow 0.2s ease;
    }

    .icon-btn:hover {
        transform: scale(1.08);
    }

    .icon-btn::after {
        content: attr(data-fullname);
        position: absolute;
        bottom: 120%;       
        left: 80%;
        transform: translateX(-50%);
        background-color: #333333e1;
        color: #fff;
        padding: 4px 8px;
        border-radius: 4px;
        font-size: 12px;
        white-space: nowrap;
        opacity: 0;
        pointer-events: none;
        transition: opacity 0.2s;
        z-index: 999;
    }

    .icon-btn.show-tooltip::after {
        opacity: 1;
    }
}
</style>