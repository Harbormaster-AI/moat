from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.PayrollItem import PayrollItem
from hrOnDjango.models.PayrollRun import PayrollRun
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PayrollItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollItemDelegate Declaration
#======================================================================
class PayrollItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, payrollItemId ):
		try:	
			payrollItem = PayrollItem.objects.filter(id=payrollItemId)
			return payrollItem.first();
		except PayrollItem.DoesNotExist:
			raise ProcessingError("PayrollItem with id " + str(payrollItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, payrollItem):
		for model in serializers.deserialize("json", payrollItem):
			model.save()
			return model;

	def create(self, payrollItem):
		payrollItem.save()
		return payrollItem;

	def saveFromJson(self, payrollItem):
		for model in serializers.deserialize("json", payrollItem):
			model.save()
			return payrollItem;
	
	def save(self, payrollItem):
		payrollItem.save()
		return payrollItem;
	
	def delete(self, payrollItemId ):
		errMsg = "Failed to delete PayrollItem from db using id " + str(payrollItemId)
		
		try:
			payrollItem = PayrollItem.objects.get(id=payrollItemId)
			payrollItem.delete()
			return True
		except PayrollItem.DoesNotExist:
			raise ProcessingError("PayrollItem with id " + str(payrollItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PayrollItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PayrollItem from db")
		except Exception:
			return None;
		
	def assignPayrollRun( self, payrollItemId, payrollRunId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PayrollRunDelegate import PayrollRunDelegate

		errMsg = "Failed to assign element " + str(payrollRunId) + " for PayrollRun on PayrollItem"

		try:
			# get the PayrollItem from db
			payrollItem = self.get( payrollItemId ).first()	
			
			# get the PayrollRun from db
			payrollRun = PayrollRunDelegate().get(payrollRunId).first();
			
			# assign the PayrollRun		
			payrollItem.payrollRun = payrollRun
			
			#save it
			payrollItem.save()

			# reload and return the appropriate version					
			return self.get( payrollItemId );
		except PayrollItem.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollItem with id " + str(payrollItemId) + " does not exist.")
		except PayrollRun.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollRun with id " + str(payrollRunId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPayrollRun( self, payrollItemId ):
		errMsg = "Failed to unassign element " + str(payrollRunId) + " for PayrollRun on PayrollItem"

		try:
			# get the PayrollItem from db
			payrollItem = self.get( payrollItemId ).first()	
			
			# assign to None for unassignment
			payrollItem.payrollRun = None			

			#save it
			payrollItem.save()

			# reload and return the appropriate version					
			return self.get( payrollItemId );
		except PayrollItem.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollItem with id " + str(payrollItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEmployee( self, payrollItemId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on PayrollItem"

		try:
			# get the PayrollItem from db
			payrollItem = self.get( payrollItemId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			payrollItem.employee = employee
			
			#save it
			payrollItem.save()

			# reload and return the appropriate version					
			return self.get( payrollItemId );
		except PayrollItem.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollItem with id " + str(payrollItemId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, payrollItemId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on PayrollItem"

		try:
			# get the PayrollItem from db
			payrollItem = self.get( payrollItemId ).first()	
			
			# assign to None for unassignment
			payrollItem.employee = None			

			#save it
			payrollItem.save()

			# reload and return the appropriate version					
			return self.get( payrollItemId );
		except PayrollItem.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollItem with id " + str(payrollItemId) + " does not exist.")
		except Exception:
			return None;
		
