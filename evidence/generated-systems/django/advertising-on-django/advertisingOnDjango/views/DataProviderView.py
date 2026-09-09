import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.DataProviderDelegate import DataProviderDelegate

 #======================================================================
# 
# Encapsulates data for View DataProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataProviderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DataProvider index.")

def get(request, dataProviderId ):
	delegate = DataProviderDelegate()
	responseData = delegate.get( dataProviderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dataProvider = json.loads(request.body)
	delegate = DataProviderDelegate()
	responseData = delegate.createFromJson( dataProvider )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dataProvider = json.loads(request.body)
	delegate = DataProviderDelegate()
	responseData = delegate.save( dataProvider )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dataProviderId ):
	delegate = DataProviderDelegate()
	responseData = delegate.delete( dataProviderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DataProviderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAudienceSegments( request, dataProviderId, AudienceSegmentsIds ):
	delegate = DataProviderDelegate()
	responseData = delegate.addAudienceSegments( dataProviderId, AudienceSegmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAudienceSegments( request, dataProviderId, AudienceSegmentsIds ):
	delegate = DataProviderDelegate()
	responseData = delegate.removeAudienceSegments( dataProviderId, AudienceSegmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

