<template>
  <v-card
    dark
    height=100%
    width=100%
    color="primary"
  >
  <v-layout fill-height align-center justify-center>
    <v-card
      flat
      color="primary"
      class="text-center"
      >
  <v-layout fill-height align-center justify-center>
      <v-img 
        src="@/assets/type1-1.svg" 
        max-height=70
        max-width=70
        position="center"
      />
  </v-layout>
    <v-card-title class="justify-center pb-10">Have you fun every day?</v-card-title>
    <v-card-text class="pb-10">
      <v-form ref="form">
        <v-text-field
          label="Name"
        ></v-text-field>
        <v-text-field
          label="E-mail"
        ></v-text-field>
        <v-text-field
          label="Password"
        ></v-text-field>
      </v-form>
    </v-card-text>
    <v-card-actions class="justify-center">
      <v-btn color="accent" outlined>Sign in</v-btn>
      <v-btn color="accent" to="/select">Sign up</v-btn>
    </v-card-actions>
    </v-card>
  </v-layout>
  </v-card>
</template>

<script>
  export default {
    name: 'home',

    components: {
    },
  }
</script>
