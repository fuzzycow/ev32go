# ev32go

** WARNING: This is a prototype / outdated project ! Use at your own risk ! **
EV3DEV sysfs layout has changed since this was written (2015) *
For ev3dev news and neater ev3dev go language bindings visit ev3dev.org *

Contents:
* clip - wrapper for easily connecting to ev3dev sysfs device and sensor "ports"
* docs
** spec - ev3dev API spec, used to generate ev3api (currently contains a copy of ev3dev spec.json file, GPL LICENSE. https://github.com/ev3dev/ev3dev-lang) 
* codegen - go code generation library for ev3dev spec.json
* cmd
** spec2go - utility used by go:generate to create go to ev3dev/sysfs language bindings, from ev3dev API description (spec.json)
* ev3api - Go API for LEGO EV3 / ev3dev
* drivers
** sysfs - ev3dev sysfs driver
** keypad - lego ev3 keypad driver
* helpers - robotics helpers
** monitor - monitor device and send change notifications over channel
** mqtt - (broken) mqtt logger
* robotics
** telemetry - telemetry logger (influxdb implementation provided)
** chassis, nav, pose -  partial port of lejos robotics framework navigation code
** pid -  experiment in user-space PID control, partially based on ev3dev kernel-level pid implementation
* examples - examples and demos
* bench - benchmarks to compare performance of different math library implementations on LEGO EV3 Platform (ARM5 non-FP CPU)


The code in this repository is (C) Fuzzycow.Org and is released under GPLv2 LICENSE
GPLv2 LICENSE was chosen to match the license of the ev3dev project, and ev3dev-lang spec.json file
