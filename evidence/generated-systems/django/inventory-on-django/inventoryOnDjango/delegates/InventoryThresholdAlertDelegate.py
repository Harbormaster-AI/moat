from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.InventoryThresholdAlert import InventoryThresholdAlert
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.models.ReplenishmentPolicy import ReplenishmentPolicy
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InventoryThresholdAlert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryThresholdAlertDelegate Declaration
#======================================================================
class InventoryThresholdAlertDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inventoryThresholdAlertId ):
		try:	
			inventoryThresholdAlert = InventoryThresholdAlert.objects.filter(id=inventoryThresholdAlertId)
			return inventoryThresholdAlert.first();
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError("InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inventoryThresholdAlert):
		for model in serializers.deserialize("json", inventoryThresholdAlert):
			model.save()
			return model;

	def create(self, inventoryThresholdAlert):
		inventoryThresholdAlert.save()
		return inventoryThresholdAlert;

	def saveFromJson(self, inventoryThresholdAlert):
		for model in serializers.deserialize("json", inventoryThresholdAlert):
			model.save()
			return inventoryThresholdAlert;
	
	def save(self, inventoryThresholdAlert):
		inventoryThresholdAlert.save()
		return inventoryThresholdAlert;
	
	def delete(self, inventoryThresholdAlertId ):
		errMsg = "Failed to delete InventoryThresholdAlert from db using id " + str(inventoryThresholdAlertId)
		
		try:
			inventoryThresholdAlert = InventoryThresholdAlert.objects.get(id=inventoryThresholdAlertId)
			inventoryThresholdAlert.delete()
			return True
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError("InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InventoryThresholdAlert.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InventoryThresholdAlert from db")
		except Exception:
			return None;
		
	def assignSku( self, inventoryThresholdAlertId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on InventoryThresholdAlert"

		try:
			# get the InventoryThresholdAlert from db
			inventoryThresholdAlert = self.get( inventoryThresholdAlertId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			inventoryThresholdAlert.sku = stockKeepingUnit
			
			#save it
			inventoryThresholdAlert.save()

			# reload and return the appropriate version					
			return self.get( inventoryThresholdAlertId );
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, inventoryThresholdAlertId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on InventoryThresholdAlert"

		try:
			# get the InventoryThresholdAlert from db
			inventoryThresholdAlert = self.get( inventoryThresholdAlertId ).first()	
			
			# assign to None for unassignment
			inventoryThresholdAlert.stockKeepingUnit = None			

			#save it
			inventoryThresholdAlert.save()

			# reload and return the appropriate version					
			return self.get( inventoryThresholdAlertId );
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWarehouse( self, inventoryThresholdAlertId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on InventoryThresholdAlert"

		try:
			# get the InventoryThresholdAlert from db
			inventoryThresholdAlert = self.get( inventoryThresholdAlertId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			inventoryThresholdAlert.warehouse = warehouse
			
			#save it
			inventoryThresholdAlert.save()

			# reload and return the appropriate version					
			return self.get( inventoryThresholdAlertId );
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, inventoryThresholdAlertId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on InventoryThresholdAlert"

		try:
			# get the InventoryThresholdAlert from db
			inventoryThresholdAlert = self.get( inventoryThresholdAlertId ).first()	
			
			# assign to None for unassignment
			inventoryThresholdAlert.warehouse = None			

			#save it
			inventoryThresholdAlert.save()

			# reload and return the appropriate version					
			return self.get( inventoryThresholdAlertId );
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, inventoryThresholdAlertId, locationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on InventoryThresholdAlert"

		try:
			# get the InventoryThresholdAlert from db
			inventoryThresholdAlert = self.get( inventoryThresholdAlertId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(locationId).first();
			
			# assign the Location		
			inventoryThresholdAlert.location = storageLocation
			
			#save it
			inventoryThresholdAlert.save()

			# reload and return the appropriate version					
			return self.get( inventoryThresholdAlertId );
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, inventoryThresholdAlertId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on InventoryThresholdAlert"

		try:
			# get the InventoryThresholdAlert from db
			inventoryThresholdAlert = self.get( inventoryThresholdAlertId ).first()	
			
			# assign to None for unassignment
			inventoryThresholdAlert.storageLocation = None			

			#save it
			inventoryThresholdAlert.save()

			# reload and return the appropriate version					
			return self.get( inventoryThresholdAlertId );
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRelatedPolicy( self, inventoryThresholdAlertId, relatedPolicyId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.ReplenishmentPolicyDelegate import ReplenishmentPolicyDelegate

		errMsg = "Failed to assign element " + str(relatedPolicyId) + " for RelatedPolicy on InventoryThresholdAlert"

		try:
			# get the InventoryThresholdAlert from db
			inventoryThresholdAlert = self.get( inventoryThresholdAlertId ).first()	
			
			# get the ReplenishmentPolicy from db
			replenishmentPolicy = ReplenishmentPolicyDelegate().get(relatedPolicyId).first();
			
			# assign the RelatedPolicy		
			inventoryThresholdAlert.relatedPolicy = replenishmentPolicy
			
			#save it
			inventoryThresholdAlert.save()

			# reload and return the appropriate version					
			return self.get( inventoryThresholdAlertId );
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ReplenishmentPolicy with id " + str(relatedPolicyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRelatedPolicy( self, inventoryThresholdAlertId ):
		errMsg = "Failed to unassign element " + str(relatedPolicyId) + " for RelatedPolicy on InventoryThresholdAlert"

		try:
			# get the InventoryThresholdAlert from db
			inventoryThresholdAlert = self.get( inventoryThresholdAlertId ).first()	
			
			# assign to None for unassignment
			inventoryThresholdAlert.replenishmentPolicy = None			

			#save it
			inventoryThresholdAlert.save()

			# reload and return the appropriate version					
			return self.get( inventoryThresholdAlertId );
		except InventoryThresholdAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryThresholdAlert with id " + str(inventoryThresholdAlertId) + " does not exist.")
		except Exception:
			return None;
		
