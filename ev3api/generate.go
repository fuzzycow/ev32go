package ev3api

//go:generate spec2go -d motor  -t device.tmpl -o device/motor.gen.go
//go:generate spec2go -d dcMotor  -t device.tmpl -o device/dcmotor.gen.go
//go:generate spec2go -d servoMotor  -t device.tmpl -o device/servomotor.gen.go
//go:generate spec2go -d sensor  -t device.tmpl -o device/sensor.gen.go

//go:generate spec2go -d legoPort  -t device.tmpl -o device/legoport.gen.go
//go:generate spec2go -d powerSupply  -t device.tmpl -o device/powersupply.gen.go
//go:generate spec2go -d led  -t device.tmpl -o device/led.gen.go

//go:generate spec2go -d i2cSensor  -t device.tmpl -o device/sensor/i2csensor.gen.go
//go:generate spec2go -d ultrasonicSensor  -t device.tmpl -o device/sensor/ultrasonic.gen.go
//go:generate spec2go -d infraredSensor  -t device.tmpl -o device/sensor/infrared.gen.go
//go:generate spec2go -d gyroSensor  -t device.tmpl -o device/sensor/gyro.gen.go
//go:generate spec2go -d colorSensor  -t device.tmpl -o device/sensor/color.gen.go
//go:generate spec2go -d lightSensor  -t device.tmpl -o device/sensor/light.gen.go
