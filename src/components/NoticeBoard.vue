<template>
    <div class="noticeboard-wrapper">
      <h3 class="noticeboard-title">Noticeboard</h3>
      <div class="noticeboard">
        <ul class="notice-list">
          <li 
            v-for="(notice, index) in notices" 
            :key="index" 
            class="notice-item"
            >
            <span class="notice-date">{{ formatDate(notice.date) }}</span>
            <div class="notice-content">
                <h4 class="notice-title">{{ notice.title }}</h4>
                <p 
                    v-if="notice.description" 
                    class="notice-description"
                >
                    {{ notice.description }}
                </p>
            </div>
          </li>
        </ul>
      </div>
    </div>
</template>
  
  
<script>
import noticesData from '/data/slug.json'

export default {
  name: "NoticeboardComponent",
  props: {
    courseSlug: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      notices: []
    };
  },
  created() {
    this.notices = noticesData[this.courseSlug] || [];
  },
  methods: {
    formatDate(dateStr) {
      const options = { month: 'short', day: 'numeric' };
      const date = new Date(dateStr);
      return date.toLocaleDateString(undefined, options); 
    }
  }
};
</script>
  
<style>
@layer noticeboard {
    .noticeboard-wrapper {
        flex: 1;
        margin: 20px;
        display: flex;
        flex-direction: column;
    }

    .noticeboard-title {
        font-size: 1.4rem;
        font-weight: 600;
        margin-bottom: 12px; 
        font-family: var(--secondary-font);
    }

    .noticeboard {
        height: 250px;
        background-color: #fff;
        border: var(--borderline);
        border-radius: 10px;
        padding: 16px;
        overflow-y: auto;
        cursor: default;
        mask-image: linear-gradient(to bottom, 
          rgba(0,0,0,1) 90%, 
          rgba(0,0,0,0) 100%
      );
    }

    .notice-list {
        list-style: none;
        padding: 0;
        margin: 0;
        border-left: 1px solid #333;
    }

    .notice-item {
        display: flex;
        align-items: flex-start;
        gap: 12px;
        padding: 10px 8px;
        cursor: default;
    }

    .notice-item:hover {
        background: transparent;
        box-shadow: none;
    }

    .notice-date {
        background-color: #f0f0f0;
        color: #333;
        font-weight: 600;
        font-size: 0.85rem;
        padding: 4px 8px;
        border-radius: 6px;
        min-width: 60px;
        text-align: center;
    }

    .notice-content {
        flex: 1;
        display: flex;
        flex-direction: column;
        gap: 4px;
    }

    .notice-title {
        font-size: 1rem;
        font-weight: 600;
        margin: 0;
    }

    .notice-description {
        font-size: 0.9rem;
        color: #555;
        margin: 0;
    }
}
</style>
  