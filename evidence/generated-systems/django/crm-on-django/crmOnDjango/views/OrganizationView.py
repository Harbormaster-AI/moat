import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

 #======================================================================
# 
# Encapsulates data for View Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrganizationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Organization index.")

def get(request, organizationId ):
	delegate = OrganizationDelegate()
	responseData = delegate.get( organizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	organization = json.loads(request.body)
	delegate = OrganizationDelegate()
	responseData = delegate.createFromJson( organization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	organization = json.loads(request.body)
	delegate = OrganizationDelegate()
	responseData = delegate.save( organization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, organizationId ):
	delegate = OrganizationDelegate()
	responseData = delegate.delete( organizationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OrganizationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addUsers( request, organizationId, UsersIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addUsers( organizationId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeUsers( request, organizationId, UsersIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeUsers( organizationId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAccounts( request, organizationId, AccountsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addAccounts( organizationId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAccounts( request, organizationId, AccountsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeAccounts( organizationId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTeams( request, organizationId, TeamsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addTeams( organizationId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTeams( request, organizationId, TeamsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeTeams( organizationId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTerritories( request, organizationId, TerritoriesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addTerritories( organizationId, TerritoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTerritories( request, organizationId, TerritoriesIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeTerritories( organizationId, TerritoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProducts( request, organizationId, ProductsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addProducts( organizationId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProducts( request, organizationId, ProductsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeProducts( organizationId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPriceBooks( request, organizationId, PriceBooksIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addPriceBooks( organizationId, PriceBooksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePriceBooks( request, organizationId, PriceBooksIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removePriceBooks( organizationId, PriceBooksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, organizationId, CampaignsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.addCampaigns( organizationId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, organizationId, CampaignsIds ):
	delegate = OrganizationDelegate()
	responseData = delegate.removeCampaigns( organizationId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

