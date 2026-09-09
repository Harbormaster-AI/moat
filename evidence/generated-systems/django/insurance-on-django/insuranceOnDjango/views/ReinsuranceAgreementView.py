import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.ReinsuranceAgreementDelegate import ReinsuranceAgreementDelegate

 #======================================================================
# 
# Encapsulates data for View ReinsuranceAgreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReinsuranceAgreementView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ReinsuranceAgreement index.")

def get(request, reinsuranceAgreementId ):
	delegate = ReinsuranceAgreementDelegate()
	responseData = delegate.get( reinsuranceAgreementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	reinsuranceAgreement = json.loads(request.body)
	delegate = ReinsuranceAgreementDelegate()
	responseData = delegate.createFromJson( reinsuranceAgreement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	reinsuranceAgreement = json.loads(request.body)
	delegate = ReinsuranceAgreementDelegate()
	responseData = delegate.save( reinsuranceAgreement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, reinsuranceAgreementId ):
	delegate = ReinsuranceAgreementDelegate()
	responseData = delegate.delete( reinsuranceAgreementId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ReinsuranceAgreementDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInsurer( request, reinsuranceAgreementId, InsurerId ):
	delegate = ReinsuranceAgreementDelegate()
	responseData = delegate.saveInsurer( reinsuranceAgreementId, InsurerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInsurer( request, reinsuranceAgreementId ):
	delegate = ReinsuranceAgreementDelegate()
	responseData = delegate.deleteInsurer( reinsuranceAgreementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, reinsuranceAgreementId, PoliciesIds ):
	delegate = ReinsuranceAgreementDelegate()
	responseData = delegate.addPolicies( reinsuranceAgreementId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, reinsuranceAgreementId, PoliciesIds ):
	delegate = ReinsuranceAgreementDelegate()
	responseData = delegate.removePolicies( reinsuranceAgreementId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

