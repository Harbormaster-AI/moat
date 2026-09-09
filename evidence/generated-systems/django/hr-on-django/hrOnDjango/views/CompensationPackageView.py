import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.CompensationPackageDelegate import CompensationPackageDelegate

 #======================================================================
# 
# Encapsulates data for View CompensationPackage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompensationPackageView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CompensationPackage index.")

def get(request, compensationPackageId ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.get( compensationPackageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	compensationPackage = json.loads(request.body)
	delegate = CompensationPackageDelegate()
	responseData = delegate.createFromJson( compensationPackage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	compensationPackage = json.loads(request.body)
	delegate = CompensationPackageDelegate()
	responseData = delegate.save( compensationPackage )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, compensationPackageId ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.delete( compensationPackageId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CompensationPackageDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignContract( request, compensationPackageId, ContractId ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.saveContract( compensationPackageId, ContractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignContract( request, compensationPackageId ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.deleteContract( compensationPackageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSalaryComponents( request, compensationPackageId, SalaryComponentsIds ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.addSalaryComponents( compensationPackageId, SalaryComponentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSalaryComponents( request, compensationPackageId, SalaryComponentsIds ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.removeSalaryComponents( compensationPackageId, SalaryComponentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBonusPlans( request, compensationPackageId, BonusPlansIds ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.addBonusPlans( compensationPackageId, BonusPlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBonusPlans( request, compensationPackageId, BonusPlansIds ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.removeBonusPlans( compensationPackageId, BonusPlansIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEquityGrants( request, compensationPackageId, EquityGrantsIds ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.addEquityGrants( compensationPackageId, EquityGrantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEquityGrants( request, compensationPackageId, EquityGrantsIds ):
	delegate = CompensationPackageDelegate()
	responseData = delegate.removeEquityGrants( compensationPackageId, EquityGrantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

