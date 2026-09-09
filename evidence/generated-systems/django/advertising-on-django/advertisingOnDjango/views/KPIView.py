import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.KPIDelegate import KPIDelegate

 #======================================================================
# 
# Encapsulates data for View KPI
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KPIView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the KPI index.")

def get(request, kPIId ):
	delegate = KPIDelegate()
	responseData = delegate.get( kPIId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	kPI = json.loads(request.body)
	delegate = KPIDelegate()
	responseData = delegate.createFromJson( kPI )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	kPI = json.loads(request.body)
	delegate = KPIDelegate()
	responseData = delegate.save( kPI )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, kPIId ):
	delegate = KPIDelegate()
	responseData = delegate.delete( kPIId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = KPIDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, kPIId, CampaignId ):
	delegate = KPIDelegate()
	responseData = delegate.saveCampaign( kPIId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, kPIId ):
	delegate = KPIDelegate()
	responseData = delegate.deleteCampaign( kPIId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

