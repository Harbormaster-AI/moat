import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.ReplenishmentPolicyDelegate import ReplenishmentPolicyDelegate

 #======================================================================
# 
# Encapsulates data for View ReplenishmentPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReplenishmentPolicyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ReplenishmentPolicy index.")

def get(request, replenishmentPolicyId ):
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.get( replenishmentPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	replenishmentPolicy = json.loads(request.body)
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.createFromJson( replenishmentPolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	replenishmentPolicy = json.loads(request.body)
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.save( replenishmentPolicy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, replenishmentPolicyId ):
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.delete( replenishmentPolicyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, replenishmentPolicyId, SkuId ):
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.saveSku( replenishmentPolicyId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, replenishmentPolicyId ):
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.deleteSku( replenishmentPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, replenishmentPolicyId, WarehouseId ):
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.saveWarehouse( replenishmentPolicyId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, replenishmentPolicyId ):
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.deleteWarehouse( replenishmentPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLocation( request, replenishmentPolicyId, LocationId ):
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.saveLocation( replenishmentPolicyId, LocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLocation( request, replenishmentPolicyId ):
	delegate = ReplenishmentPolicyDelegate()
	responseData = delegate.deleteLocation( replenishmentPolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

