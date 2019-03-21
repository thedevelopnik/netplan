import axios from './axios'

export function newNetworkMap (Name) {
  return {
    Name
  }
}

export async function getAllNetworkMaps () {
  try {
    const res = await axios.get('/networkmap')
    return res.data
  } catch (err) {
    throw new Error(err)
  }
}

export async function getNetworkMap (id) {
  try {
    const res = await axios.get(`/networkmap/${id}`)
    return res.data
  } catch (err) {
    throw new Error(err)
  }
}

export async function createNetworkMap (networkMap) {
  console.log(networkMap)
  try {
    const res = await axios.post('/networkmap', networkMap)
    return res.data
  } catch (err) {
    throw new Error(err)
  }
}
