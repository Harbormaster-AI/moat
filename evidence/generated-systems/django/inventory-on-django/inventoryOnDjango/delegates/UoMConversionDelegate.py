from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.UoMConversion import UoMConversion
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model UoMConversion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UoMConversionDelegate Declaration
#======================================================================
class UoMConversionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, uoMConversionId ):
		try:	
			uoMConversion = UoMConversion.objects.filter(id=uoMConversionId)
			return uoMConversion.first();
		except UoMConversion.DoesNotExist:
			raise ProcessingError("UoMConversion with id " + str(uoMConversionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, uoMConversion):
		for model in serializers.deserialize("json", uoMConversion):
			model.save()
			return model;

	def create(self, uoMConversion):
		uoMConversion.save()
		return uoMConversion;

	def saveFromJson(self, uoMConversion):
		for model in serializers.deserialize("json", uoMConversion):
			model.save()
			return uoMConversion;
	
	def save(self, uoMConversion):
		uoMConversion.save()
		return uoMConversion;
	
	def delete(self, uoMConversionId ):
		errMsg = "Failed to delete UoMConversion from db using id " + str(uoMConversionId)
		
		try:
			uoMConversion = UoMConversion.objects.get(id=uoMConversionId)
			uoMConversion.delete()
			return True
		except UoMConversion.DoesNotExist:
			raise ProcessingError("UoMConversion with id " + str(uoMConversionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = UoMConversion.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all UoMConversion from db")
		except Exception:
			return None;
		
	def assignSku( self, uoMConversionId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on UoMConversion"

		try:
			# get the UoMConversion from db
			uoMConversion = self.get( uoMConversionId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			uoMConversion.sku = stockKeepingUnit
			
			#save it
			uoMConversion.save()

			# reload and return the appropriate version					
			return self.get( uoMConversionId );
		except UoMConversion.DoesNotExist:
			raise ProcessingError(errMsg + " : UoMConversion with id " + str(uoMConversionId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, uoMConversionId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on UoMConversion"

		try:
			# get the UoMConversion from db
			uoMConversion = self.get( uoMConversionId ).first()	
			
			# assign to None for unassignment
			uoMConversion.stockKeepingUnit = None			

			#save it
			uoMConversion.save()

			# reload and return the appropriate version					
			return self.get( uoMConversionId );
		except UoMConversion.DoesNotExist:
			raise ProcessingError(errMsg + " : UoMConversion with id " + str(uoMConversionId) + " does not exist.")
		except Exception:
			return None;
		
