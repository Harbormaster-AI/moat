import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.GoalDelegate import GoalDelegate

 #======================================================================
# 
# Encapsulates data for View Goal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoalView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Goal index.")

def get(request, goalId ):
	delegate = GoalDelegate()
	responseData = delegate.get( goalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	goal = json.loads(request.body)
	delegate = GoalDelegate()
	responseData = delegate.createFromJson( goal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	goal = json.loads(request.body)
	delegate = GoalDelegate()
	responseData = delegate.save( goal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, goalId ):
	delegate = GoalDelegate()
	responseData = delegate.delete( goalId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = GoalDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, goalId, EmployeeId ):
	delegate = GoalDelegate()
	responseData = delegate.saveEmployee( goalId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, goalId ):
	delegate = GoalDelegate()
	responseData = delegate.deleteEmployee( goalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCycle( request, goalId, CycleId ):
	delegate = GoalDelegate()
	responseData = delegate.saveCycle( goalId, CycleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCycle( request, goalId ):
	delegate = GoalDelegate()
	responseData = delegate.deleteCycle( goalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignParentGoal( request, goalId, ParentGoalId ):
	delegate = GoalDelegate()
	responseData = delegate.saveParentGoal( goalId, ParentGoalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignParentGoal( request, goalId ):
	delegate = GoalDelegate()
	responseData = delegate.deleteParentGoal( goalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChildGoals( request, goalId, ChildGoalsIds ):
	delegate = GoalDelegate()
	responseData = delegate.addChildGoals( goalId, ChildGoalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChildGoals( request, goalId, ChildGoalsIds ):
	delegate = GoalDelegate()
	responseData = delegate.removeChildGoals( goalId, ChildGoalsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

