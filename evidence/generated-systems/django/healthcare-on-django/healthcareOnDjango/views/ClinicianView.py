import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

 #======================================================================
# 
# Encapsulates data for View Clinician
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicianView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Clinician index.")

def get(request, clinicianId ):
	delegate = ClinicianDelegate()
	responseData = delegate.get( clinicianId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	clinician = json.loads(request.body)
	delegate = ClinicianDelegate()
	responseData = delegate.createFromJson( clinician )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	clinician = json.loads(request.body)
	delegate = ClinicianDelegate()
	responseData = delegate.save( clinician )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, clinicianId ):
	delegate = ClinicianDelegate()
	responseData = delegate.delete( clinicianId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ClinicianDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCareTeams( request, clinicianId, CareTeamsIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.addCareTeams( clinicianId, CareTeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCareTeams( request, clinicianId, CareTeamsIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.removeCareTeams( clinicianId, CareTeamsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAppointments( request, clinicianId, AppointmentsIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.addAppointments( clinicianId, AppointmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAppointments( request, clinicianId, AppointmentsIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.removeAppointments( clinicianId, AppointmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEncounters( request, clinicianId, EncountersIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.addEncounters( clinicianId, EncountersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEncounters( request, clinicianId, EncountersIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.removeEncounters( clinicianId, EncountersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcedures( request, clinicianId, ProceduresIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.addProcedures( clinicianId, ProceduresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcedures( request, clinicianId, ProceduresIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.removeProcedures( clinicianId, ProceduresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addImagingReports( request, clinicianId, ImagingReportsIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.addImagingReports( clinicianId, ImagingReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeImagingReports( request, clinicianId, ImagingReportsIds ):
	delegate = ClinicianDelegate()
	responseData = delegate.removeImagingReports( clinicianId, ImagingReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

