package mongodb

import (
	"context"
	"fmt"
	"manindexer/pin"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mg *Mongodb) GetMaxMetaIdNumber() (number int64) {
	findOp := options.FindOne()
	findOp.SetSort(bson.D{{Key: "number", Value: -1}})
	var info pin.MetaIdInfo
	err := mongoClient.Collection(MetaIdInfoCollection).FindOne(context.TODO(), bson.D{}, findOp).Decode(&info)
	if err != nil && err == mongo.ErrNoDocuments {
		err = nil
		number = 1
		return
	}
	number = info.Number + 1
	return
}

func (mg *Mongodb) GetMetaIdInfo(address string, mempool bool, metaid string) (info *pin.MetaIdInfo, unconfirmed string, err error) {
	filter := bson.D{{Key: "address", Value: address}}
	if metaid != "" {
		filter = bson.D{{Key: "metaid", Value: metaid}}
	}
	var mempoolInfo pin.MetaIdInfo
	if mempool {
		if metaid != "" {
			mempoolInfo, _ = findMetaIdInfoInMempool("metaid", metaid)
		}
		if address != "" {
			mempoolInfo, _ = findMetaIdInfoInMempool("address", address)
		}
	}
	var unconfirmedList []string
	err = mongoClient.Collection(MetaIdInfoCollection).FindOne(context.TODO(), filter).Decode(&info)
	if err == mongo.ErrNoDocuments {
		err = nil
		if mempoolInfo == (pin.MetaIdInfo{}) {
			return
		}
		info = &mempoolInfo
	}
	if mempool && mempoolInfo != (pin.MetaIdInfo{}) {
		if mempoolInfo.Number == -1 {
			unconfirmedList = append(unconfirmedList, "number")
		}
		if mempoolInfo.Avatar != "" {
			info.Avatar = mempoolInfo.Avatar
			unconfirmedList = append(unconfirmedList, "avatar")
		}
		if mempoolInfo.Name != "" {
			info.Name = mempoolInfo.Name
			unconfirmedList = append(unconfirmedList, "name")
		}
		if mempoolInfo.Bio != "" {
			info.Bio = mempoolInfo.Bio
			unconfirmedList = append(unconfirmedList, "bio")
		}
	}
	if len(unconfirmedList) > 0 {
		unconfirmed = strings.Join(unconfirmedList, ",")
	}
	return
}
func findMetaIdInfoInMempool(key string, value string) (info pin.MetaIdInfo, err error) {
	result, err := mongoClient.Collection(MempoolPinsCollection).Find(context.TODO(), bson.M{key: value})
	if err != nil {
		return
	}
	var pins []pin.PinInscription
	err = result.All(context.TODO(), &pins)
	if err != nil {
		return
	}
	for _, pin := range pins {
		if pin.OriginalPath == "/info/name" {
			info.Name = string(pin.ContentBody)
		} else if pin.OriginalPath == "/info/avatar" {
			info.Avatar = fmt.Sprintf("/content/%s", pin.Id)
		} else if pin.OriginalPath == "/info/bid" {
			info.Bio = string(pin.ContentBody)
		}
	}
	return
}
func (mg *Mongodb) BatchUpsertMetaIdInfo(infoList map[string]*pin.MetaIdInfo) (err error) {
	//bT := time.Now()
	var models []mongo.WriteModel
	for _, info := range infoList {
		filter := bson.D{{Key: "address", Value: info.Address}}
		var updateInfo bson.D
		/*
			update := bson.D{{Key: "$set", Value: bson.D{
				{Key: "mumber", Value: info.Number},
				{Key: "roottxid", Value: info.RootTxId},
				{Key: "name", Value: info.Name},
				{Key: "address", Value: info.Address},
				{Key: "avatar", Value: info.Avatar},
				{Key: "bio", Value: info.Bio},
				{Key: "soulbondtoken", Value: info.SoulbondToken},
			}},
			}
		*/
		if info.Number > 0 {
			updateInfo = append(updateInfo, bson.E{Key: "number", Value: info.Number})
		}
		if info.MetaId != "" {
			updateInfo = append(updateInfo, bson.E{Key: "metaid", Value: info.MetaId})
		}
		if info.Name != "" {
			updateInfo = append(updateInfo, bson.E{Key: "name", Value: info.Name})
		}
		if info.NameId != "" {
			updateInfo = append(updateInfo, bson.E{Key: "nameid", Value: info.NameId})
		}
		if info.Address != "" {
			updateInfo = append(updateInfo, bson.E{Key: "address", Value: info.Address})
		}
		if len(info.Avatar) > 0 {
			updateInfo = append(updateInfo, bson.E{Key: "avatar", Value: info.Avatar})
		}
		if len(info.AvatarId) > 0 {
			updateInfo = append(updateInfo, bson.E{Key: "avatarid", Value: info.AvatarId})
		}
		if len(info.Bio) > 0 {
			updateInfo = append(updateInfo, bson.E{Key: "bio", Value: info.Bio})
		}
		if len(info.BioId) > 0 {
			updateInfo = append(updateInfo, bson.E{Key: "bioid", Value: info.BioId})
		}
		if len(info.SoulbondToken) > 0 {
			updateInfo = append(updateInfo, bson.E{Key: "soulbondtoken", Value: info.SoulbondToken})
		}
		update := bson.D{{Key: "$set", Value: updateInfo}}
		m := mongo.NewUpdateOneModel()
		m.SetFilter(filter).SetUpdate(update).SetUpsert(true)
		models = append(models, m)
	}
	bulkWriteOptions := options.BulkWrite().SetOrdered(false)
	_, err = mongoClient.Collection(MetaIdInfoCollection).BulkWrite(context.Background(), models, bulkWriteOptions)
	//eT := time.Since(bT)
	//fmt.Println("BatchUpsertMetaIdInfo time: ", eT)
	return
}

func (mg *Mongodb) GetMetaIdPageList(page int64, size int64) (pins []*pin.MetaIdInfo, err error) {
	cursor := (page - 1) * size
	opts := options.Find().SetSort(bson.D{{Key: "number", Value: -1}}).SetSkip(cursor).SetLimit(size)
	result, err := mongoClient.Collection(MetaIdInfoCollection).Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return
	}
	err = result.All(context.TODO(), &pins)
	return
}
func (mg *Mongodb) BatchUpsertMetaIdInfoAddition(infoList []*pin.MetaIdInfoAdditional) (err error) {
	var models []mongo.WriteModel
	for _, info := range infoList {
		filter := bson.D{{Key: "metaid", Value: info.MetaId}, {Key: "infokey", Value: info.InfoKey}}
		var updateInfo bson.D
		updateInfo = append(updateInfo, bson.E{Key: "infoValue", Value: info.InfoValue})
		updateInfo = append(updateInfo, bson.E{Key: "pinid", Value: info.PinId})
		update := bson.D{{Key: "$set", Value: updateInfo}}
		m := mongo.NewUpdateOneModel()
		m.SetFilter(filter).SetUpdate(update).SetUpsert(true)
		models = append(models, m)
	}

	bulkWriteOptions := options.BulkWrite().SetOrdered(false)
	_, err = mongoClient.Collection(InfoCollection).BulkWrite(context.Background(), models, bulkWriteOptions)
	return
}
