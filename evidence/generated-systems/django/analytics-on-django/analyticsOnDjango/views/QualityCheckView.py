import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.QualityCheckDelegate import QualityCheckDelegate

 #======================================================================
# 
# Encapsulates data for View QualityCheck
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityCheckView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the QualityCheck index.")

def get(request, qualityCheckId ):
	delegate = QualityCheckDelegate()
	responseData = delegate.get( qualityCheckId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	qualityCheck = json.loads(request.body)
	delegate = QualityCheckDelegate()
	responseData = delegate.createFromJson( qualityCheck )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	qualityCheck = json.loads(request.body)
	delegate = QualityCheckDelegate()
	responseData = delegate.save( qualityCheck )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, qualityCheckId ):
	delegate = QualityCheckDelegate()
	responseData = delegate.delete( qualityCheckId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = QualityCheckDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRule( request, qualityCheckId, RuleId ):
	delegate = QualityCheckDelegate()
	responseData = delegate.saveRule( qualityCheckId, RuleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRule( request, qualityCheckId ):
	delegate = QualityCheckDelegate()
	responseData = delegate.deleteRule( qualityCheckId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDataset( request, qualityCheckId, DatasetId ):
	delegate = QualityCheckDelegate()
	responseData = delegate.saveDataset( qualityCheckId, DatasetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDataset( request, qualityCheckId ):
	delegate = QualityCheckDelegate()
	responseData = delegate.deleteDataset( qualityCheckId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

