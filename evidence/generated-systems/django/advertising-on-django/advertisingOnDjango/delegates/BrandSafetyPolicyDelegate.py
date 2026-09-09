from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.BrandSafetyPolicy import BrandSafetyPolicy
from advertisingOnDjango.models.TargetingProfile import TargetingProfile
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BrandSafetyPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BrandSafetyPolicyDelegate Declaration
#======================================================================
class BrandSafetyPolicyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, brandSafetyPolicyId ):
		try:	
			brandSafetyPolicy = BrandSafetyPolicy.objects.filter(id=brandSafetyPolicyId)
			return brandSafetyPolicy.first();
		except BrandSafetyPolicy.DoesNotExist:
			raise ProcessingError("BrandSafetyPolicy with id " + str(brandSafetyPolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, brandSafetyPolicy):
		for model in serializers.deserialize("json", brandSafetyPolicy):
			model.save()
			return model;

	def create(self, brandSafetyPolicy):
		brandSafetyPolicy.save()
		return brandSafetyPolicy;

	def saveFromJson(self, brandSafetyPolicy):
		for model in serializers.deserialize("json", brandSafetyPolicy):
			model.save()
			return brandSafetyPolicy;
	
	def save(self, brandSafetyPolicy):
		brandSafetyPolicy.save()
		return brandSafetyPolicy;
	
	def delete(self, brandSafetyPolicyId ):
		errMsg = "Failed to delete BrandSafetyPolicy from db using id " + str(brandSafetyPolicyId)
		
		try:
			brandSafetyPolicy = BrandSafetyPolicy.objects.get(id=brandSafetyPolicyId)
			brandSafetyPolicy.delete()
			return True
		except BrandSafetyPolicy.DoesNotExist:
			raise ProcessingError("BrandSafetyPolicy with id " + str(brandSafetyPolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BrandSafetyPolicy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BrandSafetyPolicy from db")
		except Exception:
			return None;
		
	def addTargetingProfiles( self, brandSafetyPolicyId, targetingProfilesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TargetingProfileDelegate import TargetingProfileDelegate

		errMsg = "Failed to add elements " + str(targetingProfilesIds) + " for TargetingProfiles on BrandSafetyPolicy"

		try:
			# get the BrandSafetyPolicy
			brandSafetyPolicy = self.get( brandSafetyPolicyId ).first()
				
			# split on a comma with no spaces
			idList = targetingProfilesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TargetingProfile		
				targetingProfile = TargetingProfileDelegate().get(id).first();	
				# add the TargetingProfile
				brandSafetyPolicy.targetingProfiles.add(targetingProfile)
				
			# save it		
			brandSafetyPolicy.save()
			
			# reload and return the appropriate version
			return self.get( brandSafetyPolicyId );
		except BrandSafetyPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : BrandSafetyPolicy with id " + str(brandSafetyPolicyId) + " does not exist.")
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTargetingProfiles( self, brandSafetyPolicyId, targetingProfilesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TargetingProfileDelegate import TargetingProfileDelegate

		errMsg = "Failed to remove elements " + str(targetingProfilesIds) + " for TargetingProfiles on BrandSafetyPolicy"

		try:
			# get the BrandSafetyPolicy
			brandSafetyPolicy = self.get( brandSafetyPolicyId ).first()
				
			# split on a comma with no spaces
			idList = targetingProfilesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TargetingProfile		
				targetingProfile = TargetingProfileDelegate().get(id).first();	
				# add the TargetingProfile
				brandSafetyPolicy.targetingProfiles.remove(targetingProfile)
				
			# save it		
			brandSafetyPolicy.save()
			
			# reload and return the appropriate version
			return self.get( brandSafetyPolicyId );
		except BrandSafetyPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : BrandSafetyPolicy with id " + str(brandSafetyPolicyId) + " does not exist.")
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
