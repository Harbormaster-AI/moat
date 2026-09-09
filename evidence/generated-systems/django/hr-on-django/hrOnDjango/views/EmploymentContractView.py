import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.EmploymentContractDelegate import EmploymentContractDelegate

 #======================================================================
# 
# Encapsulates data for View EmploymentContract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentContractView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the EmploymentContract index.")

def get(request, employmentContractId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.get( employmentContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	employmentContract = json.loads(request.body)
	delegate = EmploymentContractDelegate()
	responseData = delegate.createFromJson( employmentContract )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	employmentContract = json.loads(request.body)
	delegate = EmploymentContractDelegate()
	responseData = delegate.save( employmentContract )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, employmentContractId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.delete( employmentContractId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EmploymentContractDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, employmentContractId, EmployeeId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.saveEmployee( employmentContractId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, employmentContractId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.deleteEmployee( employmentContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCompensationPackage( request, employmentContractId, CompensationPackageId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.saveCompensationPackage( employmentContractId, CompensationPackageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCompensationPackage( request, employmentContractId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.deleteCompensationPackage( employmentContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkSchedule( request, employmentContractId, WorkScheduleId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.saveWorkSchedule( employmentContractId, WorkScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkSchedule( request, employmentContractId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.deleteWorkSchedule( employmentContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLocation( request, employmentContractId, LocationId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.saveLocation( employmentContractId, LocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLocation( request, employmentContractId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.deleteLocation( employmentContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPayrollCalendar( request, employmentContractId, PayrollCalendarId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.savePayrollCalendar( employmentContractId, PayrollCalendarId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPayrollCalendar( request, employmentContractId ):
	delegate = EmploymentContractDelegate()
	responseData = delegate.deletePayrollCalendar( employmentContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

