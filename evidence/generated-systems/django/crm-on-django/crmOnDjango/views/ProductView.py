import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.ProductDelegate import ProductDelegate

 #======================================================================
# 
# Encapsulates data for View Product
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Product index.")

def get(request, productId ):
	delegate = ProductDelegate()
	responseData = delegate.get( productId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	product = json.loads(request.body)
	delegate = ProductDelegate()
	responseData = delegate.createFromJson( product )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	product = json.loads(request.body)
	delegate = ProductDelegate()
	responseData = delegate.save( product )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, productId ):
	delegate = ProductDelegate()
	responseData = delegate.delete( productId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProductDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, productId, OrganizationId ):
	delegate = ProductDelegate()
	responseData = delegate.saveOrganization( productId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, productId ):
	delegate = ProductDelegate()
	responseData = delegate.deleteOrganization( productId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPriceBookEntries( request, productId, PriceBookEntriesIds ):
	delegate = ProductDelegate()
	responseData = delegate.addPriceBookEntries( productId, PriceBookEntriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePriceBookEntries( request, productId, PriceBookEntriesIds ):
	delegate = ProductDelegate()
	responseData = delegate.removePriceBookEntries( productId, PriceBookEntriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOpportunityLineItems( request, productId, OpportunityLineItemsIds ):
	delegate = ProductDelegate()
	responseData = delegate.addOpportunityLineItems( productId, OpportunityLineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOpportunityLineItems( request, productId, OpportunityLineItemsIds ):
	delegate = ProductDelegate()
	responseData = delegate.removeOpportunityLineItems( productId, OpportunityLineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQuoteLineItems( request, productId, QuoteLineItemsIds ):
	delegate = ProductDelegate()
	responseData = delegate.addQuoteLineItems( productId, QuoteLineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQuoteLineItems( request, productId, QuoteLineItemsIds ):
	delegate = ProductDelegate()
	responseData = delegate.removeQuoteLineItems( productId, QuoteLineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrderItems( request, productId, OrderItemsIds ):
	delegate = ProductDelegate()
	responseData = delegate.addOrderItems( productId, OrderItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrderItems( request, productId, OrderItemsIds ):
	delegate = ProductDelegate()
	responseData = delegate.removeOrderItems( productId, OrderItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

