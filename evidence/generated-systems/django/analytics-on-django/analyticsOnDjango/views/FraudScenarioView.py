import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from analyticsOnDjango.delegates.FraudScenarioDelegate import FraudScenarioDelegate

 #======================================================================
# 
# Encapsulates data for View FraudScenario
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudScenarioView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the FraudScenario index.")

def get(request, fraudScenarioId ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.get( fraudScenarioId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	fraudScenario = json.loads(request.body)
	delegate = FraudScenarioDelegate()
	responseData = delegate.createFromJson( fraudScenario )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	fraudScenario = json.loads(request.body)
	delegate = FraudScenarioDelegate()
	responseData = delegate.save( fraudScenario )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, fraudScenarioId ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.delete( fraudScenarioId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FraudScenarioDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addModels( request, fraudScenarioId, ModelsIds ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.addModels( fraudScenarioId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeModels( request, fraudScenarioId, ModelsIds ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.removeModels( fraudScenarioId, ModelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDatasets( request, fraudScenarioId, DatasetsIds ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.addDatasets( fraudScenarioId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDatasets( request, fraudScenarioId, DatasetsIds ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.removeDatasets( fraudScenarioId, DatasetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAlerts( request, fraudScenarioId, AlertsIds ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.addAlerts( fraudScenarioId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAlerts( request, fraudScenarioId, AlertsIds ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.removeAlerts( fraudScenarioId, AlertsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSignals( request, fraudScenarioId, SignalsIds ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.addSignals( fraudScenarioId, SignalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSignals( request, fraudScenarioId, SignalsIds ):
	delegate = FraudScenarioDelegate()
	responseData = delegate.removeSignals( fraudScenarioId, SignalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

