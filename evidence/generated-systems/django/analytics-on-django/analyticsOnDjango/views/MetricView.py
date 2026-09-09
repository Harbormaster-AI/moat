import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

 #======================================================================
# 
# Encapsulates data for View Metric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MetricView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Metric index.")

def get(request, metricId ):
	delegate = MetricDelegate()
	responseData = delegate.get( metricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	metric = json.loads(request.body)
	delegate = MetricDelegate()
	responseData = delegate.createFromJson( metric )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	metric = json.loads(request.body)
	delegate = MetricDelegate()
	responseData = delegate.save( metric )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, metricId ):
	delegate = MetricDelegate()
	responseData = delegate.delete( metricId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MetricDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSemanticModel( request, metricId, SemanticModelId ):
	delegate = MetricDelegate()
	responseData = delegate.saveSemanticModel( metricId, SemanticModelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSemanticModel( request, metricId ):
	delegate = MetricDelegate()
	responseData = delegate.deleteSemanticModel( metricId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, metricId, DatasetsIds ):
	delegate = MetricDelegate()
	responseData = delegate.addDatasets( metricId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, metricId, DatasetsIds ):
	delegate = MetricDelegate()
	responseData = delegate.removeDatasets( metricId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGlossaryTerms( request, metricId, GlossaryTermsIds ):
	delegate = MetricDelegate()
	responseData = delegate.addGlossaryTerms( metricId, GlossaryTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGlossaryTerms( request, metricId, GlossaryTermsIds ):
	delegate = MetricDelegate()
	responseData = delegate.removeGlossaryTerms( metricId, GlossaryTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAlerts( request, metricId, AlertsIds ):
	delegate = MetricDelegate()
	responseData = delegate.addAlerts( metricId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAlerts( request, metricId, AlertsIds ):
	delegate = MetricDelegate()
	responseData = delegate.removeAlerts( metricId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVisualizations( request, metricId, VisualizationsIds ):
	delegate = MetricDelegate()
	responseData = delegate.addVisualizations( metricId, VisualizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVisualizations( request, metricId, VisualizationsIds ):
	delegate = MetricDelegate()
	responseData = delegate.removeVisualizations( metricId, VisualizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

