from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.CreativeAsset import CreativeAsset
from advertisingOnDjango.models.CreativeFile import CreativeFile
from advertisingOnDjango.models.CreativeApproval import CreativeApproval
from advertisingOnDjango.models.CreativeVariation import CreativeVariation
from advertisingOnDjango.models.LineItem import LineItem
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CreativeAsset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeAssetDelegate Declaration
#======================================================================
class CreativeAssetDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, creativeAssetId ):
		try:	
			creativeAsset = CreativeAsset.objects.filter(id=creativeAssetId)
			return creativeAsset.first();
		except CreativeAsset.DoesNotExist:
			raise ProcessingError("CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, creativeAsset):
		for model in serializers.deserialize("json", creativeAsset):
			model.save()
			return model;

	def create(self, creativeAsset):
		creativeAsset.save()
		return creativeAsset;

	def saveFromJson(self, creativeAsset):
		for model in serializers.deserialize("json", creativeAsset):
			model.save()
			return creativeAsset;
	
	def save(self, creativeAsset):
		creativeAsset.save()
		return creativeAsset;
	
	def delete(self, creativeAssetId ):
		errMsg = "Failed to delete CreativeAsset from db using id " + str(creativeAssetId)
		
		try:
			creativeAsset = CreativeAsset.objects.get(id=creativeAssetId)
			creativeAsset.delete()
			return True
		except CreativeAsset.DoesNotExist:
			raise ProcessingError("CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CreativeAsset.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CreativeAsset from db")
		except Exception:
			return None;
		
	def addFiles( self, creativeAssetId, filesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeFileDelegate import CreativeFileDelegate

		errMsg = "Failed to add elements " + str(filesIds) + " for Files on CreativeAsset"

		try:
			# get the CreativeAsset
			creativeAsset = self.get( creativeAssetId ).first()
				
			# split on a comma with no spaces
			idList = filesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CreativeFile		
				creativeFile = CreativeFileDelegate().get(id).first();	
				# add the CreativeFile
				creativeAsset.files.add(creativeFile)
				
			# save it		
			creativeAsset.save()
			
			# reload and return the appropriate version
			return self.get( creativeAssetId );
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except CreativeFile.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeFile does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFiles( self, creativeAssetId, filesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeFileDelegate import CreativeFileDelegate

		errMsg = "Failed to remove elements " + str(filesIds) + " for Files on CreativeAsset"

		try:
			# get the CreativeAsset
			creativeAsset = self.get( creativeAssetId ).first()
				
			# split on a comma with no spaces
			idList = filesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CreativeFile		
				creativeFile = CreativeFileDelegate().get(id).first();	
				# add the CreativeFile
				creativeAsset.files.remove(creativeFile)
				
			# save it		
			creativeAsset.save()
			
			# reload and return the appropriate version
			return self.get( creativeAssetId );
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except CreativeFile.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeFile does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addApprovals( self, creativeAssetId, approvalsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeApprovalDelegate import CreativeApprovalDelegate

		errMsg = "Failed to add elements " + str(approvalsIds) + " for Approvals on CreativeAsset"

		try:
			# get the CreativeAsset
			creativeAsset = self.get( creativeAssetId ).first()
				
			# split on a comma with no spaces
			idList = approvalsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CreativeApproval		
				creativeApproval = CreativeApprovalDelegate().get(id).first();	
				# add the CreativeApproval
				creativeAsset.approvals.add(creativeApproval)
				
			# save it		
			creativeAsset.save()
			
			# reload and return the appropriate version
			return self.get( creativeAssetId );
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except CreativeApproval.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeApproval does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeApprovals( self, creativeAssetId, approvalsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeApprovalDelegate import CreativeApprovalDelegate

		errMsg = "Failed to remove elements " + str(approvalsIds) + " for Approvals on CreativeAsset"

		try:
			# get the CreativeAsset
			creativeAsset = self.get( creativeAssetId ).first()
				
			# split on a comma with no spaces
			idList = approvalsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CreativeApproval		
				creativeApproval = CreativeApprovalDelegate().get(id).first();	
				# add the CreativeApproval
				creativeAsset.approvals.remove(creativeApproval)
				
			# save it		
			creativeAsset.save()
			
			# reload and return the appropriate version
			return self.get( creativeAssetId );
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except CreativeApproval.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeApproval does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addVariations( self, creativeAssetId, variationsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeVariationDelegate import CreativeVariationDelegate

		errMsg = "Failed to add elements " + str(variationsIds) + " for Variations on CreativeAsset"

		try:
			# get the CreativeAsset
			creativeAsset = self.get( creativeAssetId ).first()
				
			# split on a comma with no spaces
			idList = variationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CreativeVariation		
				creativeVariation = CreativeVariationDelegate().get(id).first();	
				# add the CreativeVariation
				creativeAsset.variations.add(creativeVariation)
				
			# save it		
			creativeAsset.save()
			
			# reload and return the appropriate version
			return self.get( creativeAssetId );
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except CreativeVariation.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeVariation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariations( self, creativeAssetId, variationsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeVariationDelegate import CreativeVariationDelegate

		errMsg = "Failed to remove elements " + str(variationsIds) + " for Variations on CreativeAsset"

		try:
			# get the CreativeAsset
			creativeAsset = self.get( creativeAssetId ).first()
				
			# split on a comma with no spaces
			idList = variationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CreativeVariation		
				creativeVariation = CreativeVariationDelegate().get(id).first();	
				# add the CreativeVariation
				creativeAsset.variations.remove(creativeVariation)
				
			# save it		
			creativeAsset.save()
			
			# reload and return the appropriate version
			return self.get( creativeAssetId );
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except CreativeVariation.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeVariation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLineItems( self, creativeAssetId, lineItemsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

		errMsg = "Failed to add elements " + str(lineItemsIds) + " for LineItems on CreativeAsset"

		try:
			# get the CreativeAsset
			creativeAsset = self.get( creativeAssetId ).first()
				
			# split on a comma with no spaces
			idList = lineItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LineItem		
				lineItem = LineItemDelegate().get(id).first();	
				# add the LineItem
				creativeAsset.lineItems.add(lineItem)
				
			# save it		
			creativeAsset.save()
			
			# reload and return the appropriate version
			return self.get( creativeAssetId );
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLineItems( self, creativeAssetId, lineItemsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

		errMsg = "Failed to remove elements " + str(lineItemsIds) + " for LineItems on CreativeAsset"

		try:
			# get the CreativeAsset
			creativeAsset = self.get( creativeAssetId ).first()
				
			# split on a comma with no spaces
			idList = lineItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LineItem		
				lineItem = LineItemDelegate().get(id).first();	
				# add the LineItem
				creativeAsset.lineItems.remove(lineItem)
				
			# save it		
			creativeAsset.save()
			
			# reload and return the appropriate version
			return self.get( creativeAssetId );
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
