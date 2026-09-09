import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.TimesheetDelegate import TimesheetDelegate

 #======================================================================
# 
# Encapsulates data for View Timesheet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimesheetView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Timesheet index.")

def get(request, timesheetId ):
	delegate = TimesheetDelegate()
	responseData = delegate.get( timesheetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	timesheet = json.loads(request.body)
	delegate = TimesheetDelegate()
	responseData = delegate.createFromJson( timesheet )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	timesheet = json.loads(request.body)
	delegate = TimesheetDelegate()
	responseData = delegate.save( timesheet )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, timesheetId ):
	delegate = TimesheetDelegate()
	responseData = delegate.delete( timesheetId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TimesheetDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, timesheetId, EmployeeId ):
	delegate = TimesheetDelegate()
	responseData = delegate.saveEmployee( timesheetId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, timesheetId ):
	delegate = TimesheetDelegate()
	responseData = delegate.deleteEmployee( timesheetId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTimeEntries( request, timesheetId, TimeEntriesIds ):
	delegate = TimesheetDelegate()
	responseData = delegate.addTimeEntries( timesheetId, TimeEntriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTimeEntries( request, timesheetId, TimeEntriesIds ):
	delegate = TimesheetDelegate()
	responseData = delegate.removeTimeEntries( timesheetId, TimeEntriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addApprovals( request, timesheetId, ApprovalsIds ):
	delegate = TimesheetDelegate()
	responseData = delegate.addApprovals( timesheetId, ApprovalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeApprovals( request, timesheetId, ApprovalsIds ):
	delegate = TimesheetDelegate()
	responseData = delegate.removeApprovals( timesheetId, ApprovalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

