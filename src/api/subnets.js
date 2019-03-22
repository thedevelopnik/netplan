import axios from './axios'

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

export async function createSubnet (subnet) {
  try {
    const res = await axios.post(`/networkmap/0/vpc/${subnet.VPCID}/subnet`, subnet)
    return res.data
  } catch (err) {
    throw new Error(err)
  }
}
