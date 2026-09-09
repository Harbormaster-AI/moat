import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.AccountDelegate import AccountDelegate

 #======================================================================
# 
# Encapsulates data for View Account
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccountView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Account index.")

def get(request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.get( accountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	account = json.loads(request.body)
	delegate = AccountDelegate()
	responseData = delegate.createFromJson( account )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	account = json.loads(request.body)
	delegate = AccountDelegate()
	responseData = delegate.save( account )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.delete( accountId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AccountDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, accountId, OrganizationId ):
	delegate = AccountDelegate()
	responseData = delegate.saveOrganization( accountId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.deleteOrganization( accountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignParentAccount( request, accountId, ParentAccountId ):
	delegate = AccountDelegate()
	responseData = delegate.saveParentAccount( accountId, ParentAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignParentAccount( request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.deleteParentAccount( accountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, accountId, OwnerId ):
	delegate = AccountDelegate()
	responseData = delegate.saveOwner( accountId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.deleteOwner( accountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTerritory( request, accountId, TerritoryId ):
	delegate = AccountDelegate()
	responseData = delegate.saveTerritory( accountId, TerritoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTerritory( request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.deleteTerritory( accountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChildAccounts( request, accountId, ChildAccountsIds ):
	delegate = AccountDelegate()
	responseData = delegate.addChildAccounts( accountId, ChildAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChildAccounts( request, accountId, ChildAccountsIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeChildAccounts( accountId, ChildAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContacts( request, accountId, ContactsIds ):
	delegate = AccountDelegate()
	responseData = delegate.addContacts( accountId, ContactsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContacts( request, accountId, ContactsIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeContacts( accountId, ContactsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOpportunities( request, accountId, OpportunitiesIds ):
	delegate = AccountDelegate()
	responseData = delegate.addOpportunities( accountId, OpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOpportunities( request, accountId, OpportunitiesIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeOpportunities( accountId, OpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCases( request, accountId, CasesIds ):
	delegate = AccountDelegate()
	responseData = delegate.addCases( accountId, CasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCases( request, accountId, CasesIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeCases( accountId, CasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addActivities( request, accountId, ActivitiesIds ):
	delegate = AccountDelegate()
	responseData = delegate.addActivities( accountId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeActivities( request, accountId, ActivitiesIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeActivities( accountId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, accountId, CampaignsIds ):
	delegate = AccountDelegate()
	responseData = delegate.addCampaigns( accountId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, accountId, CampaignsIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeCampaigns( accountId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQuotes( request, accountId, QuotesIds ):
	delegate = AccountDelegate()
	responseData = delegate.addQuotes( accountId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQuotes( request, accountId, QuotesIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeQuotes( accountId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, accountId, OrdersIds ):
	delegate = AccountDelegate()
	responseData = delegate.addOrders( accountId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, accountId, OrdersIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeOrders( accountId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContracts( request, accountId, ContractsIds ):
	delegate = AccountDelegate()
	responseData = delegate.addContracts( accountId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContracts( request, accountId, ContractsIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeContracts( accountId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addNotes( request, accountId, NotesIds ):
	delegate = AccountDelegate()
	responseData = delegate.addNotes( accountId, NotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeNotes( request, accountId, NotesIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeNotes( accountId, NotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmailMessages( request, accountId, EmailMessagesIds ):
	delegate = AccountDelegate()
	responseData = delegate.addEmailMessages( accountId, EmailMessagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmailMessages( request, accountId, EmailMessagesIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeEmailMessages( accountId, EmailMessagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

