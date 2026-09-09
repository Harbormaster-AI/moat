from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.MaintenancePlan import MaintenancePlan
from manufacturingOnDjango.models.Asset import Asset
from manufacturingOnDjango.models.MaintenanceOrder import MaintenanceOrder
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MaintenancePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenancePlanDelegate Declaration
#======================================================================
class MaintenancePlanDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, maintenancePlanId ):
		try:	
			maintenancePlan = MaintenancePlan.objects.filter(id=maintenancePlanId)
			return maintenancePlan.first();
		except MaintenancePlan.DoesNotExist:
			raise ProcessingError("MaintenancePlan with id " + str(maintenancePlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, maintenancePlan):
		for model in serializers.deserialize("json", maintenancePlan):
			model.save()
			return model;

	def create(self, maintenancePlan):
		maintenancePlan.save()
		return maintenancePlan;

	def saveFromJson(self, maintenancePlan):
		for model in serializers.deserialize("json", maintenancePlan):
			model.save()
			return maintenancePlan;
	
	def save(self, maintenancePlan):
		maintenancePlan.save()
		return maintenancePlan;
	
	def delete(self, maintenancePlanId ):
		errMsg = "Failed to delete MaintenancePlan from db using id " + str(maintenancePlanId)
		
		try:
			maintenancePlan = MaintenancePlan.objects.get(id=maintenancePlanId)
			maintenancePlan.delete()
			return True
		except MaintenancePlan.DoesNotExist:
			raise ProcessingError("MaintenancePlan with id " + str(maintenancePlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MaintenancePlan.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MaintenancePlan from db")
		except Exception:
			return None;
		
	def assignAsset( self, maintenancePlanId, assetId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.AssetDelegate import AssetDelegate

		errMsg = "Failed to assign element " + str(assetId) + " for Asset on MaintenancePlan"

		try:
			# get the MaintenancePlan from db
			maintenancePlan = self.get( maintenancePlanId ).first()	
			
			# get the Asset from db
			asset = AssetDelegate().get(assetId).first();
			
			# assign the Asset		
			maintenancePlan.asset = asset
			
			#save it
			maintenancePlan.save()

			# reload and return the appropriate version					
			return self.get( maintenancePlanId );
		except MaintenancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenancePlan with id " + str(maintenancePlanId) + " does not exist.")
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAsset( self, maintenancePlanId ):
		errMsg = "Failed to unassign element " + str(assetId) + " for Asset on MaintenancePlan"

		try:
			# get the MaintenancePlan from db
			maintenancePlan = self.get( maintenancePlanId ).first()	
			
			# assign to None for unassignment
			maintenancePlan.asset = None			

			#save it
			maintenancePlan.save()

			# reload and return the appropriate version					
			return self.get( maintenancePlanId );
		except MaintenancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenancePlan with id " + str(maintenancePlanId) + " does not exist.")
		except Exception:
			return None;
		
	def addMaintenanceOrders( self, maintenancePlanId, maintenanceOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.MaintenanceOrderDelegate import MaintenanceOrderDelegate

		errMsg = "Failed to add elements " + str(maintenanceOrdersIds) + " for MaintenanceOrders on MaintenancePlan"

		try:
			# get the MaintenancePlan
			maintenancePlan = self.get( maintenancePlanId ).first()
				
			# split on a comma with no spaces
			idList = maintenanceOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MaintenanceOrder		
				maintenanceOrder = MaintenanceOrderDelegate().get(id).first();	
				# add the MaintenanceOrder
				maintenancePlan.maintenanceOrders.add(maintenanceOrder)
				
			# save it		
			maintenancePlan.save()
			
			# reload and return the appropriate version
			return self.get( maintenancePlanId );
		except MaintenancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenancePlan with id " + str(maintenancePlanId) + " does not exist.")
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMaintenanceOrders( self, maintenancePlanId, maintenanceOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.MaintenanceOrderDelegate import MaintenanceOrderDelegate

		errMsg = "Failed to remove elements " + str(maintenanceOrdersIds) + " for MaintenanceOrders on MaintenancePlan"

		try:
			# get the MaintenancePlan
			maintenancePlan = self.get( maintenancePlanId ).first()
				
			# split on a comma with no spaces
			idList = maintenanceOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MaintenanceOrder		
				maintenanceOrder = MaintenanceOrderDelegate().get(id).first();	
				# add the MaintenanceOrder
				maintenancePlan.maintenanceOrders.remove(maintenanceOrder)
				
			# save it		
			maintenancePlan.save()
			
			# reload and return the appropriate version
			return self.get( maintenancePlanId );
		except MaintenancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenancePlan with id " + str(maintenancePlanId) + " does not exist.")
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
