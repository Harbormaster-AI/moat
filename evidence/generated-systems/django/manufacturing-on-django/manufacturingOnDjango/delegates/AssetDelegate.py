from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Asset import Asset
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.WorkCenter import WorkCenter
from manufacturingOnDjango.models.MaintenanceOrder import MaintenanceOrder
from manufacturingOnDjango.models.MaintenancePlan import MaintenancePlan
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Asset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AssetDelegate Declaration
#======================================================================
class AssetDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, assetId ):
		try:	
			asset = Asset.objects.filter(id=assetId)
			return asset.first();
		except Asset.DoesNotExist:
			raise ProcessingError("Asset with id " + str(assetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, asset):
		for model in serializers.deserialize("json", asset):
			model.save()
			return model;

	def create(self, asset):
		asset.save()
		return asset;

	def saveFromJson(self, asset):
		for model in serializers.deserialize("json", asset):
			model.save()
			return asset;
	
	def save(self, asset):
		asset.save()
		return asset;
	
	def delete(self, assetId ):
		errMsg = "Failed to delete Asset from db using id " + str(assetId)
		
		try:
			asset = Asset.objects.get(id=assetId)
			asset.delete()
			return True
		except Asset.DoesNotExist:
			raise ProcessingError("Asset with id " + str(assetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Asset.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Asset from db")
		except Exception:
			return None;
		
	def assignPlant( self, assetId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on Asset"

		try:
			# get the Asset from db
			asset = self.get( assetId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			asset.plant = plant
			
			#save it
			asset.save()

			# reload and return the appropriate version					
			return self.get( assetId );
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, assetId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on Asset"

		try:
			# get the Asset from db
			asset = self.get( assetId ).first()	
			
			# assign to None for unassignment
			asset.plant = None			

			#save it
			asset.save()

			# reload and return the appropriate version					
			return self.get( assetId );
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkCenter( self, assetId, workCenterId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

		errMsg = "Failed to assign element " + str(workCenterId) + " for WorkCenter on Asset"

		try:
			# get the Asset from db
			asset = self.get( assetId ).first()	
			
			# get the WorkCenter from db
			workCenter = WorkCenterDelegate().get(workCenterId).first();
			
			# assign the WorkCenter		
			asset.workCenter = workCenter
			
			#save it
			asset.save()

			# reload and return the appropriate version					
			return self.get( assetId );
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter with id " + str(workCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkCenter( self, assetId ):
		errMsg = "Failed to unassign element " + str(workCenterId) + " for WorkCenter on Asset"

		try:
			# get the Asset from db
			asset = self.get( assetId ).first()	
			
			# assign to None for unassignment
			asset.workCenter = None			

			#save it
			asset.save()

			# reload and return the appropriate version					
			return self.get( assetId );
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except Exception:
			return None;
		
	def addMaintenanceOrders( self, assetId, maintenanceOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.MaintenanceOrderDelegate import MaintenanceOrderDelegate

		errMsg = "Failed to add elements " + str(maintenanceOrdersIds) + " for MaintenanceOrders on Asset"

		try:
			# get the Asset
			asset = self.get( assetId ).first()
				
			# split on a comma with no spaces
			idList = maintenanceOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MaintenanceOrder		
				maintenanceOrder = MaintenanceOrderDelegate().get(id).first();	
				# add the MaintenanceOrder
				asset.maintenanceOrders.add(maintenanceOrder)
				
			# save it		
			asset.save()
			
			# reload and return the appropriate version
			return self.get( assetId );
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMaintenanceOrders( self, assetId, maintenanceOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.MaintenanceOrderDelegate import MaintenanceOrderDelegate

		errMsg = "Failed to remove elements " + str(maintenanceOrdersIds) + " for MaintenanceOrders on Asset"

		try:
			# get the Asset
			asset = self.get( assetId ).first()
				
			# split on a comma with no spaces
			idList = maintenanceOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MaintenanceOrder		
				maintenanceOrder = MaintenanceOrderDelegate().get(id).first();	
				# add the MaintenanceOrder
				asset.maintenanceOrders.remove(maintenanceOrder)
				
			# save it		
			asset.save()
			
			# reload and return the appropriate version
			return self.get( assetId );
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMaintenancePlans( self, assetId, maintenancePlansIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.MaintenancePlanDelegate import MaintenancePlanDelegate

		errMsg = "Failed to add elements " + str(maintenancePlansIds) + " for MaintenancePlans on Asset"

		try:
			# get the Asset
			asset = self.get( assetId ).first()
				
			# split on a comma with no spaces
			idList = maintenancePlansIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MaintenancePlan		
				maintenancePlan = MaintenancePlanDelegate().get(id).first();	
				# add the MaintenancePlan
				asset.maintenancePlans.add(maintenancePlan)
				
			# save it		
			asset.save()
			
			# reload and return the appropriate version
			return self.get( assetId );
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except MaintenancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenancePlan does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMaintenancePlans( self, assetId, maintenancePlansIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.MaintenancePlanDelegate import MaintenancePlanDelegate

		errMsg = "Failed to remove elements " + str(maintenancePlansIds) + " for MaintenancePlans on Asset"

		try:
			# get the Asset
			asset = self.get( assetId ).first()
				
			# split on a comma with no spaces
			idList = maintenancePlansIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MaintenancePlan		
				maintenancePlan = MaintenancePlanDelegate().get(id).first();	
				# add the MaintenancePlan
				asset.maintenancePlans.remove(maintenancePlan)
				
			# save it		
			asset.save()
			
			# reload and return the appropriate version
			return self.get( assetId );
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except MaintenancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenancePlan does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
