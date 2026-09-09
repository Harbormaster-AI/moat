import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.UserDelegate import UserDelegate

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

def assignOrganization( request, userId, OrganizationId ):
	delegate = UserDelegate()
	responseData = delegate.saveOrganization( userId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, userId ):
	delegate = UserDelegate()
	responseData = delegate.deleteOrganization( userId )
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

def addActivities( request, userId, ActivitiesIds ):
	delegate = UserDelegate()
	responseData = delegate.addActivities( userId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeActivities( request, userId, ActivitiesIds ):
	delegate = UserDelegate()
	responseData = delegate.removeActivities( userId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOwnedAccounts( request, userId, OwnedAccountsIds ):
	delegate = UserDelegate()
	responseData = delegate.addOwnedAccounts( userId, OwnedAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOwnedAccounts( request, userId, OwnedAccountsIds ):
	delegate = UserDelegate()
	responseData = delegate.removeOwnedAccounts( userId, OwnedAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOwnedLeads( request, userId, OwnedLeadsIds ):
	delegate = UserDelegate()
	responseData = delegate.addOwnedLeads( userId, OwnedLeadsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOwnedLeads( request, userId, OwnedLeadsIds ):
	delegate = UserDelegate()
	responseData = delegate.removeOwnedLeads( userId, OwnedLeadsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOwnedOpportunities( request, userId, OwnedOpportunitiesIds ):
	delegate = UserDelegate()
	responseData = delegate.addOwnedOpportunities( userId, OwnedOpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOwnedOpportunities( request, userId, OwnedOpportunitiesIds ):
	delegate = UserDelegate()
	responseData = delegate.removeOwnedOpportunities( userId, OwnedOpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOwnedCases( request, userId, OwnedCasesIds ):
	delegate = UserDelegate()
	responseData = delegate.addOwnedCases( userId, OwnedCasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOwnedCases( request, userId, OwnedCasesIds ):
	delegate = UserDelegate()
	responseData = delegate.removeOwnedCases( userId, OwnedCasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQuotes( request, userId, QuotesIds ):
	delegate = UserDelegate()
	responseData = delegate.addQuotes( userId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQuotes( request, userId, QuotesIds ):
	delegate = UserDelegate()
	responseData = delegate.removeQuotes( userId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, userId, OrdersIds ):
	delegate = UserDelegate()
	responseData = delegate.addOrders( userId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, userId, OrdersIds ):
	delegate = UserDelegate()
	responseData = delegate.removeOrders( userId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContracts( request, userId, ContractsIds ):
	delegate = UserDelegate()
	responseData = delegate.addContracts( userId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContracts( request, userId, ContractsIds ):
	delegate = UserDelegate()
	responseData = delegate.removeContracts( userId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmailMessages( request, userId, EmailMessagesIds ):
	delegate = UserDelegate()
	responseData = delegate.addEmailMessages( userId, EmailMessagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmailMessages( request, userId, EmailMessagesIds ):
	delegate = UserDelegate()
	responseData = delegate.removeEmailMessages( userId, EmailMessagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

