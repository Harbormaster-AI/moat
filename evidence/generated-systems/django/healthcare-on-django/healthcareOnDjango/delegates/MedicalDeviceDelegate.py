from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.MedicalDevice import MedicalDevice
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.Observation import Observation
from healthcareOnDjango.models.SoftwareUpdate import SoftwareUpdate
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MedicalDevice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicalDeviceDelegate Declaration
#======================================================================
class MedicalDeviceDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, medicalDeviceId ):
		try:	
			medicalDevice = MedicalDevice.objects.filter(id=medicalDeviceId)
			return medicalDevice.first();
		except MedicalDevice.DoesNotExist:
			raise ProcessingError("MedicalDevice with id " + str(medicalDeviceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, medicalDevice):
		for model in serializers.deserialize("json", medicalDevice):
			model.save()
			return model;

	def create(self, medicalDevice):
		medicalDevice.save()
		return medicalDevice;

	def saveFromJson(self, medicalDevice):
		for model in serializers.deserialize("json", medicalDevice):
			model.save()
			return medicalDevice;
	
	def save(self, medicalDevice):
		medicalDevice.save()
		return medicalDevice;
	
	def delete(self, medicalDeviceId ):
		errMsg = "Failed to delete MedicalDevice from db using id " + str(medicalDeviceId)
		
		try:
			medicalDevice = MedicalDevice.objects.get(id=medicalDeviceId)
			medicalDevice.delete()
			return True
		except MedicalDevice.DoesNotExist:
			raise ProcessingError("MedicalDevice with id " + str(medicalDeviceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MedicalDevice.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MedicalDevice from db")
		except Exception:
			return None;
		
	def assignPatient( self, medicalDeviceId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on MedicalDevice"

		try:
			# get the MedicalDevice from db
			medicalDevice = self.get( medicalDeviceId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			medicalDevice.patient = patient
			
			#save it
			medicalDevice.save()

			# reload and return the appropriate version					
			return self.get( medicalDeviceId );
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice with id " + str(medicalDeviceId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, medicalDeviceId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on MedicalDevice"

		try:
			# get the MedicalDevice from db
			medicalDevice = self.get( medicalDeviceId ).first()	
			
			# assign to None for unassignment
			medicalDevice.patient = None			

			#save it
			medicalDevice.save()

			# reload and return the appropriate version					
			return self.get( medicalDeviceId );
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice with id " + str(medicalDeviceId) + " does not exist.")
		except Exception:
			return None;
		
	def addObservations( self, medicalDeviceId, observationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

		errMsg = "Failed to add elements " + str(observationsIds) + " for Observations on MedicalDevice"

		try:
			# get the MedicalDevice
			medicalDevice = self.get( medicalDeviceId ).first()
				
			# split on a comma with no spaces
			idList = observationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Observation		
				observation = ObservationDelegate().get(id).first();	
				# add the Observation
				medicalDevice.observations.add(observation)
				
			# save it		
			medicalDevice.save()
			
			# reload and return the appropriate version
			return self.get( medicalDeviceId );
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice with id " + str(medicalDeviceId) + " does not exist.")
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeObservations( self, medicalDeviceId, observationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

		errMsg = "Failed to remove elements " + str(observationsIds) + " for Observations on MedicalDevice"

		try:
			# get the MedicalDevice
			medicalDevice = self.get( medicalDeviceId ).first()
				
			# split on a comma with no spaces
			idList = observationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Observation		
				observation = ObservationDelegate().get(id).first();	
				# add the Observation
				medicalDevice.observations.remove(observation)
				
			# save it		
			medicalDevice.save()
			
			# reload and return the appropriate version
			return self.get( medicalDeviceId );
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice with id " + str(medicalDeviceId) + " does not exist.")
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSoftwareUpdates( self, medicalDeviceId, softwareUpdatesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.SoftwareUpdateDelegate import SoftwareUpdateDelegate

		errMsg = "Failed to add elements " + str(softwareUpdatesIds) + " for SoftwareUpdates on MedicalDevice"

		try:
			# get the MedicalDevice
			medicalDevice = self.get( medicalDeviceId ).first()
				
			# split on a comma with no spaces
			idList = softwareUpdatesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SoftwareUpdate		
				softwareUpdate = SoftwareUpdateDelegate().get(id).first();	
				# add the SoftwareUpdate
				medicalDevice.softwareUpdates.add(softwareUpdate)
				
			# save it		
			medicalDevice.save()
			
			# reload and return the appropriate version
			return self.get( medicalDeviceId );
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice with id " + str(medicalDeviceId) + " does not exist.")
		except SoftwareUpdate.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareUpdate does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSoftwareUpdates( self, medicalDeviceId, softwareUpdatesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.SoftwareUpdateDelegate import SoftwareUpdateDelegate

		errMsg = "Failed to remove elements " + str(softwareUpdatesIds) + " for SoftwareUpdates on MedicalDevice"

		try:
			# get the MedicalDevice
			medicalDevice = self.get( medicalDeviceId ).first()
				
			# split on a comma with no spaces
			idList = softwareUpdatesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SoftwareUpdate		
				softwareUpdate = SoftwareUpdateDelegate().get(id).first();	
				# add the SoftwareUpdate
				medicalDevice.softwareUpdates.remove(softwareUpdate)
				
			# save it		
			medicalDevice.save()
			
			# reload and return the appropriate version
			return self.get( medicalDeviceId );
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice with id " + str(medicalDeviceId) + " does not exist.")
		except SoftwareUpdate.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareUpdate does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
