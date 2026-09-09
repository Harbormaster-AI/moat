import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.DataTaskDelegate import DataTaskDelegate

 #======================================================================
# 
# Encapsulates data for View DataTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataTaskView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DataTask index.")

def get(request, dataTaskId ):
	delegate = DataTaskDelegate()
	responseData = delegate.get( dataTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dataTask = json.loads(request.body)
	delegate = DataTaskDelegate()
	responseData = delegate.createFromJson( dataTask )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dataTask = json.loads(request.body)
	delegate = DataTaskDelegate()
	responseData = delegate.save( dataTask )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dataTaskId ):
	delegate = DataTaskDelegate()
	responseData = delegate.delete( dataTaskId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DataTaskDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPipeline( request, dataTaskId, PipelineId ):
	delegate = DataTaskDelegate()
	responseData = delegate.savePipeline( dataTaskId, PipelineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPipeline( request, dataTaskId ):
	delegate = DataTaskDelegate()
	responseData = delegate.deletePipeline( dataTaskId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInputDatasets( request, dataTaskId, InputDatasetsIds ):
	delegate = DataTaskDelegate()
	responseData = delegate.addInputDatasets( dataTaskId, InputDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInputDatasets( request, dataTaskId, InputDatasetsIds ):
	delegate = DataTaskDelegate()
	responseData = delegate.removeInputDatasets( dataTaskId, InputDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOutputDatasets( request, dataTaskId, OutputDatasetsIds ):
	delegate = DataTaskDelegate()
	responseData = delegate.addOutputDatasets( dataTaskId, OutputDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOutputDatasets( request, dataTaskId, OutputDatasetsIds ):
	delegate = DataTaskDelegate()
	responseData = delegate.removeOutputDatasets( dataTaskId, OutputDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

