# DB for Marshall school of Business

This is a csv parser that I have made to count the attendance, as well as other operations for the Marshall School of business. 
It uses some simple native libraries in Golang to total them up, and also a great way to read in a little more on 
concurrency!

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

To use this, simply run the executable file, and then move the finished CSV file to wherever you wish to have it located on your machine. It's as simple as that!

# Updates to Come
	
Ver 3.0 will have the following:
- Cut out the need to move the file yourself
- containerize the project using Docker.
- Use postgres to hold the stored data (because remote databases are fun to work with! :) )


