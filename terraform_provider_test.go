package main

import (
	"strings"
	"io/ioutil"
	"testing"
	"reflect"
)

const terraformStateFixture = `
{
    "version": 1,
    "serial": 5,
    "modules": [
        {
            "path": [
                "root"
            ],
            "outputs": {
                "master.1.ip": "46.101.4.15",
                "master.2.ip": "46.101.3.149",
                "master.3.ip": "46.101.3.229",
                "master_ips": "46.101.4.15,46.101.3.149,46.101.3.229",
                "slave_ips": "46.101.4.14"
            },
            "resources": {
                "atlas_artifact.mesos-master": {
                    "type": "atlas_artifact",
                    "primary": {
                        "id": "11919335",
                        "attributes": {
                            "file_url": "",
                            "id": "11919335",
                            "metadata_full.#": "2",
                            "metadata_full.created_at": "{{timestamp}}",
                            "metadata_full.version": "{{user version}}+{{timestamp}}",
                            "name": "enxebre/apollo-ubuntu-14.04-amd64",
                            "slug": "Enxebre/apollo-ubuntu-14.04-amd64/digitalocean.image/2",
                            "type": "digitalocean.image",
                            "version": "2"
                        }
                    }
                },
                "atlas_artifact.mesos-slave": {
                    "type": "atlas_artifact",
                    "primary": {
                        "id": "11919335",
                        "attributes": {
                            "file_url": "",
                            "id": "11919335",
                            "metadata_full.#": "2",
                            "metadata_full.created_at": "{{timestamp}}",
                            "metadata_full.version": "{{user version}}+{{timestamp}}",
                            "name": "enxebre/apollo-ubuntu-14.04-amd64",
                            "slug": "Enxebre/apollo-ubuntu-14.04-amd64/digitalocean.image/2",
                            "type": "digitalocean.image",
                            "version": "2"
                        }
                    }
                },
                "digitalocean_droplet.mesos-master.0": {
                    "type": "digitalocean_droplet",
                    "depends_on": [
                        "atlas_artifact.mesos-master",
                        "digitalocean_ssh_key.default",
                        "digitalocean_ssh_key.default"
                    ],
                    "primary": {
                        "id": "5506398",
                        "attributes": {
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
                    }
                },
                "digitalocean_droplet.mesos-master.1": {
                    "type": "digitalocean_droplet",
                    "depends_on": [
                        "atlas_artifact.mesos-master",
                        "digitalocean_ssh_key.default",
                        "digitalocean_ssh_key.default"
                    ],
                    "primary": {
                        "id": "5506399",
                        "attributes": {
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
                        }
                    }
                },
                "digitalocean_droplet.mesos-master.2": {
                    "type": "digitalocean_droplet",
                    "depends_on": [
                        "atlas_artifact.mesos-master",
                        "digitalocean_ssh_key.default",
                        "digitalocean_ssh_key.default"
                    ],
                    "primary": {
                        "id": "5506400",
                        "attributes": {
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
                        }
                    }
                },
                "digitalocean_droplet.mesos-slave": {
                    "type": "digitalocean_droplet",
                    "depends_on": [
                        "atlas_artifact.mesos-slave",
                        "digitalocean_droplet.mesos-master",
                        "digitalocean_ssh_key.default"
                    ],
                    "primary": {
                        "id": "5506409",
                        "attributes": {
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
                "digitalocean_ssh_key.default": {
                    "type": "digitalocean_ssh_key",
                    "primary": {
                        "id": "821202",
                        "attributes": {
                            "fingerprint": "xxx",
                            "id": "821202",
                            "name": "Apollo",
                            "public_key": "xxx"
                        }
                    }
                }
            }
        }
    ]
}
`

func TestTerraformProviderList(t *testing.T) {

	data := strings.NewReader(terraformStateFixture)
	ts, err := ioutil.ReadAll(data)
	
	if err != nil {
		panic(err)
	}

	p := terraformProvider{ Tfstate: ts }
	got := p.list()
	want := map[string]interface{} {
		  "_meta": map[string]interface{} {
		    "hostvars": map[string]interface{} {
		      "apollo-mesos-master-0": map[string]string {
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
		        "status": "active",
		      },
		      "apollo-mesos-master-1": map[string]string {
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
		        "status": "active",
		      },
		      "apollo-mesos-master-2": map[string]string {
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
		        "status": "active",
		      },
		      "apollo-mesos-slave-0": map[string]string {
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
		        "status": "active",
		      },
		    },
		  },
		  "name=apollo-mesos-master-0": map[string]interface{} {
		    "hosts": []string {
		      "apollo-mesos-master-0",
		    },
		  },
		  "name=apollo-mesos-master-1": map[string]interface{} {
		    "hosts": []string {
		      "apollo-mesos-master-1",
		    },
		  },
		  "name=apollo-mesos-master-2": map[string]interface{} {
		    "hosts": []string {
		      "apollo-mesos-master-2",
		    },
		  },
		  "name=apollo-mesos-slave-0": map[string]interface{} {
		    "hosts": []string {
		      "apollo-mesos-slave-0",
		    },
		  },
		  "region=lon1": map[string]interface{} {
		    "hosts": []string {
		      "apollo-mesos-master-0",
		      "apollo-mesos-master-1",
		      "apollo-mesos-master-2",
		      "apollo-mesos-slave-0",
		    },
		  },
		  "role=mesos-master": map[string]interface{} {
		    "hosts": []string {
		      "apollo-mesos-master-0",
		      "apollo-mesos-master-1",
		      "apollo-mesos-master-2",
		    },
		  },
		  "role=mesos-slave": map[string]interface{} {
		    "hosts": []string {
		      "apollo-mesos-slave-0",
		    },
		  },
		  "size=512mb": map[string]interface{} {
		    "hosts": []string {
		      "apollo-mesos-master-0",
		      "apollo-mesos-master-1",
		      "apollo-mesos-master-2",
		      "apollo-mesos-slave-0",
		    },
		  },
		  "type=digitalocean_droplet": map[string]interface{} {
		    "hosts": []string {
		      "apollo-mesos-master-0",
		      "apollo-mesos-master-1",
		      "apollo-mesos-master-2",
		      "apollo-mesos-slave-0",
		    },
		  },
		}

	if ! reflect.DeepEqual(got, want) {
		t.Errorf("p.list() => %q, want %q", got, want)
	}

}
