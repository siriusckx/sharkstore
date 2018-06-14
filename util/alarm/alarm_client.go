package alarm

import (
	"context"
	"google.golang.org/grpc"
	"time"
	"model/pkg/alarmpb"
	"model/pkg/taskpb"
	"util/deepcopy"
)

type AlarmClient interface {
	TaskTimeoutAlarm(clusterId int64, timeoutAlarm *alarmpb.TaskTimeout, task *taskpb.Task, desc string) error
	TaskLongTimeRunningAlarm(clusterId int64, longTimeRunningAlarm *alarmpb.TaskLongTimeRunning, task *taskpb.Task, desc string) error
	RangeNoHeartbeatAlarm(clusterId int64, rangeNoHbAlarm *alarmpb.RangeNoHeartbeatAlarm, desc string) error
	NodeNoHeartbeatAlarm(clusterId int64, nodeNoHbAlarm *alarmpb.NodeNoHeartbeatAlarm, desc string) error
	NodeDiskSizeAlarm(clusterId int64, nodeDiskSizeAlarm *alarmpb.NodeDiskSizeAlarm, desc string) error
	NodeLeaderCountAlarm(clusterId int64, nodeLeaderCountAlarm *alarmpb.NodeLeaderCountAlarm, desc string) error
	SimpleAlarm(clusterId int64, message string) error
	Close()
}

type Client struct {
	conn *grpc.ClientConn
}

func NewAlarmClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Close() {
	if c == nil {
		return
	}
	c.conn.Close()
}

func (c *Client) TaskTimeoutAlarm(clusterId int64, timeoutAlarm *alarmpb.TaskTimeout, task *taskpb.Task, desc string) error {
	if c == nil {
		return nil
	}
	cli := alarmpb.NewAlarmClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := cli.TaskAlarm(ctx, &alarmpb.TaskAlarmRequest{
		Head: &alarmpb.RequestHeader{ClusterId: clusterId},
		Type: alarmpb.TaskAlarmType_TIMEOUT,
		Task: deepcopy.Iface(task).(*taskpb.Task),
		TaskTimeoutAlarm: timeoutAlarm,
		Describe: desc,
	})
	return err
}

func (c *Client) TaskLongTimeRunningAlarm(clusterId int64, longTimeRunningAlarm *alarmpb.TaskLongTimeRunning, task *taskpb.Task, desc string) error {
	if c == nil {
		return nil
	}
	cli := alarmpb.NewAlarmClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := cli.TaskAlarm(ctx, &alarmpb.TaskAlarmRequest{
		Head: &alarmpb.RequestHeader{ClusterId: clusterId},
		Type: alarmpb.TaskAlarmType_LONG_TIME_RUNNING,
		Task: deepcopy.Iface(task).(*taskpb.Task),
		TaskLongTimeRunningAlarm: longTimeRunningAlarm,
		Describe: desc,
	})
	return err
}

func (c *Client) RangeNoHeartbeatAlarm(clusterId int64, rangeNoHbAlarm *alarmpb.RangeNoHeartbeatAlarm, desc string) error {
	if c == nil {
		return nil
	}
	cli := alarmpb.NewAlarmClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := cli.NodeRangeAlarm(ctx, &alarmpb.NodeRangeAlarmRequest{
		Head: &alarmpb.RequestHeader{ClusterId: clusterId},
		Type: alarmpb.NodeRangeAlarmType_RANGE_NO_HEARTBEAT,
		RangeNoHbAlarm: rangeNoHbAlarm,
		Describe: desc,
	})
	return err
}

func (c *Client) NodeNoHeartbeatAlarm(clusterId int64, nodeNoHbAlarm *alarmpb.NodeNoHeartbeatAlarm, desc string) error {
	if c == nil {
		return nil
	}
	cli := alarmpb.NewAlarmClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := cli.NodeRangeAlarm(ctx, &alarmpb.NodeRangeAlarmRequest{
		Head: &alarmpb.RequestHeader{ClusterId: clusterId},
		Type: alarmpb.NodeRangeAlarmType_NODE_NO_HEARTBEAT,
		NodeNoHbAlarm: nodeNoHbAlarm,
		Describe: desc,
	})
	return err
}

func (c *Client) NodeDiskSizeAlarm(clusterId int64, nodeDiskSizeAlarm *alarmpb.NodeDiskSizeAlarm, desc string) error {
	if c == nil {
		return nil
	}
	cli := alarmpb.NewAlarmClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := cli.NodeRangeAlarm(ctx, &alarmpb.NodeRangeAlarmRequest{
		Head: &alarmpb.RequestHeader{ClusterId: clusterId},
		Type: alarmpb.NodeRangeAlarmType_NODE_DISK_SIZE,
		NodeDiskSizeAlarm: nodeDiskSizeAlarm,
		Describe: desc,
	})
	return err
}

func (c *Client) NodeLeaderCountAlarm(clusterId int64, nodeLeaderCountAlarm *alarmpb.NodeLeaderCountAlarm, desc string) error {
	if c == nil {
		return nil
	}
	cli := alarmpb.NewAlarmClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := cli.NodeRangeAlarm(ctx, &alarmpb.NodeRangeAlarmRequest{
		Head: &alarmpb.RequestHeader{ClusterId: clusterId},
		Type: alarmpb.NodeRangeAlarmType_NODE_LEADER_COUNT,
		NodeLeaderCountAlarm: nodeLeaderCountAlarm,
		Describe: desc,
	})
	return err
}

func (c *Client) AliveAlarm() error {
	if c == nil {
		return nil
	}
	cli := alarmpb.NewAlarmClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := cli.AliveAlarm(ctx, &alarmpb.AliveRequest{
		// todo
	})
	return err
}

func (c *Client) SimpleAlarm(clusterId uint64, title, content string) error {
	if c == nil {
		return nil
	}
	cli := alarmpb.NewAlarmClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := cli.SimpleAlarm(ctx, &alarmpb.SimpleRequest{
		Head: &alarmpb.RequestHeader{ClusterId: int64(clusterId)},
		Title: title,
		Content: content,
	})
	return err
}

