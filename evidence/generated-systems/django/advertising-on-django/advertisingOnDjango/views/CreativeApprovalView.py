import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.CreativeApprovalDelegate import CreativeApprovalDelegate

 #======================================================================
# 
# Encapsulates data for View CreativeApproval
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeApprovalView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CreativeApproval index.")

def get(request, creativeApprovalId ):
	delegate = CreativeApprovalDelegate()
	responseData = delegate.get( creativeApprovalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	creativeApproval = json.loads(request.body)
	delegate = CreativeApprovalDelegate()
	responseData = delegate.createFromJson( creativeApproval )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	creativeApproval = json.loads(request.body)
	delegate = CreativeApprovalDelegate()
	responseData = delegate.save( creativeApproval )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, creativeApprovalId ):
	delegate = CreativeApprovalDelegate()
	responseData = delegate.delete( creativeApprovalId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CreativeApprovalDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCreativeAsset( request, creativeApprovalId, CreativeAssetId ):
	delegate = CreativeApprovalDelegate()
	responseData = delegate.saveCreativeAsset( creativeApprovalId, CreativeAssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCreativeAsset( request, creativeApprovalId ):
	delegate = CreativeApprovalDelegate()
	responseData = delegate.deleteCreativeAsset( creativeApprovalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPublisher( request, creativeApprovalId, PublisherId ):
	delegate = CreativeApprovalDelegate()
	responseData = delegate.savePublisher( creativeApprovalId, PublisherId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPublisher( request, creativeApprovalId ):
	delegate = CreativeApprovalDelegate()
	responseData = delegate.deletePublisher( creativeApprovalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

