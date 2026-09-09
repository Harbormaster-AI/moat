import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.LoanApplicationDelegate import LoanApplicationDelegate

 #======================================================================
# 
# Encapsulates data for View LoanApplication
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanApplicationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LoanApplication index.")

def get(request, loanApplicationId ):
	delegate = LoanApplicationDelegate()
	responseData = delegate.get( loanApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	loanApplication = json.loads(request.body)
	delegate = LoanApplicationDelegate()
	responseData = delegate.createFromJson( loanApplication )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	loanApplication = json.loads(request.body)
	delegate = LoanApplicationDelegate()
	responseData = delegate.save( loanApplication )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, loanApplicationId ):
	delegate = LoanApplicationDelegate()
	responseData = delegate.delete( loanApplicationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LoanApplicationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, loanApplicationId, CustomerId ):
	delegate = LoanApplicationDelegate()
	responseData = delegate.saveCustomer( loanApplicationId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, loanApplicationId ):
	delegate = LoanApplicationDelegate()
	responseData = delegate.deleteCustomer( loanApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRiskAssessment( request, loanApplicationId, RiskAssessmentId ):
	delegate = LoanApplicationDelegate()
	responseData = delegate.saveRiskAssessment( loanApplicationId, RiskAssessmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRiskAssessment( request, loanApplicationId ):
	delegate = LoanApplicationDelegate()
	responseData = delegate.deleteRiskAssessment( loanApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLoan( request, loanApplicationId, LoanId ):
	delegate = LoanApplicationDelegate()
	responseData = delegate.saveLoan( loanApplicationId, LoanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLoan( request, loanApplicationId ):
	delegate = LoanApplicationDelegate()
	responseData = delegate.deleteLoan( loanApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

