from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.BOMItem import BOMItem
from manufacturingOnDjango.models.BOM import BOM
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BOMItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BOMItemDelegate Declaration
#======================================================================
class BOMItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, bOMItemId ):
		try:	
			bOMItem = BOMItem.objects.filter(id=bOMItemId)
			return bOMItem.first();
		except BOMItem.DoesNotExist:
			raise ProcessingError("BOMItem with id " + str(bOMItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, bOMItem):
		for model in serializers.deserialize("json", bOMItem):
			model.save()
			return model;

	def create(self, bOMItem):
		bOMItem.save()
		return bOMItem;

	def saveFromJson(self, bOMItem):
		for model in serializers.deserialize("json", bOMItem):
			model.save()
			return bOMItem;
	
	def save(self, bOMItem):
		bOMItem.save()
		return bOMItem;
	
	def delete(self, bOMItemId ):
		errMsg = "Failed to delete BOMItem from db using id " + str(bOMItemId)
		
		try:
			bOMItem = BOMItem.objects.get(id=bOMItemId)
			bOMItem.delete()
			return True
		except BOMItem.DoesNotExist:
			raise ProcessingError("BOMItem with id " + str(bOMItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BOMItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BOMItem from db")
		except Exception:
			return None;
		
	def assignBom( self, bOMItemId, bomId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.BOMDelegate import BOMDelegate

		errMsg = "Failed to assign element " + str(bomId) + " for Bom on BOMItem"

		try:
			# get the BOMItem from db
			bOMItem = self.get( bOMItemId ).first()	
			
			# get the BOM from db
			bOM = BOMDelegate().get(bomId).first();
			
			# assign the Bom		
			bOMItem.bom = bOM
			
			#save it
			bOMItem.save()

			# reload and return the appropriate version					
			return self.get( bOMItemId );
		except BOMItem.DoesNotExist:
			raise ProcessingError(errMsg + " : BOMItem with id " + str(bOMItemId) + " does not exist.")
		except BOM.DoesNotExist:
			raise ProcessingError(errMsg + " : BOM with id " + str(bomId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBom( self, bOMItemId ):
		errMsg = "Failed to unassign element " + str(bomId) + " for Bom on BOMItem"

		try:
			# get the BOMItem from db
			bOMItem = self.get( bOMItemId ).first()	
			
			# assign to None for unassignment
			bOMItem.bOM = None			

			#save it
			bOMItem.save()

			# reload and return the appropriate version					
			return self.get( bOMItemId );
		except BOMItem.DoesNotExist:
			raise ProcessingError(errMsg + " : BOMItem with id " + str(bOMItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignComponent( self, bOMItemId, componentId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(componentId) + " for Component on BOMItem"

		try:
			# get the BOMItem from db
			bOMItem = self.get( bOMItemId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(componentId).first();
			
			# assign the Component		
			bOMItem.component = item
			
			#save it
			bOMItem.save()

			# reload and return the appropriate version					
			return self.get( bOMItemId );
		except BOMItem.DoesNotExist:
			raise ProcessingError(errMsg + " : BOMItem with id " + str(bOMItemId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(componentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignComponent( self, bOMItemId ):
		errMsg = "Failed to unassign element " + str(componentId) + " for Component on BOMItem"

		try:
			# get the BOMItem from db
			bOMItem = self.get( bOMItemId ).first()	
			
			# assign to None for unassignment
			bOMItem.item = None			

			#save it
			bOMItem.save()

			# reload and return the appropriate version					
			return self.get( bOMItemId );
		except BOMItem.DoesNotExist:
			raise ProcessingError(errMsg + " : BOMItem with id " + str(bOMItemId) + " does not exist.")
		except Exception:
			return None;
		
