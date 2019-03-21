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

export function newNetworkMap (Name) {
  return {
    Name
  }
}

export function newSubnet (Name, Access, Location, Provider, Env, CidrBlock, VPCID) {
  return {
    Name,
    Access,
    Location,
    Provider,
    Env,
    CidrBlock,
    VPCID
  }
}

export function newVPC (Name, Access, Location, Provider, Env, CidrBlock, Type, NetworkMapID) {
  return {
    Name,
    Access,
    Location,
    Provider,
    Env,
    CidrBlock,
    Type,
    NetworkMapID
  }
}
