import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.TeamDelegate import TeamDelegate

 #======================================================================
# 
# Encapsulates data for View Team
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TeamView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Team index.")

def get(request, teamId ):
	delegate = TeamDelegate()
	responseData = delegate.get( teamId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	team = json.loads(request.body)
	delegate = TeamDelegate()
	responseData = delegate.createFromJson( team )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	team = json.loads(request.body)
	delegate = TeamDelegate()
	responseData = delegate.save( team )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, teamId ):
	delegate = TeamDelegate()
	responseData = delegate.delete( teamId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TeamDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAgency( request, teamId, AgencyId ):
	delegate = TeamDelegate()
	responseData = delegate.saveAgency( teamId, AgencyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAgency( request, teamId ):
	delegate = TeamDelegate()
	responseData = delegate.deleteAgency( teamId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addUsers( request, teamId, UsersIds ):
	delegate = TeamDelegate()
	responseData = delegate.addUsers( teamId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeUsers( request, teamId, UsersIds ):
	delegate = TeamDelegate()
	responseData = delegate.removeUsers( teamId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAdAccounts( request, teamId, AdAccountsIds ):
	delegate = TeamDelegate()
	responseData = delegate.addAdAccounts( teamId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAdAccounts( request, teamId, AdAccountsIds ):
	delegate = TeamDelegate()
	responseData = delegate.removeAdAccounts( teamId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

