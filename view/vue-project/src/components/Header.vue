<template>
  <div
    v-if = "
      this.$route.path === '/training' ||
      this.$route.path === '/battle' ||
      this.$route.path === '/dashboard'
    "
  >
    <v-app-bar
      app
      dark
      clipped-left
      color="primary"
      class="align-center justify-center text-center"
    >
    <v-row>
      <v-col cols=2 align-self="center" class="pa-0 ma-0"></v-col>
      <v-col>
        <v-card-title class="justify-center">
          Daily Life Tree
        </v-card-title>
      </v-col>
      <v-col cols=2 align-self="center" class="pa-0 ma-0"><v-btn text dark><v-icon>mdi-logout</v-icon></v-btn></v-col>
    </v-row>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  
}
</script>
