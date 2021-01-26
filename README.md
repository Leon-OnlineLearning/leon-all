run janus locally with https

# instructions
- run docker container with `docker-compose up`
- create a certificate by following [create_cert.sh](certificate\create_cert.sh)
- add [CA.pem](certificate\myCA.pem) as trusted certificate authority to browser
- run nginx server with configuration in [nginx.conf](conf\nginx.conf)
