import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.MeasureDelegate import MeasureDelegate

 #======================================================================
# 
# Encapsulates data for View Measure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MeasureView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Measure index.")

def get(request, measureId ):
	delegate = MeasureDelegate()
	responseData = delegate.get( measureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	measure = json.loads(request.body)
	delegate = MeasureDelegate()
	responseData = delegate.createFromJson( measure )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	measure = json.loads(request.body)
	delegate = MeasureDelegate()
	responseData = delegate.save( measure )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, measureId ):
	delegate = MeasureDelegate()
	responseData = delegate.delete( measureId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MeasureDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSemanticModel( request, measureId, SemanticModelId ):
	delegate = MeasureDelegate()
	responseData = delegate.saveSemanticModel( measureId, SemanticModelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSemanticModel( request, measureId ):
	delegate = MeasureDelegate()
	responseData = delegate.deleteSemanticModel( measureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, measureId, DatasetsIds ):
	delegate = MeasureDelegate()
	responseData = delegate.addDatasets( measureId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, measureId, DatasetsIds ):
	delegate = MeasureDelegate()
	responseData = delegate.removeDatasets( measureId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGlossaryTerms( request, measureId, GlossaryTermsIds ):
	delegate = MeasureDelegate()
	responseData = delegate.addGlossaryTerms( measureId, GlossaryTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGlossaryTerms( request, measureId, GlossaryTermsIds ):
	delegate = MeasureDelegate()
	responseData = delegate.removeGlossaryTerms( measureId, GlossaryTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

