import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.BrandDelegate import BrandDelegate

 #======================================================================
# 
# Encapsulates data for View Brand
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BrandView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Brand index.")

def get(request, brandId ):
	delegate = BrandDelegate()
	responseData = delegate.get( brandId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	brand = json.loads(request.body)
	delegate = BrandDelegate()
	responseData = delegate.createFromJson( brand )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	brand = json.loads(request.body)
	delegate = BrandDelegate()
	responseData = delegate.save( brand )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, brandId ):
	delegate = BrandDelegate()
	responseData = delegate.delete( brandId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BrandDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, brandId, MerchantId ):
	delegate = BrandDelegate()
	responseData = delegate.saveMerchant( brandId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, brandId ):
	delegate = BrandDelegate()
	responseData = delegate.deleteMerchant( brandId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProducts( request, brandId, ProductsIds ):
	delegate = BrandDelegate()
	responseData = delegate.addProducts( brandId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProducts( request, brandId, ProductsIds ):
	delegate = BrandDelegate()
	responseData = delegate.removeProducts( brandId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

