# NetDog
A simple re-implementation of Necat... but `dog` 🐶. One of many side projects designed for fun and continued professional development. A core criteria  is defined outlining the fundamental target features to implement and used as the skeleton for the project/code challenge.

## Core features
The core features that must be available in the end product:

- [ ] support the listen mode using TCP (as default) from the cli with flags `-l` to specify listen and `-p` to specify port e.g:
      ```bash 
      ndog -l -p 8888
      ```
      should be supported

- [ ] support UDP severs with cli flag `-u` to specify UDP host and `-p` for port e.g:
      ```bash 
      ndog -l -p -u 8888
      ```

- [ ] support the `-z` flag to ascertain if a server is listening on `one` or a range of ports, without sending data e.g:
      ```bash
      ndog -z localhost 8888
      Connection to localhost port 8888 [tcp] succeeded!
      And with a range of ports: ...

      ndog -z localhost 8880-8890
      Connection to localhost port 8888 [tcp] succeeded!
      ```

- [ ] support the `-e` flag to execute a process and pipe the input and output to/from it to/from the connected clientand pipe the out e.g:
      ```bash
      ndog -l -p 8888 -e /bin/bash
      ```

- [ ] support a `-x` flag that provides a hex dump between client and server. With the following output expected e.g:
    ```bash
      ndog -x -l -p 8888
        Hello from the client
        Received 22 bytes from the socket
        00000000  48 65 6C 6C  6F 20 66 72  6F 6D 20 74  68 65 20 63  
        00000010  6C 69 65 6E  74 0A
        Hello from the client. 

        Response from the server
        Sent 25 bytes to the socket
        00000000  52 65 73 70  6F 6E 73 65  20 66 72 6F  6D 20 74 68  
        00000010  65 20 73 65  72 76 65 72  0A 
    ```

