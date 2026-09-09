import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

 #======================================================================
# 
# Encapsulates data for View Opportunity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Opportunity index.")

def get(request, opportunityId ):
	delegate = OpportunityDelegate()
	responseData = delegate.get( opportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	opportunity = json.loads(request.body)
	delegate = OpportunityDelegate()
	responseData = delegate.createFromJson( opportunity )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	opportunity = json.loads(request.body)
	delegate = OpportunityDelegate()
	responseData = delegate.save( opportunity )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, opportunityId ):
	delegate = OpportunityDelegate()
	responseData = delegate.delete( opportunityId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OpportunityDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, opportunityId, OrganizationId ):
	delegate = OpportunityDelegate()
	responseData = delegate.saveOrganization( opportunityId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, opportunityId ):
	delegate = OpportunityDelegate()
	responseData = delegate.deleteOrganization( opportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, opportunityId, AccountId ):
	delegate = OpportunityDelegate()
	responseData = delegate.saveAccount( opportunityId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, opportunityId ):
	delegate = OpportunityDelegate()
	responseData = delegate.deleteAccount( opportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, opportunityId, OwnerId ):
	delegate = OpportunityDelegate()
	responseData = delegate.saveOwner( opportunityId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, opportunityId ):
	delegate = OpportunityDelegate()
	responseData = delegate.deleteOwner( opportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContacts( request, opportunityId, ContactsIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.addContacts( opportunityId, ContactsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContacts( request, opportunityId, ContactsIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.removeContacts( opportunityId, ContactsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLineItems( request, opportunityId, LineItemsIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.addLineItems( opportunityId, LineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLineItems( request, opportunityId, LineItemsIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.removeLineItems( opportunityId, LineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addStageHistory( request, opportunityId, StageHistoryIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.addStageHistory( opportunityId, StageHistoryIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeStageHistory( request, opportunityId, StageHistoryIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.removeStageHistory( opportunityId, StageHistoryIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQuotes( request, opportunityId, QuotesIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.addQuotes( opportunityId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQuotes( request, opportunityId, QuotesIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.removeQuotes( opportunityId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, opportunityId, OrdersIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.addOrders( opportunityId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, opportunityId, OrdersIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.removeOrders( opportunityId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, opportunityId, CampaignsIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.addCampaigns( opportunityId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, opportunityId, CampaignsIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.removeCampaigns( opportunityId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addActivities( request, opportunityId, ActivitiesIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.addActivities( opportunityId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeActivities( request, opportunityId, ActivitiesIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.removeActivities( opportunityId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTeams( request, opportunityId, TeamsIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.addTeams( opportunityId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTeams( request, opportunityId, TeamsIds ):
	delegate = OpportunityDelegate()
	responseData = delegate.removeTeams( opportunityId, TeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

