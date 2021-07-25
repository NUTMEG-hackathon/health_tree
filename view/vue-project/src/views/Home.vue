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
          v-model="name"
          label="Name"
          required
        ></v-text-field>
        <v-text-field
          v-model="password"
          ref="password"
          :append-icon="show_pass ? 'mdi-eye-off' : 'mdi-eye'"
          :type="show_pass ? 'password' : 'text'"
          hint="8文字以上"
          counter
          @click:append="show_pass = !show_pass"
          required
          label="Password"
        ></v-text-field>
      </v-form>
    </v-card-text>
    <v-card-actions class="justify-center">
      <v-btn color="accent" to="/select" outlined>Sign in</v-btn>
      <v-btn color="accent" @click="submit">Sign up</v-btn>
    </v-card-actions>
    </v-card>
  </v-layout>
  </v-card>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {    
      name: null,
      password: null,
    }
  },
  methods: {
    submit: function () {
      const url = process.env.VUE_APP_URL + '/signup'
      axios.post(url, {
        name: this.name,
        password: this.password
      }).then(
        (response) => {
          const signinurl = process.env.VUE_APP_URL + '/signin'
          axios.post(signinurl, {
            name: this.name,
            password: this.password
          }).then(
            (response) => {
              localStorage.setItem('access-token',response.data.token)
              this.$router.push('select')
            }
          ) 
        }
      )
    }
  }
}
</script>
