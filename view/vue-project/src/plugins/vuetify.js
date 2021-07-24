import Vue from 'vue'
import Vuetify from 'vuetify/lib'

import colors from 'vuetify/lib/util/colors'

Vue.use(Vuetify)

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: '#aeea00', // #E53935
        secondary: '#79b700', // #FFCDD2
        accent: '#ff6e40', // #3F51B5
      },
    },
  },
})
