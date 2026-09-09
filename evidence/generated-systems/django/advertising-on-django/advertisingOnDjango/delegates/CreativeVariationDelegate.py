from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.CreativeVariation import CreativeVariation
from advertisingOnDjango.models.CreativeAsset import CreativeAsset
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CreativeVariation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeVariationDelegate Declaration
#======================================================================
class CreativeVariationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, creativeVariationId ):
		try:	
			creativeVariation = CreativeVariation.objects.filter(id=creativeVariationId)
			return creativeVariation.first();
		except CreativeVariation.DoesNotExist:
			raise ProcessingError("CreativeVariation with id " + str(creativeVariationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, creativeVariation):
		for model in serializers.deserialize("json", creativeVariation):
			model.save()
			return model;

	def create(self, creativeVariation):
		creativeVariation.save()
		return creativeVariation;

	def saveFromJson(self, creativeVariation):
		for model in serializers.deserialize("json", creativeVariation):
			model.save()
			return creativeVariation;
	
	def save(self, creativeVariation):
		creativeVariation.save()
		return creativeVariation;
	
	def delete(self, creativeVariationId ):
		errMsg = "Failed to delete CreativeVariation from db using id " + str(creativeVariationId)
		
		try:
			creativeVariation = CreativeVariation.objects.get(id=creativeVariationId)
			creativeVariation.delete()
			return True
		except CreativeVariation.DoesNotExist:
			raise ProcessingError("CreativeVariation with id " + str(creativeVariationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CreativeVariation.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CreativeVariation from db")
		except Exception:
			return None;
		
	def assignCreativeAsset( self, creativeVariationId, creativeAssetId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeAssetDelegate import CreativeAssetDelegate

		errMsg = "Failed to assign element " + str(creativeAssetId) + " for CreativeAsset on CreativeVariation"

		try:
			# get the CreativeVariation from db
			creativeVariation = self.get( creativeVariationId ).first()	
			
			# get the CreativeAsset from db
			creativeAsset = CreativeAssetDelegate().get(creativeAssetId).first();
			
			# assign the CreativeAsset		
			creativeVariation.creativeAsset = creativeAsset
			
			#save it
			creativeVariation.save()

			# reload and return the appropriate version					
			return self.get( creativeVariationId );
		except CreativeVariation.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeVariation with id " + str(creativeVariationId) + " does not exist.")
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCreativeAsset( self, creativeVariationId ):
		errMsg = "Failed to unassign element " + str(creativeAssetId) + " for CreativeAsset on CreativeVariation"

		try:
			# get the CreativeVariation from db
			creativeVariation = self.get( creativeVariationId ).first()	
			
			# assign to None for unassignment
			creativeVariation.creativeAsset = None			

			#save it
			creativeVariation.save()

			# reload and return the appropriate version					
			return self.get( creativeVariationId );
		except CreativeVariation.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeVariation with id " + str(creativeVariationId) + " does not exist.")
		except Exception:
			return None;
		
