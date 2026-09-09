from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.InsuranceProduct import InsuranceProduct
from insuranceOnDjango.models.Insurer import Insurer
from insuranceOnDjango.models.CoverageDefinition import CoverageDefinition
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InsuranceProduct
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsuranceProductDelegate Declaration
#======================================================================
class InsuranceProductDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, insuranceProductId ):
		try:	
			insuranceProduct = InsuranceProduct.objects.filter(id=insuranceProductId)
			return insuranceProduct.first();
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError("InsuranceProduct with id " + str(insuranceProductId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, insuranceProduct):
		for model in serializers.deserialize("json", insuranceProduct):
			model.save()
			return model;

	def create(self, insuranceProduct):
		insuranceProduct.save()
		return insuranceProduct;

	def saveFromJson(self, insuranceProduct):
		for model in serializers.deserialize("json", insuranceProduct):
			model.save()
			return insuranceProduct;
	
	def save(self, insuranceProduct):
		insuranceProduct.save()
		return insuranceProduct;
	
	def delete(self, insuranceProductId ):
		errMsg = "Failed to delete InsuranceProduct from db using id " + str(insuranceProductId)
		
		try:
			insuranceProduct = InsuranceProduct.objects.get(id=insuranceProductId)
			insuranceProduct.delete()
			return True
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError("InsuranceProduct with id " + str(insuranceProductId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InsuranceProduct.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InsuranceProduct from db")
		except Exception:
			return None;
		
	def assignInsurer( self, insuranceProductId, insurerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsurerDelegate import InsurerDelegate

		errMsg = "Failed to assign element " + str(insurerId) + " for Insurer on InsuranceProduct"

		try:
			# get the InsuranceProduct from db
			insuranceProduct = self.get( insuranceProductId ).first()	
			
			# get the Insurer from db
			insurer = InsurerDelegate().get(insurerId).first();
			
			# assign the Insurer		
			insuranceProduct.insurer = insurer
			
			#save it
			insuranceProduct.save()

			# reload and return the appropriate version					
			return self.get( insuranceProductId );
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuranceProduct with id " + str(insuranceProductId) + " does not exist.")
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInsurer( self, insuranceProductId ):
		errMsg = "Failed to unassign element " + str(insurerId) + " for Insurer on InsuranceProduct"

		try:
			# get the InsuranceProduct from db
			insuranceProduct = self.get( insuranceProductId ).first()	
			
			# assign to None for unassignment
			insuranceProduct.insurer = None			

			#save it
			insuranceProduct.save()

			# reload and return the appropriate version					
			return self.get( insuranceProductId );
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuranceProduct with id " + str(insuranceProductId) + " does not exist.")
		except Exception:
			return None;
		
	def addCoverageDefinitions( self, insuranceProductId, coverageDefinitionsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CoverageDefinitionDelegate import CoverageDefinitionDelegate

		errMsg = "Failed to add elements " + str(coverageDefinitionsIds) + " for CoverageDefinitions on InsuranceProduct"

		try:
			# get the InsuranceProduct
			insuranceProduct = self.get( insuranceProductId ).first()
				
			# split on a comma with no spaces
			idList = coverageDefinitionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CoverageDefinition		
				coverageDefinition = CoverageDefinitionDelegate().get(id).first();	
				# add the CoverageDefinition
				insuranceProduct.coverageDefinitions.add(coverageDefinition)
				
			# save it		
			insuranceProduct.save()
			
			# reload and return the appropriate version
			return self.get( insuranceProductId );
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuranceProduct with id " + str(insuranceProductId) + " does not exist.")
		except CoverageDefinition.DoesNotExist:
			raise ProcessingError(errMsg + " : CoverageDefinition does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCoverageDefinitions( self, insuranceProductId, coverageDefinitionsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CoverageDefinitionDelegate import CoverageDefinitionDelegate

		errMsg = "Failed to remove elements " + str(coverageDefinitionsIds) + " for CoverageDefinitions on InsuranceProduct"

		try:
			# get the InsuranceProduct
			insuranceProduct = self.get( insuranceProductId ).first()
				
			# split on a comma with no spaces
			idList = coverageDefinitionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CoverageDefinition		
				coverageDefinition = CoverageDefinitionDelegate().get(id).first();	
				# add the CoverageDefinition
				insuranceProduct.coverageDefinitions.remove(coverageDefinition)
				
			# save it		
			insuranceProduct.save()
			
			# reload and return the appropriate version
			return self.get( insuranceProductId );
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuranceProduct with id " + str(insuranceProductId) + " does not exist.")
		except CoverageDefinition.DoesNotExist:
			raise ProcessingError(errMsg + " : CoverageDefinition does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
