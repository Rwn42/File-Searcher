# File Search
Allows the user to upload files to the server with any amount of tags. Any files can then be searched by tag or date entered at a later time. The motivation behind building this was to add topics to my school notes so I can find them when exam period comes.

## Installation


Clone the repository, once cloned enter the repo and run the following commands
```sh
$ cd src
$ go build -o ../file_search
$ cd ..
$ ./file_search
```

the `file_search` executable is now ready to use. Edit the [config](./config.toml) to your liking.

no front end build process is required the app uses a single css, html and js file and alpine js for reactivity.

## Usage
**NOTE: Do not use this software on a server that is publically accesible.**\
**I have made no effort to add any security hosting the whole internet leaves your server vulnerbale to unlimited file uploads from anyone anywwhere**\
**If you want to allow file uploads to your server from anywhere add an authentication layer only you have the password for**

1. run the server executable
2. Navigate to the address provided to begin using the app.
3. To upload a file to the server from any device on the LAN (including the same computer running the server) fill out the form
on the right hand side.
4. To search for files on the server from any device on the LAN use the search bar on the left side of the screen

## Todo
use config file to edit index.html on first load to change host_name