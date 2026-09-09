from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Condition import Condition
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Condition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConditionDelegate Declaration
#======================================================================
class ConditionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, conditionId ):
		try:	
			condition = Condition.objects.filter(id=conditionId)
			return condition.first();
		except Condition.DoesNotExist:
			raise ProcessingError("Condition with id " + str(conditionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, condition):
		for model in serializers.deserialize("json", condition):
			model.save()
			return model;

	def create(self, condition):
		condition.save()
		return condition;

	def saveFromJson(self, condition):
		for model in serializers.deserialize("json", condition):
			model.save()
			return condition;
	
	def save(self, condition):
		condition.save()
		return condition;
	
	def delete(self, conditionId ):
		errMsg = "Failed to delete Condition from db using id " + str(conditionId)
		
		try:
			condition = Condition.objects.get(id=conditionId)
			condition.delete()
			return True
		except Condition.DoesNotExist:
			raise ProcessingError("Condition with id " + str(conditionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Condition.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Condition from db")
		except Exception:
			return None;
		
	def assignPatient( self, conditionId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on Condition"

		try:
			# get the Condition from db
			condition = self.get( conditionId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			condition.patient = patient
			
			#save it
			condition.save()

			# reload and return the appropriate version					
			return self.get( conditionId );
		except Condition.DoesNotExist:
			raise ProcessingError(errMsg + " : Condition with id " + str(conditionId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, conditionId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on Condition"

		try:
			# get the Condition from db
			condition = self.get( conditionId ).first()	
			
			# assign to None for unassignment
			condition.patient = None			

			#save it
			condition.save()

			# reload and return the appropriate version					
			return self.get( conditionId );
		except Condition.DoesNotExist:
			raise ProcessingError(errMsg + " : Condition with id " + str(conditionId) + " does not exist.")
		except Exception:
			return None;
		
