<template>
  <div v-if="windowWidth <= 768">

  <div
    v-if = "
      this.$route.path === '/training' ||
      this.$route.path === '/battle' ||
      this.$route.path === '/dashboard'
    "
  >
    <v-navigation-drawer
      expand-on-hover
      app
      clipped
      permanent
    >

      <v-list 
      dense >
        <v-list-item
          v-for="item in items"
          :key="item.title"
          :to="item.link"
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
  </div>
</template>


<script>
export default {
    data () {
      return {
        items: [
          { title: 'Training', icon: 'mdi-watering-can', link:'/training' },
          { title: 'Battle', icon: 'mdi-sword-cross' ,link:'/battle'},
          { title: 'Dachboard', icon: 'mdi-account' ,link:'/dashboard'},
        ],
      windowWidth: window.innerWidth,
      }
    },
  mounted() {
    window.addEventListener('resize', this.calculateWindowWidth);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.calculateWindowWidth);
  },
  methods: {
    calculateWindowWidth() {
      this.windowWidth = window.innerWidth;
      console.log(this.windowWidth);
    }
  }
}
</script>