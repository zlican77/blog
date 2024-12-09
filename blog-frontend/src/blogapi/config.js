import axios from 'axios'

axios.defaults.withCredentials = true

const request = axios.create({
  baseURL: 'http://localhost:8000/blogapi/v1',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json',
  },
})

export default request
