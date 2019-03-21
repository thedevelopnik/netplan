import axios from 'axios'

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

export async function createVPC (vpc) {
  try {
    const res = await axios.post(`/v1/networkmap/${vpc.NetworkMapID}/vpc`, vpc)
    return res.data
  } catch (err) {
    throw new Error(err)
  }
}
