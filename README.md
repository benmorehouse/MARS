# DB for Marshall school of Business

I implemented a MySQL database for the Marshall School of Business using Go. It keeps track of attendance for their current student
events as well as other in-house things to take care of. All that is needed from the user is an important csv file from their 
platform!

This serves as a taste of my work as a Data Intern at the school in Fall 2019. Read more below on how to run it!

# Installation

If you do not have homebrew, go here: https://brew.sh

Once you have homebrew, if you do not have go installed then enter
	
	brew install go

into your desired directory.

Once this is installed, pick the directory you wish to store this project, then 
clone this repository!

Finally, do the following command within the std file you downloaded: 
        go build 

Then you are all set to run this repository!
		
	./marshall_database_project

# Use

There are three very important flags that you will want to work with whilst using this application

-file -> what file you are parsing through. Without this flag, you will not be able too run the program

-outputFile -> The desired name of the output file. Will automatically be set to output.csv

-default -> If you are only looking for normal totals of attendance, pass this flag as true.

Example Execution:
	
	./marshall_database_project -file=input.csv -default=true -outputFile=niceOutput.csv

# Updates to Come
	
Ver 3.0 will have the following:
- Cut out the need to move the file yourself after operation with just a flag
- containerize the project using Docker.
- Use postgres to hold the stored data (because remote databases are fun to work with! :) )
- Update to count all teachers concurrently 
- use testing documentation gomock 


