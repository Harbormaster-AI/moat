from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.ProductionOrder import ProductionOrder
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.models.Plant import Plant
from aerospaceOnDjango.models.AircraftOrder import AircraftOrder
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ProductionOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionOrderDelegate Declaration
#======================================================================
class ProductionOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, productionOrderId ):
		try:	
			productionOrder = ProductionOrder.objects.filter(id=productionOrderId)
			return productionOrder.first();
		except ProductionOrder.DoesNotExist:
			raise ProcessingError("ProductionOrder with id " + str(productionOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, productionOrder):
		for model in serializers.deserialize("json", productionOrder):
			model.save()
			return model;

	def create(self, productionOrder):
		productionOrder.save()
		return productionOrder;

	def saveFromJson(self, productionOrder):
		for model in serializers.deserialize("json", productionOrder):
			model.save()
			return productionOrder;
	
	def save(self, productionOrder):
		productionOrder.save()
		return productionOrder;
	
	def delete(self, productionOrderId ):
		errMsg = "Failed to delete ProductionOrder from db using id " + str(productionOrderId)
		
		try:
			productionOrder = ProductionOrder.objects.get(id=productionOrderId)
			productionOrder.delete()
			return True
		except ProductionOrder.DoesNotExist:
			raise ProcessingError("ProductionOrder with id " + str(productionOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ProductionOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ProductionOrder from db")
		except Exception:
			return None;
		
	def assignVariant( self, productionOrderId, variantId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on ProductionOrder"

		try:
			# get the ProductionOrder from db
			productionOrder = self.get( productionOrderId ).first()	
			
			# get the AircraftVariant from db
			aircraftVariant = AircraftVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			productionOrder.variant = aircraftVariant
			
			#save it
			productionOrder.save()

			# reload and return the appropriate version					
			return self.get( productionOrderId );
		except ProductionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionOrder with id " + str(productionOrderId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, productionOrderId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on ProductionOrder"

		try:
			# get the ProductionOrder from db
			productionOrder = self.get( productionOrderId ).first()	
			
			# assign to None for unassignment
			productionOrder.aircraftVariant = None			

			#save it
			productionOrder.save()

			# reload and return the appropriate version					
			return self.get( productionOrderId );
		except ProductionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionOrder with id " + str(productionOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPlant( self, productionOrderId, plantId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on ProductionOrder"

		try:
			# get the ProductionOrder from db
			productionOrder = self.get( productionOrderId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			productionOrder.plant = plant
			
			#save it
			productionOrder.save()

			# reload and return the appropriate version					
			return self.get( productionOrderId );
		except ProductionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionOrder with id " + str(productionOrderId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, productionOrderId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on ProductionOrder"

		try:
			# get the ProductionOrder from db
			productionOrder = self.get( productionOrderId ).first()	
			
			# assign to None for unassignment
			productionOrder.plant = None			

			#save it
			productionOrder.save()

			# reload and return the appropriate version					
			return self.get( productionOrderId );
		except ProductionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionOrder with id " + str(productionOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAircraftOrder( self, productionOrderId, aircraftOrderId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOrderDelegate import AircraftOrderDelegate

		errMsg = "Failed to assign element " + str(aircraftOrderId) + " for AircraftOrder on ProductionOrder"

		try:
			# get the ProductionOrder from db
			productionOrder = self.get( productionOrderId ).first()	
			
			# get the AircraftOrder from db
			aircraftOrder = AircraftOrderDelegate().get(aircraftOrderId).first();
			
			# assign the AircraftOrder		
			productionOrder.aircraftOrder = aircraftOrder
			
			#save it
			productionOrder.save()

			# reload and return the appropriate version					
			return self.get( productionOrderId );
		except ProductionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionOrder with id " + str(productionOrderId) + " does not exist.")
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAircraftOrder( self, productionOrderId ):
		errMsg = "Failed to unassign element " + str(aircraftOrderId) + " for AircraftOrder on ProductionOrder"

		try:
			# get the ProductionOrder from db
			productionOrder = self.get( productionOrderId ).first()	
			
			# assign to None for unassignment
			productionOrder.aircraftOrder = None			

			#save it
			productionOrder.save()

			# reload and return the appropriate version					
			return self.get( productionOrderId );
		except ProductionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionOrder with id " + str(productionOrderId) + " does not exist.")
		except Exception:
			return None;
		
