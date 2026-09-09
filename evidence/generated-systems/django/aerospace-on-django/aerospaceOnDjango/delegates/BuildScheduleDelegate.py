from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.BuildSchedule import BuildSchedule
from aerospaceOnDjango.models.ProductionOrder import ProductionOrder
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BuildSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BuildScheduleDelegate Declaration
#======================================================================
class BuildScheduleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, buildScheduleId ):
		try:	
			buildSchedule = BuildSchedule.objects.filter(id=buildScheduleId)
			return buildSchedule.first();
		except BuildSchedule.DoesNotExist:
			raise ProcessingError("BuildSchedule with id " + str(buildScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, buildSchedule):
		for model in serializers.deserialize("json", buildSchedule):
			model.save()
			return model;

	def create(self, buildSchedule):
		buildSchedule.save()
		return buildSchedule;

	def saveFromJson(self, buildSchedule):
		for model in serializers.deserialize("json", buildSchedule):
			model.save()
			return buildSchedule;
	
	def save(self, buildSchedule):
		buildSchedule.save()
		return buildSchedule;
	
	def delete(self, buildScheduleId ):
		errMsg = "Failed to delete BuildSchedule from db using id " + str(buildScheduleId)
		
		try:
			buildSchedule = BuildSchedule.objects.get(id=buildScheduleId)
			buildSchedule.delete()
			return True
		except BuildSchedule.DoesNotExist:
			raise ProcessingError("BuildSchedule with id " + str(buildScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BuildSchedule.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BuildSchedule from db")
		except Exception:
			return None;
		
	def addProductionOrders( self, buildScheduleId, productionOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ProductionOrderDelegate import ProductionOrderDelegate

		errMsg = "Failed to add elements " + str(productionOrdersIds) + " for ProductionOrders on BuildSchedule"

		try:
			# get the BuildSchedule
			buildSchedule = self.get( buildScheduleId ).first()
				
			# split on a comma with no spaces
			idList = productionOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ProductionOrder		
				productionOrder = ProductionOrderDelegate().get(id).first();	
				# add the ProductionOrder
				buildSchedule.productionOrders.add(productionOrder)
				
			# save it		
			buildSchedule.save()
			
			# reload and return the appropriate version
			return self.get( buildScheduleId );
		except BuildSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : BuildSchedule with id " + str(buildScheduleId) + " does not exist.")
		except ProductionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProductionOrders( self, buildScheduleId, productionOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ProductionOrderDelegate import ProductionOrderDelegate

		errMsg = "Failed to remove elements " + str(productionOrdersIds) + " for ProductionOrders on BuildSchedule"

		try:
			# get the BuildSchedule
			buildSchedule = self.get( buildScheduleId ).first()
				
			# split on a comma with no spaces
			idList = productionOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ProductionOrder		
				productionOrder = ProductionOrderDelegate().get(id).first();	
				# add the ProductionOrder
				buildSchedule.productionOrders.remove(productionOrder)
				
			# save it		
			buildSchedule.save()
			
			# reload and return the appropriate version
			return self.get( buildScheduleId );
		except BuildSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : BuildSchedule with id " + str(buildScheduleId) + " does not exist.")
		except ProductionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
