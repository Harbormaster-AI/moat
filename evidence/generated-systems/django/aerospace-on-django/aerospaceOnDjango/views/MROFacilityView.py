import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.MROFacilityDelegate import MROFacilityDelegate

 #======================================================================
# 
# Encapsulates data for View MROFacility
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MROFacilityView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the MROFacility index.")

def get(request, mROFacilityId ):
	delegate = MROFacilityDelegate()
	responseData = delegate.get( mROFacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	mROFacility = json.loads(request.body)
	delegate = MROFacilityDelegate()
	responseData = delegate.createFromJson( mROFacility )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	mROFacility = json.loads(request.body)
	delegate = MROFacilityDelegate()
	responseData = delegate.save( mROFacility )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, mROFacilityId ):
	delegate = MROFacilityDelegate()
	responseData = delegate.delete( mROFacilityId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MROFacilityDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAppointments( request, mROFacilityId, AppointmentsIds ):
	delegate = MROFacilityDelegate()
	responseData = delegate.addAppointments( mROFacilityId, AppointmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAppointments( request, mROFacilityId, AppointmentsIds ):
	delegate = MROFacilityDelegate()
	responseData = delegate.removeAppointments( mROFacilityId, AppointmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWorkOrders( request, mROFacilityId, WorkOrdersIds ):
	delegate = MROFacilityDelegate()
	responseData = delegate.addWorkOrders( mROFacilityId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWorkOrders( request, mROFacilityId, WorkOrdersIds ):
	delegate = MROFacilityDelegate()
	responseData = delegate.removeWorkOrders( mROFacilityId, WorkOrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

