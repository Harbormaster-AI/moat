import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

 #======================================================================
# 
# Encapsulates data for View Patient
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PatientView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Patient index.")

def get(request, patientId ):
	delegate = PatientDelegate()
	responseData = delegate.get( patientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	patient = json.loads(request.body)
	delegate = PatientDelegate()
	responseData = delegate.createFromJson( patient )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	patient = json.loads(request.body)
	delegate = PatientDelegate()
	responseData = delegate.save( patient )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, patientId ):
	delegate = PatientDelegate()
	responseData = delegate.delete( patientId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PatientDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAppointments( request, patientId, AppointmentsIds ):
	delegate = PatientDelegate()
	responseData = delegate.addAppointments( patientId, AppointmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAppointments( request, patientId, AppointmentsIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeAppointments( patientId, AppointmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEncounters( request, patientId, EncountersIds ):
	delegate = PatientDelegate()
	responseData = delegate.addEncounters( patientId, EncountersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEncounters( request, patientId, EncountersIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeEncounters( patientId, EncountersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCarePlans( request, patientId, CarePlansIds ):
	delegate = PatientDelegate()
	responseData = delegate.addCarePlans( patientId, CarePlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCarePlans( request, patientId, CarePlansIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeCarePlans( patientId, CarePlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAllergies( request, patientId, AllergiesIds ):
	delegate = PatientDelegate()
	responseData = delegate.addAllergies( patientId, AllergiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAllergies( request, patientId, AllergiesIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeAllergies( patientId, AllergiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addConditions( request, patientId, ConditionsIds ):
	delegate = PatientDelegate()
	responseData = delegate.addConditions( patientId, ConditionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeConditions( request, patientId, ConditionsIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeConditions( patientId, ConditionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMedicationOrders( request, patientId, MedicationOrdersIds ):
	delegate = PatientDelegate()
	responseData = delegate.addMedicationOrders( patientId, MedicationOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMedicationOrders( request, patientId, MedicationOrdersIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeMedicationOrders( patientId, MedicationOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLabOrders( request, patientId, LabOrdersIds ):
	delegate = PatientDelegate()
	responseData = delegate.addLabOrders( patientId, LabOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLabOrders( request, patientId, LabOrdersIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeLabOrders( patientId, LabOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addImagingOrders( request, patientId, ImagingOrdersIds ):
	delegate = PatientDelegate()
	responseData = delegate.addImagingOrders( patientId, ImagingOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeImagingOrders( request, patientId, ImagingOrdersIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeImagingOrders( patientId, ImagingOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCoverages( request, patientId, CoveragesIds ):
	delegate = PatientDelegate()
	responseData = delegate.addCoverages( patientId, CoveragesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCoverages( request, patientId, CoveragesIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeCoverages( patientId, CoveragesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClaims( request, patientId, ClaimsIds ):
	delegate = PatientDelegate()
	responseData = delegate.addClaims( patientId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClaims( request, patientId, ClaimsIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeClaims( patientId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDevices( request, patientId, DevicesIds ):
	delegate = PatientDelegate()
	responseData = delegate.addDevices( patientId, DevicesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDevices( request, patientId, DevicesIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeDevices( patientId, DevicesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addObservations( request, patientId, ObservationsIds ):
	delegate = PatientDelegate()
	responseData = delegate.addObservations( patientId, ObservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeObservations( request, patientId, ObservationsIds ):
	delegate = PatientDelegate()
	responseData = delegate.removeObservations( patientId, ObservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

