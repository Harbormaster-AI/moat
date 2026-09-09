import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.CoverageDefinitionDelegate import CoverageDefinitionDelegate

 #======================================================================
# 
# Encapsulates data for View CoverageDefinition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageDefinitionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CoverageDefinition index.")

def get(request, coverageDefinitionId ):
	delegate = CoverageDefinitionDelegate()
	responseData = delegate.get( coverageDefinitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	coverageDefinition = json.loads(request.body)
	delegate = CoverageDefinitionDelegate()
	responseData = delegate.createFromJson( coverageDefinition )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	coverageDefinition = json.loads(request.body)
	delegate = CoverageDefinitionDelegate()
	responseData = delegate.save( coverageDefinition )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, coverageDefinitionId ):
	delegate = CoverageDefinitionDelegate()
	responseData = delegate.delete( coverageDefinitionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CoverageDefinitionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, coverageDefinitionId, ProductId ):
	delegate = CoverageDefinitionDelegate()
	responseData = delegate.saveProduct( coverageDefinitionId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, coverageDefinitionId ):
	delegate = CoverageDefinitionDelegate()
	responseData = delegate.deleteProduct( coverageDefinitionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

