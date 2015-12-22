package zkclient

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fsouza/go-dockerclient"
	"github.com/samuel/go-zookeeper/zk"
)

type Client struct {
	dkrClient *docker.Client
	zkConn    *zk.Conn
}

var err error

func New() *Client {
	return &Client{}
}
func (c *Client) Init(zkIP string) {
	c.dkrClient, err = docker.NewClientFromEnv()
	handleError(err)
	c.zkConn, _, err = zk.Connect([]string{zkIP}, time.Second) //TODO :: need to get ip of zookeeper from docker
	handleError(err)
	time.Sleep(time.Second)
}

func (c *Client) Connect(containerName string, zkPath string) {

	prevcontainers := 0

	for {
		containers, err := c.dkrClient.ListContainers(docker.ListContainersOptions{All: false, Filters: map[string][]string{"name": {containerName}, "status": {"running"}}})
		handleError(err)
		if len(containers) == prevcontainers {
			continue
		} else {
			if len(containers) == 1 {
				_, err := c.zkConn.Create(zkPath, []byte("localhost:"+strconv.FormatInt(containers[0].Ports[0].PrivatePort, 10)), int32(zk.FlagEphemeral)|int32(zk.FlagSequence), zk.WorldACL(zk.PermAll))
				handleError(err)
				fmt.Println("[zookeeper] : Server registered")
			} else {
				err := c.zkConn.Delete(zkPath, -1)
				handleError(err)
				fmt.Println("[zookeeper] : Server deleted")
			}
		}
		prevcontainers = len(containers)
	}

}

//handleError stops the program execution
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
