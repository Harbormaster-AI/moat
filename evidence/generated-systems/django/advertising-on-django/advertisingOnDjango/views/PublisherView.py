import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.PublisherDelegate import PublisherDelegate

 #======================================================================
# 
# Encapsulates data for View Publisher
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PublisherView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Publisher index.")

def get(request, publisherId ):
	delegate = PublisherDelegate()
	responseData = delegate.get( publisherId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	publisher = json.loads(request.body)
	delegate = PublisherDelegate()
	responseData = delegate.createFromJson( publisher )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	publisher = json.loads(request.body)
	delegate = PublisherDelegate()
	responseData = delegate.save( publisher )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, publisherId ):
	delegate = PublisherDelegate()
	responseData = delegate.delete( publisherId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PublisherDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventorySources( request, publisherId, InventorySourcesIds ):
	delegate = PublisherDelegate()
	responseData = delegate.addInventorySources( publisherId, InventorySourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventorySources( request, publisherId, InventorySourcesIds ):
	delegate = PublisherDelegate()
	responseData = delegate.removeInventorySources( publisherId, InventorySourcesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDeals( request, publisherId, DealsIds ):
	delegate = PublisherDelegate()
	responseData = delegate.addDeals( publisherId, DealsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDeals( request, publisherId, DealsIds ):
	delegate = PublisherDelegate()
	responseData = delegate.removeDeals( publisherId, DealsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCreativeApprovals( request, publisherId, CreativeApprovalsIds ):
	delegate = PublisherDelegate()
	responseData = delegate.addCreativeApprovals( publisherId, CreativeApprovalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCreativeApprovals( request, publisherId, CreativeApprovalsIds ):
	delegate = PublisherDelegate()
	responseData = delegate.removeCreativeApprovals( publisherId, CreativeApprovalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInsertionOrders( request, publisherId, InsertionOrdersIds ):
	delegate = PublisherDelegate()
	responseData = delegate.addInsertionOrders( publisherId, InsertionOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInsertionOrders( request, publisherId, InsertionOrdersIds ):
	delegate = PublisherDelegate()
	responseData = delegate.removeInsertionOrders( publisherId, InsertionOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRateCards( request, publisherId, RateCardsIds ):
	delegate = PublisherDelegate()
	responseData = delegate.addRateCards( publisherId, RateCardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRateCards( request, publisherId, RateCardsIds ):
	delegate = PublisherDelegate()
	responseData = delegate.removeRateCards( publisherId, RateCardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

