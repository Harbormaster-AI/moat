import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.CouponDelegate import CouponDelegate

 #======================================================================
# 
# Encapsulates data for View Coupon
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CouponView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Coupon index.")

def get(request, couponId ):
	delegate = CouponDelegate()
	responseData = delegate.get( couponId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	coupon = json.loads(request.body)
	delegate = CouponDelegate()
	responseData = delegate.createFromJson( coupon )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	coupon = json.loads(request.body)
	delegate = CouponDelegate()
	responseData = delegate.save( coupon )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, couponId ):
	delegate = CouponDelegate()
	responseData = delegate.delete( couponId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CouponDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPromotion( request, couponId, PromotionId ):
	delegate = CouponDelegate()
	responseData = delegate.savePromotion( couponId, PromotionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPromotion( request, couponId ):
	delegate = CouponDelegate()
	responseData = delegate.deletePromotion( couponId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRedemptions( request, couponId, RedemptionsIds ):
	delegate = CouponDelegate()
	responseData = delegate.addRedemptions( couponId, RedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRedemptions( request, couponId, RedemptionsIds ):
	delegate = CouponDelegate()
	responseData = delegate.removeRedemptions( couponId, RedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

