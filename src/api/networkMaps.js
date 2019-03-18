import axios from './axios'

export async function getAllNetworkMaps () {
  try {
    const networkMaps = await axios.get('/networkmap')
    return networkMaps.data
  } catch (err) {
    throw new Error(err)
  }
}

export async function getNetworkMap (id) {
  try {
    const networkMap = await axios.get(`/networkmap/${id}`)
    return networkMap.data
  } catch (err) {
    throw new Error(err)
  }
}
