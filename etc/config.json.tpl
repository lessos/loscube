{
  "pod_home_dir": "",
  "bigtree": {
    "peer_lan_addr": "127.0.0.1:9529",
    "peer_wan_addr": "127.0.0.1:9529",
    "peer_api_addrs": [
      {
        "address": "127.0.0.1:9529"
      }
    ],
    "srv_port": 9530
  },
  "lps_io_connectors": [
    {
      "name": "database",
      "connector": "iomix/skv/Connector",
      "driver": "lessdb/sskv",
      "items": [
        {
          "name": "data_dir",
          "value": "",
        }
      ]
    },
    {
      "name": "storage",
      "connector": "iomix/fs/Connector",
      "driver": "local/fs",
      "items": [
        {
          "name": "data_dir",
          "value": "",
        }
      ]
    }
  ]
}
