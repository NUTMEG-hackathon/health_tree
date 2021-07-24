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
        <v-card-title class="justify-center my-10">Name.</v-card-title>
        <v-card-text>
          <v-layout fill-height align-center justify-center>
            <v-form ref="form">
              <v-text-field
                label=""
              ></v-text-field>
            </v-form>
          </v-layout>
        </v-card-text>
        <v-card-actions class="justify-center">
          <v-btn color="accent" to="/training"><v-icon>mdi-chevron-down</v-icon></v-btn>
        </v-card-actions>
      </v-card>
    </v-layout>
  </v-card>

</template>
