# data representation


client browser --> backend SERVER ---> ML processing --> save back at backend server


# questions 
    - when to train ??
      - train button 
      - train after lecture
    - lecture recording will have genral languge


course 
year
processed
models

## before/during lecture
    - befor the lecture professor MAY upload one or more files related to lecture those 
    files are processed and saved as text file which can be used latter for training

    - during the lecture the audio is recorded and processed to text which can be used alongside 
    the data from pdf files 
    
    - after the year end the ML is trained on 
    [
        admin can choose general
        professor choose unrelated materials or v2, upload refrance for it
        professor choose files for his lecture 
    ] 
    ### responsibilty
        - backend
            - [x] save pdf 
            - [ ] save multible pdf for same lecture // v2
            - [ ] save audio for lecture
            - [ ] save text data of pdf
            - [ ] save text data of audio

        - ML_HA
            - [ ] create end point to recieve and process data
            - [x] process pdf files to text
            - [x] process audio files to text









### backend server responsibilty
    - save a text for lecture 
        - processing happen at ML_HA
    
    - save text for lecture generated from speach to text service 
        - processing happen at ML_HA

### ML_HA responsibilty
    - process lecture audio and genrate text from the speach
    - train the model in text data from all-lectures and pdf for course



## after all term end

    - model for every course
        - save two model one for predition , one for retraining (around 100xMB )


## during exam 

  - student send sound of at exam 
  - server save the recording and send it to ML
  |
  |
  |
  ---------------------------------------------------
  |--> seprate voices to three sounds (the whole exam) |                          
  |--> voice activity detector remove non voices       |                    
  |--> convert each files from speach to text          |                 
  |--> classification the text                         |
  |--> send report for speaches related to topic       |
  ---------------------------------------------------- 
  - save the report concatenate incident 


  __ if no one is there the seprated files will be empty







- ML server responsibilty
    - 


