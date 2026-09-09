import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

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

def assignWorkspace( request, experimentId, WorkspaceId ):
	delegate = ExperimentDelegate()
	responseData = delegate.saveWorkspace( experimentId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, experimentId ):
	delegate = ExperimentDelegate()
	responseData = delegate.deleteWorkspace( experimentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTrainingRuns( request, experimentId, TrainingRunsIds ):
	delegate = ExperimentDelegate()
	responseData = delegate.addTrainingRuns( experimentId, TrainingRunsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTrainingRuns( request, experimentId, TrainingRunsIds ):
	delegate = ExperimentDelegate()
	responseData = delegate.removeTrainingRuns( experimentId, TrainingRunsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModels( request, experimentId, ModelsIds ):
	delegate = ExperimentDelegate()
	responseData = delegate.addModels( experimentId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModels( request, experimentId, ModelsIds ):
	delegate = ExperimentDelegate()
	responseData = delegate.removeModels( experimentId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addNotebooks( request, experimentId, NotebooksIds ):
	delegate = ExperimentDelegate()
	responseData = delegate.addNotebooks( experimentId, NotebooksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeNotebooks( request, experimentId, NotebooksIds ):
	delegate = ExperimentDelegate()
	responseData = delegate.removeNotebooks( experimentId, NotebooksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

