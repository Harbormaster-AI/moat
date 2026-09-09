import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.QualityRuleDelegate import QualityRuleDelegate

 #======================================================================
# 
# Encapsulates data for View QualityRule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityRuleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the QualityRule index.")

def get(request, qualityRuleId ):
	delegate = QualityRuleDelegate()
	responseData = delegate.get( qualityRuleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	qualityRule = json.loads(request.body)
	delegate = QualityRuleDelegate()
	responseData = delegate.createFromJson( qualityRule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	qualityRule = json.loads(request.body)
	delegate = QualityRuleDelegate()
	responseData = delegate.save( qualityRule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, qualityRuleId ):
	delegate = QualityRuleDelegate()
	responseData = delegate.delete( qualityRuleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = QualityRuleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDataset( request, qualityRuleId, DatasetId ):
	delegate = QualityRuleDelegate()
	responseData = delegate.saveDataset( qualityRuleId, DatasetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDataset( request, qualityRuleId ):
	delegate = QualityRuleDelegate()
	responseData = delegate.deleteDataset( qualityRuleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChecks( request, qualityRuleId, ChecksIds ):
	delegate = QualityRuleDelegate()
	responseData = delegate.addChecks( qualityRuleId, ChecksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChecks( request, qualityRuleId, ChecksIds ):
	delegate = QualityRuleDelegate()
	responseData = delegate.removeChecks( qualityRuleId, ChecksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

