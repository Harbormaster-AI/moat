import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.ExpirationPolicyDelegate import ExpirationPolicyDelegate

 #======================================================================
# 
# Encapsulates data for View ExpirationPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExpirationPolicyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ExpirationPolicy index.")

def get(request, expirationPolicyId ):
	delegate = ExpirationPolicyDelegate()
	responseData = delegate.get( expirationPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	expirationPolicy = json.loads(request.body)
	delegate = ExpirationPolicyDelegate()
	responseData = delegate.createFromJson( expirationPolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	expirationPolicy = json.loads(request.body)
	delegate = ExpirationPolicyDelegate()
	responseData = delegate.save( expirationPolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, expirationPolicyId ):
	delegate = ExpirationPolicyDelegate()
	responseData = delegate.delete( expirationPolicyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ExpirationPolicyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, expirationPolicyId, SkuId ):
	delegate = ExpirationPolicyDelegate()
	responseData = delegate.saveSku( expirationPolicyId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, expirationPolicyId ):
	delegate = ExpirationPolicyDelegate()
	responseData = delegate.deleteSku( expirationPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, expirationPolicyId, WarehouseId ):
	delegate = ExpirationPolicyDelegate()
	responseData = delegate.saveWarehouse( expirationPolicyId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, expirationPolicyId ):
	delegate = ExpirationPolicyDelegate()
	responseData = delegate.deleteWarehouse( expirationPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

