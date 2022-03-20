## web2mgo
This test driver-connector for data exechange to MongoDB. 

This driver containts module of web-server with parser URL and driver-connector to MongoDB.
https://github.com/go-mgo/mgo/tree/v2

It's working like that:

> web-client(web browser) <---> local web-server/driver-connector <---> MongoDB

Write of data in DB execute like that:
>1. We click to our test URL.
>2. Doing request to module web-server on localhost:8080.
>3. Module web-server do parse string of URL, for getting substrings with our data.
>4. Module driver-connector is connecting to MongoDB is saving our strings to it. If strings already is there, it output warning about. 
All sended strings, we can see in web-browser.
  
To MongoDB is writting all data after symbol "/" in URL. 
Where executing parseURL, getting slice of string with Name before symbol ":" and getting slice of string with Data after symbol ":"

URL should be like that:

    http://localhost:8080/Bill:+79000000001

This module-driver one can use is working with MongoDB in electronic documentmanagement system Directum RX and etc.
