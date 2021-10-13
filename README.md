# Zoom Launcher

## Installation 

```
go install
```

This will produce a zoom-launcher binary in your go/bin directory


## Configuration

2 Config Files are required

### 1
~/.config/zoom-launcher/config.yaml

Content:

```
path_to_zooms: <path to your zooms file>
```

### 2
~/.config/zoom-launcher/zooms.yaml

Content:

```
zooms:
    meeting-1: <https url to your meeting room 1>
    meeting-2: <https url to your meeting room 2>
    meeting-3: <https url to your meeting room 3>
```

## Usage

```
zoom-lauchner start meeting-1
```

This will open zoom call with meeting room 1







