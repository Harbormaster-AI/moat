import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.Case_Delegate import Case_Delegate

 #======================================================================
# 
# Encapsulates data for View Case_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Case_View function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Case_ index.")

def get(request, case_Id ):
	delegate = Case_Delegate()
	responseData = delegate.get( case_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	case_ = json.loads(request.body)
	delegate = Case_Delegate()
	responseData = delegate.createFromJson( case_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	case_ = json.loads(request.body)
	delegate = Case_Delegate()
	responseData = delegate.save( case_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, case_Id ):
	delegate = Case_Delegate()
	responseData = delegate.delete( case_Id )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = Case_Delegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, case_Id, OrganizationId ):
	delegate = Case_Delegate()
	responseData = delegate.saveOrganization( case_Id, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, case_Id ):
	delegate = Case_Delegate()
	responseData = delegate.deleteOrganization( case_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, case_Id, AccountId ):
	delegate = Case_Delegate()
	responseData = delegate.saveAccount( case_Id, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, case_Id ):
	delegate = Case_Delegate()
	responseData = delegate.deleteAccount( case_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignContact( request, case_Id, ContactId ):
	delegate = Case_Delegate()
	responseData = delegate.saveContact( case_Id, ContactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignContact( request, case_Id ):
	delegate = Case_Delegate()
	responseData = delegate.deleteContact( case_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, case_Id, OwnerId ):
	delegate = Case_Delegate()
	responseData = delegate.saveOwner( case_Id, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, case_Id ):
	delegate = Case_Delegate()
	responseData = delegate.deleteOwner( case_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTeam( request, case_Id, TeamId ):
	delegate = Case_Delegate()
	responseData = delegate.saveTeam( case_Id, TeamId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTeam( request, case_Id ):
	delegate = Case_Delegate()
	responseData = delegate.deleteTeam( case_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addActivities( request, case_Id, ActivitiesIds ):
	delegate = Case_Delegate()
	responseData = delegate.addActivities( case_Id, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeActivities( request, case_Id, ActivitiesIds ):
	delegate = Case_Delegate()
	responseData = delegate.removeActivities( case_Id, ActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCaseComments( request, case_Id, CaseCommentsIds ):
	delegate = Case_Delegate()
	responseData = delegate.addCaseComments( case_Id, CaseCommentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCaseComments( request, case_Id, CaseCommentsIds ):
	delegate = Case_Delegate()
	responseData = delegate.removeCaseComments( case_Id, CaseCommentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmails( request, case_Id, EmailsIds ):
	delegate = Case_Delegate()
	responseData = delegate.addEmails( case_Id, EmailsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmails( request, case_Id, EmailsIds ):
	delegate = Case_Delegate()
	responseData = delegate.removeEmails( case_Id, EmailsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRelatedOpportunities( request, case_Id, RelatedOpportunitiesIds ):
	delegate = Case_Delegate()
	responseData = delegate.addRelatedOpportunities( case_Id, RelatedOpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRelatedOpportunities( request, case_Id, RelatedOpportunitiesIds ):
	delegate = Case_Delegate()
	responseData = delegate.removeRelatedOpportunities( case_Id, RelatedOpportunitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

