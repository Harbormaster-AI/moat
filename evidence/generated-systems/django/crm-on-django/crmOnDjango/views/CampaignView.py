import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

 #======================================================================
# 
# Encapsulates data for View Campaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Campaign index.")

def get(request, campaignId ):
	delegate = CampaignDelegate()
	responseData = delegate.get( campaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	campaign = json.loads(request.body)
	delegate = CampaignDelegate()
	responseData = delegate.createFromJson( campaign )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	campaign = json.loads(request.body)
	delegate = CampaignDelegate()
	responseData = delegate.save( campaign )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, campaignId ):
	delegate = CampaignDelegate()
	responseData = delegate.delete( campaignId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CampaignDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, campaignId, OrganizationId ):
	delegate = CampaignDelegate()
	responseData = delegate.saveOrganization( campaignId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, campaignId ):
	delegate = CampaignDelegate()
	responseData = delegate.deleteOrganization( campaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignParentCampaign( request, campaignId, ParentCampaignId ):
	delegate = CampaignDelegate()
	responseData = delegate.saveParentCampaign( campaignId, ParentCampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignParentCampaign( request, campaignId ):
	delegate = CampaignDelegate()
	responseData = delegate.deleteParentCampaign( campaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChildCampaigns( request, campaignId, ChildCampaignsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addChildCampaigns( campaignId, ChildCampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChildCampaigns( request, campaignId, ChildCampaignsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeChildCampaigns( campaignId, ChildCampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMembers( request, campaignId, MembersIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addMembers( campaignId, MembersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMembers( request, campaignId, MembersIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeMembers( campaignId, MembersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOpportunities( request, campaignId, OpportunitiesIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addOpportunities( campaignId, OpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOpportunities( request, campaignId, OpportunitiesIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeOpportunities( campaignId, OpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAccounts( request, campaignId, AccountsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addAccounts( campaignId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAccounts( request, campaignId, AccountsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeAccounts( campaignId, AccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLeads( request, campaignId, LeadsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addLeads( campaignId, LeadsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLeads( request, campaignId, LeadsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeLeads( campaignId, LeadsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContacts( request, campaignId, ContactsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addContacts( campaignId, ContactsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContacts( request, campaignId, ContactsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeContacts( campaignId, ContactsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTeams( request, campaignId, TeamsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addTeams( campaignId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTeams( request, campaignId, TeamsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeTeams( campaignId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addActivities( request, campaignId, ActivitiesIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addActivities( campaignId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeActivities( request, campaignId, ActivitiesIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeActivities( campaignId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

