import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.MedicationOrderDelegate import MedicationOrderDelegate

 #======================================================================
# 
# Encapsulates data for View MedicationOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicationOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MedicationOrder index.")

def get(request, medicationOrderId ):
	delegate = MedicationOrderDelegate()
	responseData = delegate.get( medicationOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	medicationOrder = json.loads(request.body)
	delegate = MedicationOrderDelegate()
	responseData = delegate.createFromJson( medicationOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	medicationOrder = json.loads(request.body)
	delegate = MedicationOrderDelegate()
	responseData = delegate.save( medicationOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, medicationOrderId ):
	delegate = MedicationOrderDelegate()
	responseData = delegate.delete( medicationOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MedicationOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, medicationOrderId, OrderId ):
	delegate = MedicationOrderDelegate()
	responseData = delegate.saveOrder( medicationOrderId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, medicationOrderId ):
	delegate = MedicationOrderDelegate()
	responseData = delegate.deleteOrder( medicationOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPharmacy( request, medicationOrderId, PharmacyId ):
	delegate = MedicationOrderDelegate()
	responseData = delegate.savePharmacy( medicationOrderId, PharmacyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPharmacy( request, medicationOrderId ):
	delegate = MedicationOrderDelegate()
	responseData = delegate.deletePharmacy( medicationOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDispenses( request, medicationOrderId, DispensesIds ):
	delegate = MedicationOrderDelegate()
	responseData = delegate.addDispenses( medicationOrderId, DispensesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDispenses( request, medicationOrderId, DispensesIds ):
	delegate = MedicationOrderDelegate()
	responseData = delegate.removeDispenses( medicationOrderId, DispensesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

