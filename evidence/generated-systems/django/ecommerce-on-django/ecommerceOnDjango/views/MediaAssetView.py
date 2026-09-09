import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.MediaAssetDelegate import MediaAssetDelegate

 #======================================================================
# 
# Encapsulates data for View MediaAsset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MediaAssetView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MediaAsset index.")

def get(request, mediaAssetId ):
	delegate = MediaAssetDelegate()
	responseData = delegate.get( mediaAssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	mediaAsset = json.loads(request.body)
	delegate = MediaAssetDelegate()
	responseData = delegate.createFromJson( mediaAsset )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	mediaAsset = json.loads(request.body)
	delegate = MediaAssetDelegate()
	responseData = delegate.save( mediaAsset )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, mediaAssetId ):
	delegate = MediaAssetDelegate()
	responseData = delegate.delete( mediaAssetId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MediaAssetDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, mediaAssetId, ProductId ):
	delegate = MediaAssetDelegate()
	responseData = delegate.saveProduct( mediaAssetId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, mediaAssetId ):
	delegate = MediaAssetDelegate()
	responseData = delegate.deleteProduct( mediaAssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, mediaAssetId, VariantId ):
	delegate = MediaAssetDelegate()
	responseData = delegate.saveVariant( mediaAssetId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, mediaAssetId ):
	delegate = MediaAssetDelegate()
	responseData = delegate.deleteVariant( mediaAssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

