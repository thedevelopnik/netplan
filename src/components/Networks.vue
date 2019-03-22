<template>
  <v-container>
    <v-layout
      align-content-space-around
      justify-space-around
      mb-3
    >
      <v-card>
        <v-card-title>Current Network Map</v-card-title>
        <v-card-content>
          {{ currentNetworkMap.Name }}
        </v-card-content>
      </v-card>
      <v-card>
        <v-card-title>Current VPC</v-card-title>
        <v-card-content>
          {{ currentVPC.Name }} - {{ currentVPC.CidrBlock }}
        </v-card-content>
      </v-card>
      <v-card>
        <v-card-title>Current Subnet</v-card-title>
        <v-card-content>
          {{ currentSubnet.Name }} - {{ currentSubnet.CidrBlock }}
        </v-card-content>
      </v-card>
    </v-layout>
    <v-layout
      text-xs-center
      align-content-space-around
      justify-space-around
      wrap
    >
      <v-flex xs12 md3 class="white pa-2">
        <h2>Add Network Map</h2>
        <v-form
          ref="vpcForm"
          lazy-validation
        >
          <v-text-field
            v-model="networkMapName"
            label="Name"
            required
          >
          </v-text-field>
          <v-btn
            color="primary"
            @click="addNewNetworkMap(networkMapName)"
          >
            Add
          </v-btn>
        </v-form>
      </v-flex>
      <v-flex xs12 md3 class="white pa-2">
        <h2>Add VPC</h2>
        <v-form
          ref="vpcForm"
          lazy-validation
        >
          <v-text-field
            v-model="vpcName"
            label="Name"
            required
          >
          </v-text-field>
          <v-select
            v-model="vpcEnv"
            label="Environment"
            :items="['dev', 'stg', 'prd']"
            required
          ></v-select>
          <v-select
            v-model="vpcProvider"
            label="Provider"
            :items="['AWS', 'GCP', 'Aiven']"
            required
          ></v-select>
          <v-text-field
            v-model="vpcLocation"
            label="Location"
            required
          >
          </v-text-field>
          <v-text-field
            v-model="vpcAccess"
            label="Access"
            required
          >
          </v-text-field>
          <v-text-field
            v-model="vpcCidrBlock"
            label="CIDR Block"
            required
          >
          </v-text-field>
          <v-btn
            color="primary"
            @click="addNewVpc"
          >
            Add
          </v-btn>
        </v-form>
      </v-flex>
      <v-flex xs12 md3 class="white pa-2">
        <h2>Add Subnet</h2>
        <v-form
          ref="subnetForm"
          lazy-validation
        >
          <v-text-field
            v-model="subnetName"
            label="Name"
            required
          >
          </v-text-field>
          <v-text-field
            v-model="subnetAccess"
            label="Access"
            required
          >
          </v-text-field>
          <v-text-field
            v-model="subnetCidrBlock"
            label="CIDR Block"
            required
          >
          </v-text-field>
          <v-btn
            color="primary"
            @click="addNewSubnet"
          >
            Add
          </v-btn>
        </v-form>
      </v-flex>
      <v-flex xs12 class="mt-5">
        <h2>Network Maps</h2>
        <v-data-table
          :headers="networkMapHeaders"
          :items="networkMaps"
          class="elevation-1"
        >
          <template slot="items" slot-scope="props">
            <td @click="setCurrentNetworkMap(props.item)">{{ props.item.ID }}</td>
            <td @click="setCurrentNetworkMap(props.item)">{{ props.item.Name }}</td>
          </template>
        </v-data-table>
      </v-flex>
      <v-flex xs12 class="mt-5">
        <h2>VPCs</h2>
        <v-data-table
          :headers="headers"
          :items="currentNetworkMap.VPCs"
          class="elevation-1"
        >
          <template slot="items" slot-scope="props">
            <td @click="setCurrentVPC(props.item)">{{ props.item.CidrBlock }}</td>
            <td @click="setCurrentVPC(props.item)">{{ props.item.Type }}</td>
            <td @click="setCurrentVPC(props.item)">{{ props.item.Env }}</td>
            <td @click="setCurrentVPC(props.item)">{{ props.item.Name }}</td>
            <td @click="setCurrentVPC(props.item)">{{ props.item.Provider }}</td>
            <td @click="setCurrentVPC(props.item)">{{ props.item.Location }}</td>
          </template>
        </v-data-table>
      </v-flex>
      <v-flex xs12 class="mt-5">
        <h2>Subnets for selected VPC</h2>
        <v-data-table
          :headers="headers"
          :items="currentVPC.Subnets"
          class="elevation-1"
        >
          <template slot="items" slot-scope="props">
            <td>{{ props.item.CidrBlock }}</td>
            <td>{{ props.item.Type }}</td>
            <td>{{ props.item.Env }}</td>
            <td>{{ props.item.Name }}</td>
            <td>{{ props.item.Provider }}</td>
            <td>{{ props.item.Location }}</td>
          </template>
        </v-data-table>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { getAllNetworkMaps, getNetworkMap, createNetworkMap, newNetworkMap } from '../api/networkMaps'
import { createVPC, newVPC } from '../api/vpcs'
import { createSubnet, newSubnet } from '../api/subnets'

export default {
  async mounted () {
    try {
      const networkMaps = await getAllNetworkMaps()
      this.networkMaps = networkMaps
    } catch (err) {
      console.error(err)
    }
  },
  methods: {
    async setCurrentNetworkMap (nm) {
      try {
        this.currentNetworkMap = await getNetworkMap(nm.ID)
      } catch (err) {
        console.error(err)
      }
    },
    async addNewNetworkMap (Name) {
      try {
        const networkMap = newNetworkMap(Name)
        const created = await createNetworkMap(networkMap)
        this.setCurrentNetworkMap(created)
      } catch (err) {
        console.error(err)
      }
    },
    setCurrentVPC (vpc) {
      this.currentVPC = vpc
    },
    async addNewVpc () {
      let vpc = newVPC(
        this.vpcName,
        this.vpcAccess,
        this.vpcLocation,
        this.vpcProvider,
        this.vpcEnv,
        this.vpcCidrBlock,
        this.type,
        this.currentNetworkMap.ID
      )
      try {
        vpc = await createVPC(vpc)
      } catch (err) {
        console.error(err)
      }
      try {
        const networkMap = await getNetworkMap(vpc.NetworkMapID)
        this.currentNetworkMap = networkMap
      } catch (err) {
        console.error(err)
      }
    },
    async addNewSubnet () {
      let subnet = newSubnet(
        this.subnetName,
        this.subnetAccess,
        this.currentVPC.Location,
        this.currentVPC.Provider,
        this.currentVPC.Env,
        this.subnetCidrBlock,
        this.currentVPC.ID
      )
      try {
        subnet = await createSubnet(subnet)
      } catch (err) {
        console.error(err)
      }
      try {
        await this.setCurrentNetworkMap(this.currentVPC.NetworkMapID)
        this.currentSubnet = subnet
      } catch (err) {
        console.error(err)
      }
    }
  },
  data: () => {
    return {
      vpcName: '',
      vpcEnv: '',
      vpcAccess: '',
      vpcProvider: '',
      vpcLocation: '',
      vpcCidrBlock: '',
      vpcNames: [],
      vpcNetworkMap: '',
      networkMapName: '',
      subnetName: '',
      subnetAccess: '',
      subnetCidrBlock: '',
      subnetVPC: '',
      currentVPC: {},
      currentNetworkMap: {},
      currentSubnet: {},
      networkMapHeaders: [
        {
          text: 'ID',
          align: 'left',
          sortable: false,
          value: 'id'
        },
        {
          text: 'Name',
          align: 'left',
          sortable: true,
          value: 'name'
        }
      ],
      headers: [
        {
          text: 'CIDR Block',
          align: 'left',
          sortable: false,
          value: 'cidrBlock'
        },
        {
          text: 'Type',
          sortable: false,
          value: 'type'
        },
        {
          text: 'Environment',
          sortable: false,
          value: 'env'
        },
        {
          text: 'Name',
          sortable: false,
          value: 'name'
        },
        {
          text: 'Provider',
          sortable: false,
          value: 'provider'
        },
        {
          text: 'Location',
          sortable: false,
          value: 'location'
        }
      ],
      networks: {},
      networkMaps: []
    }
  }
}
</script>

<style>

</style>
