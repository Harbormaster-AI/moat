from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.TargetingProfile import TargetingProfile
from advertisingOnDjango.models.AudienceSegment import AudienceSegment
from advertisingOnDjango.models.GeoRegion import GeoRegion
from advertisingOnDjango.models.ContentCategory import ContentCategory
from advertisingOnDjango.models.BrandSafetyPolicy import BrandSafetyPolicy
from advertisingOnDjango.models.DeviceCriterion import DeviceCriterion
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TargetingProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TargetingProfileDelegate Declaration
#======================================================================
class TargetingProfileDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, targetingProfileId ):
		try:	
			targetingProfile = TargetingProfile.objects.filter(id=targetingProfileId)
			return targetingProfile.first();
		except TargetingProfile.DoesNotExist:
			raise ProcessingError("TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, targetingProfile):
		for model in serializers.deserialize("json", targetingProfile):
			model.save()
			return model;

	def create(self, targetingProfile):
		targetingProfile.save()
		return targetingProfile;

	def saveFromJson(self, targetingProfile):
		for model in serializers.deserialize("json", targetingProfile):
			model.save()
			return targetingProfile;
	
	def save(self, targetingProfile):
		targetingProfile.save()
		return targetingProfile;
	
	def delete(self, targetingProfileId ):
		errMsg = "Failed to delete TargetingProfile from db using id " + str(targetingProfileId)
		
		try:
			targetingProfile = TargetingProfile.objects.get(id=targetingProfileId)
			targetingProfile.delete()
			return True
		except TargetingProfile.DoesNotExist:
			raise ProcessingError("TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TargetingProfile.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TargetingProfile from db")
		except Exception:
			return None;
		
	def assignBrandSafetyPolicy( self, targetingProfileId, brandSafetyPolicyId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.BrandSafetyPolicyDelegate import BrandSafetyPolicyDelegate

		errMsg = "Failed to assign element " + str(brandSafetyPolicyId) + " for BrandSafetyPolicy on TargetingProfile"

		try:
			# get the TargetingProfile from db
			targetingProfile = self.get( targetingProfileId ).first()	
			
			# get the BrandSafetyPolicy from db
			brandSafetyPolicy = BrandSafetyPolicyDelegate().get(brandSafetyPolicyId).first();
			
			# assign the BrandSafetyPolicy		
			targetingProfile.brandSafetyPolicy = brandSafetyPolicy
			
			#save it
			targetingProfile.save()

			# reload and return the appropriate version					
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except BrandSafetyPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : BrandSafetyPolicy with id " + str(brandSafetyPolicyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBrandSafetyPolicy( self, targetingProfileId ):
		errMsg = "Failed to unassign element " + str(brandSafetyPolicyId) + " for BrandSafetyPolicy on TargetingProfile"

		try:
			# get the TargetingProfile from db
			targetingProfile = self.get( targetingProfileId ).first()	
			
			# assign to None for unassignment
			targetingProfile.brandSafetyPolicy = None			

			#save it
			targetingProfile.save()

			# reload and return the appropriate version					
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except Exception:
			return None;
		
	def addAudienceSegments( self, targetingProfileId, audienceSegmentsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AudienceSegmentDelegate import AudienceSegmentDelegate

		errMsg = "Failed to add elements " + str(audienceSegmentsIds) + " for AudienceSegments on TargetingProfile"

		try:
			# get the TargetingProfile
			targetingProfile = self.get( targetingProfileId ).first()
				
			# split on a comma with no spaces
			idList = audienceSegmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AudienceSegment		
				audienceSegment = AudienceSegmentDelegate().get(id).first();	
				# add the AudienceSegment
				targetingProfile.audienceSegments.add(audienceSegment)
				
			# save it		
			targetingProfile.save()
			
			# reload and return the appropriate version
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAudienceSegments( self, targetingProfileId, audienceSegmentsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AudienceSegmentDelegate import AudienceSegmentDelegate

		errMsg = "Failed to remove elements " + str(audienceSegmentsIds) + " for AudienceSegments on TargetingProfile"

		try:
			# get the TargetingProfile
			targetingProfile = self.get( targetingProfileId ).first()
				
			# split on a comma with no spaces
			idList = audienceSegmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AudienceSegment		
				audienceSegment = AudienceSegmentDelegate().get(id).first();	
				# add the AudienceSegment
				targetingProfile.audienceSegments.remove(audienceSegment)
				
			# save it		
			targetingProfile.save()
			
			# reload and return the appropriate version
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGeoRegions( self, targetingProfileId, geoRegionsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.GeoRegionDelegate import GeoRegionDelegate

		errMsg = "Failed to add elements " + str(geoRegionsIds) + " for GeoRegions on TargetingProfile"

		try:
			# get the TargetingProfile
			targetingProfile = self.get( targetingProfileId ).first()
				
			# split on a comma with no spaces
			idList = geoRegionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the GeoRegion		
				geoRegion = GeoRegionDelegate().get(id).first();	
				# add the GeoRegion
				targetingProfile.geoRegions.add(geoRegion)
				
			# save it		
			targetingProfile.save()
			
			# reload and return the appropriate version
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except GeoRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : GeoRegion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGeoRegions( self, targetingProfileId, geoRegionsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.GeoRegionDelegate import GeoRegionDelegate

		errMsg = "Failed to remove elements " + str(geoRegionsIds) + " for GeoRegions on TargetingProfile"

		try:
			# get the TargetingProfile
			targetingProfile = self.get( targetingProfileId ).first()
				
			# split on a comma with no spaces
			idList = geoRegionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the GeoRegion		
				geoRegion = GeoRegionDelegate().get(id).first();	
				# add the GeoRegion
				targetingProfile.geoRegions.remove(geoRegion)
				
			# save it		
			targetingProfile.save()
			
			# reload and return the appropriate version
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except GeoRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : GeoRegion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContentCategories( self, targetingProfileId, contentCategoriesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.ContentCategoryDelegate import ContentCategoryDelegate

		errMsg = "Failed to add elements " + str(contentCategoriesIds) + " for ContentCategories on TargetingProfile"

		try:
			# get the TargetingProfile
			targetingProfile = self.get( targetingProfileId ).first()
				
			# split on a comma with no spaces
			idList = contentCategoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ContentCategory		
				contentCategory = ContentCategoryDelegate().get(id).first();	
				# add the ContentCategory
				targetingProfile.contentCategories.add(contentCategory)
				
			# save it		
			targetingProfile.save()
			
			# reload and return the appropriate version
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except ContentCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : ContentCategory does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContentCategories( self, targetingProfileId, contentCategoriesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.ContentCategoryDelegate import ContentCategoryDelegate

		errMsg = "Failed to remove elements " + str(contentCategoriesIds) + " for ContentCategories on TargetingProfile"

		try:
			# get the TargetingProfile
			targetingProfile = self.get( targetingProfileId ).first()
				
			# split on a comma with no spaces
			idList = contentCategoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ContentCategory		
				contentCategory = ContentCategoryDelegate().get(id).first();	
				# add the ContentCategory
				targetingProfile.contentCategories.remove(contentCategory)
				
			# save it		
			targetingProfile.save()
			
			# reload and return the appropriate version
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except ContentCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : ContentCategory does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDeviceCriteria( self, targetingProfileId, deviceCriteriaIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DeviceCriterionDelegate import DeviceCriterionDelegate

		errMsg = "Failed to add elements " + str(deviceCriteriaIds) + " for DeviceCriteria on TargetingProfile"

		try:
			# get the TargetingProfile
			targetingProfile = self.get( targetingProfileId ).first()
				
			# split on a comma with no spaces
			idList = deviceCriteriaIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DeviceCriterion		
				deviceCriterion = DeviceCriterionDelegate().get(id).first();	
				# add the DeviceCriterion
				targetingProfile.deviceCriteria.add(deviceCriterion)
				
			# save it		
			targetingProfile.save()
			
			# reload and return the appropriate version
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except DeviceCriterion.DoesNotExist:
			raise ProcessingError(errMsg + " : DeviceCriterion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDeviceCriteria( self, targetingProfileId, deviceCriteriaIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DeviceCriterionDelegate import DeviceCriterionDelegate

		errMsg = "Failed to remove elements " + str(deviceCriteriaIds) + " for DeviceCriteria on TargetingProfile"

		try:
			# get the TargetingProfile
			targetingProfile = self.get( targetingProfileId ).first()
				
			# split on a comma with no spaces
			idList = deviceCriteriaIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DeviceCriterion		
				deviceCriterion = DeviceCriterionDelegate().get(id).first();	
				# add the DeviceCriterion
				targetingProfile.deviceCriteria.remove(deviceCriterion)
				
			# save it		
			targetingProfile.save()
			
			# reload and return the appropriate version
			return self.get( targetingProfileId );
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except DeviceCriterion.DoesNotExist:
			raise ProcessingError(errMsg + " : DeviceCriterion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
