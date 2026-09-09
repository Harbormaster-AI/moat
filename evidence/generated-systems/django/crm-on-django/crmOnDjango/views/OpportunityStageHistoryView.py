import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.OpportunityStageHistoryDelegate import OpportunityStageHistoryDelegate

 #======================================================================
# 
# Encapsulates data for View OpportunityStageHistory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityStageHistoryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the OpportunityStageHistory index.")

def get(request, opportunityStageHistoryId ):
	delegate = OpportunityStageHistoryDelegate()
	responseData = delegate.get( opportunityStageHistoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	opportunityStageHistory = json.loads(request.body)
	delegate = OpportunityStageHistoryDelegate()
	responseData = delegate.createFromJson( opportunityStageHistory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	opportunityStageHistory = json.loads(request.body)
	delegate = OpportunityStageHistoryDelegate()
	responseData = delegate.save( opportunityStageHistory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, opportunityStageHistoryId ):
	delegate = OpportunityStageHistoryDelegate()
	responseData = delegate.delete( opportunityStageHistoryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OpportunityStageHistoryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOpportunity( request, opportunityStageHistoryId, OpportunityId ):
	delegate = OpportunityStageHistoryDelegate()
	responseData = delegate.saveOpportunity( opportunityStageHistoryId, OpportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOpportunity( request, opportunityStageHistoryId ):
	delegate = OpportunityStageHistoryDelegate()
	responseData = delegate.deleteOpportunity( opportunityStageHistoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignChangedBy( request, opportunityStageHistoryId, ChangedById ):
	delegate = OpportunityStageHistoryDelegate()
	responseData = delegate.saveChangedBy( opportunityStageHistoryId, ChangedById )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignChangedBy( request, opportunityStageHistoryId ):
	delegate = OpportunityStageHistoryDelegate()
	responseData = delegate.deleteChangedBy( opportunityStageHistoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

