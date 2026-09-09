import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

 #======================================================================
# 
# Encapsulates data for View Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Customer index.")

def get(request, customerId ):
	delegate = CustomerDelegate()
	responseData = delegate.get( customerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	customer = json.loads(request.body)
	delegate = CustomerDelegate()
	responseData = delegate.createFromJson( customer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	customer = json.loads(request.body)
	delegate = CustomerDelegate()
	responseData = delegate.save( customer )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, customerId ):
	delegate = CustomerDelegate()
	responseData = delegate.delete( customerId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CustomerDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAddresses( request, customerId, AddressesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addAddresses( customerId, AddressesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAddresses( request, customerId, AddressesIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeAddresses( customerId, AddressesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCarts( request, customerId, CartsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addCarts( customerId, CartsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCarts( request, customerId, CartsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeCarts( customerId, CartsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, customerId, OrdersIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addOrders( customerId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, customerId, OrdersIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeOrders( customerId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayments( request, customerId, PaymentsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addPayments( customerId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayments( request, customerId, PaymentsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removePayments( customerId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReviews( request, customerId, ReviewsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addReviews( customerId, ReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReviews( request, customerId, ReviewsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeReviews( customerId, ReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWishlists( request, customerId, WishlistsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addWishlists( customerId, WishlistsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWishlists( request, customerId, WishlistsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeWishlists( customerId, WishlistsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSubscriptions( request, customerId, SubscriptionsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addSubscriptions( customerId, SubscriptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSubscriptions( request, customerId, SubscriptionsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeSubscriptions( customerId, SubscriptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCouponRedemptions( request, customerId, CouponRedemptionsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addCouponRedemptions( customerId, CouponRedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCouponRedemptions( request, customerId, CouponRedemptionsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeCouponRedemptions( customerId, CouponRedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGiftCards( request, customerId, GiftCardsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.addGiftCards( customerId, GiftCardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGiftCards( request, customerId, GiftCardsIds ):
	delegate = CustomerDelegate()
	responseData = delegate.removeGiftCards( customerId, GiftCardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

