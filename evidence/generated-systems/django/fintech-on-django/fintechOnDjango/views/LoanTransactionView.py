import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.LoanTransactionDelegate import LoanTransactionDelegate

 #======================================================================
# 
# Encapsulates data for View LoanTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanTransactionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LoanTransaction index.")

def get(request, loanTransactionId ):
	delegate = LoanTransactionDelegate()
	responseData = delegate.get( loanTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	loanTransaction = json.loads(request.body)
	delegate = LoanTransactionDelegate()
	responseData = delegate.createFromJson( loanTransaction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	loanTransaction = json.loads(request.body)
	delegate = LoanTransactionDelegate()
	responseData = delegate.save( loanTransaction )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, loanTransactionId ):
	delegate = LoanTransactionDelegate()
	responseData = delegate.delete( loanTransactionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LoanTransactionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLoan( request, loanTransactionId, LoanId ):
	delegate = LoanTransactionDelegate()
	responseData = delegate.saveLoan( loanTransactionId, LoanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLoan( request, loanTransactionId ):
	delegate = LoanTransactionDelegate()
	responseData = delegate.deleteLoan( loanTransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

