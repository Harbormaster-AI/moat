import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.CollateralDelegate import CollateralDelegate

 #======================================================================
# 
# Encapsulates data for View Collateral
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CollateralView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Collateral index.")

def get(request, collateralId ):
	delegate = CollateralDelegate()
	responseData = delegate.get( collateralId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	collateral = json.loads(request.body)
	delegate = CollateralDelegate()
	responseData = delegate.createFromJson( collateral )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	collateral = json.loads(request.body)
	delegate = CollateralDelegate()
	responseData = delegate.save( collateral )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, collateralId ):
	delegate = CollateralDelegate()
	responseData = delegate.delete( collateralId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CollateralDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLoan( request, collateralId, LoanId ):
	delegate = CollateralDelegate()
	responseData = delegate.saveLoan( collateralId, LoanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLoan( request, collateralId ):
	delegate = CollateralDelegate()
	responseData = delegate.deleteLoan( collateralId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

