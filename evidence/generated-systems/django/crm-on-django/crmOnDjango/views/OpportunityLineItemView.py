import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.OpportunityLineItemDelegate import OpportunityLineItemDelegate

 #======================================================================
# 
# Encapsulates data for View OpportunityLineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityLineItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the OpportunityLineItem index.")

def get(request, opportunityLineItemId ):
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.get( opportunityLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	opportunityLineItem = json.loads(request.body)
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.createFromJson( opportunityLineItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	opportunityLineItem = json.loads(request.body)
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.save( opportunityLineItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, opportunityLineItemId ):
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.delete( opportunityLineItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOpportunity( request, opportunityLineItemId, OpportunityId ):
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.saveOpportunity( opportunityLineItemId, OpportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOpportunity( request, opportunityLineItemId ):
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.deleteOpportunity( opportunityLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, opportunityLineItemId, ProductId ):
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.saveProduct( opportunityLineItemId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, opportunityLineItemId ):
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.deleteProduct( opportunityLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPriceBookEntry( request, opportunityLineItemId, PriceBookEntryId ):
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.savePriceBookEntry( opportunityLineItemId, PriceBookEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPriceBookEntry( request, opportunityLineItemId ):
	delegate = OpportunityLineItemDelegate()
	responseData = delegate.deletePriceBookEntry( opportunityLineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

