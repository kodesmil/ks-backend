package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/go-patient-registry/pkg/pb"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~ A BRIEF DEVELOPMENT GUIDE ~~~~~~~~~~~~~~~~~~~~~~~~~
//
// TODO: Extend the GoPatientRegistry service by defining new RPCs and
// and message types in the pb/service.proto file. These RPCs and messages
// compose the API for your service. After modifying the proto schema in
// pb/service.proto, call "make protobuf" to regenerate the protobuf files.
//
// TODO: Create an implementation of the GoPatientRegistry server
// interface. This interface is generated by the protobuf compiler and exists
// inside the pb/service.pb.go file. The "server" struct already provides an
// implementation of GoPatientRegistry server interface service, but only
// for the GetVersion function. You will need to implement any new RPCs you
// add to your protobuf schema.
//
// TODO: Update the GetVersion function when newer versions of your service
// become available. Feel free to change GetVersion to better-suit how your
// versioning system, or get rid of it entirely. GetVersion helps make up a
// simple "starter" example that allows an end-to-end example. It is not
// required.
//
// TODO: Update the Publish function to better-suit your application, or get
// rid of it if your application is not using atlas pubsub.
//
// TODO: Oh yeah, delete this guide when you no longer need it.
//
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~ FAREWELL AND GOOD LUCK ~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

const (
	// version is the current version of the service
	version = "0.0.1"
)

// NewProfilesServer returns an instance of the default profiles server interface
func NewProfilesServer(database *gorm.DB) (pb.ProfilesServer, error) {
	return &profilesServer{&pb.ProfilesDefaultServer{DB: database}}, nil
}

type profilesServer struct {
	*pb.ProfilesDefaultServer
}

// NewGroupsServer returns an instance of the default groups server interface
func NewGroupsServer(database *gorm.DB) (pb.GroupsServer, error) {
	return &groupsServer{&pb.GroupsDefaultServer{DB: database}}, nil
}

type groupsServer struct {
	*pb.GroupsDefaultServer
}

// NewContactsServer returns an instance of the default contacts server interface
func NewContactsServer(database *gorm.DB) (pb.ContactsServer, error) {
	return &contactsServer{&pb.ContactsDefaultServer{DB: database}}, nil
}

type contactsServer struct {
	*pb.ContactsDefaultServer
}