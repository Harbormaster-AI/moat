import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.NotebookDelegate import NotebookDelegate

 #======================================================================
# 
# Encapsulates data for View Notebook
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NotebookView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Notebook index.")

def get(request, notebookId ):
	delegate = NotebookDelegate()
	responseData = delegate.get( notebookId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	notebook = json.loads(request.body)
	delegate = NotebookDelegate()
	responseData = delegate.createFromJson( notebook )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	notebook = json.loads(request.body)
	delegate = NotebookDelegate()
	responseData = delegate.save( notebook )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, notebookId ):
	delegate = NotebookDelegate()
	responseData = delegate.delete( notebookId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = NotebookDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, notebookId, WorkspaceId ):
	delegate = NotebookDelegate()
	responseData = delegate.saveWorkspace( notebookId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, notebookId ):
	delegate = NotebookDelegate()
	responseData = delegate.deleteWorkspace( notebookId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, notebookId, DatasetsIds ):
	delegate = NotebookDelegate()
	responseData = delegate.addDatasets( notebookId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, notebookId, DatasetsIds ):
	delegate = NotebookDelegate()
	responseData = delegate.removeDatasets( notebookId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addExperiments( request, notebookId, ExperimentsIds ):
	delegate = NotebookDelegate()
	responseData = delegate.addExperiments( notebookId, ExperimentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeExperiments( request, notebookId, ExperimentsIds ):
	delegate = NotebookDelegate()
	responseData = delegate.removeExperiments( notebookId, ExperimentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQueries( request, notebookId, QueriesIds ):
	delegate = NotebookDelegate()
	responseData = delegate.addQueries( notebookId, QueriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQueries( request, notebookId, QueriesIds ):
	delegate = NotebookDelegate()
	responseData = delegate.removeQueries( notebookId, QueriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

