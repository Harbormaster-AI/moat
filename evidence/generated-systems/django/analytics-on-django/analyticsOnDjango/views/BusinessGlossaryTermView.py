import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

 #======================================================================
# 
# Encapsulates data for View BusinessGlossaryTerm
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessGlossaryTermView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BusinessGlossaryTerm index.")

def get(request, businessGlossaryTermId ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.get( businessGlossaryTermId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	businessGlossaryTerm = json.loads(request.body)
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.createFromJson( businessGlossaryTerm )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	businessGlossaryTerm = json.loads(request.body)
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.save( businessGlossaryTerm )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, businessGlossaryTermId ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.delete( businessGlossaryTermId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRelatedTerms( request, businessGlossaryTermId, RelatedTermsIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.addRelatedTerms( businessGlossaryTermId, RelatedTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRelatedTerms( request, businessGlossaryTermId, RelatedTermsIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.removeRelatedTerms( businessGlossaryTermId, RelatedTermsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMetrics( request, businessGlossaryTermId, MetricsIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.addMetrics( businessGlossaryTermId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMetrics( request, businessGlossaryTermId, MetricsIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.removeMetrics( businessGlossaryTermId, MetricsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, businessGlossaryTermId, DatasetsIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.addDatasets( businessGlossaryTermId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, businessGlossaryTermId, DatasetsIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.removeDatasets( businessGlossaryTermId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDimensions( request, businessGlossaryTermId, DimensionsIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.addDimensions( businessGlossaryTermId, DimensionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDimensions( request, businessGlossaryTermId, DimensionsIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.removeDimensions( businessGlossaryTermId, DimensionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMeasures( request, businessGlossaryTermId, MeasuresIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.addMeasures( businessGlossaryTermId, MeasuresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMeasures( request, businessGlossaryTermId, MeasuresIds ):
	delegate = BusinessGlossaryTermDelegate()
	responseData = delegate.removeMeasures( businessGlossaryTermId, MeasuresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

