<!-- components/Navbar.vue -->
<template>
    <nav class="gc-navbar">
        <button class="menu-btn" @click="$emit('toggle-menu')">
          <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#000000"><path d="M120-240v-80h720v80H120Zm0-200v-80h720v80H120Zm0-200v-80h720v80H120Z"/></svg>
        </button>
        
        <h1 class="brand">
          <a href="/" class="brand-link">
            <span class="icon">
            <svg height="32" width="32" xmlns="http://www.w3.org/2000/svg" viewBox="0 36 528 456">
                <linearGradient id="a" x1="50.003%" x2="50.003%" y1="1.385%" y2="101.042%">
                <stop offset="0" stop-color="#bf360c" stop-opacity=".2"/>
                <stop offset="1" stop-color="#bf360c" stop-opacity=".02"/>
                </linearGradient>
                <radialGradient id="b" cx="3.407%" cy="2.344%" gradientTransform="matrix(.86364 0 0 1 .005 0)" r="477.746%">
                <stop offset="0" stop-color="#fff" stop-opacity=".1"/>
                <stop offset="1" stop-color="#fff" stop-opacity="0"/>
                </radialGradient>
                <g fill="none">
                <path d="M48 84h432v360H48z" fill="#0f9d58"/>
                <path d="M360 276c14.9 0 27-12.1 27-27s-12.1-27-27-27-27 12.1-27 27 12.1 27 27 27zm0 18c-28.9 0-60 15.3-60 34.3V348h120v-19.7c0-19-31.1-34.3-60-34.3zm-192-18c14.9 0 27-12.1 27-27s-12.1-27-27-27-27 12.1-27 27 12.1 27 27 27zm0 18c-28.9 0-60 15.3-60 34.3V348h120v-19.7c0-19-31.1-34.3-60-34.3z" fill="#57bb8a"/>
                <path d="M264 252c19.9 0 36-16.1 36-36s-16.1-36-36-36-36 16.1-36 36 16.1 36 36 36zm0 24c-40.5 0-84 21.5-84 48v24h168v-24c0-26.5-43.5-48-84-48z" fill="#f7f7f7"/>
                <path d="M312 420h108v24H312z" fill="#f1f1f1"/>
                <path d="M492 36H36C16.1 36 0 52.1 0 72v384c0 19.9 16.1 36 36 36h456c19.9 0 36-16.1 36-36V72c0-19.9-16.1-36-36-36zm-12 408H48V84h432z" fill="#f4b400"/>
                <path d="M492 36H36C16.1 36 0 52.1 0 72v3c0-19.9 16.1-36 36-36h456c19.9 0 36 16.1 36 36v-3c0-19.9-16.1-36-36-36z" fill="#fff" opacity=".2"/>
                <path d="M492 489H36c-19.9 0-36-16.1-36-36v3c0 19.9 16.1 36 36 36h456c19.9 0 36-16.1 36-36v-3c0 19.9-16.1 36-36 36z" fill="#bf360c" opacity=".2"/>
                <path d="M419.8 444h-108l48 48h107.9z" fill="url(#a)"/>
                <path d="M48 81h432v3H48z" fill="#263238" opacity=".2"/>
                <path d="M48 444h432v3H48z" fill="#fff" opacity=".2"/>
                <path d="M492 36H36C16.1 36 0 52.1 0 72v384c0 19.9 16.1 36 36 36h456c19.9 0 36-16.1 36-36V72c0-19.9-16.1-36-36-36z" fill="url(#b)"/>
                </g>
            </svg>
            </span>
            Classroom 
          </a>
        </h1>

        <div class="class-display" v-if="currentClass.name">
          <div class="arrow-icon">
            <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#000000"><path d="M504-480 320-664l56-56 240 240-240 240-56-56 184-184Z"/></svg>
          </div>
          <div class="class-metadata">
            <div class="class-name">{{ currentClass.name }}</div>
            <div class="teacher-name">{{ currentClass.teacher }}</div>
          </div>
        </div>

        <div class="spacer"></div>

        <div class="user-btn-container">
          <div class="add-btn-wrapper" v-if="!currentClass.name">
            <transition name="fade-scale">
              <button class="add-btn" @click="dropdownOpen = !dropdownOpen">
                  <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#000000">
                      <path d="M440-440H200v-80h240v-240h80v240h240v80H520v240h-80v-240Z"/>
                  </svg>
              </button>
            </transition>

            <transition name="slide-fade">
              <div v-if="dropdownOpen" class="dropdown-overlay" @click="dropdownOpen = false">
                <div class="dropdown-menu" @click.stop>
                  <button class="dropdown-item" @click="createClass">Create Class</button>
                  <button class="dropdown-item" @click="joinClass">Join Class</button>
                </div>
              </div>
            </transition>
          </div>

          <div class="user-btn">
            <img src="https://avatar.iran.liara.run/public/30" alt="Avatar">
          </div>
        </div>
    </nav>
</template>

<script>
import metadata from '/data/metadata.json';

export default {
    name: "Navbar",
    data() {
      return {
        classes: metadata,
        currentSlug: null,
        dropdownOpen: false,
      }
    },
    computed: {
      currentClass() {
        if (!this.currentSlug) return {};

        return (
          this.classes.find(cls => cls.slug === this.currentSlug) || {}
        );
      }
    },
    mounted() {
      this.detectSlug();
    },
    methods: {
      detectSlug() {
        const path = window.location.pathname;
        const filename = path.split('/').pop() || '';
        let slugCandidate = filename.replace('.html', '').toLowerCase();
        slugCandidate = slugCandidate.replace(/[^a-z0-9-_]/g, '');
        const validSlug = this.classes.some(cls => cls.slug === slugCandidate);
        this.currentSlug = validSlug ? slugCandidate : null;
      },
      createClass() {
        alert("Create Class Clicked!");
        this.dropdownOpen = false;
      },
      joinClass() {
        alert("Join Class clicked!");
        this.dropdownOpen = false;
      }
    }
};
</script>

<style>
@layer reset, layout, theme;

@layer reset {
    .gc-navbar,
    .gc-navbar * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
}

@layer layout {
  .gc-navbar {
    width: 100%;
    height: 56px;
    display: flex;
    align-items: center;
    padding: 0 16px;
    gap: 1.5rem;
  }

  .menu-btn,
  .user-btn {
    background: none;
    border: none;
    cursor: pointer;
  }

  .menu-btn:hover {
    background-color: #ebebeb;
    border-radius: 10%;
  }

  .menu-btn:active {
    background-color: #a8a8a8;
    border-radius: 10%;
  }

  .user-btn img {
    width: 35px;
    height: 35px;
    border: 1px solid var(--gray-50);
    border-radius: 50%;
    object-fit: cover;
  }

  .brand-link {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 1.5rem;
    font-weight: 400;
    line-height: 2rem;
    letter-spacing: 0;
    color: inherit;       
    text-decoration: none; 
    font-family: var(--secondary-font);
  }

  .brand-link:is(:hover, :active, :focus) {
    text-decoration: underline;
    text-decoration-thickness: 2px;
    cursor: pointer;
    color: #0b57d0;
  }

  .brand .icon svg {
    display: block;
    width: 32px;
    height: 32px;
  }

  .spacer {
    flex: 1;
  }

  .class-display {
    display: flex;
    align-items: flex-start;
    line-height: 1.1;
    gap: 1rem;
    font-family: var(--secondary-font);
  }

  .arrow-icon {
    display: flex;
  }

  .class-metadata {
    display: flex;
    flex-direction: column;
  }

  .class-name {
    font-size: 1rem;
    font-weight: 600;
    color: #222;
  }

  .teacher-name {
    font-size: 0.8rem;
    color: #555;
    opacity: 0.9;
  }

  .user-btn-container {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .add-btn {
    background: none;
    border: none;
    cursor: pointer;
    padding: 0.25rem;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .add-btn svg {
    width: 24px;
    height: 24px;
    fill: #000;
  }

  .add-btn:hover {
    border-radius: 50%;
    background-color: #ebebeb;
  }

  .add-btn:active {
    border-radius: 50%;
    background-color: #a8a8a8;
  }

  .add-btn-wrapper {
    position: relative;
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .dropdown-overlay {
    position: fixed;
    inset: 0;
    z-index: 10;
  }

  .dropdown-menu {
    position: absolute;
    top: 45px;
    right: 80px;
    background: white;
    border: 1px solid #ebebeb;
    display: flex;
    flex-direction: column;
    min-width: 120px;
    z-index: 50;
  }

  .dropdown-item {
    padding: 10px 15px;
    background: none;
    border: none;
    text-align: left;
    cursor: pointer;
    font-size: 0.9rem;
    color: #222;
  }

  .dropdown-item:hover {
    background-color: #f0f0f0;
  }

  .fade-scale-enter-active, .fade-scale-leave-active {
    transition: all 0.3s ease;
  }

  .fade-scale-enter-from, .fade-scale-leave-to {
    opacity: 0;
    transform: scale(0.5);
  }

  .fade-scale-enter-to, .fade-scale-leave-from {
    opacity: 1;
    transform: scale(1);
  }

  .slide-fade-enter-active {
    transition: all 0.25s ease-out;
  }

  .slide-fade-leave-active {
    transition: all 0.2s ease-in;
  }

  .slide-fade-enter-from {
    opacity: 0;
    transform: translateY(-10px);
  }

  .slide-fade-enter-to {
    opacity: 1;
    transform: translateY(0);
  }

  .slide-fade-leave-from {
    opacity: 1;
    transform: translateY(0);
  }

  .slide-fade-leave-to {
    opacity: 0;
    transform: translateY(-10px);
  }
}

@layer theme {
  .gc-navbar {
    background-color: var(--primary-color); 
    color: black;
  }

  .menu-btn,
  .user-btn {
    color: black;
    opacity: 0.9;
    padding: 0.5rem;
  }

  .menu-btn:hover,
  .user-btn:hover {
    opacity: 1;
  }
}

</style>
