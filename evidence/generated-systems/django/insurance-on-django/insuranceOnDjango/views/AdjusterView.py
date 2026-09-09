import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.AdjusterDelegate import AdjusterDelegate

 #======================================================================
# 
# Encapsulates data for View Adjuster
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdjusterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Adjuster index.")

def get(request, adjusterId ):
	delegate = AdjusterDelegate()
	responseData = delegate.get( adjusterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	adjuster = json.loads(request.body)
	delegate = AdjusterDelegate()
	responseData = delegate.createFromJson( adjuster )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	adjuster = json.loads(request.body)
	delegate = AdjusterDelegate()
	responseData = delegate.save( adjuster )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, adjusterId ):
	delegate = AdjusterDelegate()
	responseData = delegate.delete( adjusterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AdjusterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClaims( request, adjusterId, ClaimsIds ):
	delegate = AdjusterDelegate()
	responseData = delegate.addClaims( adjusterId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClaims( request, adjusterId, ClaimsIds ):
	delegate = AdjusterDelegate()
	responseData = delegate.removeClaims( adjusterId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addServiceProviders( request, adjusterId, ServiceProvidersIds ):
	delegate = AdjusterDelegate()
	responseData = delegate.addServiceProviders( adjusterId, ServiceProvidersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeServiceProviders( request, adjusterId, ServiceProvidersIds ):
	delegate = AdjusterDelegate()
	responseData = delegate.removeServiceProviders( adjusterId, ServiceProvidersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

