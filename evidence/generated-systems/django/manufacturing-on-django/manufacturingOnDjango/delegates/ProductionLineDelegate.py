from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.ProductionLine import ProductionLine
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.WorkCenter import WorkCenter
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ProductionLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionLineDelegate Declaration
#======================================================================
class ProductionLineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, productionLineId ):
		try:	
			productionLine = ProductionLine.objects.filter(id=productionLineId)
			return productionLine.first();
		except ProductionLine.DoesNotExist:
			raise ProcessingError("ProductionLine with id " + str(productionLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, productionLine):
		for model in serializers.deserialize("json", productionLine):
			model.save()
			return model;

	def create(self, productionLine):
		productionLine.save()
		return productionLine;

	def saveFromJson(self, productionLine):
		for model in serializers.deserialize("json", productionLine):
			model.save()
			return productionLine;
	
	def save(self, productionLine):
		productionLine.save()
		return productionLine;
	
	def delete(self, productionLineId ):
		errMsg = "Failed to delete ProductionLine from db using id " + str(productionLineId)
		
		try:
			productionLine = ProductionLine.objects.get(id=productionLineId)
			productionLine.delete()
			return True
		except ProductionLine.DoesNotExist:
			raise ProcessingError("ProductionLine with id " + str(productionLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ProductionLine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ProductionLine from db")
		except Exception:
			return None;
		
	def assignPlant( self, productionLineId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on ProductionLine"

		try:
			# get the ProductionLine from db
			productionLine = self.get( productionLineId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			productionLine.plant = plant
			
			#save it
			productionLine.save()

			# reload and return the appropriate version					
			return self.get( productionLineId );
		except ProductionLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionLine with id " + str(productionLineId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, productionLineId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on ProductionLine"

		try:
			# get the ProductionLine from db
			productionLine = self.get( productionLineId ).first()	
			
			# assign to None for unassignment
			productionLine.plant = None			

			#save it
			productionLine.save()

			# reload and return the appropriate version					
			return self.get( productionLineId );
		except ProductionLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionLine with id " + str(productionLineId) + " does not exist.")
		except Exception:
			return None;
		
	def addWorkCenters( self, productionLineId, workCentersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

		errMsg = "Failed to add elements " + str(workCentersIds) + " for WorkCenters on ProductionLine"

		try:
			# get the ProductionLine
			productionLine = self.get( productionLineId ).first()
				
			# split on a comma with no spaces
			idList = workCentersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the WorkCenter		
				workCenter = WorkCenterDelegate().get(id).first();	
				# add the WorkCenter
				productionLine.workCenters.add(workCenter)
				
			# save it		
			productionLine.save()
			
			# reload and return the appropriate version
			return self.get( productionLineId );
		except ProductionLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionLine with id " + str(productionLineId) + " does not exist.")
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWorkCenters( self, productionLineId, workCentersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

		errMsg = "Failed to remove elements " + str(workCentersIds) + " for WorkCenters on ProductionLine"

		try:
			# get the ProductionLine
			productionLine = self.get( productionLineId ).first()
				
			# split on a comma with no spaces
			idList = workCentersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the WorkCenter		
				workCenter = WorkCenterDelegate().get(id).first();	
				# add the WorkCenter
				productionLine.workCenters.remove(workCenter)
				
			# save it		
			productionLine.save()
			
			# reload and return the appropriate version
			return self.get( productionLineId );
		except ProductionLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionLine with id " + str(productionLineId) + " does not exist.")
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
