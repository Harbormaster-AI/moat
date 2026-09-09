import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.CardTokenizationDelegate import CardTokenizationDelegate

 #======================================================================
# 
# Encapsulates data for View CardTokenization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CardTokenizationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CardTokenization index.")

def get(request, cardTokenizationId ):
	delegate = CardTokenizationDelegate()
	responseData = delegate.get( cardTokenizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	cardTokenization = json.loads(request.body)
	delegate = CardTokenizationDelegate()
	responseData = delegate.createFromJson( cardTokenization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	cardTokenization = json.loads(request.body)
	delegate = CardTokenizationDelegate()
	responseData = delegate.save( cardTokenization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, cardTokenizationId ):
	delegate = CardTokenizationDelegate()
	responseData = delegate.delete( cardTokenizationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CardTokenizationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCard( request, cardTokenizationId, CardId ):
	delegate = CardTokenizationDelegate()
	responseData = delegate.saveCard( cardTokenizationId, CardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCard( request, cardTokenizationId ):
	delegate = CardTokenizationDelegate()
	responseData = delegate.deleteCard( cardTokenizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

