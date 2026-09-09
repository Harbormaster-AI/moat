import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.FinancialInstitutionDelegate import FinancialInstitutionDelegate

 #======================================================================
# 
# Encapsulates data for View FinancialInstitution
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FinancialInstitutionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the FinancialInstitution index.")

def get(request, financialInstitutionId ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.get( financialInstitutionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	financialInstitution = json.loads(request.body)
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.createFromJson( financialInstitution )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	financialInstitution = json.loads(request.body)
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.save( financialInstitution )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, financialInstitutionId ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.delete( financialInstitutionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBranches( request, financialInstitutionId, BranchesIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.addBranches( financialInstitutionId, BranchesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBranches( request, financialInstitutionId, BranchesIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.removeBranches( financialInstitutionId, BranchesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCustomers( request, financialInstitutionId, CustomersIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.addCustomers( financialInstitutionId, CustomersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCustomers( request, financialInstitutionId, CustomersIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.removeCustomers( financialInstitutionId, CustomersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProductOfferings( request, financialInstitutionId, ProductOfferingsIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.addProductOfferings( financialInstitutionId, ProductOfferingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProductOfferings( request, financialInstitutionId, ProductOfferingsIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.removeProductOfferings( financialInstitutionId, ProductOfferingsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPaymentProcessors( request, financialInstitutionId, PaymentProcessorsIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.addPaymentProcessors( financialInstitutionId, PaymentProcessorsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePaymentProcessors( request, financialInstitutionId, PaymentProcessorsIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.removePaymentProcessors( financialInstitutionId, PaymentProcessorsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCompliancePolicies( request, financialInstitutionId, CompliancePoliciesIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.addCompliancePolicies( financialInstitutionId, CompliancePoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCompliancePolicies( request, financialInstitutionId, CompliancePoliciesIds ):
	delegate = FinancialInstitutionDelegate()
	responseData = delegate.removeCompliancePolicies( financialInstitutionId, CompliancePoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

