from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Observation import Observation
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.MedicalDevice import MedicalDevice
from healthcareOnDjango.models.LabResult import LabResult
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Observation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ObservationDelegate Declaration
#======================================================================
class ObservationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, observationId ):
		try:	
			observation = Observation.objects.filter(id=observationId)
			return observation.first();
		except Observation.DoesNotExist:
			raise ProcessingError("Observation with id " + str(observationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, observation):
		for model in serializers.deserialize("json", observation):
			model.save()
			return model;

	def create(self, observation):
		observation.save()
		return observation;

	def saveFromJson(self, observation):
		for model in serializers.deserialize("json", observation):
			model.save()
			return observation;
	
	def save(self, observation):
		observation.save()
		return observation;
	
	def delete(self, observationId ):
		errMsg = "Failed to delete Observation from db using id " + str(observationId)
		
		try:
			observation = Observation.objects.get(id=observationId)
			observation.delete()
			return True
		except Observation.DoesNotExist:
			raise ProcessingError("Observation with id " + str(observationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Observation.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Observation from db")
		except Exception:
			return None;
		
	def assignEncounter( self, observationId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on Observation"

		try:
			# get the Observation from db
			observation = self.get( observationId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			observation.encounter = encounter
			
			#save it
			observation.save()

			# reload and return the appropriate version					
			return self.get( observationId );
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation with id " + str(observationId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, observationId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on Observation"

		try:
			# get the Observation from db
			observation = self.get( observationId ).first()	
			
			# assign to None for unassignment
			observation.encounter = None			

			#save it
			observation.save()

			# reload and return the appropriate version					
			return self.get( observationId );
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation with id " + str(observationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPatient( self, observationId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on Observation"

		try:
			# get the Observation from db
			observation = self.get( observationId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			observation.patient = patient
			
			#save it
			observation.save()

			# reload and return the appropriate version					
			return self.get( observationId );
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation with id " + str(observationId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, observationId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on Observation"

		try:
			# get the Observation from db
			observation = self.get( observationId ).first()	
			
			# assign to None for unassignment
			observation.patient = None			

			#save it
			observation.save()

			# reload and return the appropriate version					
			return self.get( observationId );
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation with id " + str(observationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDevice( self, observationId, deviceId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicalDeviceDelegate import MedicalDeviceDelegate

		errMsg = "Failed to assign element " + str(deviceId) + " for Device on Observation"

		try:
			# get the Observation from db
			observation = self.get( observationId ).first()	
			
			# get the MedicalDevice from db
			medicalDevice = MedicalDeviceDelegate().get(deviceId).first();
			
			# assign the Device		
			observation.device = medicalDevice
			
			#save it
			observation.save()

			# reload and return the appropriate version					
			return self.get( observationId );
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation with id " + str(observationId) + " does not exist.")
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice with id " + str(deviceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDevice( self, observationId ):
		errMsg = "Failed to unassign element " + str(deviceId) + " for Device on Observation"

		try:
			# get the Observation from db
			observation = self.get( observationId ).first()	
			
			# assign to None for unassignment
			observation.medicalDevice = None			

			#save it
			observation.save()

			# reload and return the appropriate version					
			return self.get( observationId );
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation with id " + str(observationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLabResult( self, observationId, labResultId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LabResultDelegate import LabResultDelegate

		errMsg = "Failed to assign element " + str(labResultId) + " for LabResult on Observation"

		try:
			# get the Observation from db
			observation = self.get( observationId ).first()	
			
			# get the LabResult from db
			labResult = LabResultDelegate().get(labResultId).first();
			
			# assign the LabResult		
			observation.labResult = labResult
			
			#save it
			observation.save()

			# reload and return the appropriate version					
			return self.get( observationId );
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation with id " + str(observationId) + " does not exist.")
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult with id " + str(labResultId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLabResult( self, observationId ):
		errMsg = "Failed to unassign element " + str(labResultId) + " for LabResult on Observation"

		try:
			# get the Observation from db
			observation = self.get( observationId ).first()	
			
			# assign to None for unassignment
			observation.labResult = None			

			#save it
			observation.save()

			# reload and return the appropriate version					
			return self.get( observationId );
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation with id " + str(observationId) + " does not exist.")
		except Exception:
			return None;
		
