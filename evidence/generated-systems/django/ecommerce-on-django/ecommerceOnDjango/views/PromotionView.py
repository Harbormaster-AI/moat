import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

 #======================================================================
# 
# Encapsulates data for View Promotion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PromotionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Promotion index.")

def get(request, promotionId ):
	delegate = PromotionDelegate()
	responseData = delegate.get( promotionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	promotion = json.loads(request.body)
	delegate = PromotionDelegate()
	responseData = delegate.createFromJson( promotion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	promotion = json.loads(request.body)
	delegate = PromotionDelegate()
	responseData = delegate.save( promotion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, promotionId ):
	delegate = PromotionDelegate()
	responseData = delegate.delete( promotionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PromotionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, promotionId, MerchantId ):
	delegate = PromotionDelegate()
	responseData = delegate.saveMerchant( promotionId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, promotionId ):
	delegate = PromotionDelegate()
	responseData = delegate.deleteMerchant( promotionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChannels( request, promotionId, ChannelsIds ):
	delegate = PromotionDelegate()
	responseData = delegate.addChannels( promotionId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChannels( request, promotionId, ChannelsIds ):
	delegate = PromotionDelegate()
	responseData = delegate.removeChannels( promotionId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addApplicableProducts( request, promotionId, ApplicableProductsIds ):
	delegate = PromotionDelegate()
	responseData = delegate.addApplicableProducts( promotionId, ApplicableProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeApplicableProducts( request, promotionId, ApplicableProductsIds ):
	delegate = PromotionDelegate()
	responseData = delegate.removeApplicableProducts( promotionId, ApplicableProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addApplicableCategories( request, promotionId, ApplicableCategoriesIds ):
	delegate = PromotionDelegate()
	responseData = delegate.addApplicableCategories( promotionId, ApplicableCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeApplicableCategories( request, promotionId, ApplicableCategoriesIds ):
	delegate = PromotionDelegate()
	responseData = delegate.removeApplicableCategories( promotionId, ApplicableCategoriesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCoupons( request, promotionId, CouponsIds ):
	delegate = PromotionDelegate()
	responseData = delegate.addCoupons( promotionId, CouponsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCoupons( request, promotionId, CouponsIds ):
	delegate = PromotionDelegate()
	responseData = delegate.removeCoupons( promotionId, CouponsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

