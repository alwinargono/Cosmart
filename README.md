# Cosmart
Pre-Req:
1. have golang installed in your machine
2. have postman installed in your machine


1. Open in any editor for this case i use visual studio code
2. Open terminal in visual studio code
3. In terminal type in go build
4. In terminal type in ./test
5. if all is installed correctly the program will run and shows "Welcome to The Public Library" in the terminal
6. now go to postman
7. in the link type in http://localhost:8080/pickup or http://localhost:8080/getlist
8. go to the body tab below the link and fill in the json field 
<img width="1182" alt="Screenshot 2023-04-01 at 5 32 45 PM" src="https://user-images.githubusercontent.com/29112567/229282279-6ee40461-e7e1-4630-bbb3-a07f83f61163.png">

use this for pickup:
{
    "title": "The Picture of Dorian Gray",
    "edition_number": 618,
    "date": "27/4/2023"
}

<img width="1183" alt="Screenshot 2023-04-01 at 5 41 24 PM" src="https://user-images.githubusercontent.com/29112567/229282739-eabc90e8-d2c3-4ab1-9218-60b90676256a.png">

use this for getlist:
{
    "genre": "horror"
}
