import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.AgencyDelegate import AgencyDelegate

 #======================================================================
# 
# Encapsulates data for View Agency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgencyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Agency index.")

def get(request, agencyId ):
	delegate = AgencyDelegate()
	responseData = delegate.get( agencyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	agency = json.loads(request.body)
	delegate = AgencyDelegate()
	responseData = delegate.createFromJson( agency )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	agency = json.loads(request.body)
	delegate = AgencyDelegate()
	responseData = delegate.save( agency )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, agencyId ):
	delegate = AgencyDelegate()
	responseData = delegate.delete( agencyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AgencyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAdvertisers( request, agencyId, AdvertisersIds ):
	delegate = AgencyDelegate()
	responseData = delegate.addAdvertisers( agencyId, AdvertisersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAdvertisers( request, agencyId, AdvertisersIds ):
	delegate = AgencyDelegate()
	responseData = delegate.removeAdvertisers( agencyId, AdvertisersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTeams( request, agencyId, TeamsIds ):
	delegate = AgencyDelegate()
	responseData = delegate.addTeams( agencyId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTeams( request, agencyId, TeamsIds ):
	delegate = AgencyDelegate()
	responseData = delegate.removeTeams( agencyId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addUsers( request, agencyId, UsersIds ):
	delegate = AgencyDelegate()
	responseData = delegate.addUsers( agencyId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeUsers( request, agencyId, UsersIds ):
	delegate = AgencyDelegate()
	responseData = delegate.removeUsers( agencyId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInsertionOrders( request, agencyId, InsertionOrdersIds ):
	delegate = AgencyDelegate()
	responseData = delegate.addInsertionOrders( agencyId, InsertionOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInsertionOrders( request, agencyId, InsertionOrdersIds ):
	delegate = AgencyDelegate()
	responseData = delegate.removeInsertionOrders( agencyId, InsertionOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

