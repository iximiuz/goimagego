# goimagego - how to container images in Go (showcase)

## Installation & Prerequisites

Debian 10
```bash
apt-get install pkg-config libgpgme-dev libbtrfs-dev libdevmapper-dev
```

CentOS 8:
```bash
yum install gpgme-devel libassuan-devel ostree-devel device-mapper-devel btrfs-progs-devel

yum install shadow-utils

echo "$LOGNAME:100000:65536" >> /etc/subuid
echo "$LOGNAM:100000:65536" >> /etc/subgid
```

## Usage

```bash
sudo LOG_LEVEL=trace `which go` run main.go \
    --run-root /home/vagrant/images_tmp/run \
     --root /home/vagrant/images_tmp/store  \
     pull docker://quay.io/prometheus/prometheus
```

## TODO:
- check usage on clear Debian/Centos
- document use cases

