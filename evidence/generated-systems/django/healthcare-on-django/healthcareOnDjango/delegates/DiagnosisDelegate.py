from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Diagnosis import Diagnosis
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Diagnosis
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DiagnosisDelegate Declaration
#======================================================================
class DiagnosisDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, diagnosisId ):
		try:	
			diagnosis = Diagnosis.objects.filter(id=diagnosisId)
			return diagnosis.first();
		except Diagnosis.DoesNotExist:
			raise ProcessingError("Diagnosis with id " + str(diagnosisId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, diagnosis):
		for model in serializers.deserialize("json", diagnosis):
			model.save()
			return model;

	def create(self, diagnosis):
		diagnosis.save()
		return diagnosis;

	def saveFromJson(self, diagnosis):
		for model in serializers.deserialize("json", diagnosis):
			model.save()
			return diagnosis;
	
	def save(self, diagnosis):
		diagnosis.save()
		return diagnosis;
	
	def delete(self, diagnosisId ):
		errMsg = "Failed to delete Diagnosis from db using id " + str(diagnosisId)
		
		try:
			diagnosis = Diagnosis.objects.get(id=diagnosisId)
			diagnosis.delete()
			return True
		except Diagnosis.DoesNotExist:
			raise ProcessingError("Diagnosis with id " + str(diagnosisId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Diagnosis.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Diagnosis from db")
		except Exception:
			return None;
		
	def assignEncounter( self, diagnosisId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on Diagnosis"

		try:
			# get the Diagnosis from db
			diagnosis = self.get( diagnosisId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			diagnosis.encounter = encounter
			
			#save it
			diagnosis.save()

			# reload and return the appropriate version					
			return self.get( diagnosisId );
		except Diagnosis.DoesNotExist:
			raise ProcessingError(errMsg + " : Diagnosis with id " + str(diagnosisId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, diagnosisId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on Diagnosis"

		try:
			# get the Diagnosis from db
			diagnosis = self.get( diagnosisId ).first()	
			
			# assign to None for unassignment
			diagnosis.encounter = None			

			#save it
			diagnosis.save()

			# reload and return the appropriate version					
			return self.get( diagnosisId );
		except Diagnosis.DoesNotExist:
			raise ProcessingError(errMsg + " : Diagnosis with id " + str(diagnosisId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPatient( self, diagnosisId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on Diagnosis"

		try:
			# get the Diagnosis from db
			diagnosis = self.get( diagnosisId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			diagnosis.patient = patient
			
			#save it
			diagnosis.save()

			# reload and return the appropriate version					
			return self.get( diagnosisId );
		except Diagnosis.DoesNotExist:
			raise ProcessingError(errMsg + " : Diagnosis with id " + str(diagnosisId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, diagnosisId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on Diagnosis"

		try:
			# get the Diagnosis from db
			diagnosis = self.get( diagnosisId ).first()	
			
			# assign to None for unassignment
			diagnosis.patient = None			

			#save it
			diagnosis.save()

			# reload and return the appropriate version					
			return self.get( diagnosisId );
		except Diagnosis.DoesNotExist:
			raise ProcessingError(errMsg + " : Diagnosis with id " + str(diagnosisId) + " does not exist.")
		except Exception:
			return None;
		
