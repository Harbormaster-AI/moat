import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.ChargebackDelegate import ChargebackDelegate

 #======================================================================
# 
# Encapsulates data for View Chargeback
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChargebackView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Chargeback index.")

def get(request, chargebackId ):
	delegate = ChargebackDelegate()
	responseData = delegate.get( chargebackId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	chargeback = json.loads(request.body)
	delegate = ChargebackDelegate()
	responseData = delegate.createFromJson( chargeback )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	chargeback = json.loads(request.body)
	delegate = ChargebackDelegate()
	responseData = delegate.save( chargeback )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, chargebackId ):
	delegate = ChargebackDelegate()
	responseData = delegate.delete( chargebackId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ChargebackDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDispute( request, chargebackId, DisputeId ):
	delegate = ChargebackDelegate()
	responseData = delegate.saveDispute( chargebackId, DisputeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDispute( request, chargebackId ):
	delegate = ChargebackDelegate()
	responseData = delegate.deleteDispute( chargebackId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTransaction( request, chargebackId, TransactionId ):
	delegate = ChargebackDelegate()
	responseData = delegate.saveTransaction( chargebackId, TransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTransaction( request, chargebackId ):
	delegate = ChargebackDelegate()
	responseData = delegate.deleteTransaction( chargebackId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

