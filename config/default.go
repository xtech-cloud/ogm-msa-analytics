package config

const defaultYAML string = `
service:
    name: xtc.api.ogm.analytics
    address: :9603
    ttl: 15
    interval: 10
logger:
    level: trace
    dir: /var/log/ogm/
database:
    lite: true
    mysql:
        address: 127.0.0.1:3306
        user: root
        password: mysql@OMO
        db: ogm
    sqlite:
        path: /tmp/ogm-analytics.db
`
