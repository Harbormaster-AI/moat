import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.SemanticModelDelegate import SemanticModelDelegate

 #======================================================================
# 
# Encapsulates data for View SemanticModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SemanticModelView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SemanticModel index.")

def get(request, semanticModelId ):
	delegate = SemanticModelDelegate()
	responseData = delegate.get( semanticModelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	semanticModel = json.loads(request.body)
	delegate = SemanticModelDelegate()
	responseData = delegate.createFromJson( semanticModel )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	semanticModel = json.loads(request.body)
	delegate = SemanticModelDelegate()
	responseData = delegate.save( semanticModel )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, semanticModelId ):
	delegate = SemanticModelDelegate()
	responseData = delegate.delete( semanticModelId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SemanticModelDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, semanticModelId, DatasetsIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.addDatasets( semanticModelId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, semanticModelId, DatasetsIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.removeDatasets( semanticModelId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMetrics( request, semanticModelId, MetricsIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.addMetrics( semanticModelId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMetrics( request, semanticModelId, MetricsIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.removeMetrics( semanticModelId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDimensions( request, semanticModelId, DimensionsIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.addDimensions( semanticModelId, DimensionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDimensions( request, semanticModelId, DimensionsIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.removeDimensions( semanticModelId, DimensionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMeasures( request, semanticModelId, MeasuresIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.addMeasures( semanticModelId, MeasuresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMeasures( request, semanticModelId, MeasuresIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.removeMeasures( semanticModelId, MeasuresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGlossaryTerms( request, semanticModelId, GlossaryTermsIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.addGlossaryTerms( semanticModelId, GlossaryTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGlossaryTerms( request, semanticModelId, GlossaryTermsIds ):
	delegate = SemanticModelDelegate()
	responseData = delegate.removeGlossaryTerms( semanticModelId, GlossaryTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

