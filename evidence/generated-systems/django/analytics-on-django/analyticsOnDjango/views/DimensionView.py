import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

 #======================================================================
# 
# Encapsulates data for View Dimension
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DimensionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Dimension index.")

def get(request, dimensionId ):
	delegate = DimensionDelegate()
	responseData = delegate.get( dimensionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dimension = json.loads(request.body)
	delegate = DimensionDelegate()
	responseData = delegate.createFromJson( dimension )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dimension = json.loads(request.body)
	delegate = DimensionDelegate()
	responseData = delegate.save( dimension )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dimensionId ):
	delegate = DimensionDelegate()
	responseData = delegate.delete( dimensionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DimensionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSemanticModel( request, dimensionId, SemanticModelId ):
	delegate = DimensionDelegate()
	responseData = delegate.saveSemanticModel( dimensionId, SemanticModelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSemanticModel( request, dimensionId ):
	delegate = DimensionDelegate()
	responseData = delegate.deleteSemanticModel( dimensionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, dimensionId, DatasetsIds ):
	delegate = DimensionDelegate()
	responseData = delegate.addDatasets( dimensionId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, dimensionId, DatasetsIds ):
	delegate = DimensionDelegate()
	responseData = delegate.removeDatasets( dimensionId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGlossaryTerms( request, dimensionId, GlossaryTermsIds ):
	delegate = DimensionDelegate()
	responseData = delegate.addGlossaryTerms( dimensionId, GlossaryTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGlossaryTerms( request, dimensionId, GlossaryTermsIds ):
	delegate = DimensionDelegate()
	responseData = delegate.removeGlossaryTerms( dimensionId, GlossaryTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

