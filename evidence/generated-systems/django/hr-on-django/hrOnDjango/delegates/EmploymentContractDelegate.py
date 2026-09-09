from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.EmploymentContract import EmploymentContract
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.CompensationPackage import CompensationPackage
from hrOnDjango.models.WorkSchedule import WorkSchedule
from hrOnDjango.models.Location import Location
from hrOnDjango.models.PayrollCalendar import PayrollCalendar
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model EmploymentContract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentContractDelegate Declaration
#======================================================================
class EmploymentContractDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, employmentContractId ):
		try:	
			employmentContract = EmploymentContract.objects.filter(id=employmentContractId)
			return employmentContract.first();
		except EmploymentContract.DoesNotExist:
			raise ProcessingError("EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, employmentContract):
		for model in serializers.deserialize("json", employmentContract):
			model.save()
			return model;

	def create(self, employmentContract):
		employmentContract.save()
		return employmentContract;

	def saveFromJson(self, employmentContract):
		for model in serializers.deserialize("json", employmentContract):
			model.save()
			return employmentContract;
	
	def save(self, employmentContract):
		employmentContract.save()
		return employmentContract;
	
	def delete(self, employmentContractId ):
		errMsg = "Failed to delete EmploymentContract from db using id " + str(employmentContractId)
		
		try:
			employmentContract = EmploymentContract.objects.get(id=employmentContractId)
			employmentContract.delete()
			return True
		except EmploymentContract.DoesNotExist:
			raise ProcessingError("EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = EmploymentContract.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all EmploymentContract from db")
		except Exception:
			return None;
		
	def assignEmployee( self, employmentContractId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			employmentContract.employee = employee
			
			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, employmentContractId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# assign to None for unassignment
			employmentContract.employee = None			

			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCompensationPackage( self, employmentContractId, compensationPackageId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompensationPackageDelegate import CompensationPackageDelegate

		errMsg = "Failed to assign element " + str(compensationPackageId) + " for CompensationPackage on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# get the CompensationPackage from db
			compensationPackage = CompensationPackageDelegate().get(compensationPackageId).first();
			
			# assign the CompensationPackage		
			employmentContract.compensationPackage = compensationPackage
			
			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage with id " + str(compensationPackageId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCompensationPackage( self, employmentContractId ):
		errMsg = "Failed to unassign element " + str(compensationPackageId) + " for CompensationPackage on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# assign to None for unassignment
			employmentContract.compensationPackage = None			

			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkSchedule( self, employmentContractId, workScheduleId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.WorkScheduleDelegate import WorkScheduleDelegate

		errMsg = "Failed to assign element " + str(workScheduleId) + " for WorkSchedule on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# get the WorkSchedule from db
			workSchedule = WorkScheduleDelegate().get(workScheduleId).first();
			
			# assign the WorkSchedule		
			employmentContract.workSchedule = workSchedule
			
			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except WorkSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkSchedule( self, employmentContractId ):
		errMsg = "Failed to unassign element " + str(workScheduleId) + " for WorkSchedule on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# assign to None for unassignment
			employmentContract.workSchedule = None			

			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, employmentContractId, locationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LocationDelegate import LocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# get the Location from db
			location = LocationDelegate().get(locationId).first();
			
			# assign the Location		
			employmentContract.location = location
			
			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, employmentContractId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# assign to None for unassignment
			employmentContract.location = None			

			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPayrollCalendar( self, employmentContractId, payrollCalendarId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PayrollCalendarDelegate import PayrollCalendarDelegate

		errMsg = "Failed to assign element " + str(payrollCalendarId) + " for PayrollCalendar on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# get the PayrollCalendar from db
			payrollCalendar = PayrollCalendarDelegate().get(payrollCalendarId).first();
			
			# assign the PayrollCalendar		
			employmentContract.payrollCalendar = payrollCalendar
			
			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPayrollCalendar( self, employmentContractId ):
		errMsg = "Failed to unassign element " + str(payrollCalendarId) + " for PayrollCalendar on EmploymentContract"

		try:
			# get the EmploymentContract from db
			employmentContract = self.get( employmentContractId ).first()	
			
			# assign to None for unassignment
			employmentContract.payrollCalendar = None			

			#save it
			employmentContract.save()

			# reload and return the appropriate version					
			return self.get( employmentContractId );
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(employmentContractId) + " does not exist.")
		except Exception:
			return None;
		
