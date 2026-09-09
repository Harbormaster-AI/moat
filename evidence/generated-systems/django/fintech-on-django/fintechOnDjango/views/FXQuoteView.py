import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.FXQuoteDelegate import FXQuoteDelegate

 #======================================================================
# 
# Encapsulates data for View FXQuote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXQuoteView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the FXQuote index.")

def get(request, fXQuoteId ):
	delegate = FXQuoteDelegate()
	responseData = delegate.get( fXQuoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	fXQuote = json.loads(request.body)
	delegate = FXQuoteDelegate()
	responseData = delegate.createFromJson( fXQuote )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	fXQuote = json.loads(request.body)
	delegate = FXQuoteDelegate()
	responseData = delegate.save( fXQuote )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, fXQuoteId ):
	delegate = FXQuoteDelegate()
	responseData = delegate.delete( fXQuoteId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FXQuoteDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRequestedBy( request, fXQuoteId, RequestedById ):
	delegate = FXQuoteDelegate()
	responseData = delegate.saveRequestedBy( fXQuoteId, RequestedById )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRequestedBy( request, fXQuoteId ):
	delegate = FXQuoteDelegate()
	responseData = delegate.deleteRequestedBy( fXQuoteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

