<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span class="font-weight-light">NetPlan</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>
    <v-content>
      <Networks/>
    </v-content>
  </v-app>
</template>

<script>
import Networks from './components/Networks'

export default {
  name: 'App',
  components: {
    Networks
  },
  data () {
    return {
      //
    }
  }
}
</script>
