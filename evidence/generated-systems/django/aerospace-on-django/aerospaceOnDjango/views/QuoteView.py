import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.QuoteDelegate import QuoteDelegate

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

def assignAircraftOrder( request, quoteId, AircraftOrderId ):
	delegate = QuoteDelegate()
	responseData = delegate.saveAircraftOrder( quoteId, AircraftOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAircraftOrder( request, quoteId ):
	delegate = QuoteDelegate()
	responseData = delegate.deleteAircraftOrder( quoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

