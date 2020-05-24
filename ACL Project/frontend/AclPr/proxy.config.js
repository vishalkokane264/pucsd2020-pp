const PROXY_CONFIG = {
    "/webapi/v1/*": {
        "target": "http://localhost:9091",
        "secure": false
    }
};
module.exports = PROXY_CONFIG;

