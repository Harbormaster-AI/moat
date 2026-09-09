import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.QuoteLineItemDelegate import QuoteLineItemDelegate

 #======================================================================
# 
# Encapsulates data for View QuoteLineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteLineItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the QuoteLineItem index.")

def get(request, quoteLineItemId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.get( quoteLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	quoteLineItem = json.loads(request.body)
	delegate = QuoteLineItemDelegate()
	responseData = delegate.createFromJson( quoteLineItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	quoteLineItem = json.loads(request.body)
	delegate = QuoteLineItemDelegate()
	responseData = delegate.save( quoteLineItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, quoteLineItemId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.delete( quoteLineItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignQuote( request, quoteLineItemId, QuoteId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.saveQuote( quoteLineItemId, QuoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignQuote( request, quoteLineItemId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.deleteQuote( quoteLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, quoteLineItemId, ProductId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.saveProduct( quoteLineItemId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, quoteLineItemId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.deleteProduct( quoteLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPriceBookEntry( request, quoteLineItemId, PriceBookEntryId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.savePriceBookEntry( quoteLineItemId, PriceBookEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPriceBookEntry( request, quoteLineItemId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.deletePriceBookEntry( quoteLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOpportunityLineItem( request, quoteLineItemId, OpportunityLineItemId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.saveOpportunityLineItem( quoteLineItemId, OpportunityLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOpportunityLineItem( request, quoteLineItemId ):
	delegate = QuoteLineItemDelegate()
	responseData = delegate.deleteOpportunityLineItem( quoteLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

