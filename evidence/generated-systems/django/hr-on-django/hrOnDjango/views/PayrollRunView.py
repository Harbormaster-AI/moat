import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.PayrollRunDelegate import PayrollRunDelegate

 #======================================================================
# 
# Encapsulates data for View PayrollRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollRunView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PayrollRun index.")

def get(request, payrollRunId ):
	delegate = PayrollRunDelegate()
	responseData = delegate.get( payrollRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	payrollRun = json.loads(request.body)
	delegate = PayrollRunDelegate()
	responseData = delegate.createFromJson( payrollRun )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	payrollRun = json.loads(request.body)
	delegate = PayrollRunDelegate()
	responseData = delegate.save( payrollRun )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, payrollRunId ):
	delegate = PayrollRunDelegate()
	responseData = delegate.delete( payrollRunId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PayrollRunDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPayrollCalendar( request, payrollRunId, PayrollCalendarId ):
	delegate = PayrollRunDelegate()
	responseData = delegate.savePayrollCalendar( payrollRunId, PayrollCalendarId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPayrollCalendar( request, payrollRunId ):
	delegate = PayrollRunDelegate()
	responseData = delegate.deletePayrollCalendar( payrollRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayrollItems( request, payrollRunId, PayrollItemsIds ):
	delegate = PayrollRunDelegate()
	responseData = delegate.addPayrollItems( payrollRunId, PayrollItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayrollItems( request, payrollRunId, PayrollItemsIds ):
	delegate = PayrollRunDelegate()
	responseData = delegate.removePayrollItems( payrollRunId, PayrollItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

