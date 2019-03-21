import axios from './axios'

export async function createSubnet (subnet) {
  try {
    const res = await axios.post(`/v1/networkmap/0/vpc/${subnet.VPCID}/subnet`, subnet)
    return res.data
  } catch (err) {
    throw new Error(err)
  }
}
