import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.TeamDelegate import TeamDelegate

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

def assignOrganization( request, teamId, OrganizationId ):
	delegate = TeamDelegate()
	responseData = delegate.saveOrganization( teamId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, teamId ):
	delegate = TeamDelegate()
	responseData = delegate.deleteOrganization( teamId )
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

def addAccounts( request, teamId, AccountsIds ):
	delegate = TeamDelegate()
	responseData = delegate.addAccounts( teamId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAccounts( request, teamId, AccountsIds ):
	delegate = TeamDelegate()
	responseData = delegate.removeAccounts( teamId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOpportunities( request, teamId, OpportunitiesIds ):
	delegate = TeamDelegate()
	responseData = delegate.addOpportunities( teamId, OpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOpportunities( request, teamId, OpportunitiesIds ):
	delegate = TeamDelegate()
	responseData = delegate.removeOpportunities( teamId, OpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCases( request, teamId, CasesIds ):
	delegate = TeamDelegate()
	responseData = delegate.addCases( teamId, CasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCases( request, teamId, CasesIds ):
	delegate = TeamDelegate()
	responseData = delegate.removeCases( teamId, CasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, teamId, CampaignsIds ):
	delegate = TeamDelegate()
	responseData = delegate.addCampaigns( teamId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, teamId, CampaignsIds ):
	delegate = TeamDelegate()
	responseData = delegate.removeCampaigns( teamId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

