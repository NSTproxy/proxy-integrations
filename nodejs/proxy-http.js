const axios = require("axios");

const url = "API_URL";
const resp = axios.get(url, {
    proxy: {
        host: 'gw-us.nstproxy.com',
        port: 24125,
        auth: {
            username: 'username',
            password: 'password'
        }
    }

})
console.log(resp.data);