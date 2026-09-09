import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.LaboratoryDelegate import LaboratoryDelegate

 #======================================================================
# 
# Encapsulates data for View Laboratory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LaboratoryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Laboratory index.")

def get(request, laboratoryId ):
	delegate = LaboratoryDelegate()
	responseData = delegate.get( laboratoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	laboratory = json.loads(request.body)
	delegate = LaboratoryDelegate()
	responseData = delegate.createFromJson( laboratory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	laboratory = json.loads(request.body)
	delegate = LaboratoryDelegate()
	responseData = delegate.save( laboratory )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, laboratoryId ):
	delegate = LaboratoryDelegate()
	responseData = delegate.delete( laboratoryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LaboratoryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFacility( request, laboratoryId, FacilityId ):
	delegate = LaboratoryDelegate()
	responseData = delegate.saveFacility( laboratoryId, FacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFacility( request, laboratoryId ):
	delegate = LaboratoryDelegate()
	responseData = delegate.deleteFacility( laboratoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLaboratoryOrders( request, laboratoryId, LaboratoryOrdersIds ):
	delegate = LaboratoryDelegate()
	responseData = delegate.addLaboratoryOrders( laboratoryId, LaboratoryOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLaboratoryOrders( request, laboratoryId, LaboratoryOrdersIds ):
	delegate = LaboratoryDelegate()
	responseData = delegate.removeLaboratoryOrders( laboratoryId, LaboratoryOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLabResults( request, laboratoryId, LabResultsIds ):
	delegate = LaboratoryDelegate()
	responseData = delegate.addLabResults( laboratoryId, LabResultsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLabResults( request, laboratoryId, LabResultsIds ):
	delegate = LaboratoryDelegate()
	responseData = delegate.removeLabResults( laboratoryId, LabResultsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

