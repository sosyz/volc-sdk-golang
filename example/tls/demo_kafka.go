package tls

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

func initTestProject(client tls.Client) (testProjectID string, testTopicID string, err error) {
	//新建project
	createResp, err := client.CreateProject(&tls.CreateProjectRequest{
		ProjectName: testPrefix + uuid.NewString(),
		Description: "",
		Region:      testRegion,
	})
	if err != nil {
		return "", "", err
	}
	testProjectID = createResp.ProjectID

	// 新建topic
	// TopicName Description字段规范参考api文档
	createTopicRequest := &tls.CreateTopicRequest{
		ProjectID:   testProjectID,
		TopicName:   testPrefix + uuid.NewString(),
		Ttl:         30,
		ShardCount:  2,
		Description: "topic desc",
	}
	topic, err := client.CreateTopic(createTopicRequest)
	testTopicID = topic.TopicID
	if err != nil {
		return "", "", err
	}

	//新建index，开启全文索引和kv索引
	createIndexReq := &tls.CreateIndexRequest{
		TopicID: testTopicID,
		FullText: &tls.FullTextInfo{
			CaseSensitive:  false,
			IncludeChinese: false,
			Delimiter:      ", ?",
		},
	}
	_, err = client.CreateIndex(createIndexReq)
	if err != nil {
		return "", "", err
	}
	return testProjectID, testTopicID, nil
}

func main() {
	//初始化客户端，配置AccessKeyID,AccessKeySecret,region,securityToken;securityToken可以为空
	client := tls.NewClient(testEndPoint, testAk, testSk, testSessionToken, testRegion)

	var (
		testProjectID string
		testTopicID   string
	)

	testProjectID, testTopicID, err := initTestProject(client)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("ProjectID:%s, TopicID:%s", testProjectID, testTopicID)

	// 开启 Kafka 消费特性
	openKafkaConsumerRequest := &tls.OpenKafkaConsumerRequest{
		TopicID: testTopicID,
	}
	openKafkaConsumerResponse, err := client.OpenKafkaConsumer(openKafkaConsumerRequest)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", openKafkaConsumerResponse)

	// 描述 Kafka 消费者
	describeKafkaConsumerRequest := &tls.DescribeKafkaConsumerRequest{
		TopicID: testTopicID,
	}
	describeKafkaConsumerResponse, err := client.DescribeKafkaConsumer(describeKafkaConsumerRequest)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", describeKafkaConsumerResponse)

	// 关闭 Kafka 消费特性
	closeKafkaConsumerRequest := &tls.CloseKafkaConsumerRequest{
		TopicID: testTopicID,
	}
	closeKafkaConsumerResponse, err := client.CloseKafkaConsumer(closeKafkaConsumerRequest)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", closeKafkaConsumerResponse)
}
