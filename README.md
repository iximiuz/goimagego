# goimagego - work with container images in Go (showcase)

It's a demo program showing how to pull, store, and mount container images using
<a href="https://github.com/containers/image">github.com/containers/image</a> and
<a href="https://github.com/containers/storage">github.com/containers/storage</a>
libraries.

The demo focuses on `overlay` and `vfs` storage drivers on top of `xfs`. In theory, other drivers and file systems supprted by <a href="https://github.com/containers/storage">github.com/containers/storage</a> library should work as well but some extra effort may be needed.

## Installation & Prerequisites

The following extra packages are required:

```bash
# Debian 10
apt-get install pkg-config libgpgme-dev libdevmapper-dev

# CentOS 8 (with enabled EPEL Repository)
yum install libassuan-devel gpgme-devel device-mapper-devel
```

Additionally, `policy.json` file should be created with the mininal content:

```bash
mkdir /etc/containers

cat <<EOF > /etc/containers/policy.json
{
    "default": [
        {
            "type": "insecureAcceptAnything"
        }
    ],
    "transports": {
        "docker": {}
    }
}
EOF
```

From source code (using `go modules`):

```bash
git clone https://github.com/iximiuz/goimagego
cd goimagego
go build -tags "exclude_graphdriver_btrfs"

./goimagego help
goimagego - work with container images in Go

Usage:
  goimagego [flags]
  goimagego [command]

Available Commands:
  container   create container
  containers  list local containers
  delete      delete container, image, or layer
  help        Help about any command
  images      list local images
  layers      list local layers
  mount       mount container
  pull        pull image from remote repository
  store       show store info
  unmount     unmount container
  wipe        wipe the whole storage

Flags:
  -d, --driver string     image store driver (overlay, vfs, etc) (default "overlay")
  -h, --help              help for goimagego
  -r, --root string       image store root directory (default "/var/lib/containers/storage")
  -R, --run-root string   image store run root directory (default "/var/run/containers/storage")
```

## Usage

Currently, all the examples require `sudo`.

### Show storage info

```bash
# Using `overlay` storage driver
$ ./goimagego store
root = /var/lib/containers/storage
run-root = /var/run/containers/storage
driver = overlay
driver options = []
status =
([][2]string) (len=4 cap=4) {
 ([2]string) (len=2 cap=2) {
  (string) (len=18) "Backing Filesystem",
  (string) (len=3) "xfs"
 },
 ([2]string) (len=2 cap=2) {
  (string) (len=15) "Supports d_type",
  (string) (len=4) "true"
 },
 ([2]string) (len=2 cap=2) {
  (string) (len=19) "Native Overlay Diff",
  (string) (len=4) "true"
 },
 ([2]string) (len=2 cap=2) {
  (string) (len=14) "Using metacopy",
  (string) (len=5) "false"
 }
}

# Using `vfs` storage driver
$ ./goimagego -d vfs store
root = /var/lib/containers/storage
run-root = /var/run/containers/storage
driver = vfs
driver options = []
status =
([][2]string) <nil>

# Using custom storage location
./goimagego -d vfs -r /home/vagrant/images store
root = /home/vagrant/images
run-root = /var/run/containers/storage
driver = vfs
driver options = []
status =
([][2]string) <nil>
```

### Pull image

```bash
# Using `docker` transport and default registry (docker.io)
./goimagego pull docker://alpine:latest
Pulling image docker://alpine:latest
Getting image source signatures
Copying blob c9b1b535fdd9 done
Copying config e7d92cdc71 done
Writing manifest to image destination
Storing signatures
Image pulled - {
   "schemaVersion": 2,
   "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
   "config": {
      "mediaType": "application/vnd.docker.container.image.v1+json",
      "size": 1511,
      "digest": "sha256:e7d92cdc71feacf90708cb59182d0df1b911f8ae022d29e8e95d75ca6a99776a"
   },
   "layers": [
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 2802957,
         "digest": "sha256:c9b1b535fdd91a9855fb7f82348177e5f019329a58c53c47272962dd60f71fc9"
      }
   ]
}

# Using `docker` transport and manually specified registry (quay.io):
./goimagego pull docker://quay.io/prometheus/prometheus
Pulling image docker://quay.io/prometheus/prometheus
Getting image source signatures
Copying blob 626a2a3fee8c done
Copying blob 0f8c40e1270f done
Copying blob 280e865d0f46 done
Copying blob 81d2279d1c55 done
Copying blob 1402c1f8faad done
Copying blob e7ed030afda4 done
Copying blob 40c7beb2b8e0 done
Copying blob c1be047355d9 done
Copying blob 5e958f95e7b4 done
Copying blob fb780b8f81a9 done
Copying blob d81ddb9e06a9 done
Copying blob 8b293a391a3d done
Copying config e935122ab1 done
Writing manifest to image destination
Storing signatures
Image pulled - {
   "schemaVersion": 2,
   "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
   "config": {
      "mediaType": "application/vnd.docker.container.image.v1+json",
      "size": 6669,
      "digest": "sha256:e935122ab143a64d92ed1fbb27d030cf6e2f0258207be1baf1b509c466aeeb42"
   },
   "layers": [
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 761056,
         "digest": "sha256:0f8c40e1270f10d085dda8ce12b7c5b17cd808f055df5a7222f54837ca0feae0"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 478856,
         "digest": "sha256:626a2a3fee8c6a9b5b866adc6cb15d54b5d901b6a084a2519bf7f905325b0711"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 33165393,
         "digest": "sha256:e7ed030afda42ce32fdae9cc4d86002133a3f24fc7561dee73febb04de7a58fe"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 19897006,
         "digest": "sha256:1402c1f8faadc85692e41acf6bb6f13744295af3eb09a262900f8067c98ee325"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 602,
         "digest": "sha256:81d2279d1c55bf5372cdd79788cd7760f3186c2843b9d04e8040af6a26b74a32"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 2660,
         "digest": "sha256:280e865d0f465093732ac66573d6ccf249a56939b8efcd2e8f4b96438e0ad191"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 3059,
         "digest": "sha256:40c7beb2b8e03e0cc4fd595d545b987ee8f39d680dd69122aa75d3a068c108e6"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 4019,
         "digest": "sha256:c1be047355d91ba313295e3d1d84d6e129a420c4a8d3aa02f140f41836048d08"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 1271,
         "digest": "sha256:d81ddb9e06a9f686e66f3ff903e720607b01b0d0072f290ebefb214381e3f7dd"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 114886,
         "digest": "sha256:fb780b8f81a94ad7937d9d9d2c68255d2cf7b921de821b59aeaec1fd506908de"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 183,
         "digest": "sha256:5e958f95e7b4a3b6de2f075d095267c54c69dd28da9b700e07a1d40aaee9aed3"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 691,
         "digest": "sha256:8b293a391a3d1a6544032a5f6bef74fb0cd4e421e319a72886df8a6efcc8aaa1"
      }
   ]
}
```

### List local, images, layers, containers

```bash
$ ./goimagego <images|layers|contaienrs>
```


### Create, mount, and umount container

```bash
# Pull image
$ ./goimagego pull docker://alpine:latest
Pulling image docker://alpine:latest
Getting image source signatures
Copying blob c9b1b535fdd9 skipped: already exists
Copying config e7d92cdc71 done
Writing manifest to image destination
Storing signatures
Image pulled - {
   "schemaVersion": 2,
   "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
   "config": {
      "mediaType": "application/vnd.docker.container.image.v1+json",
      "size": 1511,
      "digest": "sha256:e7d92cdc71feacf90708cb59182d0df1b911f8ae022d29e8e95d75ca6a99776a"
   },
   "layers": [
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 2802957,
         "digest": "sha256:c9b1b535fdd91a9855fb7f82348177e5f019329a58c53c47272962dd60f71fc9"
      }
   ]
}

# Create a new container using the image from above
$ ./goimagego container e7d92cdc71feacf90708cb59182d0df1b911f8ae022d29e8e95d75ca6a99776a
Container:
(*storage.Container)(0xc000089790)({
 ID: (string) (len=64) "f7c2136928fbe8e963b594833f0101b964edb5ec299444c2508ce1d3d15ef3c2",
 Names: ([]string) <nil>,
 ImageID: (string) (len=64) "e7d92cdc71feacf90708cb59182d0df1b911f8ae022d29e8e95d75ca6a99776a",
 LayerID: (string) (len=64) "5d531d0194a79ad4a40232cae1346012739380c0d7c354485d70ddc298ea0d9e",
 Metadata: (string) "",
 BigDataNames: ([]string) <nil>,
 BigDataSizes: (map[string]int64) {
 },
 BigDataDigests: (map[string]digest.Digest) {
 },
 Created: (time.Time) 2020-02-29 17:31:14.129769831 +0000 UTC,
 UIDMap: ([]idtools.IDMap) <nil>,
 GIDMap: ([]idtools.IDMap) <nil>,
 Flags: (map[string]interface {}) (len=2) {
  (string) (len=12) "ProcessLabel": (string) "",
  (string) (len=10) "MountLabel": (string) ""
 }
})

# Mount the created container
$ ./goimagego mount f7c2136928fbe8e963b594833f0101b964edb5ec299444c2508ce1d3d15ef3c2
/var/lib/containers/storage/overlay/5d531d0194a79ad4a40232cae1346012739380c0d7c354485d70ddc298ea0d9e/merged

$ df /var/lib/containers/storage/overlay/5d531d0194a79ad4a40232cae1346012739380c0d7c354485d70ddc298ea0d9e/merged
Filesystem     1K-blocks    Used Available Use% Mounted on
overlay         10474496 5666980   4807516  55% /var/lib/containers/storage/overlay/5d531d0194a79ad4a40232cae1346012739380c0d7c354485d70ddc298ea0d9e/merged

$ ./goimagego unmount f7c2136928fbe8e963b594833f0101b964edb5ec299444c2508ce1d3d15ef3c2
Unmounted!
```

