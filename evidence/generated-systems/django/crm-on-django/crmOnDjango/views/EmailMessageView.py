import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

 #======================================================================
# 
# Encapsulates data for View EmailMessage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmailMessageView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the EmailMessage index.")

def get(request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.get( emailMessageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	emailMessage = json.loads(request.body)
	delegate = EmailMessageDelegate()
	responseData = delegate.createFromJson( emailMessage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	emailMessage = json.loads(request.body)
	delegate = EmailMessageDelegate()
	responseData = delegate.save( emailMessage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.delete( emailMessageId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EmailMessageDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, emailMessageId, OrganizationId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.saveOrganization( emailMessageId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.deleteOrganization( emailMessageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, emailMessageId, OwnerId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.saveOwner( emailMessageId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.deleteOwner( emailMessageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, emailMessageId, AccountId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.saveAccount( emailMessageId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.deleteAccount( emailMessageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignContact( request, emailMessageId, ContactId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.saveContact( emailMessageId, ContactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignContact( request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.deleteContact( emailMessageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLead( request, emailMessageId, LeadId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.saveLead( emailMessageId, LeadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLead( request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.deleteLead( emailMessageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCase( request, emailMessageId, CaseId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.saveCase( emailMessageId, CaseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCase( request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.deleteCase( emailMessageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOpportunity( request, emailMessageId, OpportunityId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.saveOpportunity( emailMessageId, OpportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOpportunity( request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.deleteOpportunity( emailMessageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, emailMessageId, CampaignId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.saveCampaign( emailMessageId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, emailMessageId ):
	delegate = EmailMessageDelegate()
	responseData = delegate.deleteCampaign( emailMessageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

