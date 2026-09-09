from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.PayrollRun import PayrollRun
from hrOnDjango.models.PayrollCalendar import PayrollCalendar
from hrOnDjango.models.PayrollItem import PayrollItem
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PayrollRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollRunDelegate Declaration
#======================================================================
class PayrollRunDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, payrollRunId ):
		try:	
			payrollRun = PayrollRun.objects.filter(id=payrollRunId)
			return payrollRun.first();
		except PayrollRun.DoesNotExist:
			raise ProcessingError("PayrollRun with id " + str(payrollRunId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, payrollRun):
		for model in serializers.deserialize("json", payrollRun):
			model.save()
			return model;

	def create(self, payrollRun):
		payrollRun.save()
		return payrollRun;

	def saveFromJson(self, payrollRun):
		for model in serializers.deserialize("json", payrollRun):
			model.save()
			return payrollRun;
	
	def save(self, payrollRun):
		payrollRun.save()
		return payrollRun;
	
	def delete(self, payrollRunId ):
		errMsg = "Failed to delete PayrollRun from db using id " + str(payrollRunId)
		
		try:
			payrollRun = PayrollRun.objects.get(id=payrollRunId)
			payrollRun.delete()
			return True
		except PayrollRun.DoesNotExist:
			raise ProcessingError("PayrollRun with id " + str(payrollRunId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PayrollRun.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PayrollRun from db")
		except Exception:
			return None;
		
	def assignPayrollCalendar( self, payrollRunId, payrollCalendarId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PayrollCalendarDelegate import PayrollCalendarDelegate

		errMsg = "Failed to assign element " + str(payrollCalendarId) + " for PayrollCalendar on PayrollRun"

		try:
			# get the PayrollRun from db
			payrollRun = self.get( payrollRunId ).first()	
			
			# get the PayrollCalendar from db
			payrollCalendar = PayrollCalendarDelegate().get(payrollCalendarId).first();
			
			# assign the PayrollCalendar		
			payrollRun.payrollCalendar = payrollCalendar
			
			#save it
			payrollRun.save()

			# reload and return the appropriate version					
			return self.get( payrollRunId );
		except PayrollRun.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollRun with id " + str(payrollRunId) + " does not exist.")
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar with id " + str(payrollCalendarId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPayrollCalendar( self, payrollRunId ):
		errMsg = "Failed to unassign element " + str(payrollCalendarId) + " for PayrollCalendar on PayrollRun"

		try:
			# get the PayrollRun from db
			payrollRun = self.get( payrollRunId ).first()	
			
			# assign to None for unassignment
			payrollRun.payrollCalendar = None			

			#save it
			payrollRun.save()

			# reload and return the appropriate version					
			return self.get( payrollRunId );
		except PayrollRun.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollRun with id " + str(payrollRunId) + " does not exist.")
		except Exception:
			return None;
		
	def addPayrollItems( self, payrollRunId, payrollItemsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PayrollItemDelegate import PayrollItemDelegate

		errMsg = "Failed to add elements " + str(payrollItemsIds) + " for PayrollItems on PayrollRun"

		try:
			# get the PayrollRun
			payrollRun = self.get( payrollRunId ).first()
				
			# split on a comma with no spaces
			idList = payrollItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PayrollItem		
				payrollItem = PayrollItemDelegate().get(id).first();	
				# add the PayrollItem
				payrollRun.payrollItems.add(payrollItem)
				
			# save it		
			payrollRun.save()
			
			# reload and return the appropriate version
			return self.get( payrollRunId );
		except PayrollRun.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollRun with id " + str(payrollRunId) + " does not exist.")
		except PayrollItem.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayrollItems( self, payrollRunId, payrollItemsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PayrollItemDelegate import PayrollItemDelegate

		errMsg = "Failed to remove elements " + str(payrollItemsIds) + " for PayrollItems on PayrollRun"

		try:
			# get the PayrollRun
			payrollRun = self.get( payrollRunId ).first()
				
			# split on a comma with no spaces
			idList = payrollItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PayrollItem		
				payrollItem = PayrollItemDelegate().get(id).first();	
				# add the PayrollItem
				payrollRun.payrollItems.remove(payrollItem)
				
			# save it		
			payrollRun.save()
			
			# reload and return the appropriate version
			return self.get( payrollRunId );
		except PayrollRun.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollRun with id " + str(payrollRunId) + " does not exist.")
		except PayrollItem.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
