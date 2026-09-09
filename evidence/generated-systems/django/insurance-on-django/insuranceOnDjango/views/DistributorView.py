import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.DistributorDelegate import DistributorDelegate

 #======================================================================
# 
# Encapsulates data for View Distributor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DistributorView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Distributor index.")

def get(request, distributorId ):
	delegate = DistributorDelegate()
	responseData = delegate.get( distributorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	distributor = json.loads(request.body)
	delegate = DistributorDelegate()
	responseData = delegate.createFromJson( distributor )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	distributor = json.loads(request.body)
	delegate = DistributorDelegate()
	responseData = delegate.save( distributor )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, distributorId ):
	delegate = DistributorDelegate()
	responseData = delegate.delete( distributorId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DistributorDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInsurers( request, distributorId, InsurersIds ):
	delegate = DistributorDelegate()
	responseData = delegate.addInsurers( distributorId, InsurersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInsurers( request, distributorId, InsurersIds ):
	delegate = DistributorDelegate()
	responseData = delegate.removeInsurers( distributorId, InsurersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAgents( request, distributorId, AgentsIds ):
	delegate = DistributorDelegate()
	responseData = delegate.addAgents( distributorId, AgentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAgents( request, distributorId, AgentsIds ):
	delegate = DistributorDelegate()
	responseData = delegate.removeAgents( distributorId, AgentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, distributorId, PoliciesIds ):
	delegate = DistributorDelegate()
	responseData = delegate.addPolicies( distributorId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, distributorId, PoliciesIds ):
	delegate = DistributorDelegate()
	responseData = delegate.removePolicies( distributorId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

