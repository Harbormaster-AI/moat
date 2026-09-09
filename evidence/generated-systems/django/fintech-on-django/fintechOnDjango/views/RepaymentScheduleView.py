import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.RepaymentScheduleDelegate import RepaymentScheduleDelegate

 #======================================================================
# 
# Encapsulates data for View RepaymentSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RepaymentScheduleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the RepaymentSchedule index.")

def get(request, repaymentScheduleId ):
	delegate = RepaymentScheduleDelegate()
	responseData = delegate.get( repaymentScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	repaymentSchedule = json.loads(request.body)
	delegate = RepaymentScheduleDelegate()
	responseData = delegate.createFromJson( repaymentSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	repaymentSchedule = json.loads(request.body)
	delegate = RepaymentScheduleDelegate()
	responseData = delegate.save( repaymentSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, repaymentScheduleId ):
	delegate = RepaymentScheduleDelegate()
	responseData = delegate.delete( repaymentScheduleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RepaymentScheduleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLoan( request, repaymentScheduleId, LoanId ):
	delegate = RepaymentScheduleDelegate()
	responseData = delegate.saveLoan( repaymentScheduleId, LoanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLoan( request, repaymentScheduleId ):
	delegate = RepaymentScheduleDelegate()
	responseData = delegate.deleteLoan( repaymentScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayments( request, repaymentScheduleId, PaymentsIds ):
	delegate = RepaymentScheduleDelegate()
	responseData = delegate.addPayments( repaymentScheduleId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayments( request, repaymentScheduleId, PaymentsIds ):
	delegate = RepaymentScheduleDelegate()
	responseData = delegate.removePayments( repaymentScheduleId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

