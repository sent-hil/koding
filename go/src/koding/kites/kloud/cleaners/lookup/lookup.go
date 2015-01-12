package lookup

import (
	"fmt"
	"koding/kites/kloud/multiec2"
	"sync"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

type Lookup struct {
	// values contains a list of instance tags that are identified as test
	// instances. By default all instances are fetched.
	values []string

	clients *multiec2.Clients
}

func NewAWS(auth aws.Auth) *Lookup {
	return &Lookup{
		clients: multiec2.New(auth, []string{
			"us-east-1",
			"ap-southeast-1",
			"us-west-2",
			"eu-west-1",
		}),
	}
}

// Instances returns all instances that belongs to the given client/region if
func (l *Lookup) Instances(client *ec2.EC2) (Instances, error) {
	instances := make([]ec2.Instance, 0)

	opts := &ec2.InstancesOpts{
		MaxResults: 900,
	}

	resp, err := client.InstancesWithOpts([]string{}, opts)
	if err != nil {
		return nil, err
	}

	for _, res := range resp.Reservations {
		instances = append(instances, res.Instances...)
	}

	nextToken := resp.NextToken

	// get all results until nextToken is empty
	for nextToken != "" {
		opts := &ec2.InstancesOpts{
			NextToken: nextToken,
		}

		resp, err := client.InstancesWithOpts([]string{}, opts)
		if err != nil {
			return nil, err
		}

		for _, res := range resp.Reservations {
			instances = append(instances, res.Instances...)
		}

		nextToken = resp.NextToken
	}

	m := make(Instances, len(instances))

	for _, instance := range instances {
		m[instance.InstanceId] = instance
	}

	return m, nil
}

// FetchInstances fetches all instances from all regions
func (l *Lookup) FetchInstances() *MultiInstances {
	var wg sync.WaitGroup

	allInstances := NewMultiInstanes(l.clients)

	for region, client := range l.clients.Regions() {
		wg.Add(1)
		go func(region string, client *ec2.EC2) {
			defer wg.Done()

			instances, err := l.Instances(client)
			if err != nil {
				fmt.Printf("[%s] fetching error: %s\n", region, err)
				return
			}

			allInstances.Add(client, instances)
		}(region, client)
	}

	wg.Wait()

	return allInstances
}
