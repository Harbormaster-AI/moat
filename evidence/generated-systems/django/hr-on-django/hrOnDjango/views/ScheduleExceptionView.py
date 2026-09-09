import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.ScheduleExceptionDelegate import ScheduleExceptionDelegate

 #======================================================================
# 
# Encapsulates data for View ScheduleException
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScheduleExceptionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ScheduleException index.")

def get(request, scheduleExceptionId ):
	delegate = ScheduleExceptionDelegate()
	responseData = delegate.get( scheduleExceptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	scheduleException = json.loads(request.body)
	delegate = ScheduleExceptionDelegate()
	responseData = delegate.createFromJson( scheduleException )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	scheduleException = json.loads(request.body)
	delegate = ScheduleExceptionDelegate()
	responseData = delegate.save( scheduleException )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, scheduleExceptionId ):
	delegate = ScheduleExceptionDelegate()
	responseData = delegate.delete( scheduleExceptionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ScheduleExceptionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkSchedule( request, scheduleExceptionId, WorkScheduleId ):
	delegate = ScheduleExceptionDelegate()
	responseData = delegate.saveWorkSchedule( scheduleExceptionId, WorkScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkSchedule( request, scheduleExceptionId ):
	delegate = ScheduleExceptionDelegate()
	responseData = delegate.deleteWorkSchedule( scheduleExceptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, scheduleExceptionId, EmployeeId ):
	delegate = ScheduleExceptionDelegate()
	responseData = delegate.saveEmployee( scheduleExceptionId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, scheduleExceptionId ):
	delegate = ScheduleExceptionDelegate()
	responseData = delegate.deleteEmployee( scheduleExceptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

