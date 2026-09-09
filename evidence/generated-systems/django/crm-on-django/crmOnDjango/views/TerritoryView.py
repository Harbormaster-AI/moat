import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.TerritoryDelegate import TerritoryDelegate

 #======================================================================
# 
# Encapsulates data for View Territory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerritoryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Territory index.")

def get(request, territoryId ):
	delegate = TerritoryDelegate()
	responseData = delegate.get( territoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	territory = json.loads(request.body)
	delegate = TerritoryDelegate()
	responseData = delegate.createFromJson( territory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	territory = json.loads(request.body)
	delegate = TerritoryDelegate()
	responseData = delegate.save( territory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, territoryId ):
	delegate = TerritoryDelegate()
	responseData = delegate.delete( territoryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TerritoryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, territoryId, OrganizationId ):
	delegate = TerritoryDelegate()
	responseData = delegate.saveOrganization( territoryId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, territoryId ):
	delegate = TerritoryDelegate()
	responseData = delegate.deleteOrganization( territoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAccounts( request, territoryId, AccountsIds ):
	delegate = TerritoryDelegate()
	responseData = delegate.addAccounts( territoryId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAccounts( request, territoryId, AccountsIds ):
	delegate = TerritoryDelegate()
	responseData = delegate.removeAccounts( territoryId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addUsers( request, territoryId, UsersIds ):
	delegate = TerritoryDelegate()
	responseData = delegate.addUsers( territoryId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeUsers( request, territoryId, UsersIds ):
	delegate = TerritoryDelegate()
	responseData = delegate.removeUsers( territoryId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

