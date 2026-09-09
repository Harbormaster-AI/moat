import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.RateCardDelegate import RateCardDelegate

 #======================================================================
# 
# Encapsulates data for View RateCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RateCardView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the RateCard index.")

def get(request, rateCardId ):
	delegate = RateCardDelegate()
	responseData = delegate.get( rateCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	rateCard = json.loads(request.body)
	delegate = RateCardDelegate()
	responseData = delegate.createFromJson( rateCard )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	rateCard = json.loads(request.body)
	delegate = RateCardDelegate()
	responseData = delegate.save( rateCard )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, rateCardId ):
	delegate = RateCardDelegate()
	responseData = delegate.delete( rateCardId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RateCardDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPublisher( request, rateCardId, PublisherId ):
	delegate = RateCardDelegate()
	responseData = delegate.savePublisher( rateCardId, PublisherId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPublisher( request, rateCardId ):
	delegate = RateCardDelegate()
	responseData = delegate.deletePublisher( rateCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRates( request, rateCardId, RatesIds ):
	delegate = RateCardDelegate()
	responseData = delegate.addRates( rateCardId, RatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRates( request, rateCardId, RatesIds ):
	delegate = RateCardDelegate()
	responseData = delegate.removeRates( rateCardId, RatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

