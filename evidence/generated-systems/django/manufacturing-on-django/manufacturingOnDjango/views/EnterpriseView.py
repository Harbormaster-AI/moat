import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.EnterpriseDelegate import EnterpriseDelegate

 #======================================================================
# 
# Encapsulates data for View Enterprise
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EnterpriseView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Enterprise index.")

def get(request, enterpriseId ):
	delegate = EnterpriseDelegate()
	responseData = delegate.get( enterpriseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	enterprise = json.loads(request.body)
	delegate = EnterpriseDelegate()
	responseData = delegate.createFromJson( enterprise )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	enterprise = json.loads(request.body)
	delegate = EnterpriseDelegate()
	responseData = delegate.save( enterprise )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, enterpriseId ):
	delegate = EnterpriseDelegate()
	responseData = delegate.delete( enterpriseId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EnterpriseDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBusinessUnits( request, enterpriseId, BusinessUnitsIds ):
	delegate = EnterpriseDelegate()
	responseData = delegate.addBusinessUnits( enterpriseId, BusinessUnitsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBusinessUnits( request, enterpriseId, BusinessUnitsIds ):
	delegate = EnterpriseDelegate()
	responseData = delegate.removeBusinessUnits( enterpriseId, BusinessUnitsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPlants( request, enterpriseId, PlantsIds ):
	delegate = EnterpriseDelegate()
	responseData = delegate.addPlants( enterpriseId, PlantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePlants( request, enterpriseId, PlantsIds ):
	delegate = EnterpriseDelegate()
	responseData = delegate.removePlants( enterpriseId, PlantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSuppliers( request, enterpriseId, SuppliersIds ):
	delegate = EnterpriseDelegate()
	responseData = delegate.addSuppliers( enterpriseId, SuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSuppliers( request, enterpriseId, SuppliersIds ):
	delegate = EnterpriseDelegate()
	responseData = delegate.removeSuppliers( enterpriseId, SuppliersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCustomers( request, enterpriseId, CustomersIds ):
	delegate = EnterpriseDelegate()
	responseData = delegate.addCustomers( enterpriseId, CustomersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCustomers( request, enterpriseId, CustomersIds ):
	delegate = EnterpriseDelegate()
	responseData = delegate.removeCustomers( enterpriseId, CustomersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

