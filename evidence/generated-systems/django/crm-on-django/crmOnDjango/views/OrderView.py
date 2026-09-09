import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.OrderDelegate import OrderDelegate

 #======================================================================
# 
# Encapsulates data for View Order
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Order index.")

def get(request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.get( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	order = json.loads(request.body)
	delegate = OrderDelegate()
	responseData = delegate.createFromJson( order )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	order = json.loads(request.body)
	delegate = OrderDelegate()
	responseData = delegate.save( order )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.delete( orderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, orderId, OrganizationId ):
	delegate = OrderDelegate()
	responseData = delegate.saveOrganization( orderId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteOrganization( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, orderId, AccountId ):
	delegate = OrderDelegate()
	responseData = delegate.saveAccount( orderId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteAccount( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOpportunity( request, orderId, OpportunityId ):
	delegate = OrderDelegate()
	responseData = delegate.saveOpportunity( orderId, OpportunityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOpportunity( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteOpportunity( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignQuote( request, orderId, QuoteId ):
	delegate = OrderDelegate()
	responseData = delegate.saveQuote( orderId, QuoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignQuote( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteQuote( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, orderId, OwnerId ):
	delegate = OrderDelegate()
	responseData = delegate.saveOwner( orderId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteOwner( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignContract( request, orderId, ContractId ):
	delegate = OrderDelegate()
	responseData = delegate.saveContract( orderId, ContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignContract( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteContract( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPriceBook( request, orderId, PriceBookId ):
	delegate = OrderDelegate()
	responseData = delegate.savePriceBook( orderId, PriceBookId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPriceBook( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deletePriceBook( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addItems( request, orderId, ItemsIds ):
	delegate = OrderDelegate()
	responseData = delegate.addItems( orderId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeItems( request, orderId, ItemsIds ):
	delegate = OrderDelegate()
	responseData = delegate.removeItems( orderId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

