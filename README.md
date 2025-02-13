# 🐶 Teo

Teo is a lightweight CLI tool for managing CCTV recordings. It provides an easy way to configure, record, and manage camera streams using simple commands and systemd integration.

## Documentation

For full details, check the manual:

    man teo

Or see a quick usage guide with:

    teo -h

## Building & Installation

Teo can be built and installed using:

    make && sudo make install

Teo requires `jq` and `ffmpeg` to be installed for parsing JSON configurations and recording videos.

## Getting Started

After installation, configure Teo with:

    teo configure -c config.json

Then, start recording manually:

    teo record -c cam1 -u rtsp://example.com/stream -d /var/cctv/cam1

Or enable automatic recording via systemd:

    systemctl enable --now teo.target

Enjoy using Teo!
