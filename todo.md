## before exam at the start of year
- get one or more refrence vector 
    * one image perfect _we cannot get perfect image (complex)
    * get multible images from vedio 
        - calculate mean _
            if one of the image is not perfect it will vanach when we take mean
        - save vector for every image and compare
 

- get more refrences from lectures

## during exam
- get multible vectors and compare with refrance
    * one image every minute

- make sure face in the frame of camera



-- end point
vedio --> array of emedding
vedio , array of embedding  --> match / not match




# rtp 
* suitable for lecture 
# sctp
* complicated 
* require other server to handle saving 
# record and play send using tcp
save seconds of vedio and transmmit it
create blob 
decode the video frames








-- for lecture
- [x] test creating multiple janus objects _
    this will create multilble sessions , don't know if this any bad 
    may be it is not , but for know there is a parent that create a session and 
    will pass it to multible childs
- [x] make audio component 
- [ ] make data component
- [ ] make the data component work with generic data

-- in exam
- [ ] optain rtp packet or media in any format from Media stream
- [ ] send vedio over data channel