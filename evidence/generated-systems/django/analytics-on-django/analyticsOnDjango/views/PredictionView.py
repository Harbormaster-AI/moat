import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.PredictionDelegate import PredictionDelegate

 #======================================================================
# 
# Encapsulates data for View Prediction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PredictionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Prediction index.")

def get(request, predictionId ):
	delegate = PredictionDelegate()
	responseData = delegate.get( predictionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	prediction = json.loads(request.body)
	delegate = PredictionDelegate()
	responseData = delegate.createFromJson( prediction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	prediction = json.loads(request.body)
	delegate = PredictionDelegate()
	responseData = delegate.save( prediction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, predictionId ):
	delegate = PredictionDelegate()
	responseData = delegate.delete( predictionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PredictionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEndpoint( request, predictionId, EndpointId ):
	delegate = PredictionDelegate()
	responseData = delegate.saveEndpoint( predictionId, EndpointId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEndpoint( request, predictionId ):
	delegate = PredictionDelegate()
	responseData = delegate.deleteEndpoint( predictionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignModelVersion( request, predictionId, ModelVersionId ):
	delegate = PredictionDelegate()
	responseData = delegate.saveModelVersion( predictionId, ModelVersionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignModelVersion( request, predictionId ):
	delegate = PredictionDelegate()
	responseData = delegate.deleteModelVersion( predictionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDataset( request, predictionId, DatasetId ):
	delegate = PredictionDelegate()
	responseData = delegate.saveDataset( predictionId, DatasetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDataset( request, predictionId ):
	delegate = PredictionDelegate()
	responseData = delegate.deleteDataset( predictionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

