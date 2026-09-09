from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.SoftwareUpdate import SoftwareUpdate
from healthcareOnDjango.models.MedicalDevice import MedicalDevice
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SoftwareUpdate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareUpdateDelegate Declaration
#======================================================================
class SoftwareUpdateDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, softwareUpdateId ):
		try:	
			softwareUpdate = SoftwareUpdate.objects.filter(id=softwareUpdateId)
			return softwareUpdate.first();
		except SoftwareUpdate.DoesNotExist:
			raise ProcessingError("SoftwareUpdate with id " + str(softwareUpdateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, softwareUpdate):
		for model in serializers.deserialize("json", softwareUpdate):
			model.save()
			return model;

	def create(self, softwareUpdate):
		softwareUpdate.save()
		return softwareUpdate;

	def saveFromJson(self, softwareUpdate):
		for model in serializers.deserialize("json", softwareUpdate):
			model.save()
			return softwareUpdate;
	
	def save(self, softwareUpdate):
		softwareUpdate.save()
		return softwareUpdate;
	
	def delete(self, softwareUpdateId ):
		errMsg = "Failed to delete SoftwareUpdate from db using id " + str(softwareUpdateId)
		
		try:
			softwareUpdate = SoftwareUpdate.objects.get(id=softwareUpdateId)
			softwareUpdate.delete()
			return True
		except SoftwareUpdate.DoesNotExist:
			raise ProcessingError("SoftwareUpdate with id " + str(softwareUpdateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SoftwareUpdate.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SoftwareUpdate from db")
		except Exception:
			return None;
		
	def assignDevice( self, softwareUpdateId, deviceId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicalDeviceDelegate import MedicalDeviceDelegate

		errMsg = "Failed to assign element " + str(deviceId) + " for Device on SoftwareUpdate"

		try:
			# get the SoftwareUpdate from db
			softwareUpdate = self.get( softwareUpdateId ).first()	
			
			# get the MedicalDevice from db
			medicalDevice = MedicalDeviceDelegate().get(deviceId).first();
			
			# assign the Device		
			softwareUpdate.device = medicalDevice
			
			#save it
			softwareUpdate.save()

			# reload and return the appropriate version					
			return self.get( softwareUpdateId );
		except SoftwareUpdate.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareUpdate with id " + str(softwareUpdateId) + " does not exist.")
		except MedicalDevice.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalDevice with id " + str(deviceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDevice( self, softwareUpdateId ):
		errMsg = "Failed to unassign element " + str(deviceId) + " for Device on SoftwareUpdate"

		try:
			# get the SoftwareUpdate from db
			softwareUpdate = self.get( softwareUpdateId ).first()	
			
			# assign to None for unassignment
			softwareUpdate.medicalDevice = None			

			#save it
			softwareUpdate.save()

			# reload and return the appropriate version					
			return self.get( softwareUpdateId );
		except SoftwareUpdate.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareUpdate with id " + str(softwareUpdateId) + " does not exist.")
		except Exception:
			return None;
		
