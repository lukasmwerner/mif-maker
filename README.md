# mif-maker

## Description
This program converts images from modern file formats to a memory initialization
file used by Intel's Quartus FPGA suite.

The program works with `png` and `jpeg` files.

Each pixel in the mif will be 16 bits wide to represent 4 bits per channel in
RGBA color, regardless of whether the input image was in RGBA or not.

> [!WARNING]
> Attention: The output of the program is limited to 256x256 pixels to ensure a
> maximum of 65536 words, which is the maximum size of a ROM block using the
> Quartus II Prime ROM creation tool. Any dimensions of an image larger that 256
> pixels will be cropped to 256.

## Installation & Usage

There are a few different ways that this program can be installed.

1. On a personal computer
    1. Download the binary for your respective computer and CPU from the [release page](https://github.com/lukasmwerner/mif-maker/releases) e.g. `mif-maker-windows-amd64` for x86_64 Windows computers
    2. Open a command prompt in the folder you downloaded it to
    3. Run the program e.g. `.\mif-maker-windows-amd64` on Windows x86_64

> [!WARNING]
> Important If you are an Oregon State University student following these
> instructions, and you have never accessed the school servers before, visit
> https://teach.engr.oregonstate.edu to active your engineering resources.
2. Oregon State Engineering Servers
    1. Download the program for Linux x86_64 with the following command: `wget https://github.com/lukasmwerner/mif-maker/releases/download/v0.1.2/mif-maker-linux-amd64`
    2. Ensure the program is executable `chmod +x ./mif-maker-linux-amd64`
    3. Run the program with `./mif-maker-linux-amd64`

### Mac Demo
![](/images/mac-demo.gif)


### Windows Demo
![](/images/mac-demo.gif)

### Version Table
| CPU                              | Operating System| Binary/Release  |
| -------------------------------- | --------------- | --------------- |
| ARM CPU(e.g. Snapdragon X Elite) | Windows         | [mif-maker-windows-arm64.exe](https://github.com/lukasmwerner/mif-maker/releases/download/v0.2.0/mif-maker-windows-arm64.exe) |
| Intel/AMD CPU                    | Windows         | [mif-maker-windows-amd64.exe](https://github.com/lukasmwerner/mif-maker/releases/download/v0.2.0/mif-maker-windows-amd64.exe) |
| ARM CPU(Apple M Series)          | macOS           | [mif-maker-darwin-arm64](https://github.com/lukasmwerner/mif-maker/releases/download/v0.2.0/mif-maker-darwin-arm64) |
| Intel CPU                        | macOS           | [mif-maker-darwin-amd64](https://github.com/lukasmwerner/mif-maker/releases/download/v0.2.0/mif-maker-darwin-amd64) |
| Intel/AMD CPU                    | Linux & *Nix    | [mif-maker-linux-amd64](https://github.com/lukasmwerner/mif-maker/releases/download/v0.2.0/mif-maker-linux-amd64) |




### Acknowledgements
Thank you [p-bodson](https://github.com/p-bodson/) for making [the original program](https://github.com/p-bodson/mifMaker) that I based mine off of.

