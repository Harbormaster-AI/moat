import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

 #======================================================================
# 
# Encapsulates data for View Encounter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EncounterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Encounter index.")

def get(request, encounterId ):
	delegate = EncounterDelegate()
	responseData = delegate.get( encounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	encounter = json.loads(request.body)
	delegate = EncounterDelegate()
	responseData = delegate.createFromJson( encounter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	encounter = json.loads(request.body)
	delegate = EncounterDelegate()
	responseData = delegate.save( encounter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, encounterId ):
	delegate = EncounterDelegate()
	responseData = delegate.delete( encounterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EncounterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, encounterId, PatientId ):
	delegate = EncounterDelegate()
	responseData = delegate.savePatient( encounterId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, encounterId ):
	delegate = EncounterDelegate()
	responseData = delegate.deletePatient( encounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClinician( request, encounterId, ClinicianId ):
	delegate = EncounterDelegate()
	responseData = delegate.saveClinician( encounterId, ClinicianId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClinician( request, encounterId ):
	delegate = EncounterDelegate()
	responseData = delegate.deleteClinician( encounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFacility( request, encounterId, FacilityId ):
	delegate = EncounterDelegate()
	responseData = delegate.saveFacility( encounterId, FacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFacility( request, encounterId ):
	delegate = EncounterDelegate()
	responseData = delegate.deleteFacility( encounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAppointment( request, encounterId, AppointmentId ):
	delegate = EncounterDelegate()
	responseData = delegate.saveAppointment( encounterId, AppointmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAppointment( request, encounterId ):
	delegate = EncounterDelegate()
	responseData = delegate.deleteAppointment( encounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdmission( request, encounterId, AdmissionId ):
	delegate = EncounterDelegate()
	responseData = delegate.saveAdmission( encounterId, AdmissionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdmission( request, encounterId ):
	delegate = EncounterDelegate()
	responseData = delegate.deleteAdmission( encounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDischarge( request, encounterId, DischargeId ):
	delegate = EncounterDelegate()
	responseData = delegate.saveDischarge( encounterId, DischargeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDischarge( request, encounterId ):
	delegate = EncounterDelegate()
	responseData = delegate.deleteDischarge( encounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDiagnoses( request, encounterId, DiagnosesIds ):
	delegate = EncounterDelegate()
	responseData = delegate.addDiagnoses( encounterId, DiagnosesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDiagnoses( request, encounterId, DiagnosesIds ):
	delegate = EncounterDelegate()
	responseData = delegate.removeDiagnoses( encounterId, DiagnosesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcedures( request, encounterId, ProceduresIds ):
	delegate = EncounterDelegate()
	responseData = delegate.addProcedures( encounterId, ProceduresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcedures( request, encounterId, ProceduresIds ):
	delegate = EncounterDelegate()
	responseData = delegate.removeProcedures( encounterId, ProceduresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addObservations( request, encounterId, ObservationsIds ):
	delegate = EncounterDelegate()
	responseData = delegate.addObservations( encounterId, ObservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeObservations( request, encounterId, ObservationsIds ):
	delegate = EncounterDelegate()
	responseData = delegate.removeObservations( encounterId, ObservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, encounterId, OrdersIds ):
	delegate = EncounterDelegate()
	responseData = delegate.addOrders( encounterId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, encounterId, OrdersIds ):
	delegate = EncounterDelegate()
	responseData = delegate.removeOrders( encounterId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

