import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.PayrollItemDelegate import PayrollItemDelegate

 #======================================================================
# 
# Encapsulates data for View PayrollItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PayrollItem index.")

def get(request, payrollItemId ):
	delegate = PayrollItemDelegate()
	responseData = delegate.get( payrollItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	payrollItem = json.loads(request.body)
	delegate = PayrollItemDelegate()
	responseData = delegate.createFromJson( payrollItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	payrollItem = json.loads(request.body)
	delegate = PayrollItemDelegate()
	responseData = delegate.save( payrollItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, payrollItemId ):
	delegate = PayrollItemDelegate()
	responseData = delegate.delete( payrollItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PayrollItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPayrollRun( request, payrollItemId, PayrollRunId ):
	delegate = PayrollItemDelegate()
	responseData = delegate.savePayrollRun( payrollItemId, PayrollRunId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPayrollRun( request, payrollItemId ):
	delegate = PayrollItemDelegate()
	responseData = delegate.deletePayrollRun( payrollItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, payrollItemId, EmployeeId ):
	delegate = PayrollItemDelegate()
	responseData = delegate.saveEmployee( payrollItemId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, payrollItemId ):
	delegate = PayrollItemDelegate()
	responseData = delegate.deleteEmployee( payrollItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

