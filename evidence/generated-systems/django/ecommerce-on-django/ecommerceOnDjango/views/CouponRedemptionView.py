import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.CouponRedemptionDelegate import CouponRedemptionDelegate

 #======================================================================
# 
# Encapsulates data for View CouponRedemption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CouponRedemptionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CouponRedemption index.")

def get(request, couponRedemptionId ):
	delegate = CouponRedemptionDelegate()
	responseData = delegate.get( couponRedemptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	couponRedemption = json.loads(request.body)
	delegate = CouponRedemptionDelegate()
	responseData = delegate.createFromJson( couponRedemption )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	couponRedemption = json.loads(request.body)
	delegate = CouponRedemptionDelegate()
	responseData = delegate.save( couponRedemption )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, couponRedemptionId ):
	delegate = CouponRedemptionDelegate()
	responseData = delegate.delete( couponRedemptionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CouponRedemptionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCoupon( request, couponRedemptionId, CouponId ):
	delegate = CouponRedemptionDelegate()
	responseData = delegate.saveCoupon( couponRedemptionId, CouponId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCoupon( request, couponRedemptionId ):
	delegate = CouponRedemptionDelegate()
	responseData = delegate.deleteCoupon( couponRedemptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, couponRedemptionId, OrderId ):
	delegate = CouponRedemptionDelegate()
	responseData = delegate.saveOrder( couponRedemptionId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, couponRedemptionId ):
	delegate = CouponRedemptionDelegate()
	responseData = delegate.deleteOrder( couponRedemptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, couponRedemptionId, CustomerId ):
	delegate = CouponRedemptionDelegate()
	responseData = delegate.saveCustomer( couponRedemptionId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, couponRedemptionId ):
	delegate = CouponRedemptionDelegate()
	responseData = delegate.deleteCustomer( couponRedemptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

