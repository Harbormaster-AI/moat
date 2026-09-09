import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.InsuranceProductDelegate import InsuranceProductDelegate

 #======================================================================
# 
# Encapsulates data for View InsuranceProduct
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsuranceProductView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InsuranceProduct index.")

def get(request, insuranceProductId ):
	delegate = InsuranceProductDelegate()
	responseData = delegate.get( insuranceProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	insuranceProduct = json.loads(request.body)
	delegate = InsuranceProductDelegate()
	responseData = delegate.createFromJson( insuranceProduct )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	insuranceProduct = json.loads(request.body)
	delegate = InsuranceProductDelegate()
	responseData = delegate.save( insuranceProduct )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, insuranceProductId ):
	delegate = InsuranceProductDelegate()
	responseData = delegate.delete( insuranceProductId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InsuranceProductDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInsurer( request, insuranceProductId, InsurerId ):
	delegate = InsuranceProductDelegate()
	responseData = delegate.saveInsurer( insuranceProductId, InsurerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInsurer( request, insuranceProductId ):
	delegate = InsuranceProductDelegate()
	responseData = delegate.deleteInsurer( insuranceProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCoverageDefinitions( request, insuranceProductId, CoverageDefinitionsIds ):
	delegate = InsuranceProductDelegate()
	responseData = delegate.addCoverageDefinitions( insuranceProductId, CoverageDefinitionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCoverageDefinitions( request, insuranceProductId, CoverageDefinitionsIds ):
	delegate = InsuranceProductDelegate()
	responseData = delegate.removeCoverageDefinitions( insuranceProductId, CoverageDefinitionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

