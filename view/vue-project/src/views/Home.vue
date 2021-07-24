<template>
  <v-card
    dark
    height=100%
    width=100%
    color="primary"
  >
  <v-layout fill-height align-center justify-center>
    <v-card-title class="justify-center">Let's enjoy!</v-card-title>
        <v-card-text>
            <v-form ref="form">
              <v-text-field
                  label="フルネーム"
                  ></v-text-field>
              <v-text-field
                  label="メールアドレス"
                  ></v-text-field>
              <v-text-field
                  label="パスワード"
                  ></v-text-field>
              <v-text-field
                  label="パスワードの再入力"
                  ></v-text-field>
            </v-form>
        </v-card-text>
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
