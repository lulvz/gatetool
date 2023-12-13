# Gatetool
A tool for the raspberry pi to allow for easy creation of scripts that can be ran from a web api.

## Requirements
- rtl-sdr
- rpitx
- raspberry pi (any model)
- go

## Install
**Chances are the install script doesn't work currently, since it's an old project and I haven't tested it in a while.**
Run install.sh as a normal user and input the password when prompted to.
You can use sudo to run it but it create the files in the /root/ directory

## Functonality
Will list your scripts added to the config file in the web api, allowing for modular scripts to be added to the system.
Can be expanded with another web app to allow for a nice interface to interact with the scripts.
The scripts can also be ran manually using the template provided in the script_example.sh file.

## Raspberry pi setup
The raspberry pi needs to be setup with any copper wire in a GPIO pin, I used GPIO 4, as described in the [rpitx documentation](https://github.com/F5OEO/rpitx).
The rtl-sdr dongle is not used in conjunction with the raspberry pi unless you want to capture the signal from your remote and save it as a .wav file using the raspberry pi as well, which I do not reccomend since it's a slower process and it's much easier to just use a computer to capture the signal and save it as a .wav file.

### Soldering the copper cable to the raspberry pi
It's not necessary to solder the copper cable, you can just wrap it around the GPIO pin, but I reccomend soldering it to the pin to make sure it doesn't come loose, I also gained a bit of range by soldering it to the pin.

### Use of another antennas
I was able to open a gate at 100m range with a random television antenna I had lying around, I was pretty impressed with the range, but since this setup emits a lot of noise in other frequencies, I don't reccomend using it with a big antenna, since it will probably interfere with other devices.

## Capture the signal
To capture the signal you need to use a computer with the rtl-sdr dongle and any software that allows you to capture the signal and save it as a .wav file.

## Considerations about signal transmission using rpitx
Our configuration is very crude and it emits a lot of noise in other frequencies, I used the rtl-sdr dongle to check and besides emitting the initial signal of my gates at 433.92mhz it also emits a lot of noise in other frequencies.
This is possibly not a big problem, since the emittions are fairly small and the signal is not very strong, but it's something to consider.

### Disclaimer
The golang part of this project was mostly an old learning project that I decided to just recently pickup again and polish a little bit.
I had to build on top of old code and It's not very pretty, but the most important part of this project was the learning experience I got using a raspberry pi 0 w heaadless, the rtl-sdr dongle and just general knowledge about linux, go, bash scripting and also a bit of wireless communication.
Overall I'm very happy with the result and I hope to expand on this project in the future.
