package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoUtils 提供MongoDB常用操作的工具类
type MongoUtils struct {
	Collection *mongo.Collection
}

// InsertOne 插入单个文档
func (m *MongoUtils) InsertOne(ctx context.Context, document interface{}) (primitive.ObjectID, error) {
	result, err := m.Collection.InsertOne(ctx, document)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("无法获取插入的ID")
	}
	return id, nil
}

// FindOne 查找单个文档
func (m *MongoUtils) FindOne(ctx context.Context, filter interface{}, result interface{}) error {
	return m.Collection.FindOne(ctx, filter).Decode(result)
}

// FindByID 通过ID查找文档
func (m *MongoUtils) FindByID(ctx context.Context, id string, result interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return m.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(result)
}

// Find 查找多个文档
func (m *MongoUtils) Find(ctx context.Context, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	cursor, err := m.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, results)
}

// UpdateOne 更新单个文档
func (m *MongoUtils) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return m.Collection.UpdateOne(ctx, filter, update)
}

// UpdateByID 通过ID更新文档
func (m *MongoUtils) UpdateByID(ctx context.Context, id string, update interface{}) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.Collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
}

// DeleteOne 删除单个文档
func (m *MongoUtils) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return m.Collection.DeleteOne(ctx, filter)
}

// DeleteByID 通过ID删除文档
func (m *MongoUtils) DeleteByID(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.Collection.DeleteOne(ctx, bson.M{"_id": objectID})
}

// Count 计数
func (m *MongoUtils) Count(ctx context.Context, filter interface{}) (int64, error) {
	return m.Collection.CountDocuments(ctx, filter)
}

// NewMongoUtils 创建MongoDB工具类实例
func NewMongoUtils(collectionName string) *MongoUtils {
	return &MongoUtils{
		Collection: MongoDB.Collection(collectionName),
	}
}
