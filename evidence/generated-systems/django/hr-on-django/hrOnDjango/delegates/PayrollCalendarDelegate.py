from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.PayrollCalendar import PayrollCalendar
from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.PayrollRun import PayrollRun
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PayrollCalendar
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollCalendarDelegate Declaration
#======================================================================
class PayrollCalendarDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, payrollCalendarId ):
		try:	
			payrollCalendar = PayrollCalendar.objects.filter(id=payrollCalendarId)
			return payrollCalendar.first();
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError("PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, payrollCalendar):
		for model in serializers.deserialize("json", payrollCalendar):
			model.save()
			return model;

	def create(self, payrollCalendar):
		payrollCalendar.save()
		return payrollCalendar;

	def saveFromJson(self, payrollCalendar):
		for model in serializers.deserialize("json", payrollCalendar):
			model.save()
			return payrollCalendar;
	
	def save(self, payrollCalendar):
		payrollCalendar.save()
		return payrollCalendar;
	
	def delete(self, payrollCalendarId ):
		errMsg = "Failed to delete PayrollCalendar from db using id " + str(payrollCalendarId)
		
		try:
			payrollCalendar = PayrollCalendar.objects.get(id=payrollCalendarId)
			payrollCalendar.delete()
			return True
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError("PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PayrollCalendar.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PayrollCalendar from db")
		except Exception:
			return None;
		
	def assignOrganization( self, payrollCalendarId, organizationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on PayrollCalendar"

		try:
			# get the PayrollCalendar from db
			payrollCalendar = self.get( payrollCalendarId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			payrollCalendar.organization = organization
			
			#save it
			payrollCalendar.save()

			# reload and return the appropriate version					
			return self.get( payrollCalendarId );
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, payrollCalendarId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on PayrollCalendar"

		try:
			# get the PayrollCalendar from db
			payrollCalendar = self.get( payrollCalendarId ).first()	
			
			# assign to None for unassignment
			payrollCalendar.organization = None			

			#save it
			payrollCalendar.save()

			# reload and return the appropriate version					
			return self.get( payrollCalendarId );
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except Exception:
			return None;
		
	def addPayrollRuns( self, payrollCalendarId, payrollRunsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PayrollRunDelegate import PayrollRunDelegate

		errMsg = "Failed to add elements " + str(payrollRunsIds) + " for PayrollRuns on PayrollCalendar"

		try:
			# get the PayrollCalendar
			payrollCalendar = self.get( payrollCalendarId ).first()
				
			# split on a comma with no spaces
			idList = payrollRunsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PayrollRun		
				payrollRun = PayrollRunDelegate().get(id).first();	
				# add the PayrollRun
				payrollCalendar.payrollRuns.add(payrollRun)
				
			# save it		
			payrollCalendar.save()
			
			# reload and return the appropriate version
			return self.get( payrollCalendarId );
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except PayrollRun.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollRun does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayrollRuns( self, payrollCalendarId, payrollRunsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PayrollRunDelegate import PayrollRunDelegate

		errMsg = "Failed to remove elements " + str(payrollRunsIds) + " for PayrollRuns on PayrollCalendar"

		try:
			# get the PayrollCalendar
			payrollCalendar = self.get( payrollCalendarId ).first()
				
			# split on a comma with no spaces
			idList = payrollRunsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PayrollRun		
				payrollRun = PayrollRunDelegate().get(id).first();	
				# add the PayrollRun
				payrollCalendar.payrollRuns.remove(payrollRun)
				
			# save it		
			payrollCalendar.save()
			
			# reload and return the appropriate version
			return self.get( payrollCalendarId );
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except PayrollRun.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollRun does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmployees( self, payrollCalendarId, employeesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to add elements " + str(employeesIds) + " for Employees on PayrollCalendar"

		try:
			# get the PayrollCalendar
			payrollCalendar = self.get( payrollCalendarId ).first()
				
			# split on a comma with no spaces
			idList = employeesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				payrollCalendar.employees.add(employee)
				
			# save it		
			payrollCalendar.save()
			
			# reload and return the appropriate version
			return self.get( payrollCalendarId );
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmployees( self, payrollCalendarId, employeesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to remove elements " + str(employeesIds) + " for Employees on PayrollCalendar"

		try:
			# get the PayrollCalendar
			payrollCalendar = self.get( payrollCalendarId ).first()
				
			# split on a comma with no spaces
			idList = employeesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				payrollCalendar.employees.remove(employee)
				
			# save it		
			payrollCalendar.save()
			
			# reload and return the appropriate version
			return self.get( payrollCalendarId );
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
