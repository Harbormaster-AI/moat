import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.RateDelegate import RateDelegate

 #======================================================================
# 
# Encapsulates data for View Rate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RateView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Rate index.")

def get(request, rateId ):
	delegate = RateDelegate()
	responseData = delegate.get( rateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	rate = json.loads(request.body)
	delegate = RateDelegate()
	responseData = delegate.createFromJson( rate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	rate = json.loads(request.body)
	delegate = RateDelegate()
	responseData = delegate.save( rate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, rateId ):
	delegate = RateDelegate()
	responseData = delegate.delete( rateId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RateDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRateCard( request, rateId, RateCardId ):
	delegate = RateDelegate()
	responseData = delegate.saveRateCard( rateId, RateCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRateCard( request, rateId ):
	delegate = RateDelegate()
	responseData = delegate.deleteRateCard( rateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdSlot( request, rateId, AdSlotId ):
	delegate = RateDelegate()
	responseData = delegate.saveAdSlot( rateId, AdSlotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdSlot( request, rateId ):
	delegate = RateDelegate()
	responseData = delegate.deleteAdSlot( rateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

