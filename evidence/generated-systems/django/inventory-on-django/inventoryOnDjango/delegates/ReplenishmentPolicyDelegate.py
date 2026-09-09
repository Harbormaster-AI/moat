from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.ReplenishmentPolicy import ReplenishmentPolicy
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ReplenishmentPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReplenishmentPolicyDelegate Declaration
#======================================================================
class ReplenishmentPolicyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, replenishmentPolicyId ):
		try:	
			replenishmentPolicy = ReplenishmentPolicy.objects.filter(id=replenishmentPolicyId)
			return replenishmentPolicy.first();
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError("ReplenishmentPolicy with id " + str(replenishmentPolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, replenishmentPolicy):
		for model in serializers.deserialize("json", replenishmentPolicy):
			model.save()
			return model;

	def create(self, replenishmentPolicy):
		replenishmentPolicy.save()
		return replenishmentPolicy;

	def saveFromJson(self, replenishmentPolicy):
		for model in serializers.deserialize("json", replenishmentPolicy):
			model.save()
			return replenishmentPolicy;
	
	def save(self, replenishmentPolicy):
		replenishmentPolicy.save()
		return replenishmentPolicy;
	
	def delete(self, replenishmentPolicyId ):
		errMsg = "Failed to delete ReplenishmentPolicy from db using id " + str(replenishmentPolicyId)
		
		try:
			replenishmentPolicy = ReplenishmentPolicy.objects.get(id=replenishmentPolicyId)
			replenishmentPolicy.delete()
			return True
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError("ReplenishmentPolicy with id " + str(replenishmentPolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ReplenishmentPolicy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ReplenishmentPolicy from db")
		except Exception:
			return None;
		
	def assignSku( self, replenishmentPolicyId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on ReplenishmentPolicy"

		try:
			# get the ReplenishmentPolicy from db
			replenishmentPolicy = self.get( replenishmentPolicyId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			replenishmentPolicy.sku = stockKeepingUnit
			
			#save it
			replenishmentPolicy.save()

			# reload and return the appropriate version					
			return self.get( replenishmentPolicyId );
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ReplenishmentPolicy with id " + str(replenishmentPolicyId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, replenishmentPolicyId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on ReplenishmentPolicy"

		try:
			# get the ReplenishmentPolicy from db
			replenishmentPolicy = self.get( replenishmentPolicyId ).first()	
			
			# assign to None for unassignment
			replenishmentPolicy.stockKeepingUnit = None			

			#save it
			replenishmentPolicy.save()

			# reload and return the appropriate version					
			return self.get( replenishmentPolicyId );
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ReplenishmentPolicy with id " + str(replenishmentPolicyId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWarehouse( self, replenishmentPolicyId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on ReplenishmentPolicy"

		try:
			# get the ReplenishmentPolicy from db
			replenishmentPolicy = self.get( replenishmentPolicyId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			replenishmentPolicy.warehouse = warehouse
			
			#save it
			replenishmentPolicy.save()

			# reload and return the appropriate version					
			return self.get( replenishmentPolicyId );
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ReplenishmentPolicy with id " + str(replenishmentPolicyId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, replenishmentPolicyId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on ReplenishmentPolicy"

		try:
			# get the ReplenishmentPolicy from db
			replenishmentPolicy = self.get( replenishmentPolicyId ).first()	
			
			# assign to None for unassignment
			replenishmentPolicy.warehouse = None			

			#save it
			replenishmentPolicy.save()

			# reload and return the appropriate version					
			return self.get( replenishmentPolicyId );
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ReplenishmentPolicy with id " + str(replenishmentPolicyId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, replenishmentPolicyId, locationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on ReplenishmentPolicy"

		try:
			# get the ReplenishmentPolicy from db
			replenishmentPolicy = self.get( replenishmentPolicyId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(locationId).first();
			
			# assign the Location		
			replenishmentPolicy.location = storageLocation
			
			#save it
			replenishmentPolicy.save()

			# reload and return the appropriate version					
			return self.get( replenishmentPolicyId );
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ReplenishmentPolicy with id " + str(replenishmentPolicyId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, replenishmentPolicyId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on ReplenishmentPolicy"

		try:
			# get the ReplenishmentPolicy from db
			replenishmentPolicy = self.get( replenishmentPolicyId ).first()	
			
			# assign to None for unassignment
			replenishmentPolicy.storageLocation = None			

			#save it
			replenishmentPolicy.save()

			# reload and return the appropriate version					
			return self.get( replenishmentPolicyId );
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ReplenishmentPolicy with id " + str(replenishmentPolicyId) + " does not exist.")
		except Exception:
			return None;
		
