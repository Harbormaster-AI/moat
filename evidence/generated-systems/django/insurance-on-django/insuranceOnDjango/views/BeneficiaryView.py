import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.BeneficiaryDelegate import BeneficiaryDelegate

 #======================================================================
# 
# Encapsulates data for View Beneficiary
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BeneficiaryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Beneficiary index.")

def get(request, beneficiaryId ):
	delegate = BeneficiaryDelegate()
	responseData = delegate.get( beneficiaryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	beneficiary = json.loads(request.body)
	delegate = BeneficiaryDelegate()
	responseData = delegate.createFromJson( beneficiary )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	beneficiary = json.loads(request.body)
	delegate = BeneficiaryDelegate()
	responseData = delegate.save( beneficiary )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, beneficiaryId ):
	delegate = BeneficiaryDelegate()
	responseData = delegate.delete( beneficiaryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BeneficiaryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, beneficiaryId, PolicyId ):
	delegate = BeneficiaryDelegate()
	responseData = delegate.savePolicy( beneficiaryId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, beneficiaryId ):
	delegate = BeneficiaryDelegate()
	responseData = delegate.deletePolicy( beneficiaryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, beneficiaryId, CustomerId ):
	delegate = BeneficiaryDelegate()
	responseData = delegate.saveCustomer( beneficiaryId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, beneficiaryId ):
	delegate = BeneficiaryDelegate()
	responseData = delegate.deleteCustomer( beneficiaryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

