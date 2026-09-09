import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.LaboratoryOrderDelegate import LaboratoryOrderDelegate

 #======================================================================
# 
# Encapsulates data for View LaboratoryOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LaboratoryOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LaboratoryOrder index.")

def get(request, laboratoryOrderId ):
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.get( laboratoryOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	laboratoryOrder = json.loads(request.body)
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.createFromJson( laboratoryOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	laboratoryOrder = json.loads(request.body)
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.save( laboratoryOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, laboratoryOrderId ):
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.delete( laboratoryOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, laboratoryOrderId, OrderId ):
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.saveOrder( laboratoryOrderId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, laboratoryOrderId ):
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.deleteOrder( laboratoryOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLaboratory( request, laboratoryOrderId, LaboratoryId ):
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.saveLaboratory( laboratoryOrderId, LaboratoryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLaboratory( request, laboratoryOrderId ):
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.deleteLaboratory( laboratoryOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addResults( request, laboratoryOrderId, ResultsIds ):
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.addResults( laboratoryOrderId, ResultsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeResults( request, laboratoryOrderId, ResultsIds ):
	delegate = LaboratoryOrderDelegate()
	responseData = delegate.removeResults( laboratoryOrderId, ResultsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

