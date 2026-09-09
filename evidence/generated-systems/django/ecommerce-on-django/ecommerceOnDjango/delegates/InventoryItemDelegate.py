from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.InventoryItem import InventoryItem
from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.models.FulfillmentCenter import FulfillmentCenter
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InventoryItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryItemDelegate Declaration
#======================================================================
class InventoryItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inventoryItemId ):
		try:	
			inventoryItem = InventoryItem.objects.filter(id=inventoryItemId)
			return inventoryItem.first();
		except InventoryItem.DoesNotExist:
			raise ProcessingError("InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inventoryItem):
		for model in serializers.deserialize("json", inventoryItem):
			model.save()
			return model;

	def create(self, inventoryItem):
		inventoryItem.save()
		return inventoryItem;

	def saveFromJson(self, inventoryItem):
		for model in serializers.deserialize("json", inventoryItem):
			model.save()
			return inventoryItem;
	
	def save(self, inventoryItem):
		inventoryItem.save()
		return inventoryItem;
	
	def delete(self, inventoryItemId ):
		errMsg = "Failed to delete InventoryItem from db using id " + str(inventoryItemId)
		
		try:
			inventoryItem = InventoryItem.objects.get(id=inventoryItemId)
			inventoryItem.delete()
			return True
		except InventoryItem.DoesNotExist:
			raise ProcessingError("InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InventoryItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InventoryItem from db")
		except Exception:
			return None;
		
	def assignVariant( self, inventoryItemId, variantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the ProductVariant from db
			productVariant = ProductVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			inventoryItem.variant = productVariant
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.productVariant = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignFulfillmentCenter( self, inventoryItemId, fulfillmentCenterId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.FulfillmentCenterDelegate import FulfillmentCenterDelegate

		errMsg = "Failed to assign element " + str(fulfillmentCenterId) + " for FulfillmentCenter on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the FulfillmentCenter from db
			fulfillmentCenter = FulfillmentCenterDelegate().get(fulfillmentCenterId).first();
			
			# assign the FulfillmentCenter		
			inventoryItem.fulfillmentCenter = fulfillmentCenter
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFulfillmentCenter( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(fulfillmentCenterId) + " for FulfillmentCenter on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.fulfillmentCenter = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
