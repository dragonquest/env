# env 
I ported the env linux shell command to go so it can be used under Windows as well.

More info: https://en.wikipedia.org/wiki/Env

# example
````shell
env GOOS=linux GOARCH=arm GOARM=5 go build
```` 

# limitations 
Only very basic version 