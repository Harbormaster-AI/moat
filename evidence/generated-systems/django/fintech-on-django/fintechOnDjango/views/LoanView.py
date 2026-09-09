import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.LoanDelegate import LoanDelegate

 #======================================================================
# 
# Encapsulates data for View Loan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Loan index.")

def get(request, loanId ):
	delegate = LoanDelegate()
	responseData = delegate.get( loanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	loan = json.loads(request.body)
	delegate = LoanDelegate()
	responseData = delegate.createFromJson( loan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	loan = json.loads(request.body)
	delegate = LoanDelegate()
	responseData = delegate.save( loan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, loanId ):
	delegate = LoanDelegate()
	responseData = delegate.delete( loanId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LoanDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, loanId, CustomerId ):
	delegate = LoanDelegate()
	responseData = delegate.saveCustomer( loanId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, loanId ):
	delegate = LoanDelegate()
	responseData = delegate.deleteCustomer( loanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSchedule( request, loanId, ScheduleIds ):
	delegate = LoanDelegate()
	responseData = delegate.addSchedule( loanId, ScheduleIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSchedule( request, loanId, ScheduleIds ):
	delegate = LoanDelegate()
	responseData = delegate.removeSchedule( loanId, ScheduleIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCollateral( request, loanId, CollateralIds ):
	delegate = LoanDelegate()
	responseData = delegate.addCollateral( loanId, CollateralIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCollateral( request, loanId, CollateralIds ):
	delegate = LoanDelegate()
	responseData = delegate.removeCollateral( loanId, CollateralIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, loanId, TransactionsIds ):
	delegate = LoanDelegate()
	responseData = delegate.addTransactions( loanId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, loanId, TransactionsIds ):
	delegate = LoanDelegate()
	responseData = delegate.removeTransactions( loanId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

