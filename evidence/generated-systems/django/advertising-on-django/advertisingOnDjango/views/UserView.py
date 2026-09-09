import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.UserDelegate import UserDelegate

 #======================================================================
# 
# Encapsulates data for View User
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UserView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the User index.")

def get(request, userId ):
	delegate = UserDelegate()
	responseData = delegate.get( userId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	user = json.loads(request.body)
	delegate = UserDelegate()
	responseData = delegate.createFromJson( user )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	user = json.loads(request.body)
	delegate = UserDelegate()
	responseData = delegate.save( user )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, userId ):
	delegate = UserDelegate()
	responseData = delegate.delete( userId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = UserDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAgency( request, userId, AgencyId ):
	delegate = UserDelegate()
	responseData = delegate.saveAgency( userId, AgencyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAgency( request, userId ):
	delegate = UserDelegate()
	responseData = delegate.deleteAgency( userId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTeams( request, userId, TeamsIds ):
	delegate = UserDelegate()
	responseData = delegate.addTeams( userId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTeams( request, userId, TeamsIds ):
	delegate = UserDelegate()
	responseData = delegate.removeTeams( userId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAdAccounts( request, userId, AdAccountsIds ):
	delegate = UserDelegate()
	responseData = delegate.addAdAccounts( userId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAdAccounts( request, userId, AdAccountsIds ):
	delegate = UserDelegate()
	responseData = delegate.removeAdAccounts( userId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

