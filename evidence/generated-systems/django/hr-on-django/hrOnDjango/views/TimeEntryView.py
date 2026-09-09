import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.TimeEntryDelegate import TimeEntryDelegate

 #======================================================================
# 
# Encapsulates data for View TimeEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeEntryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TimeEntry index.")

def get(request, timeEntryId ):
	delegate = TimeEntryDelegate()
	responseData = delegate.get( timeEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	timeEntry = json.loads(request.body)
	delegate = TimeEntryDelegate()
	responseData = delegate.createFromJson( timeEntry )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	timeEntry = json.loads(request.body)
	delegate = TimeEntryDelegate()
	responseData = delegate.save( timeEntry )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, timeEntryId ):
	delegate = TimeEntryDelegate()
	responseData = delegate.delete( timeEntryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TimeEntryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTimesheet( request, timeEntryId, TimesheetId ):
	delegate = TimeEntryDelegate()
	responseData = delegate.saveTimesheet( timeEntryId, TimesheetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTimesheet( request, timeEntryId ):
	delegate = TimeEntryDelegate()
	responseData = delegate.deleteTimesheet( timeEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, timeEntryId, EmployeeId ):
	delegate = TimeEntryDelegate()
	responseData = delegate.saveEmployee( timeEntryId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, timeEntryId ):
	delegate = TimeEntryDelegate()
	responseData = delegate.deleteEmployee( timeEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCostCenter( request, timeEntryId, CostCenterId ):
	delegate = TimeEntryDelegate()
	responseData = delegate.saveCostCenter( timeEntryId, CostCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCostCenter( request, timeEntryId ):
	delegate = TimeEntryDelegate()
	responseData = delegate.deleteCostCenter( timeEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

