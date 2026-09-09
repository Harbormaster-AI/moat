import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.InferenceEndpointDelegate import InferenceEndpointDelegate

 #======================================================================
# 
# Encapsulates data for View InferenceEndpoint
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InferenceEndpointView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InferenceEndpoint index.")

def get(request, inferenceEndpointId ):
	delegate = InferenceEndpointDelegate()
	responseData = delegate.get( inferenceEndpointId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inferenceEndpoint = json.loads(request.body)
	delegate = InferenceEndpointDelegate()
	responseData = delegate.createFromJson( inferenceEndpoint )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inferenceEndpoint = json.loads(request.body)
	delegate = InferenceEndpointDelegate()
	responseData = delegate.save( inferenceEndpoint )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inferenceEndpointId ):
	delegate = InferenceEndpointDelegate()
	responseData = delegate.delete( inferenceEndpointId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InferenceEndpointDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignModelVersion( request, inferenceEndpointId, ModelVersionId ):
	delegate = InferenceEndpointDelegate()
	responseData = delegate.saveModelVersion( inferenceEndpointId, ModelVersionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignModelVersion( request, inferenceEndpointId ):
	delegate = InferenceEndpointDelegate()
	responseData = delegate.deleteModelVersion( inferenceEndpointId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, inferenceEndpointId, WorkspaceId ):
	delegate = InferenceEndpointDelegate()
	responseData = delegate.saveWorkspace( inferenceEndpointId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, inferenceEndpointId ):
	delegate = InferenceEndpointDelegate()
	responseData = delegate.deleteWorkspace( inferenceEndpointId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPredictions( request, inferenceEndpointId, PredictionsIds ):
	delegate = InferenceEndpointDelegate()
	responseData = delegate.addPredictions( inferenceEndpointId, PredictionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePredictions( request, inferenceEndpointId, PredictionsIds ):
	delegate = InferenceEndpointDelegate()
	responseData = delegate.removePredictions( inferenceEndpointId, PredictionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

