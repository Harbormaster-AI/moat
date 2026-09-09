import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.InsertionOrderDelegate import InsertionOrderDelegate

 #======================================================================
# 
# Encapsulates data for View InsertionOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsertionOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InsertionOrder index.")

def get(request, insertionOrderId ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.get( insertionOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	insertionOrder = json.loads(request.body)
	delegate = InsertionOrderDelegate()
	responseData = delegate.createFromJson( insertionOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	insertionOrder = json.loads(request.body)
	delegate = InsertionOrderDelegate()
	responseData = delegate.save( insertionOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, insertionOrderId ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.delete( insertionOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InsertionOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdvertiser( request, insertionOrderId, AdvertiserId ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.saveAdvertiser( insertionOrderId, AdvertiserId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdvertiser( request, insertionOrderId ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.deleteAdvertiser( insertionOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAgency( request, insertionOrderId, AgencyId ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.saveAgency( insertionOrderId, AgencyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAgency( request, insertionOrderId ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.deleteAgency( insertionOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPublisher( request, insertionOrderId, PublisherId ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.savePublisher( insertionOrderId, PublisherId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPublisher( request, insertionOrderId ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.deletePublisher( insertionOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, insertionOrderId, CampaignsIds ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.addCampaigns( insertionOrderId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, insertionOrderId, CampaignsIds ):
	delegate = InsertionOrderDelegate()
	responseData = delegate.removeCampaigns( insertionOrderId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

