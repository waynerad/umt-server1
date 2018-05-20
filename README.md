umt-server1 is a web server written in Golang that handles the Ultimate Music Toy (UMT) web app. The composition engine in JavaScript is incorporated (as of 2016-03-27) from the umt repository where it was originally developed as a standalone web page into umt.go here. The file ws.go in the main directory is a web server that simply picks up the web requests and hands them off to the relevant app; each app has its own package in Go. Web requests for UMT are simply handed off to umt.go, and everything else is handled from there. The only other package needed is the login package which identifies who logged in and assigns everyone user IDs.

This is called server1 because the plan is to have 2 servers: one that runs on the internet that handles saving and loading of tunes (or more precisely, of parameter sets), and basic administrative chores like that. It will create MIDI files and maybe someday handle things like creation of sheet music. The other server, server2, will run on the same machine as the web interface, and will handle real-time instrument playback over non-WebAudio API systems, such as MIDI and OSC. As of right now, server2 hasn't been started yet, so only server1 exists.

The Ultimate Music Toy (UMT) is a music composition system that lets you interactively change the composition while it's playing, and immediately hear the results!

It runs in the web browser and is written entirely in Javascript and uses the WebAudio API in HTML5 to actually play the music.

Tested to work on Chrome but should work on any HTML5-compliant browser.

by Wayne Radinsky

Demo

There is a live demo at http://ultimatemusictoy.org/ .

Conventions

Goal of doing everything the simplest possible way.

- All UMT code must pass JSLint. No exceptions. JSLint is treated as the target language. This includes stylistic conventions. None of the default options for disabling checks are allowed. The comment line at the beginning of the code specifies all the JSLint flags that will be used, which is currently only one: browser: true.

- Modularity is achieved by prefixing all the functions with umt. Yes, this is more tedious than other methods, but it is the simplest possible way I could think of and unambiguous. All data is put into a single global variable called gUmt. This should eliminate variable name collisions if UMT is used as part of another project. In general, I tried to design the system using the simplest possible design decisions at each stage, so the code is as unambiguous as possible and as easy to understand as possible. (Not saying it's *actually* easy to understand, since there is a certain level of unavoidable complexity, only that an effort was made to make it easy and reduce ambiguity at every stage as much as possible.)

- The UI is done with only jQuery sliders. (And built-in HTML controls such as listboxes). There are no custom UI elements and no other UI elements in the system. This was partly for simplicity and partly because my primary focus has been getting the actual music composition working properly.

- Since jQuery sliders are shadowed in data variables (stored in gUmt.UIParams), the decision was made to shadow all UI imputs, and implement a function that transfers data *from* the shadowed variables to the UI -- this will make it possible to "load" settings saved earlier into the UI. If you develop an alternate UI for UMT you will need to implement this.


