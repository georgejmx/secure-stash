# Secure Stash

Encrypted key value store for secure information

## Build and Run

_Requires go at the version in `go.mod` and redis version 7.* installed_

### Development

- Run `echo 'REDIS_PASSWORD=examplepass\nREDIS_PORT=5680' > .env` to initialise the app in dev mode
- `docker run -d -p 6380:6379 redis` to spin up a dev cache
- `make run-source` build and run source code

### Production

#### Initial Setup

##### Linux

1. Set up Redis locally on your system following [this guide](https://linuxhint.com/install-configure-redis-ubuntu-22-04/)
2. Set environment variables in this directory with the above password and the default redis port; `echo 'REDIS_PASSWORD=examplepass\nREDIS_PORT=6379' > .env`

##### Mac

1. Ensure that redis is installed and running locally on your computer by first executing `brew install redis`. Then add a password to the _/opt/homebrew/etc/redis.conf_ file on your system by running `echo 'requirepass examplepass' >> redis.conf`. Now run `brew services start redis`
2. Set environment variables in this directory with the above password and the default redis port; `echo 'REDIS_PASSWORD=examplepass\nREDIS_PORT=5679' > .env`

#### Start program

Run `make run` to start **Secure Stash**

#### Stopping and restarting cache

- Optionally run `brew service stop redis` to stop before logging out
  - Before restarting **Secure Stash**, will need to run `brew service start redis`
- Optionally run `redis-cli SAVE` to save a backup of the program that can be loaded into a different device

#### Deleting cash

Execute `redis-cli -a examplepass FLUSHALL`
