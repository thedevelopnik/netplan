import Axios from 'axios'

const requester = Axios.create({
  baseURL: '/v1',
  timeout: 1000
})

export default requester
