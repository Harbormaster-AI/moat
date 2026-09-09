import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.ServiceBulletinDelegate import ServiceBulletinDelegate

 #======================================================================
# 
# Encapsulates data for View ServiceBulletin
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceBulletinView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ServiceBulletin index.")

def get(request, serviceBulletinId ):
	delegate = ServiceBulletinDelegate()
	responseData = delegate.get( serviceBulletinId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	serviceBulletin = json.loads(request.body)
	delegate = ServiceBulletinDelegate()
	responseData = delegate.createFromJson( serviceBulletin )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	serviceBulletin = json.loads(request.body)
	delegate = ServiceBulletinDelegate()
	responseData = delegate.save( serviceBulletin )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, serviceBulletinId ):
	delegate = ServiceBulletinDelegate()
	responseData = delegate.delete( serviceBulletinId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ServiceBulletinDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWorkOrders( request, serviceBulletinId, WorkOrdersIds ):
	delegate = ServiceBulletinDelegate()
	responseData = delegate.addWorkOrders( serviceBulletinId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWorkOrders( request, serviceBulletinId, WorkOrdersIds ):
	delegate = ServiceBulletinDelegate()
	responseData = delegate.removeWorkOrders( serviceBulletinId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariants( request, serviceBulletinId, VariantsIds ):
	delegate = ServiceBulletinDelegate()
	responseData = delegate.addVariants( serviceBulletinId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariants( request, serviceBulletinId, VariantsIds ):
	delegate = ServiceBulletinDelegate()
	responseData = delegate.removeVariants( serviceBulletinId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

