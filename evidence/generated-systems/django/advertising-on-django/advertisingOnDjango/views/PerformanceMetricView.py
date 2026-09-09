import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.PerformanceMetricDelegate import PerformanceMetricDelegate

 #======================================================================
# 
# Encapsulates data for View PerformanceMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceMetricView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PerformanceMetric index.")

def get(request, performanceMetricId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.get( performanceMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	performanceMetric = json.loads(request.body)
	delegate = PerformanceMetricDelegate()
	responseData = delegate.createFromJson( performanceMetric )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	performanceMetric = json.loads(request.body)
	delegate = PerformanceMetricDelegate()
	responseData = delegate.save( performanceMetric )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, performanceMetricId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.delete( performanceMetricId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdAccount( request, performanceMetricId, AdAccountId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.saveAdAccount( performanceMetricId, AdAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdAccount( request, performanceMetricId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.deleteAdAccount( performanceMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, performanceMetricId, CampaignId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.saveCampaign( performanceMetricId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, performanceMetricId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.deleteCampaign( performanceMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLineItem( request, performanceMetricId, LineItemId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.saveLineItem( performanceMetricId, LineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLineItem( request, performanceMetricId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.deleteLineItem( performanceMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlacement( request, performanceMetricId, PlacementId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.savePlacement( performanceMetricId, PlacementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlacement( request, performanceMetricId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.deletePlacement( performanceMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCreativeAsset( request, performanceMetricId, CreativeAssetId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.saveCreativeAsset( performanceMetricId, CreativeAssetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCreativeAsset( request, performanceMetricId ):
	delegate = PerformanceMetricDelegate()
	responseData = delegate.deleteCreativeAsset( performanceMetricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

