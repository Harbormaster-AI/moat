from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.WorkCenter import WorkCenter
from aerospaceOnDjango.models.ProductionLine import ProductionLine
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model WorkCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkCenterDelegate Declaration
#======================================================================
class WorkCenterDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, workCenterId ):
		try:	
			workCenter = WorkCenter.objects.filter(id=workCenterId)
			return workCenter.first();
		except WorkCenter.DoesNotExist:
			raise ProcessingError("WorkCenter with id " + str(workCenterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, workCenter):
		for model in serializers.deserialize("json", workCenter):
			model.save()
			return model;

	def create(self, workCenter):
		workCenter.save()
		return workCenter;

	def saveFromJson(self, workCenter):
		for model in serializers.deserialize("json", workCenter):
			model.save()
			return workCenter;
	
	def save(self, workCenter):
		workCenter.save()
		return workCenter;
	
	def delete(self, workCenterId ):
		errMsg = "Failed to delete WorkCenter from db using id " + str(workCenterId)
		
		try:
			workCenter = WorkCenter.objects.get(id=workCenterId)
			workCenter.delete()
			return True
		except WorkCenter.DoesNotExist:
			raise ProcessingError("WorkCenter with id " + str(workCenterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = WorkCenter.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all WorkCenter from db")
		except Exception:
			return None;
		
	def assignProductionLine( self, workCenterId, productionLineId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ProductionLineDelegate import ProductionLineDelegate

		errMsg = "Failed to assign element " + str(productionLineId) + " for ProductionLine on WorkCenter"

		try:
			# get the WorkCenter from db
			workCenter = self.get( workCenterId ).first()	
			
			# get the ProductionLine from db
			productionLine = ProductionLineDelegate().get(productionLineId).first();
			
			# assign the ProductionLine		
			workCenter.productionLine = productionLine
			
			#save it
			workCenter.save()

			# reload and return the appropriate version					
			return self.get( workCenterId );
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter with id " + str(workCenterId) + " does not exist.")
		except ProductionLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionLine with id " + str(productionLineId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProductionLine( self, workCenterId ):
		errMsg = "Failed to unassign element " + str(productionLineId) + " for ProductionLine on WorkCenter"

		try:
			# get the WorkCenter from db
			workCenter = self.get( workCenterId ).first()	
			
			# assign to None for unassignment
			workCenter.productionLine = None			

			#save it
			workCenter.save()

			# reload and return the appropriate version					
			return self.get( workCenterId );
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter with id " + str(workCenterId) + " does not exist.")
		except Exception:
			return None;
		
