# Dummy StatsD

This is a simple dummy StatsD server.

This is useful in that it lets you inspect the raw StatsD messages a system is sending as they are sent,
and without having to deal with the whole Graphite, etc setup.

You run this from the command line, and then can see raw StatsD messages coming in, as they are coming in.


## Usage

The TL;DR version is, run the command:

    go run dummy-statsd.go

(More details below.)

This Dummy StatsD server is written in programming language Go (which is also called Golang).
Thus if you want to compile this dummy StatsD server from the source code, you will need Go on your system.

On some Linux systems, if Go is not alreaedy installed you can install it with:

    sudo apt-get install golang-go

Other systems will have other means of installing Go.

Once Go is installed on your system, the quickest way to run this dummy StatsD server is to run it
straight from the source code with:

    go run dummy-statsd.go

(This method is probably sufficient for software engineers and other developers looking to inspect the StatsD
messages being sent from their application.)

Alternatively, you can also compile the source code, into a binary executable file, with the command:

    go build dummy-statsd.go

On Linux based systems, this will create a binary executable file named: "dummy-statsd"

And then run the resultant binary executable file, with a command such as:

    ./dummy-statsd

