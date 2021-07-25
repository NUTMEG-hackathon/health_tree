<template>
  <div>
    <v-row class="pa-0">
      <v-col class="pa-0">
        <v-img 
          src="@/assets/type2-1.jpg" 
        />
        <v-card
          flat
          class="ma-10"
        >
        {{ aaaa }}
          <v-progress-linear 
            v-model="this.total_point"
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
                  <v-card-title class="px-5 pt-6">
                    今日は何時に起きましたか？
                  </v-card-title>
                    <v-overflow-btn
                     v-model="point1"
                     :items="wake_up_time"
                     item-text="text"
                     item-value="point"
                    ></v-overflow-btn>
                    <v-card-title class="px-5 py-0">
                    朝食は食べましたか？
                  </v-card-title>
                    <v-overflow-btn
                     v-model="point2"
                     :items="eat_breakfast"
                     item-text="text"
                     item-value="point"
                     ></v-overflow-btn>
                    <v-card-title class="px-5 py-0">
                    水をどのくらい飲みましたか？
                  </v-card-title>
                    <v-overflow-btn
                     v-model="point3"
                     :items="drink_water"
                     item-text="text"
                     item-value="point"
                     ></v-overflow-btn>
                  <v-card-title class="px-5 py-0">
                    歯を磨く時間はどのくらいでしたか？
                  </v-card-title>
                    <v-overflow-btn
                     v-model="point4"
                     :items="brashing_time"
                     item-text="text"
                     item-value="point"
                     ></v-overflow-btn>
                    <v-card-title class="px-5 py-0">
                    何時に寝ましたか？
                  </v-card-title>
                    <v-overflow-btn
                     v-model="point5"
                     :items="retiring_time"
                     item-text="text"
                     item-value="point"
                     ></v-overflow-btn>
                  <v-card-actions class="justify-space-around">
                     <v-btn
                        text
                        to="/training2"
                        >Submit</v-btn>
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


<script>
  export default {
    data(){
      return{
        point1: 0,
        point2: 0,
        point3: 0,
        point4: 0,
        point5: 0,
        total_point: 75,
        
      }

    },
    data: () => ({

      wake_up_time: [
        { text: '5:00～5:59' ,point: 15},
        { text: '6:00～6:59' ,point: 10},
        { text: '7:00～7:59' ,point: 5},
        { text: '8:00～8:59' ,point: 2},
        { text: '9:00～9:59' ,point: 1},
      ],
      eat_breakfast: [
        { text:'はい',point: 10},
        { text:'いいえ',point: 0},
      ],
      drink_water: [
        { text: '0～499ml' ,point: 0},
        { text: '500ml～999mL' ,point: 1},
        { text: '1L～1.49L' ,point: 5},
        { text: '1.5L～1.99L' ,point: 10},
        { text: '2L～2.49L' ,point: 15},
        { text: '2.5L～' ,point: 20},
      ],
      brashing_time: [
        { text: '磨いていない' ,point: 0},
        { text: '1分くらい' ,point: 5},
        { text: '2分くらい' ,point: 10},
        { text: '3分くらい' ,point: 15},
        { text: '3分以上' ,point: 20},
      ],
      retiring_time: [
        { text: '～22:59' ,point:  20},
        { text: '23:00～23:59' ,point: 15},
        { text: '00:00～00:59' ,point: 10},
        { text: '1:00～1:59' ,point: 5},
        { text: '2:00～' ,point: 1},
      ],
    }),
    methods:{
      submit: function(){
        this.total_point=this.point1+this.point2+this.point3+this.point4+this.point5;
        console.log(this.total_point);
        
      },
      reload: function(){
        
      }
    }
  }
</script>