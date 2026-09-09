from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.Plant import Plant
from aerospaceOnDjango.models.AerospaceManufacturer import AerospaceManufacturer
from aerospaceOnDjango.models.ProductionLine import ProductionLine
from aerospaceOnDjango.models.Warehouse import Warehouse
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Plant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlantDelegate Declaration
#======================================================================
class PlantDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, plantId ):
		try:	
			plant = Plant.objects.filter(id=plantId)
			return plant.first();
		except Plant.DoesNotExist:
			raise ProcessingError("Plant with id " + str(plantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, plant):
		for model in serializers.deserialize("json", plant):
			model.save()
			return model;

	def create(self, plant):
		plant.save()
		return plant;

	def saveFromJson(self, plant):
		for model in serializers.deserialize("json", plant):
			model.save()
			return plant;
	
	def save(self, plant):
		plant.save()
		return plant;
	
	def delete(self, plantId ):
		errMsg = "Failed to delete Plant from db using id " + str(plantId)
		
		try:
			plant = Plant.objects.get(id=plantId)
			plant.delete()
			return True
		except Plant.DoesNotExist:
			raise ProcessingError("Plant with id " + str(plantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Plant.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Plant from db")
		except Exception:
			return None;
		
	def assignManufacturer( self, plantId, manufacturerId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AerospaceManufacturerDelegate import AerospaceManufacturerDelegate

		errMsg = "Failed to assign element " + str(manufacturerId) + " for Manufacturer on Plant"

		try:
			# get the Plant from db
			plant = self.get( plantId ).first()	
			
			# get the AerospaceManufacturer from db
			aerospaceManufacturer = AerospaceManufacturerDelegate().get(manufacturerId).first();
			
			# assign the Manufacturer		
			plant.manufacturer = aerospaceManufacturer
			
			#save it
			plant.save()

			# reload and return the appropriate version					
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(manufacturerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignManufacturer( self, plantId ):
		errMsg = "Failed to unassign element " + str(manufacturerId) + " for Manufacturer on Plant"

		try:
			# get the Plant from db
			plant = self.get( plantId ).first()	
			
			# assign to None for unassignment
			plant.aerospaceManufacturer = None			

			#save it
			plant.save()

			# reload and return the appropriate version					
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
		
	def addProductionLines( self, plantId, productionLinesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ProductionLineDelegate import ProductionLineDelegate

		errMsg = "Failed to add elements " + str(productionLinesIds) + " for ProductionLines on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = productionLinesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ProductionLine		
				productionLine = ProductionLineDelegate().get(id).first();	
				# add the ProductionLine
				plant.productionLines.add(productionLine)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except ProductionLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProductionLines( self, plantId, productionLinesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ProductionLineDelegate import ProductionLineDelegate

		errMsg = "Failed to remove elements " + str(productionLinesIds) + " for ProductionLines on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = productionLinesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ProductionLine		
				productionLine = ProductionLineDelegate().get(id).first();	
				# add the ProductionLine
				plant.productionLines.remove(productionLine)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except ProductionLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addWarehouses( self, plantId, warehousesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to add elements " + str(warehousesIds) + " for Warehouses on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = warehousesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Warehouse		
				warehouse = WarehouseDelegate().get(id).first();	
				# add the Warehouse
				plant.warehouses.add(warehouse)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWarehouses( self, plantId, warehousesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to remove elements " + str(warehousesIds) + " for Warehouses on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = warehousesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Warehouse		
				warehouse = WarehouseDelegate().get(id).first();	
				# add the Warehouse
				plant.warehouses.remove(warehouse)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
