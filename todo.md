## before exam at the start of year
- get one or more refrence vector 
    * one image perfect _we cannot get perfect image (complex)
    * get multible images from vedio 
        - calculate mean _
            if one of the image is not perfect it will vanach when we take mean
        - save vector for every image and compare
- [x] save embeding in database
- [ ] make sure face in region during record

- [ ] create admin ui to check recorded video

- get more refrences from lectures

## during exam
- get multible vectors and compare with refrance
    * one image every minute
- [x] send recorded video from client to server
- [x] send from server to ML
- [x] save result came back from ML
- [ ] view result after exam




# fix 
-- for exam 
- [ ] use faceapi to make sure someone in front of camera


-- for lecture
- [x] test creating multiple janus objects _
    this will create multilble sessions , don't know if this any bad 
    may be it is not , but for know there is a parent that create a session and 
    will pass it to multible childs
- [x] make audio component 
- [x] make data component
- [x] make the data component work with generic data
- [x] use data component with pdf viewer
- [ ] only send data when channel is open
- [ ] create room associated with lecture
- [ ] destroy the room when lecture ends

-- in exam
- [ ] optain rtp packet or media in any format from Media stream
- [ ] send vedio over data channel


# integrate
- [x] sepration
- [x] speach to text
- [x] topic recognition
- [ ] continues training


# current abdo
- [x] populate data for exam
- [x] use exam from db
- [x] exam creation ui
- [x] report page after exam


- [ ] upload solution