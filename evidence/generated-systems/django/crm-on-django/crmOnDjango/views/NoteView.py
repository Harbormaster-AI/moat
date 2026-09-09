import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.NoteDelegate import NoteDelegate

 #======================================================================
# 
# Encapsulates data for View Note
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NoteView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Note index.")

def get(request, noteId ):
	delegate = NoteDelegate()
	responseData = delegate.get( noteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	note = json.loads(request.body)
	delegate = NoteDelegate()
	responseData = delegate.createFromJson( note )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	note = json.loads(request.body)
	delegate = NoteDelegate()
	responseData = delegate.save( note )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, noteId ):
	delegate = NoteDelegate()
	responseData = delegate.delete( noteId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = NoteDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, noteId, OrganizationId ):
	delegate = NoteDelegate()
	responseData = delegate.saveOrganization( noteId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, noteId ):
	delegate = NoteDelegate()
	responseData = delegate.deleteOrganization( noteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, noteId, OwnerId ):
	delegate = NoteDelegate()
	responseData = delegate.saveOwner( noteId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, noteId ):
	delegate = NoteDelegate()
	responseData = delegate.deleteOwner( noteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, noteId, AccountId ):
	delegate = NoteDelegate()
	responseData = delegate.saveAccount( noteId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, noteId ):
	delegate = NoteDelegate()
	responseData = delegate.deleteAccount( noteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignContact( request, noteId, ContactId ):
	delegate = NoteDelegate()
	responseData = delegate.saveContact( noteId, ContactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignContact( request, noteId ):
	delegate = NoteDelegate()
	responseData = delegate.deleteContact( noteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOpportunity( request, noteId, OpportunityId ):
	delegate = NoteDelegate()
	responseData = delegate.saveOpportunity( noteId, OpportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOpportunity( request, noteId ):
	delegate = NoteDelegate()
	responseData = delegate.deleteOpportunity( noteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCase( request, noteId, CaseId ):
	delegate = NoteDelegate()
	responseData = delegate.saveCase( noteId, CaseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCase( request, noteId ):
	delegate = NoteDelegate()
	responseData = delegate.deleteCase( noteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLead( request, noteId, LeadId ):
	delegate = NoteDelegate()
	responseData = delegate.saveLead( noteId, LeadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLead( request, noteId ):
	delegate = NoteDelegate()
	responseData = delegate.deleteLead( noteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

