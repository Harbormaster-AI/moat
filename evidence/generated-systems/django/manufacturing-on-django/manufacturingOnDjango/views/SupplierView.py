import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.SupplierDelegate import SupplierDelegate

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

def addEnterprises( request, supplierId, EnterprisesIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addEnterprises( supplierId, EnterprisesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEnterprises( request, supplierId, EnterprisesIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeEnterprises( supplierId, EnterprisesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addItems( request, supplierId, ItemsIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addItems( supplierId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeItems( request, supplierId, ItemsIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removeItems( supplierId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPurchaseOrders( request, supplierId, PurchaseOrdersIds ):
	delegate = SupplierDelegate()
	responseData = delegate.addPurchaseOrders( supplierId, PurchaseOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePurchaseOrders( request, supplierId, PurchaseOrdersIds ):
	delegate = SupplierDelegate()
	responseData = delegate.removePurchaseOrders( supplierId, PurchaseOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

