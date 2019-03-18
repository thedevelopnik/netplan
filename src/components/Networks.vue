<template>
  <v-container>
    <v-layout
      text-xs-center
      wrap
    >
      <v-flex xs12 md5 class="white pa-2">
        <h2>Add New VPC</h2>
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
      <v-flex xs0 md1>
      </v-flex>
      <v-flex xs12 md5 class="white pa-2">
        <h2>Add New Subnet</h2>
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
          <v-select
            v-model="subnetVPC"
            label="Parent VPC"
            required
            :items="vpcNames"
          >
          </v-select>
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
          :items="currentVPC.subnets"
          class="elevation-1"
        >
          <template slot="items" slot-scope="props" @click="setCurrentVPC(props.item)">
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
// import { initializeExistingNetworks } from '../models/networks'
import { addSubnet, addVPC, createMetadata } from '../models/helpers'
import { getAllNetworkMaps, getNetworkMap } from '../api/networkMaps'

export default {
  async mounted () {
    try {
      const networkMaps = await getAllNetworkMaps()
      this.networkMaps = networkMaps
      // this.networks = networks
      // this.vpcNames = networks.vpcs.map((vpc) => {
      //   return `${vpc.metadata.name}-${vpc.metadata.env}`
      // })
    } catch (err) {
      console.error(err)
    }
  },
  methods: {
    async setCurrentNetworkMap (nm) {
      try {
        const networkMap = await getNetworkMap(nm.ID)
        this.currentNetworkMap = networkMap
      } catch (err) {
        console.error(err)
      }
    },
    setCurrentVPC (vpc) {
      this.currentVPC = vpc
    },
    async addNewVpc () {
      const metadata = createMetadata(
        this.vpcName,
        this.vpcAccess,
        this.vpcLocation,
        this.vpcProvider,
        this.vpcEnv,
        'vpc',
        this.vpcCidrBlock
      )
      try {
        const networks = await addVPC(this.networks, metadata)
        this.networks = networks
        this.vpcNames = networks.vpcs.map((vpc) => {
          return `${vpc.metadata.name}-${vpc.metadata.env}`
        })
      } catch (err) {
        console.error(err)
      }
    },
    addNewSubnet () {
      const vpcIdSplit = this.subnetVPC.split('-')
      const vpc = this.networks.vpcs.filter(vpc => {
        return vpc.metadata.name === vpcIdSplit[0] && vpc.metadata.env === vpcIdSplit[1]
      })[0]
      const metadata = createMetadata(
        this.subnetName,
        this.subnetAccess,
        vpc.metadata.location,
        vpc.metadata.provider,
        vpc.metadata.env,
        'subnet',
        this.subnetCidrBlock
      )
      const networks = addSubnet(metadata, vpcIdSplit[0], vpcIdSplit[1], this.networks)
      this.networks = networks
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
      subnetName: '',
      subnetAccess: '',
      subnetCidrBlock: '',
      subnetVPC: '',
      currentVPC: {},
      currentNetworkMap: {},
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
