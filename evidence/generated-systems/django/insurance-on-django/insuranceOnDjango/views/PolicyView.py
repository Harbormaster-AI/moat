import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

 #======================================================================
# 
# Encapsulates data for View Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Policy index.")

def get(request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.get( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	policy = json.loads(request.body)
	delegate = PolicyDelegate()
	responseData = delegate.createFromJson( policy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	policy = json.loads(request.body)
	delegate = PolicyDelegate()
	responseData = delegate.save( policy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.delete( policyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PolicyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInsurer( request, policyId, InsurerId ):
	delegate = PolicyDelegate()
	responseData = delegate.saveInsurer( policyId, InsurerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInsurer( request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.deleteInsurer( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, policyId, CustomerId ):
	delegate = PolicyDelegate()
	responseData = delegate.saveCustomer( policyId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.deleteCustomer( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, policyId, ProductId ):
	delegate = PolicyDelegate()
	responseData = delegate.saveProduct( policyId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.deleteProduct( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAgent( request, policyId, AgentId ):
	delegate = PolicyDelegate()
	responseData = delegate.saveAgent( policyId, AgentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAgent( request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.deleteAgent( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBillingAccount( request, policyId, BillingAccountId ):
	delegate = PolicyDelegate()
	responseData = delegate.saveBillingAccount( policyId, BillingAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBillingAccount( request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.deleteBillingAccount( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCoverages( request, policyId, CoveragesIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addCoverages( policyId, CoveragesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCoverages( request, policyId, CoveragesIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeCoverages( policyId, CoveragesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInsuredObjects( request, policyId, InsuredObjectsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addInsuredObjects( policyId, InsuredObjectsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInsuredObjects( request, policyId, InsuredObjectsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeInsuredObjects( policyId, InsuredObjectsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEndorsements( request, policyId, EndorsementsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addEndorsements( policyId, EndorsementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEndorsements( request, policyId, EndorsementsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeEndorsements( policyId, EndorsementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBeneficiaries( request, policyId, BeneficiariesIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addBeneficiaries( policyId, BeneficiariesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBeneficiaries( request, policyId, BeneficiariesIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeBeneficiaries( policyId, BeneficiariesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClaims( request, policyId, ClaimsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addClaims( policyId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClaims( request, policyId, ClaimsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeClaims( policyId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReinsuranceAgreements( request, policyId, ReinsuranceAgreementsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addReinsuranceAgreements( policyId, ReinsuranceAgreementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReinsuranceAgreements( request, policyId, ReinsuranceAgreementsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeReinsuranceAgreements( policyId, ReinsuranceAgreementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

