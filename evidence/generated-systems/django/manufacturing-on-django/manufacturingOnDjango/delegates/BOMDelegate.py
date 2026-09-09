from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.BOM import BOM
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.BOMItem import BOMItem
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BOM
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BOMDelegate Declaration
#======================================================================
class BOMDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, bOMId ):
		try:	
			bOM = BOM.objects.filter(id=bOMId)
			return bOM.first();
		except BOM.DoesNotExist:
			raise ProcessingError("BOM with id " + str(bOMId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, bOM):
		for model in serializers.deserialize("json", bOM):
			model.save()
			return model;

	def create(self, bOM):
		bOM.save()
		return bOM;

	def saveFromJson(self, bOM):
		for model in serializers.deserialize("json", bOM):
			model.save()
			return bOM;
	
	def save(self, bOM):
		bOM.save()
		return bOM;
	
	def delete(self, bOMId ):
		errMsg = "Failed to delete BOM from db using id " + str(bOMId)
		
		try:
			bOM = BOM.objects.get(id=bOMId)
			bOM.delete()
			return True
		except BOM.DoesNotExist:
			raise ProcessingError("BOM with id " + str(bOMId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BOM.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BOM from db")
		except Exception:
			return None;
		
	def assignParentItem( self, bOMId, parentItemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(parentItemId) + " for ParentItem on BOM"

		try:
			# get the BOM from db
			bOM = self.get( bOMId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(parentItemId).first();
			
			# assign the ParentItem		
			bOM.parentItem = item
			
			#save it
			bOM.save()

			# reload and return the appropriate version					
			return self.get( bOMId );
		except BOM.DoesNotExist:
			raise ProcessingError(errMsg + " : BOM with id " + str(bOMId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(parentItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignParentItem( self, bOMId ):
		errMsg = "Failed to unassign element " + str(parentItemId) + " for ParentItem on BOM"

		try:
			# get the BOM from db
			bOM = self.get( bOMId ).first()	
			
			# assign to None for unassignment
			bOM.item = None			

			#save it
			bOM.save()

			# reload and return the appropriate version					
			return self.get( bOMId );
		except BOM.DoesNotExist:
			raise ProcessingError(errMsg + " : BOM with id " + str(bOMId) + " does not exist.")
		except Exception:
			return None;
		
	def addBomItems( self, bOMId, bomItemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.BOMItemDelegate import BOMItemDelegate

		errMsg = "Failed to add elements " + str(bomItemsIds) + " for BomItems on BOM"

		try:
			# get the BOM
			bOM = self.get( bOMId ).first()
				
			# split on a comma with no spaces
			idList = bomItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BOMItem		
				bOMItem = BOMItemDelegate().get(id).first();	
				# add the BOMItem
				bOM.bomItems.add(bOMItem)
				
			# save it		
			bOM.save()
			
			# reload and return the appropriate version
			return self.get( bOMId );
		except BOM.DoesNotExist:
			raise ProcessingError(errMsg + " : BOM with id " + str(bOMId) + " does not exist.")
		except BOMItem.DoesNotExist:
			raise ProcessingError(errMsg + " : BOMItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBomItems( self, bOMId, bomItemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.BOMItemDelegate import BOMItemDelegate

		errMsg = "Failed to remove elements " + str(bomItemsIds) + " for BomItems on BOM"

		try:
			# get the BOM
			bOM = self.get( bOMId ).first()
				
			# split on a comma with no spaces
			idList = bomItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BOMItem		
				bOMItem = BOMItemDelegate().get(id).first();	
				# add the BOMItem
				bOM.bomItems.remove(bOMItem)
				
			# save it		
			bOM.save()
			
			# reload and return the appropriate version
			return self.get( bOMId );
		except BOM.DoesNotExist:
			raise ProcessingError(errMsg + " : BOM with id " + str(bOMId) + " does not exist.")
		except BOMItem.DoesNotExist:
			raise ProcessingError(errMsg + " : BOMItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
