import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.SupplierDelegate import SupplierDelegate

 #======================================================================
# 
# Encapsulates data for View Supplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SupplierView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Supplier index.")

def get(request, supplierId ):
	delegate = SupplierDelegate()
	responseData = delegate.get( supplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	supplier = json.loads(request.body)
	delegate = SupplierDelegate()
	responseData = delegate.createFromJson( supplier )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	supplier = json.loads(request.body)
	delegate = SupplierDelegate()
	responseData = delegate.save( supplier )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, supplierId ):
	delegate = SupplierDelegate()
	responseData = delegate.delete( supplierId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SupplierDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, supplierId, MerchantId ):
	delegate = SupplierDelegate()
	responseData = delegate.saveMerchant( supplierId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, supplierId ):
	delegate = SupplierDelegate()
	responseData = delegate.deleteMerchant( supplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProducts( request, supplierId, ProductsIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addProducts( supplierId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProducts( request, supplierId, ProductsIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeProducts( supplierId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFulfillmentCenters( request, supplierId, FulfillmentCentersIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addFulfillmentCenters( supplierId, FulfillmentCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFulfillmentCenters( request, supplierId, FulfillmentCentersIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeFulfillmentCenters( supplierId, FulfillmentCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

