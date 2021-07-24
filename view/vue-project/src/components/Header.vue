<template>
  <div
    v-if="this.$route.path !== '/'"
  >
    <v-app-bar
      app
      dark
      clipped-left
      color="primary"
      class="align-center justify-center text-center"
    >
    <v-row>
      <v-col>
        <v-card-title class="justify-center">
          <v-icon>mdi-tree-outline</v-icon>
          Health Tree
        </v-card-title>
      </v-col>
    </v-row>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  
}
</script>
