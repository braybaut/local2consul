package consul

import (
	"bufio"
	"bulk-upload-to-consul/structs"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"io/ioutil"
	"os"
	"strings"
)

func Unmarshalconfig(file string) structs.Configurationvalues {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	var keyvalues structs.Configurationvalues

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &keyvalues)
	return keyvalues
}

func PutKeyValueJson(keyvalues structs.Configurationvalues, domain string) error {
	for i := 0; i < len(keyvalues); i++ {
		fmt.Println(keyvalues[i].ConsulAddress)
		fmt.Println(keyvalues[i].Environment)

		for j := 0; j < len(keyvalues[i].Keyvalues); j++ {
			fmt.Println(keyvalues[i].Keyvalues[j].Key)
			fmt.Println(keyvalues[i].Keyvalues[j].Value)

			err := PutKeyValuev2(keyvalues[i].ConsulAddress, domain, keyvalues[i].Keyvalues[j].Key, keyvalues[i].Keyvalues[j].Value)
			if err != nil {
				return fmt.Errorf("error to Put the values: %d", err)
			}
		}
	}
	return nil
}

func consulClient(url string) (*api.Client, error) {
	config := api.DefaultConfig()
	config.Address = url

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func PutKeyValuev2(url string, domain string, key string, value string) error {
	fmt.Printf("url: %s \n", url)
	fmt.Printf("domain: %s \n", domain)
	fmt.Printf("file: %s \n", key)
	fmt.Printf("file: %s \n", value)

	client, err := consulClient(url)
	if err != nil {
		return fmt.Errorf("Error to create Consul Client %d", client)
	}
	kv := client.KV()

	p := &api.KVPair{Key: fmt.Sprint(domain + key), Value: []byte(value)}
	_, err = kv.Put(p, nil)
	if err != nil {
		return err
	}
	fmt.Printf("Key %v, loaded \n ", value)
	fmt.Print("\n")
	return nil
}

// PutKeyValue : main function to upload
func PutKeyValue(url string, domain string, filewithKeyValues string) error {

	fmt.Printf("url: %s \n", url)
	fmt.Printf("domain: %s \n", domain)
	fmt.Printf("file: %s \n", filewithKeyValues)

	client, err := consulClient(url)
	if err != nil {
		return fmt.Errorf("Error to create Consul client %d", client)
	}

	kv := client.KV()

	file, err := os.Open(filewithKeyValues)
	if err != nil {
		return fmt.Errorf("Error to get file with key and values %d", file)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "=")
		fmt.Printf("Trying to put the key %v, and value %v \n ", split[0], split[1])
		p := &api.KVPair{Key: fmt.Sprint(domain + split[0]), Value: []byte(split[1])}
		_, err = kv.Put(p, nil)
		if err != nil {
			return err
		}
		fmt.Printf("Key %v, loaded \n ", split[0])
		fmt.Print("\n")

	}
	return nil
}
