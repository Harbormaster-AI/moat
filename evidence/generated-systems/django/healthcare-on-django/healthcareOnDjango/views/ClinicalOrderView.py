import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ClinicalOrderDelegate import ClinicalOrderDelegate

 #======================================================================
# 
# Encapsulates data for View ClinicalOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicalOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ClinicalOrder index.")

def get(request, clinicalOrderId ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.get( clinicalOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	clinicalOrder = json.loads(request.body)
	delegate = ClinicalOrderDelegate()
	responseData = delegate.createFromJson( clinicalOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	clinicalOrder = json.loads(request.body)
	delegate = ClinicalOrderDelegate()
	responseData = delegate.save( clinicalOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, clinicalOrderId ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.delete( clinicalOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, clinicalOrderId, PatientId ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.savePatient( clinicalOrderId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, clinicalOrderId ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.deletePatient( clinicalOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEncounter( request, clinicalOrderId, EncounterId ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.saveEncounter( clinicalOrderId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, clinicalOrderId ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.deleteEncounter( clinicalOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrderingClinician( request, clinicalOrderId, OrderingClinicianId ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.saveOrderingClinician( clinicalOrderId, OrderingClinicianId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrderingClinician( request, clinicalOrderId ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.deleteOrderingClinician( clinicalOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMedicationOrders( request, clinicalOrderId, MedicationOrdersIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.addMedicationOrders( clinicalOrderId, MedicationOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMedicationOrders( request, clinicalOrderId, MedicationOrdersIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.removeMedicationOrders( clinicalOrderId, MedicationOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLaboratoryOrders( request, clinicalOrderId, LaboratoryOrdersIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.addLaboratoryOrders( clinicalOrderId, LaboratoryOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLaboratoryOrders( request, clinicalOrderId, LaboratoryOrdersIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.removeLaboratoryOrders( clinicalOrderId, LaboratoryOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addImagingOrders( request, clinicalOrderId, ImagingOrdersIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.addImagingOrders( clinicalOrderId, ImagingOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeImagingOrders( request, clinicalOrderId, ImagingOrdersIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.removeImagingOrders( clinicalOrderId, ImagingOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcedureOrders( request, clinicalOrderId, ProcedureOrdersIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.addProcedureOrders( clinicalOrderId, ProcedureOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcedureOrders( request, clinicalOrderId, ProcedureOrdersIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.removeProcedureOrders( clinicalOrderId, ProcedureOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAuthorizations( request, clinicalOrderId, AuthorizationsIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.addAuthorizations( clinicalOrderId, AuthorizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAuthorizations( request, clinicalOrderId, AuthorizationsIds ):
	delegate = ClinicalOrderDelegate()
	responseData = delegate.removeAuthorizations( clinicalOrderId, AuthorizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

