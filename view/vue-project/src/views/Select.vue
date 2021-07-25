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
      >
        <v-card-title class="justify-center my-10">Choose.</v-card-title>
        <v-card-text>
          <v-layout fill-height align-center justify-center>
            <v-card
                height=100
                @click="submit"
                color="primary"
                >
                <v-img 
                src="@/assets/type1-3.svg" 
                max-height=60
                max-width=60
                class="ma-5"
                />
            </v-card>
              <v-card
                  height=100
                  @click="submit"
                  color="primary"
                  >
                  <v-img 
                  src="@/assets/type2-3.svg" 
                  max-height=60
                  max-width=60
                  class="ma-5"
                  />
              </v-card>
                <v-card
                    height=100
                    @click="submit"
                    color="primary"
                    >
                    <v-img 
                    src="@/assets/type3-3.svg" 
                    max-height=60
                    max-width=60
                    class="ma-5"
                    />
                </v-card>
          </v-layout>
        </v-card-text>
        <v-card-title class="justify-center my-10">Name.</v-card-title>
        <v-card-text>
          <v-layout fill-height align-center justify-center>
            <v-form ref="form">
              <v-text-field
                v-model="name"
                label=""
              ></v-text-field>
            </v-form>
          </v-layout>
        </v-card-text>
        <v-card-actions class="justify-center">
          <v-btn color="accent" @click="submit"><v-icon>mdi-chevron-down</v-icon></v-btn>
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
      type: 1,
      point: 100,
      token: localStorage.getItem('access-token'),
    }

  },
  methods: {
    submit: function () {
      const url = process.env.VUE_APP_URL + '/api/tree'
      axios.post(url, {
        name: this.name,
        type: this.type,
        point: this.point
      },{
        headers: {
          Authorization: `Bearer ${this.token}`
        }
      }).then(
        (response) => {
          this.$router.push('training')
        }
      )
    }
  }
}
</script>
