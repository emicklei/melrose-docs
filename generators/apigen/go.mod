module generators

go 1.22

toolchain go1.22.5

require github.com/emicklei/melrose v0.0.0

require (
	github.com/Try431/EasyMIDI v1.0.3 // indirect
	github.com/emicklei/tre v1.7.0 // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	gitlab.com/gomidi/midi/v2/drivers/rtmididrv v0.15.0 // indirect
)

replace github.com/emicklei/melrose => ../../melrose
