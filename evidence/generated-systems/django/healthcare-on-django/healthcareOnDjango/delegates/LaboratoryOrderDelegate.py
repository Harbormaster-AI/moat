from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.LaboratoryOrder import LaboratoryOrder
from healthcareOnDjango.models.ClinicalOrder import ClinicalOrder
from healthcareOnDjango.models.Laboratory import Laboratory
from healthcareOnDjango.models.LabResult import LabResult
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LaboratoryOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LaboratoryOrderDelegate Declaration
#======================================================================
class LaboratoryOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, laboratoryOrderId ):
		try:	
			laboratoryOrder = LaboratoryOrder.objects.filter(id=laboratoryOrderId)
			return laboratoryOrder.first();
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError("LaboratoryOrder with id " + str(laboratoryOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, laboratoryOrder):
		for model in serializers.deserialize("json", laboratoryOrder):
			model.save()
			return model;

	def create(self, laboratoryOrder):
		laboratoryOrder.save()
		return laboratoryOrder;

	def saveFromJson(self, laboratoryOrder):
		for model in serializers.deserialize("json", laboratoryOrder):
			model.save()
			return laboratoryOrder;
	
	def save(self, laboratoryOrder):
		laboratoryOrder.save()
		return laboratoryOrder;
	
	def delete(self, laboratoryOrderId ):
		errMsg = "Failed to delete LaboratoryOrder from db using id " + str(laboratoryOrderId)
		
		try:
			laboratoryOrder = LaboratoryOrder.objects.get(id=laboratoryOrderId)
			laboratoryOrder.delete()
			return True
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError("LaboratoryOrder with id " + str(laboratoryOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LaboratoryOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LaboratoryOrder from db")
		except Exception:
			return None;
		
	def assignOrder( self, laboratoryOrderId, orderId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicalOrderDelegate import ClinicalOrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on LaboratoryOrder"

		try:
			# get the LaboratoryOrder from db
			laboratoryOrder = self.get( laboratoryOrderId ).first()	
			
			# get the ClinicalOrder from db
			clinicalOrder = ClinicalOrderDelegate().get(orderId).first();
			
			# assign the Order		
			laboratoryOrder.order = clinicalOrder
			
			#save it
			laboratoryOrder.save()

			# reload and return the appropriate version					
			return self.get( laboratoryOrderId );
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder with id " + str(laboratoryOrderId) + " does not exist.")
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, laboratoryOrderId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on LaboratoryOrder"

		try:
			# get the LaboratoryOrder from db
			laboratoryOrder = self.get( laboratoryOrderId ).first()	
			
			# assign to None for unassignment
			laboratoryOrder.clinicalOrder = None			

			#save it
			laboratoryOrder.save()

			# reload and return the appropriate version					
			return self.get( laboratoryOrderId );
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder with id " + str(laboratoryOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLaboratory( self, laboratoryOrderId, laboratoryId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryDelegate import LaboratoryDelegate

		errMsg = "Failed to assign element " + str(laboratoryId) + " for Laboratory on LaboratoryOrder"

		try:
			# get the LaboratoryOrder from db
			laboratoryOrder = self.get( laboratoryOrderId ).first()	
			
			# get the Laboratory from db
			laboratory = LaboratoryDelegate().get(laboratoryId).first();
			
			# assign the Laboratory		
			laboratoryOrder.laboratory = laboratory
			
			#save it
			laboratoryOrder.save()

			# reload and return the appropriate version					
			return self.get( laboratoryOrderId );
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder with id " + str(laboratoryOrderId) + " does not exist.")
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory with id " + str(laboratoryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLaboratory( self, laboratoryOrderId ):
		errMsg = "Failed to unassign element " + str(laboratoryId) + " for Laboratory on LaboratoryOrder"

		try:
			# get the LaboratoryOrder from db
			laboratoryOrder = self.get( laboratoryOrderId ).first()	
			
			# assign to None for unassignment
			laboratoryOrder.laboratory = None			

			#save it
			laboratoryOrder.save()

			# reload and return the appropriate version					
			return self.get( laboratoryOrderId );
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder with id " + str(laboratoryOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addResults( self, laboratoryOrderId, resultsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LabResultDelegate import LabResultDelegate

		errMsg = "Failed to add elements " + str(resultsIds) + " for Results on LaboratoryOrder"

		try:
			# get the LaboratoryOrder
			laboratoryOrder = self.get( laboratoryOrderId ).first()
				
			# split on a comma with no spaces
			idList = resultsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LabResult		
				labResult = LabResultDelegate().get(id).first();	
				# add the LabResult
				laboratoryOrder.results.add(labResult)
				
			# save it		
			laboratoryOrder.save()
			
			# reload and return the appropriate version
			return self.get( laboratoryOrderId );
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder with id " + str(laboratoryOrderId) + " does not exist.")
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeResults( self, laboratoryOrderId, resultsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LabResultDelegate import LabResultDelegate

		errMsg = "Failed to remove elements " + str(resultsIds) + " for Results on LaboratoryOrder"

		try:
			# get the LaboratoryOrder
			laboratoryOrder = self.get( laboratoryOrderId ).first()
				
			# split on a comma with no spaces
			idList = resultsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LabResult		
				labResult = LabResultDelegate().get(id).first();	
				# add the LabResult
				laboratoryOrder.results.remove(labResult)
				
			# save it		
			laboratoryOrder.save()
			
			# reload and return the appropriate version
			return self.get( laboratoryOrderId );
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder with id " + str(laboratoryOrderId) + " does not exist.")
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
