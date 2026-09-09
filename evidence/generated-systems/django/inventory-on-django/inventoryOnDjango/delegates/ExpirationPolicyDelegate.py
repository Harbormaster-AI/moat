from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.ExpirationPolicy import ExpirationPolicy
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ExpirationPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExpirationPolicyDelegate Declaration
#======================================================================
class ExpirationPolicyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, expirationPolicyId ):
		try:	
			expirationPolicy = ExpirationPolicy.objects.filter(id=expirationPolicyId)
			return expirationPolicy.first();
		except ExpirationPolicy.DoesNotExist:
			raise ProcessingError("ExpirationPolicy with id " + str(expirationPolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, expirationPolicy):
		for model in serializers.deserialize("json", expirationPolicy):
			model.save()
			return model;

	def create(self, expirationPolicy):
		expirationPolicy.save()
		return expirationPolicy;

	def saveFromJson(self, expirationPolicy):
		for model in serializers.deserialize("json", expirationPolicy):
			model.save()
			return expirationPolicy;
	
	def save(self, expirationPolicy):
		expirationPolicy.save()
		return expirationPolicy;
	
	def delete(self, expirationPolicyId ):
		errMsg = "Failed to delete ExpirationPolicy from db using id " + str(expirationPolicyId)
		
		try:
			expirationPolicy = ExpirationPolicy.objects.get(id=expirationPolicyId)
			expirationPolicy.delete()
			return True
		except ExpirationPolicy.DoesNotExist:
			raise ProcessingError("ExpirationPolicy with id " + str(expirationPolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ExpirationPolicy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ExpirationPolicy from db")
		except Exception:
			return None;
		
	def assignSku( self, expirationPolicyId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on ExpirationPolicy"

		try:
			# get the ExpirationPolicy from db
			expirationPolicy = self.get( expirationPolicyId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			expirationPolicy.sku = stockKeepingUnit
			
			#save it
			expirationPolicy.save()

			# reload and return the appropriate version					
			return self.get( expirationPolicyId );
		except ExpirationPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ExpirationPolicy with id " + str(expirationPolicyId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, expirationPolicyId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on ExpirationPolicy"

		try:
			# get the ExpirationPolicy from db
			expirationPolicy = self.get( expirationPolicyId ).first()	
			
			# assign to None for unassignment
			expirationPolicy.stockKeepingUnit = None			

			#save it
			expirationPolicy.save()

			# reload and return the appropriate version					
			return self.get( expirationPolicyId );
		except ExpirationPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ExpirationPolicy with id " + str(expirationPolicyId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWarehouse( self, expirationPolicyId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on ExpirationPolicy"

		try:
			# get the ExpirationPolicy from db
			expirationPolicy = self.get( expirationPolicyId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			expirationPolicy.warehouse = warehouse
			
			#save it
			expirationPolicy.save()

			# reload and return the appropriate version					
			return self.get( expirationPolicyId );
		except ExpirationPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ExpirationPolicy with id " + str(expirationPolicyId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, expirationPolicyId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on ExpirationPolicy"

		try:
			# get the ExpirationPolicy from db
			expirationPolicy = self.get( expirationPolicyId ).first()	
			
			# assign to None for unassignment
			expirationPolicy.warehouse = None			

			#save it
			expirationPolicy.save()

			# reload and return the appropriate version					
			return self.get( expirationPolicyId );
		except ExpirationPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ExpirationPolicy with id " + str(expirationPolicyId) + " does not exist.")
		except Exception:
			return None;
		
