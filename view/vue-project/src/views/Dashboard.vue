<template>
  <div>
    <v-row>
      <v-col class="text-center pa-0">
        <v-card 
          flat
          class="mx-10 my-5 align-center"
        >
           <v-card-title>Profile</v-card-title>
            <v-card-text>
              <v-simple-table>
                <template v-slot:default>
                  <tbody>
                    <tr>
                      <td class="font-weight-bold">User name</td>
                      <td>Jiro GIDAI</td>
                    </tr>
                    <tr>
                      <td class="font-weight-bold">Tree name</td>
                      <td>mashimashi</td>
                    </tr>
                    <tr>
                      <td class="font-weight-bold">Bio</td>
                      <td>対戦よろしくお願いします．</td>
                    </tr>
                    <tr>
                      <td class="font-weight-bold">Email</td>
                      <td>nutfes-jiro@email.com</td>
                    </tr>
                  </tbody>
                </template>
              </v-simple-table>
            </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-divider />
    <v-row>
      <v-col class="text-center pa-0">
        <v-card 
          flat
          class="mx-10 my-5 align-center"
        >
        <v-card-title>Log</v-card-title>
        <v-card-text>
          <v-simple-table>
            <template v-slot:default>
              <tbody>
                <tr>
                  <td class="font-weight-bold">Battle</td>
                  <td>You win</td>
                </tr>
                <tr>
                  <td class="font-weight-bold">Training</td>
                  <td>Evolution</td>
                </tr>
              </tbody>
            </template>
          </v-simple-table>
        </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-divider />
  </div>
</template>
