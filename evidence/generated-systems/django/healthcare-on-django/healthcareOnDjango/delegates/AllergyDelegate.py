from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Allergy import Allergy
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Allergy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AllergyDelegate Declaration
#======================================================================
class AllergyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, allergyId ):
		try:	
			allergy = Allergy.objects.filter(id=allergyId)
			return allergy.first();
		except Allergy.DoesNotExist:
			raise ProcessingError("Allergy with id " + str(allergyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, allergy):
		for model in serializers.deserialize("json", allergy):
			model.save()
			return model;

	def create(self, allergy):
		allergy.save()
		return allergy;

	def saveFromJson(self, allergy):
		for model in serializers.deserialize("json", allergy):
			model.save()
			return allergy;
	
	def save(self, allergy):
		allergy.save()
		return allergy;
	
	def delete(self, allergyId ):
		errMsg = "Failed to delete Allergy from db using id " + str(allergyId)
		
		try:
			allergy = Allergy.objects.get(id=allergyId)
			allergy.delete()
			return True
		except Allergy.DoesNotExist:
			raise ProcessingError("Allergy with id " + str(allergyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Allergy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Allergy from db")
		except Exception:
			return None;
		
	def assignPatient( self, allergyId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on Allergy"

		try:
			# get the Allergy from db
			allergy = self.get( allergyId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			allergy.patient = patient
			
			#save it
			allergy.save()

			# reload and return the appropriate version					
			return self.get( allergyId );
		except Allergy.DoesNotExist:
			raise ProcessingError(errMsg + " : Allergy with id " + str(allergyId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, allergyId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on Allergy"

		try:
			# get the Allergy from db
			allergy = self.get( allergyId ).first()	
			
			# assign to None for unassignment
			allergy.patient = None			

			#save it
			allergy.save()

			# reload and return the appropriate version					
			return self.get( allergyId );
		except Allergy.DoesNotExist:
			raise ProcessingError(errMsg + " : Allergy with id " + str(allergyId) + " does not exist.")
		except Exception:
			return None;
		
