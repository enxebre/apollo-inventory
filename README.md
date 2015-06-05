Apollo Inventory
================

[![wercker status](https://app.wercker.com/status/dae34b479f1643eb3a43bfc38b3e69eb/s "wercker status")](https://app.wercker.com/project/bykey/dae34b479f1643eb3a43bfc38b3e69eb)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

This is WIP. Currently only Digitalocean is supported.

This is a [Ansible Dynamic Inventory](http://docs.ansible.com/intro_dynamic_inventory.html) for [Apollo](https://github.com/Capgemini/Apollo) allowing pluggable clouds and system providers.

It use [Terraform](https://github.com/hashicorp/terraform) as a default provider for generating the inventory by parsing a given Terraform state file.

It allows to create new providers by implementing inventoryProvider interface:

```go
type inventoryProvider interface {
    list() map[string]interface{}
	host(string) map[string]string
}
````


## Usage

```bash
apollo-inventory --pretty --path=/path_to/terraform.tfstate
```

The default value for --path can be set as an environment variable via:

```
export TF_VAR_state_path=/path_to/terraform.tfstate
```

For Digital ocean you will see something like:

```json
{
  "_meta": {
    "hostvars": {
      "apollo-mesos-master-0": {
        "ansible_ssh_host": "46.101.4.15",
        "ansible_ssh_user": "root",
        "id": "5506398",
        "image": "11919335",
        "ipv4_address": "46.101.4.15",
        "ipv4_address_private": "10.131.166.113",
        "locked": "false",
        "name": "apollo-mesos-master-0",
        "private_networking": "true",
        "region": "lon1",
        "size": "512mb",
        "ssh_keys.#": "1",
        "ssh_keys.0": "821202",
        "status": "active"
      },
      "apollo-mesos-master-1": {
        "ansible_ssh_host": "46.101.3.149",
        "ansible_ssh_user": "root",
        "id": "5506399",
        "image": "11919335",
        "ipv4_address": "46.101.3.149",
        "ipv4_address_private": "10.131.163.45",
        "locked": "false",
        "name": "apollo-mesos-master-1",
        "private_networking": "true",
        "region": "lon1",
        "size": "512mb",
        "ssh_keys.#": "1",
        "ssh_keys.0": "821202",
        "status": "active"
      },
      "apollo-mesos-master-2": {
        "ansible_ssh_host": "46.101.3.229",
        "ansible_ssh_user": "root",
        "id": "5506400",
        "image": "11919335",
        "ipv4_address": "46.101.3.229",
        "ipv4_address_private": "10.131.163.57",
        "locked": "false",
        "name": "apollo-mesos-master-2",
        "private_networking": "true",
        "region": "lon1",
        "size": "512mb",
        "ssh_keys.#": "1",
        "ssh_keys.0": "821202",
        "status": "active"
      },
      "apollo-mesos-slave-0": {
        "ansible_ssh_host": "46.101.4.14",
        "ansible_ssh_user": "root",
        "id": "5506409",
        "image": "11919335",
        "ipv4_address": "46.101.4.14",
        "ipv4_address_private": "10.131.166.115",
        "locked": "false",
        "name": "apollo-mesos-slave-0",
        "private_networking": "true",
        "region": "lon1",
        "size": "512mb",
        "ssh_keys.#": "1",
        "ssh_keys.0": "821202",
        "status": "active"
      }
    }
  },
  "name=apollo-mesos-master-0": {
    "hosts": [
      "apollo-mesos-master-0"
    ]
  },
  "name=apollo-mesos-master-1": {
    "hosts": [
      "apollo-mesos-master-1"
    ]
  },
  "name=apollo-mesos-master-2": {
    "hosts": [
      "apollo-mesos-master-2"
    ]
  },
  "name=apollo-mesos-slave-0": {
    "hosts": [
      "apollo-mesos-slave-0"
    ]
  },
  "region=lon1": {
    "hosts": [
      "apollo-mesos-master-0",
      "apollo-mesos-master-1",
      "apollo-mesos-master-2",
      "apollo-mesos-slave-0"
    ]
  },
  "role=mesos-master": {
    "hosts": [
      "apollo-mesos-master-0",
      "apollo-mesos-master-1",
      "apollo-mesos-master-2"
    ]
  },
  "role=mesos-slave": {
    "hosts": [
      "apollo-mesos-slave-0"
    ]
  },
  "size=512mb": {
    "hosts": [
      "apollo-mesos-master-0",
      "apollo-mesos-master-1",
      "apollo-mesos-master-2",
      "apollo-mesos-slave-0"
    ]
  },
  "type=digitalocean_droplet": {
    "hosts": [
      "apollo-mesos-master-0",
      "apollo-mesos-master-1",
      "apollo-mesos-master-2",
      "apollo-mesos-slave-0"
    ]
  }
```

```bash
apollo-inventory --host=hostname --pretty --path=/path_to/terraform.tfstate
```

For Digitalocean you will see somehting like:

```json
{
  "id": "5506398",
  "image": "11919335",
  "ipv4_address": "46.101.4.15",
  "ipv4_address_private": "10.131.166.113",
  "locked": "false",
  "name": "apollo-mesos-master-0",
  "private_networking": "true",
  "region": "lon1",
  "size": "512mb",
  "ssh_keys.#": "1",
  "ssh_keys.0": "821202",
  "status": "active"
}
```