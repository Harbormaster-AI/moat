import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

 #======================================================================
# 
# Encapsulates data for View Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Customer index.")

def get(request, customerId ):
	delegate = CustomerDelegate()
	responseData = delegate.get( customerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	customer = json.loads(request.body)
	delegate = CustomerDelegate()
	responseData = delegate.createFromJson( customer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	customer = json.loads(request.body)
	delegate = CustomerDelegate()
	responseData = delegate.save( customer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, customerId ):
	delegate = CustomerDelegate()
	responseData = delegate.delete( customerId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CustomerDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addApplications( request, customerId, ApplicationsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addApplications( customerId, ApplicationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeApplications( request, customerId, ApplicationsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeApplications( customerId, ApplicationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, customerId, PoliciesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addPolicies( customerId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, customerId, PoliciesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removePolicies( customerId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClaims( request, customerId, ClaimsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addClaims( customerId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClaims( request, customerId, ClaimsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeClaims( customerId, ClaimsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAgents( request, customerId, AgentsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addAgents( customerId, AgentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAgents( request, customerId, AgentsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeAgents( customerId, AgentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBeneficiaries( request, customerId, BeneficiariesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addBeneficiaries( customerId, BeneficiariesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBeneficiaries( request, customerId, BeneficiariesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeBeneficiaries( customerId, BeneficiariesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

