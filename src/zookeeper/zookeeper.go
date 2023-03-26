package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-zookeeper/zk"
)

type ZookeeperConfig struct {
	Servers    []string
	RootPath   string
	MasterPath string
}

type ElectionManager struct {
	ZKClient *zk.Conn
	ZKConfig *ZookeeperConfig
	IsMaster chan bool
}

func NewElectionManager(zkConfig *ZookeeperConfig, isMaster chan bool) *ElectionManager {
	electionManager := &ElectionManager{
		nil,
		zkConfig,
		isMaster,
	}
	electionManager.initConnection()
	return electionManager
}

func (e *ElectionManager) Run() {
	err := e.electMaster()
	if err != nil {
		fmt.Println("elect master error, ", err.Error())
	}
	e.watchMaster()
}

func (e *ElectionManager) isConnected() bool {
	if e.ZKClient == nil {
		return false
	} else if e.ZKClient.State() != zk.StateConnected {
		return false
	}
	return true
}

func (e *ElectionManager) initConnection() error {
	if !e.isConnected() {
		conn, connChan, err := zk.Connect(e.ZKConfig.Servers, time.Second)
		if err != nil {
			return err
		}

		for {
			isConnected := false
			select {
			case connEvent := <-connChan:
				if connEvent.State == zk.StateConnected {
					isConnected = true
					fmt.Println("connect to zookeeper server success!")
				}
			case _ = <-time.After(time.Second * 3):
				return errors.New("connect to zookeeper server timeout!")
			}
			if isConnected {
				break
			}
		}
		e.ZKClient = conn
	}
	return nil
}

func (e *ElectionManager) electMaster() error {
	err := e.initConnection()
	if err != nil {
		return err
	}

	isExist, _, err := e.ZKClient.Exists(e.ZKConfig.RootPath)
	if err != nil {
		return err
	}
	if !isExist {
		path, err := e.ZKClient.Create(e.ZKConfig.RootPath, nil, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			return err
		}

		if e.ZKConfig.RootPath != path {
			return errors.New("Create returned different path " + e.ZKConfig.RootPath + " != " + path)
		}
	}

	masterPath := e.ZKConfig.RootPath + e.ZKConfig.MasterPath
	path, err := e.ZKClient.Create(masterPath, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err == nil {
		if path == masterPath {
			fmt.Println("elct master success!")
			e.IsMaster <- true
		} else {
			return errors.New("Create returned different path " + masterPath + " != " + path)
		}
	} else {
		fmt.Println("elect master failed, ", err.Error())
		e.IsMaster <- false
	}
	return nil
}

func (e *ElectionManager) watchMaster() {
	for {
		children, state, childCh, err := e.ZKClient.ChildrenW(e.ZKConfig.RootPath + e.ZKConfig.MasterPath)
		if err != nil {
			fmt.Print("watch children error, ", err.Error())
		}
		fmt.Println("watch children result, ", children, state)
		select {
		case childEvent := <-childCh:
			if childEvent.Type == zk.EventNodeDeleted {
				fmt.Println("receive znode delete event, ", childEvent)

				fmt.Println("start elect new master ...")
				err = e.electMaster()
				if err != nil {
					fmt.Println("elect new master error, ", err)
				}
			}

		}
	}
}

func main() {
	config := &ZookeeperConfig{
		Servers:    []string{"127.0.0.1"},
		RootPath:   "/ElectMasterDemo",
		MasterPath: "/master",
	}

	isMasterChan := make(chan bool)

	var isMaster bool

	em := NewElectionManager(config, isMasterChan)
	go em.Run()

	for {
		select {
		case isMaster = <-isMasterChan:
			if isMaster {
				fmt.Println("do some job in master")
			}

		}
	}

}
