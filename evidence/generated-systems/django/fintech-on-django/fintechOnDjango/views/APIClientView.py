import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.APIClientDelegate import APIClientDelegate

 #======================================================================
# 
# Encapsulates data for View APIClient
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class APIClientView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the APIClient index.")

def get(request, aPIClientId ):
	delegate = APIClientDelegate()
	responseData = delegate.get( aPIClientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aPIClient = json.loads(request.body)
	delegate = APIClientDelegate()
	responseData = delegate.createFromJson( aPIClient )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aPIClient = json.loads(request.body)
	delegate = APIClientDelegate()
	responseData = delegate.save( aPIClient )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aPIClientId ):
	delegate = APIClientDelegate()
	responseData = delegate.delete( aPIClientId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = APIClientDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addConsents( request, aPIClientId, ConsentsIds ):
	delegate = APIClientDelegate()
	responseData = delegate.addConsents( aPIClientId, ConsentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeConsents( request, aPIClientId, ConsentsIds ):
	delegate = APIClientDelegate()
	responseData = delegate.removeConsents( aPIClientId, ConsentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

