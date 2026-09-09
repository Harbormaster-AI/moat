import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.PayrollCalendarDelegate import PayrollCalendarDelegate

 #======================================================================
# 
# Encapsulates data for View PayrollCalendar
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollCalendarView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PayrollCalendar index.")

def get(request, payrollCalendarId ):
	delegate = PayrollCalendarDelegate()
	responseData = delegate.get( payrollCalendarId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	payrollCalendar = json.loads(request.body)
	delegate = PayrollCalendarDelegate()
	responseData = delegate.createFromJson( payrollCalendar )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	payrollCalendar = json.loads(request.body)
	delegate = PayrollCalendarDelegate()
	responseData = delegate.save( payrollCalendar )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, payrollCalendarId ):
	delegate = PayrollCalendarDelegate()
	responseData = delegate.delete( payrollCalendarId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PayrollCalendarDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, payrollCalendarId, OrganizationId ):
	delegate = PayrollCalendarDelegate()
	responseData = delegate.saveOrganization( payrollCalendarId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, payrollCalendarId ):
	delegate = PayrollCalendarDelegate()
	responseData = delegate.deleteOrganization( payrollCalendarId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayrollRuns( request, payrollCalendarId, PayrollRunsIds ):
	delegate = PayrollCalendarDelegate()
	responseData = delegate.addPayrollRuns( payrollCalendarId, PayrollRunsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayrollRuns( request, payrollCalendarId, PayrollRunsIds ):
	delegate = PayrollCalendarDelegate()
	responseData = delegate.removePayrollRuns( payrollCalendarId, PayrollRunsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmployees( request, payrollCalendarId, EmployeesIds ):
	delegate = PayrollCalendarDelegate()
	responseData = delegate.addEmployees( payrollCalendarId, EmployeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmployees( request, payrollCalendarId, EmployeesIds ):
	delegate = PayrollCalendarDelegate()
	responseData = delegate.removeEmployees( payrollCalendarId, EmployeesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

