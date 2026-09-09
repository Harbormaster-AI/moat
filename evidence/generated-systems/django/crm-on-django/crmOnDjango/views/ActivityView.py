import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

 #======================================================================
# 
# Encapsulates data for View Activity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ActivityView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Activity index.")

def get(request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.get( activityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	activity = json.loads(request.body)
	delegate = ActivityDelegate()
	responseData = delegate.createFromJson( activity )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	activity = json.loads(request.body)
	delegate = ActivityDelegate()
	responseData = delegate.save( activity )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.delete( activityId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ActivityDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, activityId, OrganizationId ):
	delegate = ActivityDelegate()
	responseData = delegate.saveOrganization( activityId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.deleteOrganization( activityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, activityId, OwnerId ):
	delegate = ActivityDelegate()
	responseData = delegate.saveOwner( activityId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.deleteOwner( activityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, activityId, AccountId ):
	delegate = ActivityDelegate()
	responseData = delegate.saveAccount( activityId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.deleteAccount( activityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignContact( request, activityId, ContactId ):
	delegate = ActivityDelegate()
	responseData = delegate.saveContact( activityId, ContactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignContact( request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.deleteContact( activityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLead( request, activityId, LeadId ):
	delegate = ActivityDelegate()
	responseData = delegate.saveLead( activityId, LeadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLead( request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.deleteLead( activityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOpportunity( request, activityId, OpportunityId ):
	delegate = ActivityDelegate()
	responseData = delegate.saveOpportunity( activityId, OpportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOpportunity( request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.deleteOpportunity( activityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCase( request, activityId, CaseId ):
	delegate = ActivityDelegate()
	responseData = delegate.saveCase( activityId, CaseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCase( request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.deleteCase( activityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, activityId, CampaignId ):
	delegate = ActivityDelegate()
	responseData = delegate.saveCampaign( activityId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, activityId ):
	delegate = ActivityDelegate()
	responseData = delegate.deleteCampaign( activityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

