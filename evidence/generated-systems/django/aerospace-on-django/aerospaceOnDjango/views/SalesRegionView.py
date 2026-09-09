import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.SalesRegionDelegate import SalesRegionDelegate

 #======================================================================
# 
# Encapsulates data for View SalesRegion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesRegionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SalesRegion index.")

def get(request, salesRegionId ):
	delegate = SalesRegionDelegate()
	responseData = delegate.get( salesRegionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	salesRegion = json.loads(request.body)
	delegate = SalesRegionDelegate()
	responseData = delegate.createFromJson( salesRegion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	salesRegion = json.loads(request.body)
	delegate = SalesRegionDelegate()
	responseData = delegate.save( salesRegion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, salesRegionId ):
	delegate = SalesRegionDelegate()
	responseData = delegate.delete( salesRegionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SalesRegionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOperators( request, salesRegionId, OperatorsIds ):
	delegate = SalesRegionDelegate()
	responseData = delegate.addOperators( salesRegionId, OperatorsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOperators( request, salesRegionId, OperatorsIds ):
	delegate = SalesRegionDelegate()
	responseData = delegate.removeOperators( salesRegionId, OperatorsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSalesCampaigns( request, salesRegionId, SalesCampaignsIds ):
	delegate = SalesRegionDelegate()
	responseData = delegate.addSalesCampaigns( salesRegionId, SalesCampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSalesCampaigns( request, salesRegionId, SalesCampaignsIds ):
	delegate = SalesRegionDelegate()
	responseData = delegate.removeSalesCampaigns( salesRegionId, SalesCampaignsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

