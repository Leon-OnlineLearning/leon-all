



# revirse ports from hyper v
some type hyper-v take needed ports and prevent docker from starting
to prevent this from happening again

symtomps
```log
docker: error response from daemon: Ports are not available: listen tcp 0.0.0.0/xxxx: bind: An attempt was made to access a socket in a way forbidden by its access permissions.
```

1. make sure it is the problem 

```powershell
netsh interface ipv4 show excludedportrange protocol=tcp
```

2. diable hyper-v 
```powershell
dism.exe /Online /Disable-Feature:Microsoft-Hyper-V
```

3. restart

4. exclude port
```powershell
netsh int ipv4 add excludedportrange protocol=tcp startport=<port_you_want> numberofports=1
```
ports we want 
|service|port|
|---|---|
|postgres|5432|
|redis|6379|
|node|3333|
|next|3000|
|janus|8088|
|janus|7088|
|python|50000|
|python|5000|




5. reenable hyper-v
```powershell
dism.exe /Online /Enable-Feature:Microsoft-Hyper-V /All
```


ref : https://stackoverflow.com/questions/65272764/ports-are-not-available-listen-tcp-0-0-0-0-50070-bind-an-attempt-was-made-to
