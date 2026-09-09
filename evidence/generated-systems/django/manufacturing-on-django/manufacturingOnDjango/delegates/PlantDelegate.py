from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.Enterprise import Enterprise
from manufacturingOnDjango.models.ProductionLine import ProductionLine
from manufacturingOnDjango.models.WorkCenter import WorkCenter
from manufacturingOnDjango.models.Warehouse import Warehouse
from manufacturingOnDjango.models.Asset import Asset
from manufacturingOnDjango.models.ProductionSchedule import ProductionSchedule
from manufacturingOnDjango.exceptions import Exceptions

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
		
	def assignEnterprise( self, plantId, enterpriseId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.EnterpriseDelegate import EnterpriseDelegate

		errMsg = "Failed to assign element " + str(enterpriseId) + " for Enterprise on Plant"

		try:
			# get the Plant from db
			plant = self.get( plantId ).first()	
			
			# get the Enterprise from db
			enterprise = EnterpriseDelegate().get(enterpriseId).first();
			
			# assign the Enterprise		
			plant.enterprise = enterprise
			
			#save it
			plant.save()

			# reload and return the appropriate version					
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEnterprise( self, plantId ):
		errMsg = "Failed to unassign element " + str(enterpriseId) + " for Enterprise on Plant"

		try:
			# get the Plant from db
			plant = self.get( plantId ).first()	
			
			# assign to None for unassignment
			plant.enterprise = None			

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
		from manufacturingOnDjango.delegates.ProductionLineDelegate import ProductionLineDelegate

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
		from manufacturingOnDjango.delegates.ProductionLineDelegate import ProductionLineDelegate

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
		
	def addWorkCenters( self, plantId, workCentersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

		errMsg = "Failed to add elements " + str(workCentersIds) + " for WorkCenters on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = workCentersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the WorkCenter		
				workCenter = WorkCenterDelegate().get(id).first();	
				# add the WorkCenter
				plant.workCenters.add(workCenter)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWorkCenters( self, plantId, workCentersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

		errMsg = "Failed to remove elements " + str(workCentersIds) + " for WorkCenters on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = workCentersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the WorkCenter		
				workCenter = WorkCenterDelegate().get(id).first();	
				# add the WorkCenter
				plant.workCenters.remove(workCenter)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addWarehouses( self, plantId, warehousesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

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
		from manufacturingOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

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
		
	def addAssets( self, plantId, assetsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.AssetDelegate import AssetDelegate

		errMsg = "Failed to add elements " + str(assetsIds) + " for Assets on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = assetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Asset		
				asset = AssetDelegate().get(id).first();	
				# add the Asset
				plant.assets.add(asset)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAssets( self, plantId, assetsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.AssetDelegate import AssetDelegate

		errMsg = "Failed to remove elements " + str(assetsIds) + " for Assets on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = assetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Asset		
				asset = AssetDelegate().get(id).first();	
				# add the Asset
				plant.assets.remove(asset)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProductionSchedules( self, plantId, productionSchedulesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ProductionScheduleDelegate import ProductionScheduleDelegate

		errMsg = "Failed to add elements " + str(productionSchedulesIds) + " for ProductionSchedules on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = productionSchedulesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ProductionSchedule		
				productionSchedule = ProductionScheduleDelegate().get(id).first();	
				# add the ProductionSchedule
				plant.productionSchedules.add(productionSchedule)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except ProductionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionSchedule does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProductionSchedules( self, plantId, productionSchedulesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ProductionScheduleDelegate import ProductionScheduleDelegate

		errMsg = "Failed to remove elements " + str(productionSchedulesIds) + " for ProductionSchedules on Plant"

		try:
			# get the Plant
			plant = self.get( plantId ).first()
				
			# split on a comma with no spaces
			idList = productionSchedulesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ProductionSchedule		
				productionSchedule = ProductionScheduleDelegate().get(id).first();	
				# add the ProductionSchedule
				plant.productionSchedules.remove(productionSchedule)
				
			# save it		
			plant.save()
			
			# reload and return the appropriate version
			return self.get( plantId );
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except ProductionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionSchedule does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
