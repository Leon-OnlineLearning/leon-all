run janus locally with https

# instructions
- create a certificate [nginx/cert/README.md](https://github.com/Leon-OnlineLearning/leon-all/blob/main/nginx/cert/README.md)
- configure firewall to accept incoming connections at port `443`
- create `.env` file from `.env.docker` and add you secrets
- run npm install at `leon-ClientSide` & `Leon-ServerSide`
- run docker containers with `docker-compose up`

# Design
## During exam sequence
![during exam sequenec](design/DuringExamSequence.png)