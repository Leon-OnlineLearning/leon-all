certificated were created with certbot 
steps to regenerate 
- first you may need to install cert bot for examle in debian based
`sudo apt install certbot`

- then genrate the certificates with 
```bash
certbot certonly --standalone
```

in production those certificates may bee moved to container it self where it 
is renewed whenever they expire
