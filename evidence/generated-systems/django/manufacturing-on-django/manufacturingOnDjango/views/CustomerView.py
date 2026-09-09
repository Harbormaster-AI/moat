import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.CustomerDelegate import CustomerDelegate

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

def addEnterprises( request, customerId, EnterprisesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addEnterprises( customerId, EnterprisesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEnterprises( request, customerId, EnterprisesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeEnterprises( customerId, EnterprisesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSalesOrders( request, customerId, SalesOrdersIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addSalesOrders( customerId, SalesOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSalesOrders( request, customerId, SalesOrdersIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeSalesOrders( customerId, SalesOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

