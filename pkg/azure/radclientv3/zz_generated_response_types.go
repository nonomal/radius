// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclientv3

import "net/http"

// ApplicationListResponse is the response envelope for operations that return a ApplicationList type.
type ApplicationListResponse struct {
	// List of Application resources.
	ApplicationList *ApplicationList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ApplicationResourceResponse is the response envelope for operations that return a ApplicationResource type.
type ApplicationResourceResponse struct {
	// Application resource.
	ApplicationResource *ApplicationResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AzureCosmosDBMongoComponentListResponse is the response envelope for operations that return a AzureCosmosDBMongoComponentList type.
type AzureCosmosDBMongoComponentListResponse struct {
	// List of azure.com.CosmosDBMongoComponent resources.
	AzureCosmosDBMongoComponentList *AzureCosmosDBMongoComponentList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AzureCosmosDBMongoComponentResourceResponse is the response envelope for operations that return a AzureCosmosDBMongoComponentResource type.
type AzureCosmosDBMongoComponentResourceResponse struct {
	// Component for Azure CosmosDB with Mongo
	AzureCosmosDBMongoComponentResource *AzureCosmosDBMongoComponentResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AzureCosmosDBSQLComponentListResponse is the response envelope for operations that return a AzureCosmosDBSQLComponentList type.
type AzureCosmosDBSQLComponentListResponse struct {
	// List of azure.com.CosmosDBSQLComponent resources.
	AzureCosmosDBSQLComponentList *AzureCosmosDBSQLComponentList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AzureCosmosDBSQLComponentResourceResponse is the response envelope for operations that return a AzureCosmosDBSQLComponentResource type.
type AzureCosmosDBSQLComponentResourceResponse struct {
	// Component for Azure CosmosDB with SQL
	AzureCosmosDBSQLComponentResource *AzureCosmosDBSQLComponentResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AzureKeyVaultComponentListResponse is the response envelope for operations that return a AzureKeyVaultComponentList type.
type AzureKeyVaultComponentListResponse struct {
	// List of azure.com.KeyVaultComponent resources.
	AzureKeyVaultComponentList *AzureKeyVaultComponentList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AzureKeyVaultComponentResourceResponse is the response envelope for operations that return a AzureKeyVaultComponentResource type.
type AzureKeyVaultComponentResourceResponse struct {
	// Component for Azure KeyVault
	AzureKeyVaultComponentResource *AzureKeyVaultComponentResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AzureServiceBusComponentListResponse is the response envelope for operations that return a AzureServiceBusComponentList type.
type AzureServiceBusComponentListResponse struct {
	// List of azure.com.ServiceBusQueueComponent resources.
	AzureServiceBusComponentList *AzureServiceBusComponentList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AzureServiceBusComponentResourceResponse is the response envelope for operations that return a AzureServiceBusComponentResource type.
type AzureServiceBusComponentResourceResponse struct {
	// Component for Azure ServiceBus
	AzureServiceBusComponentResource *AzureServiceBusComponentResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ContainerComponentListResponse is the response envelope for operations that return a ContainerComponentList type.
type ContainerComponentListResponse struct {
	// List of ContainerComponent resources.
	ContainerComponentList *ContainerComponentList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ContainerComponentResourceResponse is the response envelope for operations that return a ContainerComponentResource type.
type ContainerComponentResourceResponse struct {
	// The radius.dev/Container component provides an abstraction for a container workload that can be run on any Radius platform
	ContainerComponentResource *ContainerComponentResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprPubSubComponentListResponse is the response envelope for operations that return a DaprPubSubComponentList type.
type DaprPubSubComponentListResponse struct {
	// List of dapr.io.PubSubComponent resources.
	DaprPubSubComponentList *DaprPubSubComponentList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprPubSubComponentResourceResponse is the response envelope for operations that return a DaprPubSubComponentResource type.
type DaprPubSubComponentResourceResponse struct {
	// Component for Dapr Pub/Sub
	DaprPubSubComponentResource *DaprPubSubComponentResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoreComponentListResponse is the response envelope for operations that return a DaprStateStoreComponentList type.
type DaprStateStoreComponentListResponse struct {
	// List of dapr.io.StateStoreComponent resources.
	DaprStateStoreComponentList *DaprStateStoreComponentList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoreComponentResourceResponse is the response envelope for operations that return a DaprStateStoreComponentResource type.
type DaprStateStoreComponentResourceResponse struct {
	// Component for Dapr state store
	DaprStateStoreComponentResource *DaprStateStoreComponentResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDBComponentListResponse is the response envelope for operations that return a MongoDBComponentList type.
type MongoDBComponentListResponse struct {
	// List of mongodb.com.MongoDBComponent resources.
	MongoDBComponentList *MongoDBComponentList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDBComponentResourceResponse is the response envelope for operations that return a MongoDBComponentResource type.
type MongoDBComponentResourceResponse struct {
	// The mongodb.com/MongoDB component is a portable component which can be deployed to any Radius platform.
	MongoDBComponentResource *MongoDBComponentResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RabbitMQComponentListResponse is the response envelope for operations that return a RabbitMQComponentList type.
type RabbitMQComponentListResponse struct {
	// List of rabbitmq.com.MessageQueue resources.
	RabbitMQComponentList *RabbitMQComponentList

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RabbitMQComponentResourceResponse is the response envelope for operations that return a RabbitMQComponentResource type.
type RabbitMQComponentResourceResponse struct {
	// The rabbitmq.com/MessageQueue component is a Kubernetes specific component for message brokering.
	RabbitMQComponentResource *RabbitMQComponentResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisComponentListResponse is the response envelope for operations that return a RedisComponentList type.
type RedisComponentListResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// List of redislabs.com.Redis resources.
	RedisComponentList *RedisComponentList
}

// RedisComponentResourceResponse is the response envelope for operations that return a RedisComponentResource type.
type RedisComponentResourceResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The redislabs.com/Redis component is a portable component which can be deployed to any Radius platform.
	RedisComponentResource *RedisComponentResource
}

