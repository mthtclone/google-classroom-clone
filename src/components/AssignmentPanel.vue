<template>
<div class="assignment-panel">
    <div class="left-panel">
    <h3>Pending Assignments</h3>
    <ul>
        <li 
            v-for="assignment in pendingAssignments" 
            :key="assignment.id" 
            @click="selectAssignment(assignment)"
            :class="{ selected: selectedAssignment && selectedAssignment.id === assignment.id }"
        >
        {{ assignment.title }}
        </li>
    </ul>
    </div>

    <div class="right-panel">
        <div class="tabs">
            <button
                :class="{ active: activeTab === 'overview' }"
                @click="activeTab = 'overview'"
            >
                Overview
            </button>

            <button
                :class="{ active: activeTab === 'resources' }"
                @click="activeTab = 'resources'"
            >
                Resources
            </button>

            <button
                :class="{ active: activeTab === 'works' }"
                @click="activeTab = 'works'"
            >
                Works
            </button>

            <button
                :class="{ active: activeTab === 'people' }"
                @click="activeTab = 'people'"
            >
                People
            </button>
        </div>

        <div
            v-if="activeTab === 'overview'"
            class="overview-section"
        >
            <div class="overview-card full-width">
                <h2 class="course-title">{{ overview.classTitle }}</h2>
                <p class="course-description">{{ overview.classDescription }}</p>
            </div>

            <div class="overview-card full-width teacher-section">
                <h3 class="section-title">Instructor</h3>

                <div class="teacher-card">
                    <div class="avatar" :style="{ backgroundImage: `url(${overview.teacher.avatar})` }"></div>
                
                    <div class="teacher-info">
                        <h4>{{ overview.teacher.name }}</h4>
                        <p class="bio">{{ overview.teacher.bio }}</p>

                        <div class="contact-info">
                            <p><strong>Email:</strong>{{ overview.teacher.email }}</p>
                            <p><strong>Office Hours:</strong>{{ overview.teacher.officeHours }}</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="overview-card full-width">
                <h3 class="section-title">Syllabus</h3>
                <p class="syllabus-text" v-html="overview.syllabus.replace(/\\n/g, '<br>')"></p>
            </div>

            <div class="overview-card full-width">
                <h3 class="section-title">Class Schedule</h3>
                <ul class="schedule-list">
                    <li v-for="item in overview.schedule" :key="item.week">
                        <strong>Week {{ item.week }}</strong> {{ item.topic }}
                    </li>
                </ul>
            </div>
        </div>

        <div
            v-if="activeTab === 'resources'"
            class="resource-section"
        >
            <a 
                v-for="resource in resources"
                :key="resource.id"
                :href="resource.link"
                target="_blank"
                class="resource-card"
            >
                <div class="card-header">
                    <div class="icon-circle resource-icon">
                        <svg width="24" height="24" viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg">
                            <path d="M141.18,56.938l30.484,33.531v157.594c0,2.563,1.422,4.938,3.688,6.141c2.281,1.203,5.031,1.063,7.156-0.391
                                l36.406-24.656l36.391,24.656c2.141,1.453,4.891,1.594,7.172,0.391c2.25-1.203,3.688-3.578,3.688-6.141V90.469l-30.5-33.531H141.18
                                z"/>
                            <path d="M436.008,93.344l-25.875-62.563c9.188-0.563,14.719-8.156,14.719-14.078C424.852,7.469,417.383,0,408.164,0
                                H109.477C92.086,0,76.195,7.094,64.836,18.5C53.43,29.859,46.32,45.75,46.336,63.125V470.75c0,22.781,18.469,41.25,41.25,41.25
                                h343.359c19.188,0,34.719-15.547,34.719-34.734V127.578C465.664,110.125,452.789,95.844,436.008,93.344z M290.664,92.844v155.219
                                c0,11.672-6.406,22.328-16.719,27.797c-4.531,2.391-9.625,3.672-14.75,3.672c-6.313,0-12.422-1.875-17.641-5.438l-22.641-15.344
                                l-22.656,15.344c-5.219,3.563-11.313,5.438-17.641,5.438c-5.109,0-10.219-1.281-14.75-3.688
                                c-10.297-5.453-16.703-16.109-16.703-27.781V99.938l-6.469-7.094h-31.219c-8.266,0-15.594-3.313-21.016-8.703
                                c-5.406-5.453-8.719-12.766-8.719-21.016s3.313-15.578,8.719-21c5.422-5.406,12.75-8.719,21.016-8.719H383.57l26.688,59.438
                                H290.664z"/>
                        </svg>
                        </div>
                        <div class="card-content">
                            <h5>{{ resource.title }}</h5>
                            <span class="view-link">View Resource</span>
                        </div>
                </div>
            </a>
        </div>

        <div
            v-if="activeTab === 'works'"
            class="work-section"
        >

            <div
                v-if="selectedAssignment"
                class="view-assignment-card polished"
            >
                <div class="header-row">
                    <h3>{{ selectedAssignment.title }}</h3>
                    <div class="badge-row">
                        <span :class="['status-badge', selectedAssignment.completed ? 'completed' : 'pending']">
                            {{ selectedAssignment.completed ? 'Completed' : 'Pending' }}
                        </span>
                        <span v-if="dueBadgeText(selectedAssignment)" :class="['status-badge', 'due-badge']">
                            {{ dueBadgeText(selectedAssignment) }}
                        </span>
                    </div>
                </div>
                <p class="description">{{ selectedAssignment.description }}</p>
                <div class="meta-row">
                    <span class="due-date">Due: {{ selectedAssignment.dueDate }}</span>
                </div>
            </div>
            <a
                v-for="assignment in assignments"
                :key="assignment.id"
                class="assignment-card horizontal-card"
                :href="`http://localhost:5173/src/c/3d-graphic-and-animation.html/assignment.html?id=${assignment.id}`"
              >   
                <div class="left-info">
                    <div class="icon-circle assignment-icon" :style="{ backgroundColor: assignment.completed ? '#34a853' : '#ffb347' }">
                        <svg width="24" height="24" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
                            <title>assignment-text</title>
                            <g id="Layer_2" data-name="Layer 2">
                                <g id="invisible_box" data-name="invisible box">
                                <rect width="48" height="48" fill="none"/>
                                </g>
                                <g id="icons_Q2" data-name="icons Q2">
                                <g>
                                    <path d="M16,16a2,2,0,0,0,0,4H32a2,2,0,0,0,0-4Z"/>
                                    <path d="M32,24H16a2,2,0,0,0,0,4H32a2,2,0,0,0,0-4Z"/>
                                    <path d="M24,32H16a2,2,0,0,0,0,4h8a2,2,0,0,0,0-4Z"/>
                                    <path d="M40,6H36V4a2,2,0,0,0-2-2H14a2,2,0,0,0-2,2V6H8A2,2,0,0,0,6,8V44a2,2,0,0,0,2,2H40a2,2,0,0,0,2-2V8A2,2,0,0,0,40,6ZM38,42H10V10h6V6H32v4h6Z"/>
                                </g>
                                </g>
                            </g>
                            </svg>
                        </div>
                        <div class="text-info">
                            <div class="title-column">
                                <h3 class="title">{{ assignment.title }}</h3>
                                <span class="upload-date">{{ formatDate(assignment.uploadDate) }}</span>
                            </div>
                            <p class="description">{{ assignment.description }}</p>
                        </div>
                </div>

                <div class="right-info">
                    <span class="due-label">Due: {{ assignment.dueDate }}</span>
                    <div class="badge-row">
                        <span :class="['status-badge', assignment.completed ? 'completed' : 'pending']">
                            {{ assignment.completed ? 'Completed' : 'Pending' }}
                        </span>
                        <span v-if="dueBadgeText(assignment)" :class="['status-badge', 'due-badge']">
                            {{ dueBadgeText(assignment) }}
                        </span>
                    </div>
                </div>
            </a>

            <div
                v-if="assignments.length === 0"
                class="no-selection"
            >
                <p>No assignments have been posted yet.</p>
            </div>

            <!-- - Upload Area 
                    - Submission Status
                        - Grade Area 
            -->
        </div>

        <div
            v-if="activeTab === 'people'"
            class="people-section"
        >
            <div
                v-for="(peopleList, className) in peopleByClass"
                :key="className"
                class="class-group"
            >   
                <h4 class="class-name">{{ className }}</h4>

                <div class="role-group">
                    <h5 class="role-title">Teacher</h5>
                    <div 
                        v-for="person in peopleList.filter(p => p.role == 'Teacher')"
                        :key="person.id" class="person-card full-width"
                    >
                        <div class="icon-circle" :style="{ backgroundImage: `url(${person.avatar})`, backgroundSize: 'cover' }"></div>
                        <div class="person-info">
                            <h5>{{ person.name }}</h5>
                        </div>
                    </div>
                </div>

                <div class="role-group">
                    <h5 class="role-title">Students</h5>
                    <div 
                        v-for="person in peopleList.filter(p => p.role == 'Student')"
                        :key="person.id" class="person-card full-width"
                    >
                        <div class="icon-circle" :style="{ backgroundImage: `url(${person.avatar})`, backgroundSize: 'cover' }"></div>
                        <div class="person-info">
                            <h5>{{ person.name }}</h5>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</template>
  
<script>
import data from '/data/assignment.json';
import peopleData from '/data/people.json';
import overviewData from '/data/overview.json';

export default {
    name: "AssignmentPanel",
    data() {
        return {
            currentSlug: null,
            assignments: [],
            resources: [],
            people: peopleData.people,
            selectedAssignment: null,
            overview: [],
            activeTab: 'overview'
        };
    },
    computed: {
        pendingAssignments() {
            return this.assignments.filter(a => !a.completed);
        },
        peopleByClass() {
            const grouped = {};
            this.people.forEach(person => {
                person.classes.forEach(cls => {
                    if (!grouped[cls]) 
                        grouped[cls] = [];
                    grouped[cls].push(person);
                });
            });
            return grouped;
        }
    },
    created() {
        this.detectSlug();

        this.assignments = data.assignments.filter(a => a.courseSlug === this.currentSlug);
        this.resources = data.resources.filter(r => r.courseSlug === this.currentSlug);
        this.overview = overviewData.find(o => o.courseSlug === this.currentSlug);
    },  
    methods: {
        detectSlug() {
            const path = window.location.pathname;
            const filename = path.split('/').pop();
            this.currentSlug = filename.replace('.html', '');
        },
        selectAssignment(assignment) {
            this.selectedAssignment = assignment;
            this.activeTab = 'works';
        },
        formatDate(dateStr) {
            const options = { month: 'short', day: 'numeric' };
            const date = new Date(dateStr);
            return `${date.toLocaleDateString('en-US', options)}`;
        },
        dueBadgeText(assignment) {
            const today = new Date();
            const due = new Date(assignment.dueDate);
            today.setHours(0,0,0,0);
            due.setHours(0,0,0,0);
            
            
            if (due.getTime() === today.getTime()) {
                return 'Due Today';
            } else if (due < today) {
                return 'Overdue';
            } else {
                return '';
            }
        }
    },
};
</script>
  
<style>
@layer assignment-panel {
    .assignment-panel {
        display: flex;
        padding: 20px;
        gap: 1rem;
        font-family: var(--secondary-font);
    }

    .left-panel {
        width: 15%;
        padding: 1rem;
        background: #fafafa;
        border: 1px solid #dcdcdc;
        border-radius: 10px;
        height: 20vh;
        overflow-y: auto;
    }

    .left-panel h3 {
        margin-top: 0;
        font-size: 0.9rem;
        margin-bottom: 1rem;
    }

    .left-panel ul {
        list-style: none;
        padding: 0;
        margin: 0;
    }

    .left-panel li {
        padding: 0.6rem 0.4rem;
        cursor: pointer;
        border-radius: 6px;
        transition: background 0.2s;
        font-size: 0.75rem;
    }

    .left-panel li:hover {
        text-decoration: underline;
        background-color: transparent;
        box-shadow: none;
    }

    .left-panel li.selected {
        background-color: #d4e3ff;
    }

    .right-panel {
        width: 70%;
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
    }

    .tabs {
        display: flex;
    }

    .tabs button {
        padding: 0.6rem 1.2rem;
        background: #f5f5f5;
        border: none;
        cursor: pointer;
        transition: 0.2s;
    }

    .tabs button:first-child {
        border-top-left-radius: 12px;
        border-bottom-left-radius: 12px;
    }

    .tabs button:last-child {
        border-top-right-radius: 12px;
        border-bottom-right-radius: 12px;
    }

    .tabs button.active {
        background: #ffffff;
        border: 1px solid #909090;
    }

    .tabs button:hover {
        background: #ebebeb;
    }

    .overview-section {
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
    }

    .overview-card {
        background: #fff;
        padding: 1.3rem 1.4rem;
        border: 1px solid #e1e1e1;
        border-radius: 12px;
        transition: box-shadow 0.2s ease;
    }

    .course-title {
        margin: 0;
        font-size: 1.4rem;
        font-weight: 700;
    }

    .course-description {
        margin-top: 0.6rem;
        font-size: 0.95rem;
        color: #555;
        line-height: 1.5;
    }

    .teacher-card {
        display: flex;
        gap: 1rem;
    }

    .teacher-card .avatar {
        width: 100%;
        height: 200px;
        background-size: cover;
        background-position: center;
    }

    .teacher-info h4 {
        margin: 0;
        font-size: 1.1rem;
        font-weight: 600;
    }

    .teacher-info .bio {
        font-size: 0.9rem;
        color: #555;
        margin: 0.4rem 0;
        line-height: 1.4;
    }

    .contact-info p {
        margin: 0.1rem 0;
        font-size: 0.85rem;
        color: #444;
    }

    .syllabus-text {
        white-space: pre-wrap;
        font-size: 0.9rem;
        line-height: 1.55;
        color: #444;
    }

    .schedule-list {
        padding-left: 1rem;
        margin: 0.5rem 0;
    }

    .schedule-list li {
        margin-bottom: 0.4rem;
        font-size: 0.9rem;
        color: #444;
    }

    .assignment-card {
        background: white;
        padding: 0.6rem;
        border-radius: 12px;
        border: 1px solid #e0e0e0;
        margin-bottom: 1rem;
    }

    .view-assignment-card.polished {
        background: #fdfdfd;
        border: 1px solid #c1c1c1;
        border-radius: 12px;
        padding: 1.5rem 1.8rem;
        margin-bottom: 1rem;
        display: flex;
        flex-direction: column;
        gap: 1rem;
        transition: box-shadow 0.2s ease;
    }

    .view-assignment-card.polished:hover {
        box-shadow: 0 6px 18px rgba(0, 0, 0, 0.075);
    }

    .view-assignment-card.polished .header-row {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .view-assignment-card.polished h3 {
        margin: 0;
        font-size: 1.3rem;
        font-weight: 700;
        color: #222;
    }

    .view-assignment-card.polished .description {
        font-size: 0.9rem;
        color: #555;
        line-height: 1.5;
        white-space: pre-wrap;
    }

    .view-assignment-card.polished .meta-row {
        display: flex;
        justify-content: flex-end;
        font-size: 0.85rem;
        color: #777;
    }

    .assignment-card.horizontal-card {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0.7rem 1rem;
        background: white;
        transition: box-shadow 0.2s ease, transform 0.2s ease;
        margin-bottom: 0.7rem;
        text-decoration: none;
    }

    .assignment-card.horizontal-card:hover {
        background-color: #cfcfcf72;
    }

    .resources h4 {
        margin-bottom: 1rem;
    }

    .resource-section {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .resource-card {
        display: flex;
        align-items: center;
        gap: 0.8rem;
        padding: 1rem 1.2rem;
        border-radius: 12px;
        border: 1px solid #e0e0e0;
        background: #fff;
        text-decoration: none;
        transition: box-shadow 0.2s ease, transform 0.2s ease;
    }

    .resource-card:hover {
        background-color: #cfcfcf72;
    }

    .resource-card .card-header {
        display: flex;
        align-items: center;
        gap: 0.8rem;
    }

    .resource-card .card-content h5 {
        margin: 0 0 0.2rem 0;
        font-size: 1rem;
        font-weight: 600;
        color: #222;
    }

    .resource-card .view-link {
        font-size: 0.85rem;
        color: #1a73e8;
        font-weight: 500;
    }

    .no-selection {
        padding: 2rem;
        text-align: center;
        color: #666;
        opacity: 0.8;
    }

    .card-header {
        display: flex;
        align-items: center;
        gap: 0.7rem;
    }

    .left-info {
        display: flex;
        align-items: center;
        gap: 0.7rem;
        min-width: 60%;
    }

    .left-info svg {
        flex-shrink: 0;
        fill: #1a73e8;
    }

    .text-info .title {
        margin: 0;
        font-size: 1rem;
        font-weight: 600;
        color: var(--black);
    }

    .text-info .description {
        margin: 0;
        font-size: 0.8rem;
        max-width: 90%;
        color: #666;
        line-clamp: 2;
        overflow: hidden;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
    }

    .text-info .title-column {
        display: flex;
        flex-direction: column;
        margin-bottom: 0.4rem;
    }

    .text-info .upload-date {
        font-size: 0.75rem;
        color: #888;
        font-weight: 500;
    }

    .icon-circle {
        width: 40px;
        height: 40px;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-shrink: 0;
        margin: 0.5rem;
    }

    .right-info {
        display: flex;
        flex-direction: column;
        align-items: flex-end;
    }

    .due-label {
        font-size: 0.8rem;
        margin-bottom: 0.2rem;
        color: #444;
        margin-bottom: 0.4rem;
    }

    .badge-row {
        display: flex;
        gap: 0.5rem;
    }

    .status-badge {
        font-size: 0.7rem;
        padding: 0.3rem 0.6rem;
        border-radius: 20px;
        font-weight: 600;
        text-transform: uppercase;
        white-space: nowrap;
    }

    .status-badge.completed {
        background: #d4f5d9;
        color: #1d7a35;
    }

    .status-badge.pending {
        background: #fff3cd;
        color: #856404;
    }

    .status-badge.due-badge {
        background: #ffeaea;
        color: #ff4d4f;
    }

    .assignment-icon {
        background-color: #34a853; 
    }

    .resource-icon {
        background-color: #000000; 
    }

    .icon-circle svg {
        width: 20px;
        height: 20px;
        fill: #ffffff;
    }

    .people-section {
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
    }

    .class-group {
        display: flex;
        flex-direction: column;
        gap: 0.8rem;
    }

    .class-name {
        font-size: 1rem;
        font-weight: 600;
        color: #222;
    }

    .role-group {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        margin-bottom: 1rem;
    }

    .role-title {
        font-size: 0.95rem;
        font-weight: 600;
        color: #222;
        margin-bottom: 0.5rem;
    }

    .person-card.full-width {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        padding: 0.6rem 1rem;
        border-radius: 12px;
        border: 1px solid #e0e0e0;
        background: #fff;
        width: 100%;
        box-sizing: border-box;
    }

    .person-info h5 {
        margin: 0;
        font-size: 0.9rem;
        font-weight: 600;
    }

    .icon-circle {
        width: 40px;
        height: 40px;
        border-radius: 50%;
        flex-shrink: 0;
    }
}
</style>
  