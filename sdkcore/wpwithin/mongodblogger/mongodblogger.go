package mongodblogger

import (
	"os"
	"time"
	"gopkg.in/mgo.v2"
)

//
// Joe.Drumgoole@mongodb.com
//
// A Simple logging package for WPWWithin Framework
//

// must point at a running writeable MongoDB Server
const WPW_MONGODB_SERVER   string = "WPW_MONGODB_SERVER"

// Set to anything to turn logging on
const WPW_MONGODB_LOGGING  string = "WPW_MONGODB_LOGGING" 

const WPW_EVENT_DATABASE   string = "WPWEvents" 

const WPW_EVENT_COLLECTION string = "WPWLog" 


type MongoDBLogger struct {
	
	MgoSession  *mgo.Session
	DB          *mgo.Database
	Collection  *mgo.Collection
	loggingOn   bool
	
}

type StringEvent struct {

	Name             string
	Msg              string
	Timestamp        time.Time
}

type IntEvent struct {

	Name             string
	Msg              string
	Amount			 int64
	Timestamp        time.Time
}

type FloatEvent struct {

	Name             string
	Msg              string
	Amount			 float64
	Timestamp        time.Time
}

type PaymentEvent struct {

	Name             string
	Msg              string
	Payment		     interface{}
	Timestamp        time.Time
}

type DocEvent struct {

	Name             string
	Msg              string
	Doc 		     interface{}
	Timestamp        time.Time
}
func ( m *MongoDBLogger ) Initialise() error {
	
	m.loggingOn = true
	
	serverUrl := os.Getenv( WPW_MONGODB_SERVER ) 
	if serverUrl == "" {
		m.TurnLoggingOff()
	} else {

		session, err := mgo.Dial( serverUrl )
		
		if err != nil {
			return err
		} else {
			m.MgoSession = session
		}
		
		m.DB         = m.MgoSession.DB( WPW_EVENT_DATABASE )
		m.Collection = m.DB.C( WPW_EVENT_COLLECTION )
	
	}	
	
	loggingOn := os.Getenv( WPW_MONGODB_LOGGING ) 
	
	if ( loggingOn == "ON" ) && ( serverUrl != "" ){
		m.TurnLoggingOn()
	} else {
		m.TurnLoggingOff()
	}
	
	return nil
}

func ( m *MongoDBLogger ) TurnLoggingOn() bool {
	m.loggingOn = true
	return m.loggingOn
	
}

func ( m *MongoDBLogger )  TurnLoggingOff() bool {
	m.loggingOn = false
	return m.loggingOn
	
}

func ( m *MongoDBLogger ) LogEventStr( name string, message string ) error {
	
	if m.loggingOn {
		
		err := m.Collection.Insert( &StringEvent{ Name      : name, 
										          Msg       : message, 
										          Timestamp : time.Now()})
		return err
	}
	return nil
}

func ( m *MongoDBLogger ) LogEventInt( name string, message string, amount int64 ) error {
	
	if m.loggingOn {
		
		err := m.Collection.Insert( &IntEvent{ Name      : name, 
										       Msg       : message,
										       Amount    : amount,
										       Timestamp : time.Now()})
		return err
	}
	return nil
}

func ( m *MongoDBLogger ) LogEventFloat( name string, message string, amount float64) error {
	
	if m.loggingOn {
		
		err := m.Collection.Insert( &FloatEvent{ Name      : name, 
										         Msg       : message, 
										         Amount    : amount,
										         Timestamp : time.Now()})

		return err
	}
	return nil
}

func ( m *MongoDBLogger ) LogEventDoc( name string, message string, doc interface{}) error {
	
	if m.loggingOn {
	
		err := m.Collection.Insert( &DocEvent{ Name      : name, 
										       Msg       : message, 
										       Doc       : doc,
										       Timestamp : time.Now()})

		return err
	}
	return nil
}

func ( m *MongoDBLogger ) LogEventPayment( name string, message string, payment interface{}) error {
	
	if m.loggingOn {
	
		err := m.Collection.Insert( &PaymentEvent{ Name      : name, 
										           Msg       : message, 
										           Payment   : payment,
										           Timestamp : time.Now()})

		return err
	}
	return nil
}