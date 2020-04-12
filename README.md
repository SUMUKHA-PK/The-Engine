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

### What will it do?
* Take input for a command, return a popup notification on the screen / push notifications.
* Types of "taking commands in" :
    * A curl command to a general API endpoint of The Engine.
    * Periodic updates.
* Inputs can come from:
    * Reddit
    * The web in general
    * My email (maybe google API?)
    * ....
* Log all notifications!
* Should know what devices are online (when its implemented remotely) and push notifications to those/ all or some, as desired. Client in device is necessary.
* Play media! Find a structured way for input queries. Say "Play House MD S1 E2"


## Implementation

* A good notification pusher for ubuntu - https://www.maketecheasier.com/desktop-notifications-for-linux-command/
* Since I want to make concurrent command processing, some kind of garbage collection should happen based on the timestamps of the commands running.
* Above demand needs to have a well structured logger of all the commands being run and must be able to cancel them, report their errors. (maybe use context?)
