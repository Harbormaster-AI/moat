import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.ContactDelegate import ContactDelegate

 #======================================================================
# 
# Encapsulates data for View Contact
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContactView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Contact index.")

def get(request, contactId ):
	delegate = ContactDelegate()
	responseData = delegate.get( contactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	contact = json.loads(request.body)
	delegate = ContactDelegate()
	responseData = delegate.createFromJson( contact )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	contact = json.loads(request.body)
	delegate = ContactDelegate()
	responseData = delegate.save( contact )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, contactId ):
	delegate = ContactDelegate()
	responseData = delegate.delete( contactId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ContactDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, contactId, OrganizationId ):
	delegate = ContactDelegate()
	responseData = delegate.saveOrganization( contactId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, contactId ):
	delegate = ContactDelegate()
	responseData = delegate.deleteOrganization( contactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, contactId, AccountId ):
	delegate = ContactDelegate()
	responseData = delegate.saveAccount( contactId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, contactId ):
	delegate = ContactDelegate()
	responseData = delegate.deleteAccount( contactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, contactId, OwnerId ):
	delegate = ContactDelegate()
	responseData = delegate.saveOwner( contactId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, contactId ):
	delegate = ContactDelegate()
	responseData = delegate.deleteOwner( contactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addActivities( request, contactId, ActivitiesIds ):
	delegate = ContactDelegate()
	responseData = delegate.addActivities( contactId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeActivities( request, contactId, ActivitiesIds ):
	delegate = ContactDelegate()
	responseData = delegate.removeActivities( contactId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOpportunities( request, contactId, OpportunitiesIds ):
	delegate = ContactDelegate()
	responseData = delegate.addOpportunities( contactId, OpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOpportunities( request, contactId, OpportunitiesIds ):
	delegate = ContactDelegate()
	responseData = delegate.removeOpportunities( contactId, OpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCases( request, contactId, CasesIds ):
	delegate = ContactDelegate()
	responseData = delegate.addCases( contactId, CasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCases( request, contactId, CasesIds ):
	delegate = ContactDelegate()
	responseData = delegate.removeCases( contactId, CasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, contactId, CampaignsIds ):
	delegate = ContactDelegate()
	responseData = delegate.addCampaigns( contactId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, contactId, CampaignsIds ):
	delegate = ContactDelegate()
	responseData = delegate.removeCampaigns( contactId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addNotes( request, contactId, NotesIds ):
	delegate = ContactDelegate()
	responseData = delegate.addNotes( contactId, NotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeNotes( request, contactId, NotesIds ):
	delegate = ContactDelegate()
	responseData = delegate.removeNotes( contactId, NotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmailMessages( request, contactId, EmailMessagesIds ):
	delegate = ContactDelegate()
	responseData = delegate.addEmailMessages( contactId, EmailMessagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmailMessages( request, contactId, EmailMessagesIds ):
	delegate = ContactDelegate()
	responseData = delegate.removeEmailMessages( contactId, EmailMessagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

