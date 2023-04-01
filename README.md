# Cosmart
Pre-Req:
1. have golang installed in your machine
2. have postman installed in your machine


Open in any editor for this case i use visual studio code
Open terminal in visual studio code
In terminal type in go build
then type in ./test
if all is installed correctly the program will run and shows "Welcome to The Public Library" in the terminal
now go to postman
in the link type in http://localhost:8080/pickup or http://localhost:8080/getlist
go to the body tab below the link and fill in the json field 
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
