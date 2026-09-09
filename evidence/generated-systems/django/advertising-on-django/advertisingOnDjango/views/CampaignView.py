import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

 #======================================================================
# 
# Encapsulates data for View Campaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Campaign index.")

def get(request, campaignId ):
	delegate = CampaignDelegate()
	responseData = delegate.get( campaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	campaign = json.loads(request.body)
	delegate = CampaignDelegate()
	responseData = delegate.createFromJson( campaign )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	campaign = json.loads(request.body)
	delegate = CampaignDelegate()
	responseData = delegate.save( campaign )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, campaignId ):
	delegate = CampaignDelegate()
	responseData = delegate.delete( campaignId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CampaignDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdAccount( request, campaignId, AdAccountId ):
	delegate = CampaignDelegate()
	responseData = delegate.saveAdAccount( campaignId, AdAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdAccount( request, campaignId ):
	delegate = CampaignDelegate()
	responseData = delegate.deleteAdAccount( campaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInsertionOrder( request, campaignId, InsertionOrderId ):
	delegate = CampaignDelegate()
	responseData = delegate.saveInsertionOrder( campaignId, InsertionOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInsertionOrder( request, campaignId ):
	delegate = CampaignDelegate()
	responseData = delegate.deleteInsertionOrder( campaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLineItems( request, campaignId, LineItemsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addLineItems( campaignId, LineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLineItems( request, campaignId, LineItemsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeLineItems( campaignId, LineItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addKpis( request, campaignId, KpisIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addKpis( campaignId, KpisIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeKpis( request, campaignId, KpisIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeKpis( campaignId, KpisIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTrackingPixels( request, campaignId, TrackingPixelsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addTrackingPixels( campaignId, TrackingPixelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTrackingPixels( request, campaignId, TrackingPixelsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeTrackingPixels( campaignId, TrackingPixelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAudiences( request, campaignId, AudiencesIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addAudiences( campaignId, AudiencesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAudiences( request, campaignId, AudiencesIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeAudiences( campaignId, AudiencesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReports( request, campaignId, ReportsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.addReports( campaignId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReports( request, campaignId, ReportsIds ):
	delegate = CampaignDelegate()
	responseData = delegate.removeReports( campaignId, ReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

