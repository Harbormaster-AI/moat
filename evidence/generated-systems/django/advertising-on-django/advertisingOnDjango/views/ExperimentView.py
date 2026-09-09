import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

 #======================================================================
# 
# Encapsulates data for View Experiment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExperimentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Experiment index.")

def get(request, experimentId ):
	delegate = ExperimentDelegate()
	responseData = delegate.get( experimentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	experiment = json.loads(request.body)
	delegate = ExperimentDelegate()
	responseData = delegate.createFromJson( experiment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	experiment = json.loads(request.body)
	delegate = ExperimentDelegate()
	responseData = delegate.save( experiment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, experimentId ):
	delegate = ExperimentDelegate()
	responseData = delegate.delete( experimentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ExperimentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, experimentId, CampaignId ):
	delegate = ExperimentDelegate()
	responseData = delegate.saveCampaign( experimentId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, experimentId ):
	delegate = ExperimentDelegate()
	responseData = delegate.deleteCampaign( experimentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariants( request, experimentId, VariantsIds ):
	delegate = ExperimentDelegate()
	responseData = delegate.addVariants( experimentId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariants( request, experimentId, VariantsIds ):
	delegate = ExperimentDelegate()
	responseData = delegate.removeVariants( experimentId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

