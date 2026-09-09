import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.PriceBookDelegate import PriceBookDelegate

 #======================================================================
# 
# Encapsulates data for View PriceBook
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PriceBookView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PriceBook index.")

def get(request, priceBookId ):
	delegate = PriceBookDelegate()
	responseData = delegate.get( priceBookId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	priceBook = json.loads(request.body)
	delegate = PriceBookDelegate()
	responseData = delegate.createFromJson( priceBook )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	priceBook = json.loads(request.body)
	delegate = PriceBookDelegate()
	responseData = delegate.save( priceBook )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, priceBookId ):
	delegate = PriceBookDelegate()
	responseData = delegate.delete( priceBookId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PriceBookDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, priceBookId, OrganizationId ):
	delegate = PriceBookDelegate()
	responseData = delegate.saveOrganization( priceBookId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, priceBookId ):
	delegate = PriceBookDelegate()
	responseData = delegate.deleteOrganization( priceBookId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEntries( request, priceBookId, EntriesIds ):
	delegate = PriceBookDelegate()
	responseData = delegate.addEntries( priceBookId, EntriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEntries( request, priceBookId, EntriesIds ):
	delegate = PriceBookDelegate()
	responseData = delegate.removeEntries( priceBookId, EntriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQuotes( request, priceBookId, QuotesIds ):
	delegate = PriceBookDelegate()
	responseData = delegate.addQuotes( priceBookId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQuotes( request, priceBookId, QuotesIds ):
	delegate = PriceBookDelegate()
	responseData = delegate.removeQuotes( priceBookId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, priceBookId, OrdersIds ):
	delegate = PriceBookDelegate()
	responseData = delegate.addOrders( priceBookId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, priceBookId, OrdersIds ):
	delegate = PriceBookDelegate()
	responseData = delegate.removeOrders( priceBookId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

