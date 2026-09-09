import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

 #======================================================================
# 
# Encapsulates data for View AdAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdAccountView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AdAccount index.")

def get(request, adAccountId ):
	delegate = AdAccountDelegate()
	responseData = delegate.get( adAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	adAccount = json.loads(request.body)
	delegate = AdAccountDelegate()
	responseData = delegate.createFromJson( adAccount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	adAccount = json.loads(request.body)
	delegate = AdAccountDelegate()
	responseData = delegate.save( adAccount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, adAccountId ):
	delegate = AdAccountDelegate()
	responseData = delegate.delete( adAccountId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AdAccountDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdvertiser( request, adAccountId, AdvertiserId ):
	delegate = AdAccountDelegate()
	responseData = delegate.saveAdvertiser( adAccountId, AdvertiserId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdvertiser( request, adAccountId ):
	delegate = AdAccountDelegate()
	responseData = delegate.deleteAdvertiser( adAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBillingProfile( request, adAccountId, BillingProfileId ):
	delegate = AdAccountDelegate()
	responseData = delegate.saveBillingProfile( adAccountId, BillingProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBillingProfile( request, adAccountId ):
	delegate = AdAccountDelegate()
	responseData = delegate.deleteBillingProfile( adAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDsp( request, adAccountId, DspId ):
	delegate = AdAccountDelegate()
	responseData = delegate.saveDsp( adAccountId, DspId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDsp( request, adAccountId ):
	delegate = AdAccountDelegate()
	responseData = delegate.deleteDsp( adAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addUsers( request, adAccountId, UsersIds ):
	delegate = AdAccountDelegate()
	responseData = delegate.addUsers( adAccountId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeUsers( request, adAccountId, UsersIds ):
	delegate = AdAccountDelegate()
	responseData = delegate.removeUsers( adAccountId, UsersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, adAccountId, CampaignsIds ):
	delegate = AdAccountDelegate()
	responseData = delegate.addCampaigns( adAccountId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, adAccountId, CampaignsIds ):
	delegate = AdAccountDelegate()
	responseData = delegate.removeCampaigns( adAccountId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPerformanceMetrics( request, adAccountId, PerformanceMetricsIds ):
	delegate = AdAccountDelegate()
	responseData = delegate.addPerformanceMetrics( adAccountId, PerformanceMetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePerformanceMetrics( request, adAccountId, PerformanceMetricsIds ):
	delegate = AdAccountDelegate()
	responseData = delegate.removePerformanceMetrics( adAccountId, PerformanceMetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

