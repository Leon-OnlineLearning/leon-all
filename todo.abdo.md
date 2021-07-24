# high priority
- [ ] use mobile as secodary camera
- [ ] add start lecture at student
- [ ] fix memory error in forbidden object
- [ ] better look for lecture page
- [ ]fix no embedding
- [ ] add before exam prep page



# fix 
-- for exam 
## creating exam
- [x] send created exam to db
## during exam
- [ ] prevent changing tab during exam
- [ ] make sure face in region during record
- [ ] use face-api to make sure someone in front of camera
- [x] make question come one by one
- [ ] create admin ui to check recorded video


-- for lecture
- [x] test creating multiple janus objects _
    this will create multilble sessions , don't know if this any bad 
    may be it is not , but for know there is a parent that create a session and 
    will pass it to multible childs
- [x] make audio component 
- [x] make data component
- [x] make the data component work with generic data
- [x] use data component with pdf viewer
- [x] only send data when channel is open
- [x] create room associated with lecture
    - [ ] keep the calender working
    - [ ] make get all events for professor
    - [x] create room when professor access its page
    - [ ] add wait messsage for students if room isn't created yet
- [x] destroy the room when lecture ends
- [ ] get more refrences from lectures



# current 
- [ ] prevent changing tab during exam
- [ ] support secondary camera by phone
- [ ] make sure face in region during record




# done 
- [X] upload solution

- [x] populate data for exam
- [x] use exam from db
- [x] exam creation ui
- [x] report page after exam

- [x] send recorded video from client to server
- [x] send from server to ML
- [x] save result came back from ML
- [x] view result after exam


