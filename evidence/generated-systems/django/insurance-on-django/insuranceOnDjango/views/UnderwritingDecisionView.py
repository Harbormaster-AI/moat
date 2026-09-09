import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.UnderwritingDecisionDelegate import UnderwritingDecisionDelegate

 #======================================================================
# 
# Encapsulates data for View UnderwritingDecision
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnderwritingDecisionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the UnderwritingDecision index.")

def get(request, underwritingDecisionId ):
	delegate = UnderwritingDecisionDelegate()
	responseData = delegate.get( underwritingDecisionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	underwritingDecision = json.loads(request.body)
	delegate = UnderwritingDecisionDelegate()
	responseData = delegate.createFromJson( underwritingDecision )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	underwritingDecision = json.loads(request.body)
	delegate = UnderwritingDecisionDelegate()
	responseData = delegate.save( underwritingDecision )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, underwritingDecisionId ):
	delegate = UnderwritingDecisionDelegate()
	responseData = delegate.delete( underwritingDecisionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = UnderwritingDecisionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignQuote( request, underwritingDecisionId, QuoteId ):
	delegate = UnderwritingDecisionDelegate()
	responseData = delegate.saveQuote( underwritingDecisionId, QuoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignQuote( request, underwritingDecisionId ):
	delegate = UnderwritingDecisionDelegate()
	responseData = delegate.deleteQuote( underwritingDecisionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignUnderwriter( request, underwritingDecisionId, UnderwriterId ):
	delegate = UnderwritingDecisionDelegate()
	responseData = delegate.saveUnderwriter( underwritingDecisionId, UnderwriterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignUnderwriter( request, underwritingDecisionId ):
	delegate = UnderwritingDecisionDelegate()
	responseData = delegate.deleteUnderwriter( underwritingDecisionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

