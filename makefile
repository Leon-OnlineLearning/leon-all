public_ip=$(shell dig +short myip.opendns.com @resolver1.opendns.com)


all: start_container update_ip
	
start_container:;docker-compose up -d
stop_container:;docker-compose down
update_ip: .env
	sed s/nat_1_1_mapping.*/nat_1_1_mapping=\"$(public_ip)\"/g janus/janus.jcfg