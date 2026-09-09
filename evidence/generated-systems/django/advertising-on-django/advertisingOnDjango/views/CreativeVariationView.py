import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.CreativeVariationDelegate import CreativeVariationDelegate

 #======================================================================
# 
# Encapsulates data for View CreativeVariation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeVariationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CreativeVariation index.")

def get(request, creativeVariationId ):
	delegate = CreativeVariationDelegate()
	responseData = delegate.get( creativeVariationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	creativeVariation = json.loads(request.body)
	delegate = CreativeVariationDelegate()
	responseData = delegate.createFromJson( creativeVariation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	creativeVariation = json.loads(request.body)
	delegate = CreativeVariationDelegate()
	responseData = delegate.save( creativeVariation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, creativeVariationId ):
	delegate = CreativeVariationDelegate()
	responseData = delegate.delete( creativeVariationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CreativeVariationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCreativeAsset( request, creativeVariationId, CreativeAssetId ):
	delegate = CreativeVariationDelegate()
	responseData = delegate.saveCreativeAsset( creativeVariationId, CreativeAssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCreativeAsset( request, creativeVariationId ):
	delegate = CreativeVariationDelegate()
	responseData = delegate.deleteCreativeAsset( creativeVariationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

