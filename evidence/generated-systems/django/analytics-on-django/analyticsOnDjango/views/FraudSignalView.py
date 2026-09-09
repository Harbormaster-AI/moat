import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.FraudSignalDelegate import FraudSignalDelegate

 #======================================================================
# 
# Encapsulates data for View FraudSignal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudSignalView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the FraudSignal index.")

def get(request, fraudSignalId ):
	delegate = FraudSignalDelegate()
	responseData = delegate.get( fraudSignalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	fraudSignal = json.loads(request.body)
	delegate = FraudSignalDelegate()
	responseData = delegate.createFromJson( fraudSignal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	fraudSignal = json.loads(request.body)
	delegate = FraudSignalDelegate()
	responseData = delegate.save( fraudSignal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, fraudSignalId ):
	delegate = FraudSignalDelegate()
	responseData = delegate.delete( fraudSignalId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FraudSignalDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignScenario( request, fraudSignalId, ScenarioId ):
	delegate = FraudSignalDelegate()
	responseData = delegate.saveScenario( fraudSignalId, ScenarioId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignScenario( request, fraudSignalId ):
	delegate = FraudSignalDelegate()
	responseData = delegate.deleteScenario( fraudSignalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDataset( request, fraudSignalId, DatasetId ):
	delegate = FraudSignalDelegate()
	responseData = delegate.saveDataset( fraudSignalId, DatasetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDataset( request, fraudSignalId ):
	delegate = FraudSignalDelegate()
	responseData = delegate.deleteDataset( fraudSignalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignModelVersion( request, fraudSignalId, ModelVersionId ):
	delegate = FraudSignalDelegate()
	responseData = delegate.saveModelVersion( fraudSignalId, ModelVersionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignModelVersion( request, fraudSignalId ):
	delegate = FraudSignalDelegate()
	responseData = delegate.deleteModelVersion( fraudSignalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

