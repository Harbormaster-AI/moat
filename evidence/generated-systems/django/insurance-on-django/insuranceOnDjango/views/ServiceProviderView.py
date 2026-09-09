import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.ServiceProviderDelegate import ServiceProviderDelegate

 #======================================================================
# 
# Encapsulates data for View ServiceProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceProviderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ServiceProvider index.")

def get(request, serviceProviderId ):
	delegate = ServiceProviderDelegate()
	responseData = delegate.get( serviceProviderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	serviceProvider = json.loads(request.body)
	delegate = ServiceProviderDelegate()
	responseData = delegate.createFromJson( serviceProvider )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	serviceProvider = json.loads(request.body)
	delegate = ServiceProviderDelegate()
	responseData = delegate.save( serviceProvider )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, serviceProviderId ):
	delegate = ServiceProviderDelegate()
	responseData = delegate.delete( serviceProviderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ServiceProviderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClaims( request, serviceProviderId, ClaimsIds ):
	delegate = ServiceProviderDelegate()
	responseData = delegate.addClaims( serviceProviderId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClaims( request, serviceProviderId, ClaimsIds ):
	delegate = ServiceProviderDelegate()
	responseData = delegate.removeClaims( serviceProviderId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

