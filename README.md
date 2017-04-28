# Custom pam module.

This is a pam module skeleton used as base to develop custom authentication methods.

It is based on [Uber's SSH certificate pam module](https://github.com/uber/pam-ussh) and simplified to get the basic
of the pam structure.

Creates a docker container with an OpenSSH server and the pam module configured to be used on ssh authentication.
Default pam authentication developed only accepts user with 'victor' as username.

## Building:

1. Clone the repo and run 'docker-compose up'
```
  $ git clone github.com/uber/pam-ussh
```

2. Modify the 'pamAuthenticate' method in 'pam_custom.go' to your own authentication method.
Current simple authentication checks if the username is 'victor'.

3. Modify docker-compose.yml file line 9 if you want to use ubuntu docker container.

4. Run docker image
```
$ docker-compose up --build
```

## Usage:

A docker container will be up with ssh server listening on port 2222 of the docker host.
You just need to add the user 'victor' with a password and connect using ssh to check the pam module is working.

1. Enter the container
```
$ docker exec -ti pamcustom_ssh-server_1 bash
```
2. Add the user 'victor' and set its password and exit the container
```
$ useradd victor
$ passwd victor
$ exit
```

3. From outside the container connect to it using ssh
```
$ ssh victor@localhost -p 2222
```

4. You can create another user to check authentication fails.
