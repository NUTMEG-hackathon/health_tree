<template>
  <v-footer
    v-if = "
      this.$route.path === '/training' ||
      this.$route.path === '/battle' ||
      this.$route.path === '/dashboard'
    "
    app
    color="primary"
    class="pa-0 ma-0 justify-center"
  >
  <v-card
    flat
    color="primary"
    class="pa-0"
    >
    <v-card-text class="pa-0">
      <v-btn-toggle
        tile
        dark
        color="accent"
        group
      >
        <v-btn
          class="px-7"
          to="/training"
        >
          <v-icon>mdi-watering-can</v-icon>
        </v-btn>
        <v-btn 
          class="px-7"
          to="/battle"
        >
          <v-icon>mdi-sword-cross</v-icon>
        </v-btn>       
        <v-btn 
          class="px-7"
          to="/dashboard"
        >
          <v-icon>mdi-account</v-icon>
        </v-btn>
      </v-btn-toggle>
    </v-card-text>
  </v-card>
</v-footer>
</template>
