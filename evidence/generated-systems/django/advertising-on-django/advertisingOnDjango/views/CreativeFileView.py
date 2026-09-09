import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.CreativeFileDelegate import CreativeFileDelegate

 #======================================================================
# 
# Encapsulates data for View CreativeFile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeFileView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CreativeFile index.")

def get(request, creativeFileId ):
	delegate = CreativeFileDelegate()
	responseData = delegate.get( creativeFileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	creativeFile = json.loads(request.body)
	delegate = CreativeFileDelegate()
	responseData = delegate.createFromJson( creativeFile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	creativeFile = json.loads(request.body)
	delegate = CreativeFileDelegate()
	responseData = delegate.save( creativeFile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, creativeFileId ):
	delegate = CreativeFileDelegate()
	responseData = delegate.delete( creativeFileId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CreativeFileDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCreativeAsset( request, creativeFileId, CreativeAssetId ):
	delegate = CreativeFileDelegate()
	responseData = delegate.saveCreativeAsset( creativeFileId, CreativeAssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCreativeAsset( request, creativeFileId ):
	delegate = CreativeFileDelegate()
	responseData = delegate.deleteCreativeAsset( creativeFileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

