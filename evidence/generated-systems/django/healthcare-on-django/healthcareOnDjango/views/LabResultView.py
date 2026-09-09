import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.LabResultDelegate import LabResultDelegate

 #======================================================================
# 
# Encapsulates data for View LabResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LabResultView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LabResult index.")

def get(request, labResultId ):
	delegate = LabResultDelegate()
	responseData = delegate.get( labResultId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	labResult = json.loads(request.body)
	delegate = LabResultDelegate()
	responseData = delegate.createFromJson( labResult )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	labResult = json.loads(request.body)
	delegate = LabResultDelegate()
	responseData = delegate.save( labResult )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, labResultId ):
	delegate = LabResultDelegate()
	responseData = delegate.delete( labResultId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LabResultDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLaboratoryOrder( request, labResultId, LaboratoryOrderId ):
	delegate = LabResultDelegate()
	responseData = delegate.saveLaboratoryOrder( labResultId, LaboratoryOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLaboratoryOrder( request, labResultId ):
	delegate = LabResultDelegate()
	responseData = delegate.deleteLaboratoryOrder( labResultId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLaboratory( request, labResultId, LaboratoryId ):
	delegate = LabResultDelegate()
	responseData = delegate.saveLaboratory( labResultId, LaboratoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLaboratory( request, labResultId ):
	delegate = LabResultDelegate()
	responseData = delegate.deleteLaboratory( labResultId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addObservations( request, labResultId, ObservationsIds ):
	delegate = LabResultDelegate()
	responseData = delegate.addObservations( labResultId, ObservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeObservations( request, labResultId, ObservationsIds ):
	delegate = LabResultDelegate()
	responseData = delegate.removeObservations( labResultId, ObservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

