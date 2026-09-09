from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.DeviceCriterion import DeviceCriterion
from advertisingOnDjango.models.TargetingProfile import TargetingProfile
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DeviceCriterion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DeviceCriterionDelegate Declaration
#======================================================================
class DeviceCriterionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, deviceCriterionId ):
		try:	
			deviceCriterion = DeviceCriterion.objects.filter(id=deviceCriterionId)
			return deviceCriterion.first();
		except DeviceCriterion.DoesNotExist:
			raise ProcessingError("DeviceCriterion with id " + str(deviceCriterionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, deviceCriterion):
		for model in serializers.deserialize("json", deviceCriterion):
			model.save()
			return model;

	def create(self, deviceCriterion):
		deviceCriterion.save()
		return deviceCriterion;

	def saveFromJson(self, deviceCriterion):
		for model in serializers.deserialize("json", deviceCriterion):
			model.save()
			return deviceCriterion;
	
	def save(self, deviceCriterion):
		deviceCriterion.save()
		return deviceCriterion;
	
	def delete(self, deviceCriterionId ):
		errMsg = "Failed to delete DeviceCriterion from db using id " + str(deviceCriterionId)
		
		try:
			deviceCriterion = DeviceCriterion.objects.get(id=deviceCriterionId)
			deviceCriterion.delete()
			return True
		except DeviceCriterion.DoesNotExist:
			raise ProcessingError("DeviceCriterion with id " + str(deviceCriterionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DeviceCriterion.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DeviceCriterion from db")
		except Exception:
			return None;
		
	def assignTargetingProfile( self, deviceCriterionId, targetingProfileId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TargetingProfileDelegate import TargetingProfileDelegate

		errMsg = "Failed to assign element " + str(targetingProfileId) + " for TargetingProfile on DeviceCriterion"

		try:
			# get the DeviceCriterion from db
			deviceCriterion = self.get( deviceCriterionId ).first()	
			
			# get the TargetingProfile from db
			targetingProfile = TargetingProfileDelegate().get(targetingProfileId).first();
			
			# assign the TargetingProfile		
			deviceCriterion.targetingProfile = targetingProfile
			
			#save it
			deviceCriterion.save()

			# reload and return the appropriate version					
			return self.get( deviceCriterionId );
		except DeviceCriterion.DoesNotExist:
			raise ProcessingError(errMsg + " : DeviceCriterion with id " + str(deviceCriterionId) + " does not exist.")
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTargetingProfile( self, deviceCriterionId ):
		errMsg = "Failed to unassign element " + str(targetingProfileId) + " for TargetingProfile on DeviceCriterion"

		try:
			# get the DeviceCriterion from db
			deviceCriterion = self.get( deviceCriterionId ).first()	
			
			# assign to None for unassignment
			deviceCriterion.targetingProfile = None			

			#save it
			deviceCriterion.save()

			# reload and return the appropriate version					
			return self.get( deviceCriterionId );
		except DeviceCriterion.DoesNotExist:
			raise ProcessingError(errMsg + " : DeviceCriterion with id " + str(deviceCriterionId) + " does not exist.")
		except Exception:
			return None;
		
