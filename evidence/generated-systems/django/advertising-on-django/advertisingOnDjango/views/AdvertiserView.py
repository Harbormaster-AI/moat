import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.AdvertiserDelegate import AdvertiserDelegate

 #======================================================================
# 
# Encapsulates data for View Advertiser
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdvertiserView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Advertiser index.")

def get(request, advertiserId ):
	delegate = AdvertiserDelegate()
	responseData = delegate.get( advertiserId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	advertiser = json.loads(request.body)
	delegate = AdvertiserDelegate()
	responseData = delegate.createFromJson( advertiser )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	advertiser = json.loads(request.body)
	delegate = AdvertiserDelegate()
	responseData = delegate.save( advertiser )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, advertiserId ):
	delegate = AdvertiserDelegate()
	responseData = delegate.delete( advertiserId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AdvertiserDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAgency( request, advertiserId, AgencyId ):
	delegate = AdvertiserDelegate()
	responseData = delegate.saveAgency( advertiserId, AgencyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAgency( request, advertiserId ):
	delegate = AdvertiserDelegate()
	responseData = delegate.deleteAgency( advertiserId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAdAccounts( request, advertiserId, AdAccountsIds ):
	delegate = AdvertiserDelegate()
	responseData = delegate.addAdAccounts( advertiserId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAdAccounts( request, advertiserId, AdAccountsIds ):
	delegate = AdvertiserDelegate()
	responseData = delegate.removeAdAccounts( advertiserId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBillingProfiles( request, advertiserId, BillingProfilesIds ):
	delegate = AdvertiserDelegate()
	responseData = delegate.addBillingProfiles( advertiserId, BillingProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBillingProfiles( request, advertiserId, BillingProfilesIds ):
	delegate = AdvertiserDelegate()
	responseData = delegate.removeBillingProfiles( advertiserId, BillingProfilesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCampaigns( request, advertiserId, CampaignsIds ):
	delegate = AdvertiserDelegate()
	responseData = delegate.addCampaigns( advertiserId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCampaigns( request, advertiserId, CampaignsIds ):
	delegate = AdvertiserDelegate()
	responseData = delegate.removeCampaigns( advertiserId, CampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTrackingPixels( request, advertiserId, TrackingPixelsIds ):
	delegate = AdvertiserDelegate()
	responseData = delegate.addTrackingPixels( advertiserId, TrackingPixelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTrackingPixels( request, advertiserId, TrackingPixelsIds ):
	delegate = AdvertiserDelegate()
	responseData = delegate.removeTrackingPixels( advertiserId, TrackingPixelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

