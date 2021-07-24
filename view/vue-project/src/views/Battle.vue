<template>
  <div>
    <v-card 
      flat
      class="mx-10 my-5 align-center"
    >
      <v-card-title>BattleList</v-card-title>
      <v-card-text>
        <v-simple-table>
          <template v-slot:default>
            <tbody>
              <tr>
                <td>
                  <v-img 
                    src="@/assets/01_sample_tree.jpg" 
                    max-height=30
                    max-width=30
                  />
                </td>
                <td>User name</td>
                <td>Jiro GIDAI</td>
              </tr>
              <tr>
                <td>
                  <v-img 
                    src="@/assets/01_sample_tree.jpg" 
                    max-height=30
                    max-width=30
                  />
                </td>
                <td>User name</td>
                <td>Jiro GIDAI</td>
              </tr>
              <tr>
                <td>
                  <v-img 
                    src="@/assets/01_sample_tree.jpg" 
                    max-height=30
                    max-width=30
                  />
                </td>
                <td>User name</td>
                <td>Jiro GIDAI</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-card-text>
    </v-card>
  </div>

</template>


