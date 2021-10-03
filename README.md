# genpwd
A CLI tool to print alot of possible combinations (approx 67k using default list) of passwords with words like `dev`, `admin` and some special characters that could've been used by the developer as the password of an admin panel. You can use your own custom list using `-cl` flag to get a list of possible passwords, but make sure that your custom list is not too big (mine is only of 12 words and makes a list of 67k possible passwords) otherwise you may run out of storage.

## Installation
```
$ go install github.com/meispi/genpwd@latest
```

## How to use
```
Usage of genpwd:
  -cc string
    	camelCase verison of company (default: camelCase on the middle character eg: comPany)
  -cl string
    	custom list for combination
  -l int
    	min length of password (default 6)
  -w string
    	Enter your word
```
e.g.: `genpwd -w meispi` will make a list named meispi.txt containing possible password combinations

Default list of words:
```
admin
Admin
ADMIN
administrator
Administrator
ADMINISTRATOR
dev
Dev
DEV
password
Password
PASSWORD
```


*WARNING*: Admin panels are generally not meant to be public, hence they lack rate limiting, so if you are going to try this tool to get the wordlist and fuzz the admin panel make sure that you have permission to do so else you can get into a lot of trouble.
