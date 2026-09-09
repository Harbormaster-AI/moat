import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

 #======================================================================
# 
# Encapsulates data for View Order
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Order index.")

def get(request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.get( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	order = json.loads(request.body)
	delegate = OrderDelegate()
	responseData = delegate.createFromJson( order )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	order = json.loads(request.body)
	delegate = OrderDelegate()
	responseData = delegate.save( order )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.delete( orderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, orderId, CustomerId ):
	delegate = OrderDelegate()
	responseData = delegate.saveCustomer( orderId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteCustomer( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignChannel( request, orderId, ChannelId ):
	delegate = OrderDelegate()
	responseData = delegate.saveChannel( orderId, ChannelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignChannel( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteChannel( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSeller( request, orderId, SellerId ):
	delegate = OrderDelegate()
	responseData = delegate.saveSeller( orderId, SellerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSeller( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteSeller( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInvoice( request, orderId, InvoiceId ):
	delegate = OrderDelegate()
	responseData = delegate.saveInvoice( orderId, InvoiceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInvoice( request, orderId ):
	delegate = OrderDelegate()
	responseData = delegate.deleteInvoice( orderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrderLines( request, orderId, OrderLinesIds ):
	delegate = OrderDelegate()
	responseData = delegate.addOrderLines( orderId, OrderLinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrderLines( request, orderId, OrderLinesIds ):
	delegate = OrderDelegate()
	responseData = delegate.removeOrderLines( orderId, OrderLinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayments( request, orderId, PaymentsIds ):
	delegate = OrderDelegate()
	responseData = delegate.addPayments( orderId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayments( request, orderId, PaymentsIds ):
	delegate = OrderDelegate()
	responseData = delegate.removePayments( orderId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addShipments( request, orderId, ShipmentsIds ):
	delegate = OrderDelegate()
	responseData = delegate.addShipments( orderId, ShipmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeShipments( request, orderId, ShipmentsIds ):
	delegate = OrderDelegate()
	responseData = delegate.removeShipments( orderId, ShipmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRefunds( request, orderId, RefundsIds ):
	delegate = OrderDelegate()
	responseData = delegate.addRefunds( orderId, RefundsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRefunds( request, orderId, RefundsIds ):
	delegate = OrderDelegate()
	responseData = delegate.removeRefunds( orderId, RefundsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAppliedPromotions( request, orderId, AppliedPromotionsIds ):
	delegate = OrderDelegate()
	responseData = delegate.addAppliedPromotions( orderId, AppliedPromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAppliedPromotions( request, orderId, AppliedPromotionsIds ):
	delegate = OrderDelegate()
	responseData = delegate.removeAppliedPromotions( orderId, AppliedPromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addGiftCardRedemptions( request, orderId, GiftCardRedemptionsIds ):
	delegate = OrderDelegate()
	responseData = delegate.addGiftCardRedemptions( orderId, GiftCardRedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeGiftCardRedemptions( request, orderId, GiftCardRedemptionsIds ):
	delegate = OrderDelegate()
	responseData = delegate.removeGiftCardRedemptions( orderId, GiftCardRedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCouponRedemptions( request, orderId, CouponRedemptionsIds ):
	delegate = OrderDelegate()
	responseData = delegate.addCouponRedemptions( orderId, CouponRedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCouponRedemptions( request, orderId, CouponRedemptionsIds ):
	delegate = OrderDelegate()
	responseData = delegate.removeCouponRedemptions( orderId, CouponRedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReturnRequests( request, orderId, ReturnRequestsIds ):
	delegate = OrderDelegate()
	responseData = delegate.addReturnRequests( orderId, ReturnRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReturnRequests( request, orderId, ReturnRequestsIds ):
	delegate = OrderDelegate()
	responseData = delegate.removeReturnRequests( orderId, ReturnRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

