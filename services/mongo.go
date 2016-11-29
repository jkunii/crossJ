package services

import (
	"fmt"
	"strings"

	"github.com/jkunii/crossJ/global"

	"gopkg.in/mgo.v2/bson"
)

func (ac AnalyticsController) PutAnalytics(analytic Analytic) {
	if global.Cfg.MongoAnalyticsActive {

		// Stub an Analytic to be populated from the body
		a := analytic

		// try get log in graylog
		info(fmt.Sprintf("{mongo-analytics: %v}", a))

		// Add an Id
		a.Id = bson.NewObjectId()
		// Write analytics to mongo
		ac.session.DB(global.Cfg.MongoDatabase).C(global.Cfg.MongoCertificateCollection).Insert(a)
		debug("Put Analytics to MongoDB")
	}

}

func (ac AnalyticsController) GetAnalytics(query Analytic) Analytic {
	// Grab id
	id := "someId"

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		errorr("Not Mongo Id")
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub analytics
	a := Analytic{}

	// Fetch analytics
	if err := ac.session.DB(global.Cfg.MongoDatabase).C(global.Cfg.MongoCertificateCollection).FindId(oid).One(&a); err != nil {
		errorr("NOT FOUND or Another error")
	}

	// Marshal provided interface into JSON structure
	// aj, _ := json.Marshal(a)
	return a
}

type (
	Analytic struct {
		Id              bson.ObjectId `json:"id" bson:"_id"`
		TransactionType string        `json:"transactionType"`
		Certificate     Certificate   `json:"certificate"`
		Step            string        `json:"step"`
	}

	// AnalyticsController represents the controller for operating on the Analytics resource
	AnalyticsController struct {
		session *mgo.Session
	}
)

// NewAnalyticsController provides a reference to a AnalyticsController with provided mongo session
func NewAnalyticsController(s *mgo.Session) *AnalyticsController {
	return &AnalyticsController{s}
}

func GetSession() *mgo.Session {
	Host := strings.Split(global.Cfg.MongoHosts, ",")

	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Username: global.Cfg.MongoUser,
		Password: global.Cfg.MongoUserSecret,
		Database: global.Cfg.MongoDatabase,
	})

	// s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	s.SetMode(mgo.Monotonic, true)

	fmt.Printf("Connected to replica set %v!\n", s.LiveServers())
	return s
}
