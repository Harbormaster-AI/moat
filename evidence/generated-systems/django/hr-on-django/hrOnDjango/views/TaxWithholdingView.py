import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.TaxWithholdingDelegate import TaxWithholdingDelegate

 #======================================================================
# 
# Encapsulates data for View TaxWithholding
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaxWithholdingView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TaxWithholding index.")

def get(request, taxWithholdingId ):
	delegate = TaxWithholdingDelegate()
	responseData = delegate.get( taxWithholdingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	taxWithholding = json.loads(request.body)
	delegate = TaxWithholdingDelegate()
	responseData = delegate.createFromJson( taxWithholding )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	taxWithholding = json.loads(request.body)
	delegate = TaxWithholdingDelegate()
	responseData = delegate.save( taxWithholding )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, taxWithholdingId ):
	delegate = TaxWithholdingDelegate()
	responseData = delegate.delete( taxWithholdingId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TaxWithholdingDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, taxWithholdingId, EmployeeId ):
	delegate = TaxWithholdingDelegate()
	responseData = delegate.saveEmployee( taxWithholdingId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, taxWithholdingId ):
	delegate = TaxWithholdingDelegate()
	responseData = delegate.deleteEmployee( taxWithholdingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

