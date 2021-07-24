<template>
  <v-app class="overflow-hidden">
    <Header />
    <v-navigation-drawer
      app
      clipped
    >
    <v-sheet>
    <v-card-title>Pages</v-card-title>
      <v-divider />
    <v-list dense class="ma-3">
        <v-list-item>
          Dashboard
        </v-list-item>
        <v-list-item>
          Home
        </v-list-item>
        <v-list-item>
          About
        </v-list-item>
    </v-list>
    </v-sheet>
    </v-navigation-drawer>

    <v-main class="overflow-hidden">
      <transition mode='out-in'>
        <router-view/>
      </transition>
    </v-main>
    <Footer />
  </v-app>
</template>

<script>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
export default {
  name: 'App',
  components: {
    Header,
    Footer
  },
};
</script>

<style>
#app{
  background-color: white;
}
.v-enter {
  transform: translate(-100px, 0);
  opacity: 0;
}
.v-enter-to {
  opacity: 1;
}
.v-enter-active {
  transition: all .5s .2s ease;
}
.v-leave {
  transform: translate(0, 0);
  opacity: 1;
}
.v-leave-to {
  transform: translate(100px, 0);
  opacity: 0;
}
.v-leave-active {
  transition: all .4s .2s ease;
}
</style>
