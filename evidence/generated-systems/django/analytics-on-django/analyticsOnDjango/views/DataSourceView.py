import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.DataSourceDelegate import DataSourceDelegate

 #======================================================================
# 
# Encapsulates data for View DataSource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSourceView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DataSource index.")

def get(request, dataSourceId ):
	delegate = DataSourceDelegate()
	responseData = delegate.get( dataSourceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dataSource = json.loads(request.body)
	delegate = DataSourceDelegate()
	responseData = delegate.createFromJson( dataSource )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dataSource = json.loads(request.body)
	delegate = DataSourceDelegate()
	responseData = delegate.save( dataSource )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dataSourceId ):
	delegate = DataSourceDelegate()
	responseData = delegate.delete( dataSourceId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DataSourceDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkspace( request, dataSourceId, WorkspaceId ):
	delegate = DataSourceDelegate()
	responseData = delegate.saveWorkspace( dataSourceId, WorkspaceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkspace( request, dataSourceId ):
	delegate = DataSourceDelegate()
	responseData = delegate.deleteWorkspace( dataSourceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProducedDatasets( request, dataSourceId, ProducedDatasetsIds ):
	delegate = DataSourceDelegate()
	responseData = delegate.addProducedDatasets( dataSourceId, ProducedDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProducedDatasets( request, dataSourceId, ProducedDatasetsIds ):
	delegate = DataSourceDelegate()
	responseData = delegate.removeProducedDatasets( dataSourceId, ProducedDatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPipelines( request, dataSourceId, PipelinesIds ):
	delegate = DataSourceDelegate()
	responseData = delegate.addPipelines( dataSourceId, PipelinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePipelines( request, dataSourceId, PipelinesIds ):
	delegate = DataSourceDelegate()
	responseData = delegate.removePipelines( dataSourceId, PipelinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

