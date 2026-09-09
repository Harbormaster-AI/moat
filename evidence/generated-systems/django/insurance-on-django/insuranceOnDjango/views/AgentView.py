import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.AgentDelegate import AgentDelegate

 #======================================================================
# 
# Encapsulates data for View Agent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Agent index.")

def get(request, agentId ):
	delegate = AgentDelegate()
	responseData = delegate.get( agentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	agent = json.loads(request.body)
	delegate = AgentDelegate()
	responseData = delegate.createFromJson( agent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	agent = json.loads(request.body)
	delegate = AgentDelegate()
	responseData = delegate.save( agent )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, agentId ):
	delegate = AgentDelegate()
	responseData = delegate.delete( agentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AgentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDistributor( request, agentId, DistributorId ):
	delegate = AgentDelegate()
	responseData = delegate.saveDistributor( agentId, DistributorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDistributor( request, agentId ):
	delegate = AgentDelegate()
	responseData = delegate.deleteDistributor( agentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, agentId, PoliciesIds ):
	delegate = AgentDelegate()
	responseData = delegate.addPolicies( agentId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, agentId, PoliciesIds ):
	delegate = AgentDelegate()
	responseData = delegate.removePolicies( agentId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCustomers( request, agentId, CustomersIds ):
	delegate = AgentDelegate()
	responseData = delegate.addCustomers( agentId, CustomersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCustomers( request, agentId, CustomersIds ):
	delegate = AgentDelegate()
	responseData = delegate.removeCustomers( agentId, CustomersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

