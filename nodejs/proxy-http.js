const axios = require("axios");

const url = "API_URL";
axios.get(url, {
    proxy: {
        protocol: 'http',
        host: 'gw-us.nstproxy.com',
        port: 24125,
        auth: {
            username: 'username',
            password: 'password',
        },
    },
})
    .then((res) => {
        console.log(res.data);
    })
    .catch((err) => {
        console.log('[err]:', err);
    });