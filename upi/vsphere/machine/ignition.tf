provider "ignition" {
  version = "1.1.0"
}

locals {
  mask = "${element(split("/", var.machine_cidr), 1)}"
  gw   = "${cidrhost(var.machine_cidr,1)}"

  ignition_encoded = "data:text/plain;charset=utf-8;base64,${base64encode(var.ignition)}"
}

data "ignition_file" "hostname" {
  count = "${var.instance_count}"

  filesystem = "root"
  path       = "/etc/hostname"
  mode       = "420"

  content {
    content = "${var.name}0${count.index + 1}l"
  }
}

data "ignition_file" "static_ip" {
  count = "${var.instance_count}"

  filesystem = "root"
  path       = "/etc/sysconfig/network-scripts/ifcfg-Wired_connection_1"
  mode       = "420"

  content {
    content = <<EOF
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=none
IPADDR=${local.ip_addresses[count.index]}
PREFIX=23
GATEWAY=10.6.202.1
DNS1=10.6.202.20
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=stable-privacy
NAME="Wired connection 1"
DEVICE=ens192
ONBOOT=yes
AUTOCONNECT_PRIORITY=-999
EOF
  }
}

data "ignition_systemd_unit" "restart" {
  count = "${var.instance_count}"

  name = "restart.service"

  content = <<EOF
[Unit]
ConditionFirstBoot=yes
[Service]
Type=idle
ExecStart=/sbin/reboot
[Install]
WantedBy=multi-user.target
EOF
}

data "ignition_config" "ign" {
  count = "${var.instance_count}"

  append {
    source = "${var.ignition_url != "" ? var.ignition_url : local.ignition_encoded}"
  }

  systemd = [
    "${data.ignition_systemd_unit.restart.*.id[count.index]}",
  ]

  files = [
    "${data.ignition_file.hostname.*.id[count.index]}",
    "${data.ignition_file.static_ip.*.id[count.index]}",
  ]
}
