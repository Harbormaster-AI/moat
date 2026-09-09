from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.CoverageDefinition import CoverageDefinition
from insuranceOnDjango.models.InsuranceProduct import InsuranceProduct
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CoverageDefinition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageDefinitionDelegate Declaration
#======================================================================
class CoverageDefinitionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, coverageDefinitionId ):
		try:	
			coverageDefinition = CoverageDefinition.objects.filter(id=coverageDefinitionId)
			return coverageDefinition.first();
		except CoverageDefinition.DoesNotExist:
			raise ProcessingError("CoverageDefinition with id " + str(coverageDefinitionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, coverageDefinition):
		for model in serializers.deserialize("json", coverageDefinition):
			model.save()
			return model;

	def create(self, coverageDefinition):
		coverageDefinition.save()
		return coverageDefinition;

	def saveFromJson(self, coverageDefinition):
		for model in serializers.deserialize("json", coverageDefinition):
			model.save()
			return coverageDefinition;
	
	def save(self, coverageDefinition):
		coverageDefinition.save()
		return coverageDefinition;
	
	def delete(self, coverageDefinitionId ):
		errMsg = "Failed to delete CoverageDefinition from db using id " + str(coverageDefinitionId)
		
		try:
			coverageDefinition = CoverageDefinition.objects.get(id=coverageDefinitionId)
			coverageDefinition.delete()
			return True
		except CoverageDefinition.DoesNotExist:
			raise ProcessingError("CoverageDefinition with id " + str(coverageDefinitionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CoverageDefinition.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CoverageDefinition from db")
		except Exception:
			return None;
		
	def assignProduct( self, coverageDefinitionId, productId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuranceProductDelegate import InsuranceProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on CoverageDefinition"

		try:
			# get the CoverageDefinition from db
			coverageDefinition = self.get( coverageDefinitionId ).first()	
			
			# get the InsuranceProduct from db
			insuranceProduct = InsuranceProductDelegate().get(productId).first();
			
			# assign the Product		
			coverageDefinition.product = insuranceProduct
			
			#save it
			coverageDefinition.save()

			# reload and return the appropriate version					
			return self.get( coverageDefinitionId );
		except CoverageDefinition.DoesNotExist:
			raise ProcessingError(errMsg + " : CoverageDefinition with id " + str(coverageDefinitionId) + " does not exist.")
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuranceProduct with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, coverageDefinitionId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on CoverageDefinition"

		try:
			# get the CoverageDefinition from db
			coverageDefinition = self.get( coverageDefinitionId ).first()	
			
			# assign to None for unassignment
			coverageDefinition.insuranceProduct = None			

			#save it
			coverageDefinition.save()

			# reload and return the appropriate version					
			return self.get( coverageDefinitionId );
		except CoverageDefinition.DoesNotExist:
			raise ProcessingError(errMsg + " : CoverageDefinition with id " + str(coverageDefinitionId) + " does not exist.")
		except Exception:
			return None;
		
