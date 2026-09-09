import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.BOMDelegate import BOMDelegate

 #======================================================================
# 
# Encapsulates data for View BOM
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BOMView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BOM index.")

def get(request, bOMId ):
	delegate = BOMDelegate()
	responseData = delegate.get( bOMId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	bOM = json.loads(request.body)
	delegate = BOMDelegate()
	responseData = delegate.createFromJson( bOM )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	bOM = json.loads(request.body)
	delegate = BOMDelegate()
	responseData = delegate.save( bOM )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, bOMId ):
	delegate = BOMDelegate()
	responseData = delegate.delete( bOMId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BOMDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignParentItem( request, bOMId, ParentItemId ):
	delegate = BOMDelegate()
	responseData = delegate.saveParentItem( bOMId, ParentItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignParentItem( request, bOMId ):
	delegate = BOMDelegate()
	responseData = delegate.deleteParentItem( bOMId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBomItems( request, bOMId, BomItemsIds ):
	delegate = BOMDelegate()
	responseData = delegate.addBomItems( bOMId, BomItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBomItems( request, bOMId, BomItemsIds ):
	delegate = BOMDelegate()
	responseData = delegate.removeBomItems( bOMId, BomItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

