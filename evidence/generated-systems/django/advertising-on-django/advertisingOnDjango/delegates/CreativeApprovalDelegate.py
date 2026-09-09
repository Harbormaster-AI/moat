from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.CreativeApproval import CreativeApproval
from advertisingOnDjango.models.CreativeAsset import CreativeAsset
from advertisingOnDjango.models.Publisher import Publisher
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CreativeApproval
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeApprovalDelegate Declaration
#======================================================================
class CreativeApprovalDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, creativeApprovalId ):
		try:	
			creativeApproval = CreativeApproval.objects.filter(id=creativeApprovalId)
			return creativeApproval.first();
		except CreativeApproval.DoesNotExist:
			raise ProcessingError("CreativeApproval with id " + str(creativeApprovalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, creativeApproval):
		for model in serializers.deserialize("json", creativeApproval):
			model.save()
			return model;

	def create(self, creativeApproval):
		creativeApproval.save()
		return creativeApproval;

	def saveFromJson(self, creativeApproval):
		for model in serializers.deserialize("json", creativeApproval):
			model.save()
			return creativeApproval;
	
	def save(self, creativeApproval):
		creativeApproval.save()
		return creativeApproval;
	
	def delete(self, creativeApprovalId ):
		errMsg = "Failed to delete CreativeApproval from db using id " + str(creativeApprovalId)
		
		try:
			creativeApproval = CreativeApproval.objects.get(id=creativeApprovalId)
			creativeApproval.delete()
			return True
		except CreativeApproval.DoesNotExist:
			raise ProcessingError("CreativeApproval with id " + str(creativeApprovalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CreativeApproval.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CreativeApproval from db")
		except Exception:
			return None;
		
	def assignCreativeAsset( self, creativeApprovalId, creativeAssetId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeAssetDelegate import CreativeAssetDelegate

		errMsg = "Failed to assign element " + str(creativeAssetId) + " for CreativeAsset on CreativeApproval"

		try:
			# get the CreativeApproval from db
			creativeApproval = self.get( creativeApprovalId ).first()	
			
			# get the CreativeAsset from db
			creativeAsset = CreativeAssetDelegate().get(creativeAssetId).first();
			
			# assign the CreativeAsset		
			creativeApproval.creativeAsset = creativeAsset
			
			#save it
			creativeApproval.save()

			# reload and return the appropriate version					
			return self.get( creativeApprovalId );
		except CreativeApproval.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeApproval with id " + str(creativeApprovalId) + " does not exist.")
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCreativeAsset( self, creativeApprovalId ):
		errMsg = "Failed to unassign element " + str(creativeAssetId) + " for CreativeAsset on CreativeApproval"

		try:
			# get the CreativeApproval from db
			creativeApproval = self.get( creativeApprovalId ).first()	
			
			# assign to None for unassignment
			creativeApproval.creativeAsset = None			

			#save it
			creativeApproval.save()

			# reload and return the appropriate version					
			return self.get( creativeApprovalId );
		except CreativeApproval.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeApproval with id " + str(creativeApprovalId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPublisher( self, creativeApprovalId, publisherId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PublisherDelegate import PublisherDelegate

		errMsg = "Failed to assign element " + str(publisherId) + " for Publisher on CreativeApproval"

		try:
			# get the CreativeApproval from db
			creativeApproval = self.get( creativeApprovalId ).first()	
			
			# get the Publisher from db
			publisher = PublisherDelegate().get(publisherId).first();
			
			# assign the Publisher		
			creativeApproval.publisher = publisher
			
			#save it
			creativeApproval.save()

			# reload and return the appropriate version					
			return self.get( creativeApprovalId );
		except CreativeApproval.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeApproval with id " + str(creativeApprovalId) + " does not exist.")
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPublisher( self, creativeApprovalId ):
		errMsg = "Failed to unassign element " + str(publisherId) + " for Publisher on CreativeApproval"

		try:
			# get the CreativeApproval from db
			creativeApproval = self.get( creativeApprovalId ).first()	
			
			# assign to None for unassignment
			creativeApproval.publisher = None			

			#save it
			creativeApproval.save()

			# reload and return the appropriate version					
			return self.get( creativeApprovalId );
		except CreativeApproval.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeApproval with id " + str(creativeApprovalId) + " does not exist.")
		except Exception:
			return None;
		
