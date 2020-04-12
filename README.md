# The-Engine

## Design

### What am I doing?

Something I'd like to call a bot. This will always be active when my Linux is booted up (maybe deploy on an RPi later)

* It should do on demand activities.
* It should also push timely notifications.

### How will I do it?

* A frickin backend server on Golang -> packed as a docker image.
* Multiple modules attach and listen on.
* A notification queue (data will be pushed to it, it'll doll it up and send out to the user) (Use channels!)
* Should scrape the web for data / hook on to APIs / webhooks.
* Should know what time it is rn. Obviously.

### What should I keep in mind?
* Should be extremely lightweight. (How? Maybe run benchmarks and mem logs)
* Should be fast. (Go will handle it?)
* Should adapt and be modular.