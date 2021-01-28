#source https://stackoverflow.com/questions/7580508/getting-chrome-to-accept-self-signed-localhost-certificate

######################
# Become a Certificate Authority
######################

# Generate private key
# openssl genrsa -des3 -out myCA.key 2048
# Generate root certificate
# openssl req -x509 -new -nodes -key myCA.key -sha256 -days 825 -out myCA.pem

######################
# Create CA-signed certs
######################
localIP=192.168.18.19 #use your ip here
publicIP=156.209.166.123 # use public ip if you want

NAME=leon-test.com # Use your own domain name
# Generate a private key
openssl genrsa -out $NAME.key 2048
# Create a certificate-signing request
openssl req -new -key $NAME.key -out $NAME.csr
# Create a config file for the extensions
>$NAME.ext cat <<-EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names
[alt_names]
DNS.1 = $NAME # Be sure to include the domain name here because Common Name is not so commonly honoured by itself
DNS.2 = bar.$NAME # Optionally, add additional domains (I've added a subdomain here)
IP.1 = $localIP  # Optionally, add an IP address (if the connection which you have planned requires it)
IP.2 = $publicIP 
EOF
# Create the signed certificate
openssl x509 -req -in $NAME.csr -CA myCA.pem -CAkey myCA.key -CAcreateserial \
-out $NAME.crt -days 825 -sha256 -extfile $NAME.ext


# openssl verify -CAfile myCA.pem -verify_hostname bar.mydomain.com mydomain.com.crt
