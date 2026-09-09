import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.ProductionCertificateDelegate import ProductionCertificateDelegate

 #======================================================================
# 
# Encapsulates data for View ProductionCertificate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionCertificateView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ProductionCertificate index.")

def get(request, productionCertificateId ):
	delegate = ProductionCertificateDelegate()
	responseData = delegate.get( productionCertificateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	productionCertificate = json.loads(request.body)
	delegate = ProductionCertificateDelegate()
	responseData = delegate.createFromJson( productionCertificate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	productionCertificate = json.loads(request.body)
	delegate = ProductionCertificateDelegate()
	responseData = delegate.save( productionCertificate )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, productionCertificateId ):
	delegate = ProductionCertificateDelegate()
	responseData = delegate.delete( productionCertificateId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProductionCertificateDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignManufacturer( request, productionCertificateId, ManufacturerId ):
	delegate = ProductionCertificateDelegate()
	responseData = delegate.saveManufacturer( productionCertificateId, ManufacturerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignManufacturer( request, productionCertificateId ):
	delegate = ProductionCertificateDelegate()
	responseData = delegate.deleteManufacturer( productionCertificateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

