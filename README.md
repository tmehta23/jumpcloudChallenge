README.md
Author: Tarj Mehta

***Directions - To create executable***
Cd to source/ directory from the jumpcloudChallenge directory. This directory includes all source code for the Jumpcloud coding challenge.

To compile, from source directory, run command:
-go build -o <desired_name_for_exe> *

ie, "go build -o passwordSrvr.exe * ". This will create the
executable with all the desired functionality.

***Directions - To run/check executable:***
1. Run executable from one terminal
2. Open second terminal window (to be used for curl commands) and enter the following commands-

  - **For password hashing/encoding:**
    - In your "curl commands" terminal window, type in the command "curl --data "password=angryMonkey" http://localhost:8080/hash
    - The hashed and encoded password should display after about 5 seconds.
    - Try to send a request without a password specified.
    - An error message should display to the console.

  - **For stats:**
    - From your "curl commands" terminal window, run a handful of hash requests through and then type in the command "curl http://localhost:8080/stats". You should see the total number of requests sent and the average time per request (in microseconds).

  - **For shutdown:**
    - For shutdown, in your "curl commands" window, type in the command line "curl http://localhost:8080/shutdown"
    - In the terminal window used for curl commands, you will see "Recv failure: Connection was reset" error and in the terminal window used for running the executable, you will see the message "Shutdown has been successfully completed. The program will not take new requests at this time. "
