import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.SalesCampaignDelegate import SalesCampaignDelegate

 #======================================================================
# 
# Encapsulates data for View SalesCampaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesCampaignView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SalesCampaign index.")

def get(request, salesCampaignId ):
	delegate = SalesCampaignDelegate()
	responseData = delegate.get( salesCampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	salesCampaign = json.loads(request.body)
	delegate = SalesCampaignDelegate()
	responseData = delegate.createFromJson( salesCampaign )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	salesCampaign = json.loads(request.body)
	delegate = SalesCampaignDelegate()
	responseData = delegate.save( salesCampaign )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, salesCampaignId ):
	delegate = SalesCampaignDelegate()
	responseData = delegate.delete( salesCampaignId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SalesCampaignDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRegion( request, salesCampaignId, RegionId ):
	delegate = SalesCampaignDelegate()
	responseData = delegate.saveRegion( salesCampaignId, RegionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRegion( request, salesCampaignId ):
	delegate = SalesCampaignDelegate()
	responseData = delegate.deleteRegion( salesCampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOperator( request, salesCampaignId, OperatorId ):
	delegate = SalesCampaignDelegate()
	responseData = delegate.saveOperator( salesCampaignId, OperatorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOperator( request, salesCampaignId ):
	delegate = SalesCampaignDelegate()
	responseData = delegate.deleteOperator( salesCampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addQuotes( request, salesCampaignId, QuotesIds ):
	delegate = SalesCampaignDelegate()
	responseData = delegate.addQuotes( salesCampaignId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeQuotes( request, salesCampaignId, QuotesIds ):
	delegate = SalesCampaignDelegate()
	responseData = delegate.removeQuotes( salesCampaignId, QuotesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

