import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

 #======================================================================
# 
# Encapsulates data for View Issue
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IssueView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Issue index.")

def get(request, issueId ):
	delegate = IssueDelegate()
	responseData = delegate.get( issueId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	issue = json.loads(request.body)
	delegate = IssueDelegate()
	responseData = delegate.createFromJson( issue )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	issue = json.loads(request.body)
	delegate = IssueDelegate()
	responseData = delegate.save( issue )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, issueId ):
	delegate = IssueDelegate()
	responseData = delegate.delete( issueId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = IssueDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRisk( request, issueId, RiskId ):
	delegate = IssueDelegate()
	responseData = delegate.saveRisk( issueId, RiskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRisk( request, issueId ):
	delegate = IssueDelegate()
	responseData = delegate.deleteRisk( issueId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFinding( request, issueId, FindingId ):
	delegate = IssueDelegate()
	responseData = delegate.saveFinding( issueId, FindingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFinding( request, issueId ):
	delegate = IssueDelegate()
	responseData = delegate.deleteFinding( issueId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignControl( request, issueId, ControlId ):
	delegate = IssueDelegate()
	responseData = delegate.saveControl( issueId, ControlId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignControl( request, issueId ):
	delegate = IssueDelegate()
	responseData = delegate.deleteControl( issueId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCorrectiveActions( request, issueId, CorrectiveActionsIds ):
	delegate = IssueDelegate()
	responseData = delegate.addCorrectiveActions( issueId, CorrectiveActionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCorrectiveActions( request, issueId, CorrectiveActionsIds ):
	delegate = IssueDelegate()
	responseData = delegate.removeCorrectiveActions( issueId, CorrectiveActionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

