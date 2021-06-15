# Deputy 
[![GitHub release](https://img.shields.io/github/tag/techievee/deputy.svg?label=latest)](https://github.com/techievee/deputy/relesases)
[![Go Report Card](https://goreportcard.com/badge/github.com/techievee/deputy)](https://goreportcard.com/report/github.com/techievee/deputy)

> Golang 1.16  | cobra , promptui CLI   | interactive and non-interactive mode 

* [Data Structure](#data-structure)
* [Algorithm](#algorithm)  
* [Complexity](#complexity)
* [Running Pre-built Image](#running-pre-built-image)
* [Building image](#building-image)
* [Program Components](#program-components)
* [Command reference](#command-reference)
* [Program Specification](#program-specification)
* [Assumptions](#assumptions)
* [Libraries used](#libraries-used)


## Data structure
```
User and roles
--------------
map[roleId]-> [userId]->Pointer to user details

Role Hierarchy
--------------
 map [parent roleId] -> [subordinate roleId] -> Pointer to role details                                                                   
```


## Algorithm
```
Calculating Indexes
-------------------
for each users in the users file
    Add the user to the map of user roles
    Add the user to the map of userId
end

for each of the roles in the roles file
      Add the roles in the hash of the parent role key
end   


Finding Subordinates
--------------------
For the given user, for his role
   - Find all the subordinate roleid, from the role hierarchy map
   - Add all the roleid to the queue
   Repeat, till the queue is empty,
     For each role present in the queue, find the subordinates roleid and add them to the queue and process list
        While adding check if that role is already processed earlier
        Processed roles are kept in Set processed or not)
   end
  
  For all the subordinateid roles from the processlist sets
      Find the user for the role from the userlist and print them
  end
                                                              
```

## Complexity
```
Runtime complexity
------------------- 
Indexing: O(n), n-> number of users and roles
Subordinate calculaton: O(1), Amortized value, constant time, as the data is parsed from the index (ideally total users that are matching those subordinates)
```


## Running Pre-built Image
To Run the pre-built image, the image can be downloaded from the following releases link
- Download the release corresponding to your Operating System architecture and run the command<br/>The files need to be copied to the working directory
```
https://github.com/techievee/deputy/releases/tag/v1.0.0
```
## Building image
There are multiple ways to build an image
- custom build
```
git clone https://github.com/techievee/deputy.git
cd deputy
go build -trimpath -ldflags "-X main.Build=v1.0.0" -o ./deputy ./cmd
```
- using makefile (linux and mac)
```
git clone https://github.com/techievee/deputy.git
cd deputy
make all
make bin-darwin
make bin-windows
```
## Program Components
The program consists of the following components
* `console`: Helps in invoking the program and to run in either interactive or non-interactive CLI mode
* `cmd`: Command processor that helps in executing commands
* `indexer`: Index processor, which stores and process index for various file such as roles and users. Helps in finding the subordinates through index in constant time
* `data`: Data store that reads and process the json files from disk to memory structures.
* `indexedData`: Data structure that stores the index as a inverted index.
* `sets`: Data structure that stores and processes sets.
* `utilities`: Helper functions.

## Command reference
### Interactive Console
```
./deputy --help          

deputy is a CLI library that indexes and search json file for finding all the subordinates for the user.
This application is a tool that reads the csv files, indexes it and 
calculates the total licences required for the application.

Usage:
  deputy [flags]
  deputy [command]

Available Commands:
  help         Help about any command
  subordinates Deputy CLI Application to Index and Search Subordinates

Flags:
      --config string   config file (default is $HOME/deputy.yaml)
  -h, --help            help for deputy
  -R, --role string     Path of the JSON file to load Roles, Defaults to currentPath/roles.json 
  -U, --user string     Path of the JSON file to load Users, Defaults to currentPath/users.json


```
### Non-Interactive Console
```
 ./deputy subordinates --help

Deputy subordinates is a CLI library that indexes and search JSON file.
This application is a tool that reads the json file, indexes it and 
searches for subordinates without any user interactions.

Usage:
  deputy subordinates [flags]

Flags:
  -h, --help             help for subordinates
  -u, --user-id string   ID of the user, for whom the subordinates need to be calculated

Global Flags:
      --config string   config file (default is $HOME/deputy.yaml)
  -R, --roles string    Path of the JSON file to load Roles, Defaults to currentPath/roles.json 
  -U, --users string    Path of the JSON file to load Users, Defaults to currentPath/users.json


```
## Program Specification
### Interactive Console
1. Running the ```./deputy```command without any verb will invoke the Interactive CLI, the roles and the users input files need to be present in the directory where the command is present or need to override using 
    - -R or --role flags or DEPUTY_ROLE environment variable for roles
    - -U or --user flags or DEPUTY_USER environment variable for users
   <br/>Invoking interactive command without any CLI flags<br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/interactive_cli.png?raw=true)
   <br/>Invoking interactive command using the short flags<br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/shortflag.png?raw=true)
   <br/>Invoking interactive command using the full flags<br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/longflag.png?raw=true)
   <br/>Invoking interactive command after setting environment variable<br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/environment.png?raw=true)
   
2. Press Enter to pass the entry screen <br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/screen1.png?raw=true)
3. Use arrow keys and select the Find subordinates option <br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/screen2.png?raw=true)

4. Enter the user id for whom the subordinates need to be calculated<br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/screen3.png?raw=true)
   
While entering the user id string, it gets validated to passthrough the subordinates command, the following rules apply
 - If enter is selected without any value it is considered as empty search, and it won't be validated
- If alphanumeric or text is entered for field, it won't allow to passthroughs, until the input is corrected

  <br/>Invalid inputs<br/>
  ![alt text](https://github.com/techievee/deputy/blob/master/_images/invalidoutput.png?raw=true)


5. Result gets displayed after the input is validated, along with runtime statistics<br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/output.png?raw=true)

### Non-Interactive Console
1. Running the ```./deputy subordinates -u 1``` command (deputy with additional subordinates verb) will invoke the Non-interactive CLI for one time run or calculation, the users and roles input files need to be present in the directory where the command is present or need to override using
    - -R or --role flags or DEPUTY_ROLE environment variable for roles
    - -U or --user flags or DEPUTY_USER environment variable for users
   - -u or --user-id flags or DEPUTY_USER_ID environment variable for usersId 
   <br/><br/>Invoking non-interactive command with full flag name <br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/noncli_2.png?raw=true)
   <br/>Invoking non-interactive command using the short flags<br/>
   ![alt text](https://github.com/techievee/deputy/blob/master/_images/noninteractive_cli.png.png?raw=true)
   
## Assumptions
1. Each file is considered as a valid json file containing the array of json data, if data is noisy the parsing is halted

2. There can be a user without corresponding roles and roles without user, that doesn't impact loading and indexing

3. If the file contains noisy data  such as string for integer field are considered as invalid

4. The uniqueness for roles and user were not considered, when multiple user or roles with same id when present in the file, the data that was present first were considered

5. When the parent is not present for the roles, they are considered as root or parent=0, and the data is valid
   ```
   {
   "Id": 1,
   "Name": "System Administrator"
   }
   ```
6. When the other tags such as Id and Name for roles and all fields in the users when missing from the file, the file is considered as invalid

## Libraries used

* `cobra`- [CLI Engine] ([https://github.com/spf13/cobra])(https://github.com/spf13/cobra)
* `promptui`: [CLI Prompts] ([https://github.com/manifoldco/promptui.git])(https://github.com/manifoldco/promptui)
* `viper`: [configuration manager] ([https://github.com/spf13/viper])(https://github.com/spf13/viper)
