import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.ContractDelegate import ContractDelegate

 #======================================================================
# 
# Encapsulates data for View Contract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContractView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Contract index.")

def get(request, contractId ):
	delegate = ContractDelegate()
	responseData = delegate.get( contractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	contract = json.loads(request.body)
	delegate = ContractDelegate()
	responseData = delegate.createFromJson( contract )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	contract = json.loads(request.body)
	delegate = ContractDelegate()
	responseData = delegate.save( contract )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, contractId ):
	delegate = ContractDelegate()
	responseData = delegate.delete( contractId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ContractDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, contractId, OrganizationId ):
	delegate = ContractDelegate()
	responseData = delegate.saveOrganization( contractId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, contractId ):
	delegate = ContractDelegate()
	responseData = delegate.deleteOrganization( contractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAccount( request, contractId, AccountId ):
	delegate = ContractDelegate()
	responseData = delegate.saveAccount( contractId, AccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAccount( request, contractId ):
	delegate = ContractDelegate()
	responseData = delegate.deleteAccount( contractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOwner( request, contractId, OwnerId ):
	delegate = ContractDelegate()
	responseData = delegate.saveOwner( contractId, OwnerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOwner( request, contractId ):
	delegate = ContractDelegate()
	responseData = delegate.deleteOwner( contractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, contractId, OrdersIds ):
	delegate = ContractDelegate()
	responseData = delegate.addOrders( contractId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, contractId, OrdersIds ):
	delegate = ContractDelegate()
	responseData = delegate.removeOrders( contractId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCases( request, contractId, CasesIds ):
	delegate = ContractDelegate()
	responseData = delegate.addCases( contractId, CasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCases( request, contractId, CasesIds ):
	delegate = ContractDelegate()
	responseData = delegate.removeCases( contractId, CasesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

