package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"sort"
)

type terraformState struct {
	Modules []moduleState `json:"modules"`
}

type moduleState struct {
	Outputs map[string]string `json:"outputs,omitempty"`
	Resources map[string]resourceState `json:"resources"`
}

type resourceState struct {
	Type    string        `json:"type"`
	Primary instanceState `json:"primary"`
}

type instanceState struct {
	ID         string            `json:"id"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type terraformProvider struct {
	Path string
	Tfstate []byte
}

func (tp terraformProvider) list() map[string]interface{} {

	ts := tp.loadStateSource()

	inventory := make(map[string]interface{}, 0)
	meta := make(map[string]interface{}, 0)
	meta["hostvars"] = make(map[string]interface{}, 0)

	for name, value := range ts.Modules[0].Resources {
		switch {
			case value.Type == "digitalocean_droplet":
			host, attributes, groups := tp.digitaloceanList(name, value, ts)
			for _, group := range groups {
				if inventory[group] == nil {
					inventory[group] = make(map[string]interface{}, 0)					
				}
				if inventory[group].(map[string]interface{})["hosts"] == nil {
					inventory[group].(map[string]interface{})["hosts"] = make([]string, 0)					
				}
				inventory[group].(map[string]interface{})["hosts"] = append(inventory[group].(map[string]interface{})["hosts"].([]string), host)
				meta["hostvars"].(map[string]interface{})[host] = attributes
			}
		}
	}

	// TODO: This is so we can test consistently. Find a better way.
	for name := range inventory {
		sort.Strings(inventory[name].(map[string]interface{})["hosts"].([]string))		
	}

	inventory["_meta"] = meta
	return inventory
}
func (tp terraformProvider) digitaloceanList(name string, value resourceState, ts terraformState) (string, map[string]string, []string) {

	host := value.Primary.Attributes["name"]
	attributes := value.Primary.Attributes
	attributes["ansible_ssh_host"] = value.Primary.Attributes["ipv4_address"]
	attributes["ansible_ssh_user"] = "root"

	groups := make([]string, 0)
	groups = append(groups, "region=" + value.Primary.Attributes["region"])
	groups = append(groups, "size=" + value.Primary.Attributes["size"])
	//groups = append(groups, "public_ip=" + value.Primary.Attributes["ipv4_address"])
	groups = append(groups, "name=" + value.Primary.Attributes["name"])
	groups = append(groups, "type=" + value.Type)
	groups = append(groups, "role=" + strings.Split(name, ".")[1])

	return host, attributes, groups
}

func (tp terraformProvider) loadStateSource() terraformState {
	var err error

	if tp.Tfstate == nil {
		tp.Tfstate, err = ioutil.ReadFile(tp.Path)		
	}

	if err != nil {
		panic(err)
	}

	var ts terraformState
	err = json.Unmarshal(tp.Tfstate, &ts)

	if err != nil {
		panic(err)
	}

	return ts
} 

func (tp terraformProvider) host(hostName string) map[string]string {
	ts := tp.loadStateSource()
	attributes := make(map[string]string, 0)

	for _, value := range ts.Modules[0].Resources {
		if hostName == value.Primary.Attributes["name"] {
			attributes = value.Primary.Attributes
		}
	}

	return attributes
}