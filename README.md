# ev32go

__WARNING:__ This is a prototype / outdated project ! Use at your own risk.

For ev3dev news and neater ev3dev go language bindings visit ev3dev.org.
EV3DEV sysfs layout / spec has changed since this was written(2015). 
Use the provided code generator (spec2go) to regenerate ev3api go code for the updated ev3dev spec file.


Project goals:
* Provide ev3dev go language bindings, and ability to rebuild the bindings from spec
* Help novice ev3dev and go developers - use code generation to enable use of code-completion and contextual help IDE features
* Support navigation for Differential (2 motor) and Holonomic (Omni / 3 motor) Robitic chassis
* Publish telemetry for live display and storage (e.g.: via InfluxDB and Grafana)
* Enable the user to write CSP / State Machine -based robots with ease, and minimal ev3-specific boiler plate code
* Try to avoid programming techniques which are slow on Go and ARM5 CPU


Contents:
* clip - wrapper for easily opening ev3dev devices/ports. Use this to create new device/port objects.
* docs
    * spec - ev3dev API spec, used to generate ev3api (currently contains a copy of ev3dev spec.json file, GPL LICENSE. See [ev3dev-lang repo](https://github.com/ev3dev/ev3dev-lang) ) 
* cmd/spec2go - utility used to generate ev3api from ev3dev spec.json, called via go:generate, uses codegen library
* codegen - library and templates for generating device and sensor classes from ev3dev spec.json
* ev3api - Generated Go API for LEGO EV3 / ev3dev
* drivers
    * sysfs - ev3dev sysfs driver. Includes two implementations: 
        * direct - sysfs files are open/closed on each access
        * cached - sysfs files are cached after opening. caching/open/close is handled automatically. Provides significant performance boost to sysfs access on EV3 platform.
    * keypad - lego ev3 keypad driver
* helpers - robotics helpers
    * monitor - monitor device and send change notifications over channel
    * mqtt - (broken) mqtt logger
* robotics
    * telemetry - telemetry logger (influxdb implementation provided)
    * chassis, nav, pose -  partial port of lejos robotics framework navigation code. Supports both diffirential and holonomic (omni) robotic chassis.
    * pid -  experiments in user-space PID control, partially based on ev3dev kernel-level pid implementation
* examples - examples and demos
    * holocontrol / diffcontrol - state machine -based robot control example. Robot follows IR remote while IR remote is in beacon/seek mode, but switches to IR direct control mode, if directional control command is recieved.      
* bench - benchmarks to compare performance of different math library implementations on LEGO EV3 Platform (ARM5 non-FP CPU)

To build on Linux x86 for LEGO EV3 ARM5 no-FP target:
`GOARM=5 GOARCH=arm go build github.com/fuzzycow/ev32go/....`



__Pictures of robots and grabbers used with this project can be found [at my mocpages page](http://www.moc-pages.com/home.php/114748)__
and [home page](http://www.fuzzycow.org./fuzzybots/holonomic/overview/gallery1)


The code in this repository is (C) Fuzzycow.Org and is released under GPLv2 LICENSE

GPLv2 LICENSE was chosen to match the license of the ev3dev project, and ev3dev-lang spec.json file


