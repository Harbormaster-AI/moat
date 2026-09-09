import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.LeadDelegate import LeadDelegate

 #======================================================================
# 
# Encapsulates data for View Lead
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeadView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Lead index.")

def get(request, leadId ):
	delegate = LeadDelegate()
	responseData = delegate.get( leadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	lead = json.loads(request.body)
	delegate = LeadDelegate()
	responseData = delegate.createFromJson( lead )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	lead = json.loads(request.body)
	delegate = LeadDelegate()
	responseData = delegate.save( lead )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, leadId ):
	delegate = LeadDelegate()
	responseData = delegate.delete( leadId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LeadDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, leadId, OrganizationId ):
	delegate = LeadDelegate()
	responseData = delegate.saveOrganization( leadId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, leadId ):
	delegate = LeadDelegate()
	responseData = delegate.deleteOrganization( leadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, leadId, OwnerId ):
	delegate = LeadDelegate()
	responseData = delegate.saveOwner( leadId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, leadId ):
	delegate = LeadDelegate()
	responseData = delegate.deleteOwner( leadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignConvertedAccount( request, leadId, ConvertedAccountId ):
	delegate = LeadDelegate()
	responseData = delegate.saveConvertedAccount( leadId, ConvertedAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignConvertedAccount( request, leadId ):
	delegate = LeadDelegate()
	responseData = delegate.deleteConvertedAccount( leadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignConvertedContact( request, leadId, ConvertedContactId ):
	delegate = LeadDelegate()
	responseData = delegate.saveConvertedContact( leadId, ConvertedContactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignConvertedContact( request, leadId ):
	delegate = LeadDelegate()
	responseData = delegate.deleteConvertedContact( leadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignConvertedOpportunity( request, leadId, ConvertedOpportunityId ):
	delegate = LeadDelegate()
	responseData = delegate.saveConvertedOpportunity( leadId, ConvertedOpportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignConvertedOpportunity( request, leadId ):
	delegate = LeadDelegate()
	responseData = delegate.deleteConvertedOpportunity( leadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addActivities( request, leadId, ActivitiesIds ):
	delegate = LeadDelegate()
	responseData = delegate.addActivities( leadId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeActivities( request, leadId, ActivitiesIds ):
	delegate = LeadDelegate()
	responseData = delegate.removeActivities( leadId, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, leadId, CampaignsIds ):
	delegate = LeadDelegate()
	responseData = delegate.addCampaigns( leadId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, leadId, CampaignsIds ):
	delegate = LeadDelegate()
	responseData = delegate.removeCampaigns( leadId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addNotes( request, leadId, NotesIds ):
	delegate = LeadDelegate()
	responseData = delegate.addNotes( leadId, NotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeNotes( request, leadId, NotesIds ):
	delegate = LeadDelegate()
	responseData = delegate.removeNotes( leadId, NotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmailMessages( request, leadId, EmailMessagesIds ):
	delegate = LeadDelegate()
	responseData = delegate.addEmailMessages( leadId, EmailMessagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmailMessages( request, leadId, EmailMessagesIds ):
	delegate = LeadDelegate()
	responseData = delegate.removeEmailMessages( leadId, EmailMessagesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

