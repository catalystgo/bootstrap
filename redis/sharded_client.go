package redis

import (
	"context"
	"errors"
	"sync"

	"github.com/catalystgo/logger/logger"
	"github.com/redis/go-redis/v9"
	"github.com/serialx/hashring"
)

var (
	ErrNoNodes      = errors.New("no nodes available")
	ErrNodeNotFound = errors.New("node not found")
)

type ShardedClient interface {
	AddNode(ctx context.Context, address string) error
	RemoveNode(ctx context.Context, address string)
	GetShard(key string) (*redis.Client, error)
	Nodes() []string
	Close()
}

// shardedClient is a Redis client wrapper supporting sharding with consistent hashing
type shardedClient struct {
	hashRing   *hashring.HashRing
	nodesMap   map[string]*redis.Client
	nodesMutex sync.RWMutex
}

func NewShardedClient(ctx context.Context, shardAddresses []string) (ShardedClient, error) {
	client := &shardedClient{
		nodesMap: make(map[string]*redis.Client),
		hashRing: hashring.New(nil),
	}

	for _, addr := range shardAddresses {
		if err := client.AddNode(ctx, addr); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// AddNode adds a new node to the client
func (rsc *shardedClient) AddNode(ctx context.Context, address string) error {
	opt, err := redis.ParseURL(address)
	if err != nil {
		return err
	}

	client := redis.NewClient(opt)
	if _, err := client.Ping(ctx).Result(); err != nil {
		return err
	}

	rsc.nodesMutex.Lock()
	defer rsc.nodesMutex.Unlock()

	rsc.nodesMap[opt.Addr] = client
	rsc.hashRing = rsc.hashRing.AddNode(opt.Addr)

	return nil
}

// RemoveNode removes a node from the client
// Notice that address must be in the form of "host:port"
func (rsc *shardedClient) RemoveNode(ctx context.Context, address string) {
	rsc.nodesMutex.Lock()
	defer rsc.nodesMutex.Unlock()

	client := rsc.nodesMap[address]
	if client == nil {
		return
	}

	if err := client.Close(); err != nil {
		logger.Errorf(ctx, "error closing client: %v", err)
	}

	rsc.hashRing.RemoveNode(address)
}

// GetShard returns the shard responsible for a given key
func (rsc *shardedClient) GetShard(key string) (*redis.Client, error) {
	rsc.nodesMutex.RLock()
	defer rsc.nodesMutex.RUnlock()

	node, ok := rsc.hashRing.GetNode(key)
	if !ok {
		return nil, ErrNoNodes
	}

	client, ok := rsc.nodesMap[node]
	if !ok {
		return nil, ErrNodeNotFound
	}

	return client, nil
}

func (rsc *shardedClient) Nodes() []string {
	rsc.nodesMutex.RLock()
	defer rsc.nodesMutex.RUnlock()

	nodes := make([]string, 0, len(rsc.nodesMap))
	for node := range rsc.nodesMap {
		nodes = append(nodes, node)
	}

	return nodes
}

func (rsc *shardedClient) Close() {
	rsc.nodesMutex.Lock()
	defer rsc.nodesMutex.Unlock()

	for _, client := range rsc.nodesMap {
		if err := client.Close(); err != nil {
			logger.Errorf(context.Background(), "error closing client: %v", err)
		}
	}

	rsc.nodesMap = nil
	rsc.hashRing = nil
}
