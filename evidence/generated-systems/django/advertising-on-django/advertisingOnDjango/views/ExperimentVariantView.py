import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.ExperimentVariantDelegate import ExperimentVariantDelegate

 #======================================================================
# 
# Encapsulates data for View ExperimentVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExperimentVariantView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ExperimentVariant index.")

def get(request, experimentVariantId ):
	delegate = ExperimentVariantDelegate()
	responseData = delegate.get( experimentVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	experimentVariant = json.loads(request.body)
	delegate = ExperimentVariantDelegate()
	responseData = delegate.createFromJson( experimentVariant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	experimentVariant = json.loads(request.body)
	delegate = ExperimentVariantDelegate()
	responseData = delegate.save( experimentVariant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, experimentVariantId ):
	delegate = ExperimentVariantDelegate()
	responseData = delegate.delete( experimentVariantId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ExperimentVariantDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignExperiment( request, experimentVariantId, ExperimentId ):
	delegate = ExperimentVariantDelegate()
	responseData = delegate.saveExperiment( experimentVariantId, ExperimentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignExperiment( request, experimentVariantId ):
	delegate = ExperimentVariantDelegate()
	responseData = delegate.deleteExperiment( experimentVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCreativeVariation( request, experimentVariantId, CreativeVariationId ):
	delegate = ExperimentVariantDelegate()
	responseData = delegate.saveCreativeVariation( experimentVariantId, CreativeVariationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCreativeVariation( request, experimentVariantId ):
	delegate = ExperimentVariantDelegate()
	responseData = delegate.deleteCreativeVariation( experimentVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLineItem( request, experimentVariantId, LineItemId ):
	delegate = ExperimentVariantDelegate()
	responseData = delegate.saveLineItem( experimentVariantId, LineItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLineItem( request, experimentVariantId ):
	delegate = ExperimentVariantDelegate()
	responseData = delegate.deleteLineItem( experimentVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

