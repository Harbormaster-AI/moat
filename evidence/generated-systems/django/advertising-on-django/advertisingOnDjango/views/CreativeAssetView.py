import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.CreativeAssetDelegate import CreativeAssetDelegate

 #======================================================================
# 
# Encapsulates data for View CreativeAsset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeAssetView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CreativeAsset index.")

def get(request, creativeAssetId ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.get( creativeAssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	creativeAsset = json.loads(request.body)
	delegate = CreativeAssetDelegate()
	responseData = delegate.createFromJson( creativeAsset )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	creativeAsset = json.loads(request.body)
	delegate = CreativeAssetDelegate()
	responseData = delegate.save( creativeAsset )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, creativeAssetId ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.delete( creativeAssetId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CreativeAssetDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFiles( request, creativeAssetId, FilesIds ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.addFiles( creativeAssetId, FilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFiles( request, creativeAssetId, FilesIds ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.removeFiles( creativeAssetId, FilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addApprovals( request, creativeAssetId, ApprovalsIds ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.addApprovals( creativeAssetId, ApprovalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeApprovals( request, creativeAssetId, ApprovalsIds ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.removeApprovals( creativeAssetId, ApprovalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariations( request, creativeAssetId, VariationsIds ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.addVariations( creativeAssetId, VariationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariations( request, creativeAssetId, VariationsIds ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.removeVariations( creativeAssetId, VariationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLineItems( request, creativeAssetId, LineItemsIds ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.addLineItems( creativeAssetId, LineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLineItems( request, creativeAssetId, LineItemsIds ):
	delegate = CreativeAssetDelegate()
	responseData = delegate.removeLineItems( creativeAssetId, LineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

