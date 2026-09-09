import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.ApplicationDelegate import ApplicationDelegate

 #======================================================================
# 
# Encapsulates data for View Application
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ApplicationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Application index.")

def get(request, applicationId ):
	delegate = ApplicationDelegate()
	responseData = delegate.get( applicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	application = json.loads(request.body)
	delegate = ApplicationDelegate()
	responseData = delegate.createFromJson( application )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	application = json.loads(request.body)
	delegate = ApplicationDelegate()
	responseData = delegate.save( application )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, applicationId ):
	delegate = ApplicationDelegate()
	responseData = delegate.delete( applicationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ApplicationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, applicationId, CustomerId ):
	delegate = ApplicationDelegate()
	responseData = delegate.saveCustomer( applicationId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, applicationId ):
	delegate = ApplicationDelegate()
	responseData = delegate.deleteCustomer( applicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, applicationId, ProductId ):
	delegate = ApplicationDelegate()
	responseData = delegate.saveProduct( applicationId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, applicationId ):
	delegate = ApplicationDelegate()
	responseData = delegate.deleteProduct( applicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDistributor( request, applicationId, DistributorId ):
	delegate = ApplicationDelegate()
	responseData = delegate.saveDistributor( applicationId, DistributorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDistributor( request, applicationId ):
	delegate = ApplicationDelegate()
	responseData = delegate.deleteDistributor( applicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSelectedQuote( request, applicationId, SelectedQuoteId ):
	delegate = ApplicationDelegate()
	responseData = delegate.saveSelectedQuote( applicationId, SelectedQuoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSelectedQuote( request, applicationId ):
	delegate = ApplicationDelegate()
	responseData = delegate.deleteSelectedQuote( applicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQuotes( request, applicationId, QuotesIds ):
	delegate = ApplicationDelegate()
	responseData = delegate.addQuotes( applicationId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQuotes( request, applicationId, QuotesIds ):
	delegate = ApplicationDelegate()
	responseData = delegate.removeQuotes( applicationId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

