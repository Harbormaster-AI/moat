from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.CreativeFile import CreativeFile
from advertisingOnDjango.models.CreativeAsset import CreativeAsset
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CreativeFile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeFileDelegate Declaration
#======================================================================
class CreativeFileDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, creativeFileId ):
		try:	
			creativeFile = CreativeFile.objects.filter(id=creativeFileId)
			return creativeFile.first();
		except CreativeFile.DoesNotExist:
			raise ProcessingError("CreativeFile with id " + str(creativeFileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, creativeFile):
		for model in serializers.deserialize("json", creativeFile):
			model.save()
			return model;

	def create(self, creativeFile):
		creativeFile.save()
		return creativeFile;

	def saveFromJson(self, creativeFile):
		for model in serializers.deserialize("json", creativeFile):
			model.save()
			return creativeFile;
	
	def save(self, creativeFile):
		creativeFile.save()
		return creativeFile;
	
	def delete(self, creativeFileId ):
		errMsg = "Failed to delete CreativeFile from db using id " + str(creativeFileId)
		
		try:
			creativeFile = CreativeFile.objects.get(id=creativeFileId)
			creativeFile.delete()
			return True
		except CreativeFile.DoesNotExist:
			raise ProcessingError("CreativeFile with id " + str(creativeFileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CreativeFile.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CreativeFile from db")
		except Exception:
			return None;
		
	def assignCreativeAsset( self, creativeFileId, creativeAssetId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeAssetDelegate import CreativeAssetDelegate

		errMsg = "Failed to assign element " + str(creativeAssetId) + " for CreativeAsset on CreativeFile"

		try:
			# get the CreativeFile from db
			creativeFile = self.get( creativeFileId ).first()	
			
			# get the CreativeAsset from db
			creativeAsset = CreativeAssetDelegate().get(creativeAssetId).first();
			
			# assign the CreativeAsset		
			creativeFile.creativeAsset = creativeAsset
			
			#save it
			creativeFile.save()

			# reload and return the appropriate version					
			return self.get( creativeFileId );
		except CreativeFile.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeFile with id " + str(creativeFileId) + " does not exist.")
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCreativeAsset( self, creativeFileId ):
		errMsg = "Failed to unassign element " + str(creativeAssetId) + " for CreativeAsset on CreativeFile"

		try:
			# get the CreativeFile from db
			creativeFile = self.get( creativeFileId ).first()	
			
			# assign to None for unassignment
			creativeFile.creativeAsset = None			

			#save it
			creativeFile.save()

			# reload and return the appropriate version					
			return self.get( creativeFileId );
		except CreativeFile.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeFile with id " + str(creativeFileId) + " does not exist.")
		except Exception:
			return None;
		
