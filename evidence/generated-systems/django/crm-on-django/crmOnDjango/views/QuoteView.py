import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

 #======================================================================
# 
# Encapsulates data for View Quote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Quote index.")

def get(request, quoteId ):
	delegate = QuoteDelegate()
	responseData = delegate.get( quoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	quote = json.loads(request.body)
	delegate = QuoteDelegate()
	responseData = delegate.createFromJson( quote )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	quote = json.loads(request.body)
	delegate = QuoteDelegate()
	responseData = delegate.save( quote )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, quoteId ):
	delegate = QuoteDelegate()
	responseData = delegate.delete( quoteId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = QuoteDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, quoteId, OrganizationId ):
	delegate = QuoteDelegate()
	responseData = delegate.saveOrganization( quoteId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, quoteId ):
	delegate = QuoteDelegate()
	responseData = delegate.deleteOrganization( quoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, quoteId, AccountId ):
	delegate = QuoteDelegate()
	responseData = delegate.saveAccount( quoteId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, quoteId ):
	delegate = QuoteDelegate()
	responseData = delegate.deleteAccount( quoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOpportunity( request, quoteId, OpportunityId ):
	delegate = QuoteDelegate()
	responseData = delegate.saveOpportunity( quoteId, OpportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOpportunity( request, quoteId ):
	delegate = QuoteDelegate()
	responseData = delegate.deleteOpportunity( quoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, quoteId, OwnerId ):
	delegate = QuoteDelegate()
	responseData = delegate.saveOwner( quoteId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, quoteId ):
	delegate = QuoteDelegate()
	responseData = delegate.deleteOwner( quoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPriceBook( request, quoteId, PriceBookId ):
	delegate = QuoteDelegate()
	responseData = delegate.savePriceBook( quoteId, PriceBookId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPriceBook( request, quoteId ):
	delegate = QuoteDelegate()
	responseData = delegate.deletePriceBook( quoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, quoteId, OrderId ):
	delegate = QuoteDelegate()
	responseData = delegate.saveOrder( quoteId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, quoteId ):
	delegate = QuoteDelegate()
	responseData = delegate.deleteOrder( quoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLineItems( request, quoteId, LineItemsIds ):
	delegate = QuoteDelegate()
	responseData = delegate.addLineItems( quoteId, LineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLineItems( request, quoteId, LineItemsIds ):
	delegate = QuoteDelegate()
	responseData = delegate.removeLineItems( quoteId, LineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

