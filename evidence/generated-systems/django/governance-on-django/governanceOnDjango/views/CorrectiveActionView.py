import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

 #======================================================================
# 
# Encapsulates data for View CorrectiveAction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CorrectiveActionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CorrectiveAction index.")

def get(request, correctiveActionId ):
	delegate = CorrectiveActionDelegate()
	responseData = delegate.get( correctiveActionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	correctiveAction = json.loads(request.body)
	delegate = CorrectiveActionDelegate()
	responseData = delegate.createFromJson( correctiveAction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	correctiveAction = json.loads(request.body)
	delegate = CorrectiveActionDelegate()
	responseData = delegate.save( correctiveAction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, correctiveActionId ):
	delegate = CorrectiveActionDelegate()
	responseData = delegate.delete( correctiveActionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CorrectiveActionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFinding( request, correctiveActionId, FindingId ):
	delegate = CorrectiveActionDelegate()
	responseData = delegate.saveFinding( correctiveActionId, FindingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFinding( request, correctiveActionId ):
	delegate = CorrectiveActionDelegate()
	responseData = delegate.deleteFinding( correctiveActionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignIssue( request, correctiveActionId, IssueId ):
	delegate = CorrectiveActionDelegate()
	responseData = delegate.saveIssue( correctiveActionId, IssueId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignIssue( request, correctiveActionId ):
	delegate = CorrectiveActionDelegate()
	responseData = delegate.deleteIssue( correctiveActionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

