import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.TaxRuleDelegate import TaxRuleDelegate

 #======================================================================
# 
# Encapsulates data for View TaxRule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaxRuleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the TaxRule index.")

def get(request, taxRuleId ):
	delegate = TaxRuleDelegate()
	responseData = delegate.get( taxRuleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	taxRule = json.loads(request.body)
	delegate = TaxRuleDelegate()
	responseData = delegate.createFromJson( taxRule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	taxRule = json.loads(request.body)
	delegate = TaxRuleDelegate()
	responseData = delegate.save( taxRule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, taxRuleId ):
	delegate = TaxRuleDelegate()
	responseData = delegate.delete( taxRuleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TaxRuleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, taxRuleId, MerchantId ):
	delegate = TaxRuleDelegate()
	responseData = delegate.saveMerchant( taxRuleId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, taxRuleId ):
	delegate = TaxRuleDelegate()
	responseData = delegate.deleteMerchant( taxRuleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChannels( request, taxRuleId, ChannelsIds ):
	delegate = TaxRuleDelegate()
	responseData = delegate.addChannels( taxRuleId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChannels( request, taxRuleId, ChannelsIds ):
	delegate = TaxRuleDelegate()
	responseData = delegate.removeChannels( taxRuleId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

