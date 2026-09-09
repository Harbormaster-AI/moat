import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

 #======================================================================
# 
# Encapsulates data for View Employee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmployeeView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Employee index.")

def get(request, employeeId ):
	delegate = EmployeeDelegate()
	responseData = delegate.get( employeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	employee = json.loads(request.body)
	delegate = EmployeeDelegate()
	responseData = delegate.createFromJson( employee )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	employee = json.loads(request.body)
	delegate = EmployeeDelegate()
	responseData = delegate.save( employee )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, employeeId ):
	delegate = EmployeeDelegate()
	responseData = delegate.delete( employeeId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EmployeeDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignManager( request, employeeId, ManagerId ):
	delegate = EmployeeDelegate()
	responseData = delegate.saveManager( employeeId, ManagerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignManager( request, employeeId ):
	delegate = EmployeeDelegate()
	responseData = delegate.deleteManager( employeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDepartment( request, employeeId, DepartmentId ):
	delegate = EmployeeDelegate()
	responseData = delegate.saveDepartment( employeeId, DepartmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDepartment( request, employeeId ):
	delegate = EmployeeDelegate()
	responseData = delegate.deleteDepartment( employeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPrimaryLocation( request, employeeId, PrimaryLocationId ):
	delegate = EmployeeDelegate()
	responseData = delegate.savePrimaryLocation( employeeId, PrimaryLocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPrimaryLocation( request, employeeId ):
	delegate = EmployeeDelegate()
	responseData = delegate.deletePrimaryLocation( employeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCostCenter( request, employeeId, CostCenterId ):
	delegate = EmployeeDelegate()
	responseData = delegate.saveCostCenter( employeeId, CostCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCostCenter( request, employeeId ):
	delegate = EmployeeDelegate()
	responseData = delegate.deleteCostCenter( employeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDirectReports( request, employeeId, DirectReportsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addDirectReports( employeeId, DirectReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDirectReports( request, employeeId, DirectReportsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeDirectReports( employeeId, DirectReportsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEmploymentAssignments( request, employeeId, EmploymentAssignmentsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addEmploymentAssignments( employeeId, EmploymentAssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEmploymentAssignments( request, employeeId, EmploymentAssignmentsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeEmploymentAssignments( employeeId, EmploymentAssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContracts( request, employeeId, ContractsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addContracts( employeeId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContracts( request, employeeId, ContractsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeContracts( employeeId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBenefitEnrollments( request, employeeId, BenefitEnrollmentsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addBenefitEnrollments( employeeId, BenefitEnrollmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBenefitEnrollments( request, employeeId, BenefitEnrollmentsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeBenefitEnrollments( employeeId, BenefitEnrollmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTimesheets( request, employeeId, TimesheetsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addTimesheets( employeeId, TimesheetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTimesheets( request, employeeId, TimesheetsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeTimesheets( employeeId, TimesheetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLeaveRequests( request, employeeId, LeaveRequestsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addLeaveRequests( employeeId, LeaveRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLeaveRequests( request, employeeId, LeaveRequestsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeLeaveRequests( employeeId, LeaveRequestsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPerformanceReviews( request, employeeId, PerformanceReviewsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addPerformanceReviews( employeeId, PerformanceReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePerformanceReviews( request, employeeId, PerformanceReviewsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removePerformanceReviews( employeeId, PerformanceReviewsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTrainingEnrollments( request, employeeId, TrainingEnrollmentsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addTrainingEnrollments( employeeId, TrainingEnrollmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTrainingEnrollments( request, employeeId, TrainingEnrollmentsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeTrainingEnrollments( employeeId, TrainingEnrollmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWorkAuthorizations( request, employeeId, WorkAuthorizationsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addWorkAuthorizations( employeeId, WorkAuthorizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWorkAuthorizations( request, employeeId, WorkAuthorizationsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeWorkAuthorizations( employeeId, WorkAuthorizationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

