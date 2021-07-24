<template>
  <div>
    <v-row class="pa-0">
      <v-col class="pa-0">
        <v-img 
          src="@/assets/type1-1.jpg" 
        />
        <v-card
          flat
          class="ma-10"
        >
          <v-progress-linear 
            value="75"
            color="secondary"
          />
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card
          flat
        >
        <v-card-title class="justify-center">Task
          <v-dialog
              transition="dialog-bottom-transition"
              max-width="600"
              >
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                    small
                    fab
                    elevation="1"
                    color="accent"
                    class="ml-2"
                    v-bind="attrs"
                    v-on="on"
                    >
                    <v-icon>mdi-chevron-down</v-icon>
                </v-btn>
               </template>
              <template v-slot:default="dialog">
                <v-card>
                  <v-toolbar
                    dark
                    color="primary"
                  >Today's task</v-toolbar>
                  <v-card-text>
                    <v-overflow-btn></v-overflow-btn>
                    <v-overflow-btn></v-overflow-btn>
                    <v-overflow-btn></v-overflow-btn>
                    <v-overflow-btn></v-overflow-btn>
                    <v-overflow-btn></v-overflow-btn>
                  </v-card-text>
                  <v-card-actions class="justify-end">
                    <v-btn
                        text
                        @click="dialog.value = false"
                        >Close</v-btn>
                  </v-card-actions>
                </v-card>
              </template>
          </v-dialog>
        </v-card-title>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>
