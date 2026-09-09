import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.DeviceCriterionDelegate import DeviceCriterionDelegate

 #======================================================================
# 
# Encapsulates data for View DeviceCriterion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DeviceCriterionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DeviceCriterion index.")

def get(request, deviceCriterionId ):
	delegate = DeviceCriterionDelegate()
	responseData = delegate.get( deviceCriterionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	deviceCriterion = json.loads(request.body)
	delegate = DeviceCriterionDelegate()
	responseData = delegate.createFromJson( deviceCriterion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	deviceCriterion = json.loads(request.body)
	delegate = DeviceCriterionDelegate()
	responseData = delegate.save( deviceCriterion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, deviceCriterionId ):
	delegate = DeviceCriterionDelegate()
	responseData = delegate.delete( deviceCriterionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DeviceCriterionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTargetingProfile( request, deviceCriterionId, TargetingProfileId ):
	delegate = DeviceCriterionDelegate()
	responseData = delegate.saveTargetingProfile( deviceCriterionId, TargetingProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTargetingProfile( request, deviceCriterionId ):
	delegate = DeviceCriterionDelegate()
	responseData = delegate.deleteTargetingProfile( deviceCriterionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

