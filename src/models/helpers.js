import { Netmask } from 'netmask'
import { IPv4 } from 'ip-num'

export function createMetadata (name, access, location, provider, env, type, cidrBlock) {
  return {
    name,
    access,
    location,
    provider,
    env,
    cidrBlock,
    type
  }
}

export function newSubnet (metadata) {
  const netmask = new Netmask(metadata.cidrBlock)
  const first = new IPv4(netmask.first)
  const last = new IPv4(netmask.last)
  return {
    metadata,
    netmask,
    first,
    last
  }
}

export function addVPC (networkMap, vpcMetadata, subnetMetadatas) {
  return new Promise((resolve, reject) => {
    const vpc = newVPC(vpcMetadata, subnetMetadatas)
    if (!networkHasOverlap(vpcMetadata.netmask, networkMap.vpcs)) {
      let vpcs = networkMap.vpcs
      vpcs.push(vpc)
      vpcs.sort(sortNetworks)
      networkMap.vpcs = vpcs
      resolve(networkMap)
    }
    reject(new Error('vpc had overlap'))
  })
}

function sortNetworks (net1, net2) {
  if (net1.first.isLessThan(net2.first)) {
    return -1
  }
  return 1
}

function newVPC (vpcMetadata, subnetMetadatas) {
  const netmask = new Netmask(vpcMetadata.cidrBlock)
  const first = new IPv4(netmask.first)
  const last = new IPv4(netmask.last)

  let subnets = []
  if (subnetMetadatas) {
    subnets = subnetMetadatas.map(subnet => {
      if (netmask.contains(subnet.cidrBlock)) {
        return newSubnet(subnet)
      }
    })
  }
  subnets.sort(sortNetworks)

  return {
    metadata: vpcMetadata,
    netmask,
    first,
    last,
    subnets
  }
}

export function networkHasOverlap (netmask, existingNets) {
  let overlap = false
  existingNets.forEach(net => {
    if (net.netmask.contains(netmask.first) || net.netmask.contains(netmask.last)) {
      overlap = true
    }
  })
  return overlap
}

export function addSubnet (subnetMetadata, vpcName, vpcEnv, networkMap) {
  let vpcIndex
  networkMap.vpcs.forEach((vpc, ind) => {
    if (vpc.metadata.name === vpcName && vpc.metadata.env === vpcEnv) {
      vpcIndex = ind
    }
  })
  const subnet = newSubnet(subnetMetadata)
  if (!networkMap.vpcs[vpcIndex].netmask.contains(subnet.netmask.first) || !networkMap.vpcs[vpcIndex].netmask.contains(subnet.netmask.last)) {
    throw new Error('subnet range was not in the specified vpc')
  }
  if (networkHasOverlap(subnet.netmask, networkMap.vpcs[vpcIndex].subnets)) {
    throw new Error('subnet has overlap with another subnet in the specified vpc')
  }
  networkMap.vpcs[vpcIndex].subnets.push(subnet)
  networkMap.vpcs[vpcIndex].subnets.sort(sortNetworks)
  return networkMap
}
