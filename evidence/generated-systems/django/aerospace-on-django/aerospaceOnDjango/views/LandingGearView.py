import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.LandingGearDelegate import LandingGearDelegate

 #======================================================================
# 
# Encapsulates data for View LandingGear
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LandingGearView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the LandingGear index.")

def get(request, landingGearId ):
	delegate = LandingGearDelegate()
	responseData = delegate.get( landingGearId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	landingGear = json.loads(request.body)
	delegate = LandingGearDelegate()
	responseData = delegate.createFromJson( landingGear )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	landingGear = json.loads(request.body)
	delegate = LandingGearDelegate()
	responseData = delegate.save( landingGear )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, landingGearId ):
	delegate = LandingGearDelegate()
	responseData = delegate.delete( landingGearId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = LandingGearDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSupplier( request, landingGearId, SupplierId ):
	delegate = LandingGearDelegate()
	responseData = delegate.saveSupplier( landingGearId, SupplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSupplier( request, landingGearId ):
	delegate = LandingGearDelegate()
	responseData = delegate.deleteSupplier( landingGearId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariants( request, landingGearId, VariantsIds ):
	delegate = LandingGearDelegate()
	responseData = delegate.addVariants( landingGearId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariants( request, landingGearId, VariantsIds ):
	delegate = LandingGearDelegate()
	responseData = delegate.removeVariants( landingGearId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

