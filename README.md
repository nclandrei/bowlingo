## BowlinGO structure

The backend code is stored inside the root repository, being part of the *bowlingo* package. The command used for running 
the server is stored inside the *cmd* folder. The backend has its dependency management handled by *dep* and it is dockerized 
(please see Dockerfile). There are as few dependencies as possible used. Also, for reasoning on how to structure the app and 
various other conventions, there are many comments scattered across the files explaining these decision (together with cool 
links as well :-)). It is built using Go 1.10.3 (not on 1.11 as it is on beta at the moment :-().

All the Angular frontend is stored inside the web folder. It was built using Angular 5, TypeScript and less.js.

## How to run it

For the backend to run, you can simply run (from the root folder):
```bash 
	$ docker build -t <image_name_here> .
	$ docker run -p <host_port_here>:80 <image_name_here>
```

Otherwise, you can issue the following commands (from the root folder):
```bash
	$ dep ensure
	$ go build cmd/main.go
	$ ./main -port=<host_port_here> -addr=<hostname_here>
```

**Note**: The server port used by the frontend is hardcoded to 8000. This is one of the possible improvements of the app: store the 
hostname and server port inside environment variables. This can be easily done using either .env files (which are not version 
controlled) or an encrypted JSON file (using, for example, gpg symmetrical encryption). 

For the frontend to run, you can simply run (from the root folder):
```bash
	$ cd web/
	$ yarn
	$ ng serve
```

## How to test it

Simply run, from the root folder:
```bash
	$ go test -bench=. -cover -v
```

Hope you like it and I am really looking forward for feedback and improvement ideas :-). Have a great week!