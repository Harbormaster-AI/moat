from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.CorrectiveAction import CorrectiveAction
from manufacturingOnDjango.models.Nonconformance import Nonconformance
from manufacturingOnDjango.models.Employee import Employee
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CorrectiveAction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CorrectiveActionDelegate Declaration
#======================================================================
class CorrectiveActionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, correctiveActionId ):
		try:	
			correctiveAction = CorrectiveAction.objects.filter(id=correctiveActionId)
			return correctiveAction.first();
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError("CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, correctiveAction):
		for model in serializers.deserialize("json", correctiveAction):
			model.save()
			return model;

	def create(self, correctiveAction):
		correctiveAction.save()
		return correctiveAction;

	def saveFromJson(self, correctiveAction):
		for model in serializers.deserialize("json", correctiveAction):
			model.save()
			return correctiveAction;
	
	def save(self, correctiveAction):
		correctiveAction.save()
		return correctiveAction;
	
	def delete(self, correctiveActionId ):
		errMsg = "Failed to delete CorrectiveAction from db using id " + str(correctiveActionId)
		
		try:
			correctiveAction = CorrectiveAction.objects.get(id=correctiveActionId)
			correctiveAction.delete()
			return True
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError("CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CorrectiveAction.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CorrectiveAction from db")
		except Exception:
			return None;
		
	def assignNonconformance( self, correctiveActionId, nonconformanceId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.NonconformanceDelegate import NonconformanceDelegate

		errMsg = "Failed to assign element " + str(nonconformanceId) + " for Nonconformance on CorrectiveAction"

		try:
			# get the CorrectiveAction from db
			correctiveAction = self.get( correctiveActionId ).first()	
			
			# get the Nonconformance from db
			nonconformance = NonconformanceDelegate().get(nonconformanceId).first();
			
			# assign the Nonconformance		
			correctiveAction.nonconformance = nonconformance
			
			#save it
			correctiveAction.save()

			# reload and return the appropriate version					
			return self.get( correctiveActionId );
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except Nonconformance.DoesNotExist:
			raise ProcessingError(errMsg + " : Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignNonconformance( self, correctiveActionId ):
		errMsg = "Failed to unassign element " + str(nonconformanceId) + " for Nonconformance on CorrectiveAction"

		try:
			# get the CorrectiveAction from db
			correctiveAction = self.get( correctiveActionId ).first()	
			
			# assign to None for unassignment
			correctiveAction.nonconformance = None			

			#save it
			correctiveAction.save()

			# reload and return the appropriate version					
			return self.get( correctiveActionId );
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, correctiveActionId, ownerId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on CorrectiveAction"

		try:
			# get the CorrectiveAction from db
			correctiveAction = self.get( correctiveActionId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(ownerId).first();
			
			# assign the Owner		
			correctiveAction.owner = employee
			
			#save it
			correctiveAction.save()

			# reload and return the appropriate version					
			return self.get( correctiveActionId );
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, correctiveActionId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on CorrectiveAction"

		try:
			# get the CorrectiveAction from db
			correctiveAction = self.get( correctiveActionId ).first()	
			
			# assign to None for unassignment
			correctiveAction.employee = None			

			#save it
			correctiveAction.save()

			# reload and return the appropriate version					
			return self.get( correctiveActionId );
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except Exception:
			return None;
		
