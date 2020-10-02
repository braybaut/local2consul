package consul

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/consul/api"
)

func consulClient(url string) (*api.Client, error) {
	config := api.DefaultConfig()
	config.Address = url

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

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
