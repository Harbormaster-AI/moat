from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.ProductOffering import ProductOffering
from fintechOnDjango.models.FinancialInstitution import FinancialInstitution
from fintechOnDjango.models.PricingPlan import PricingPlan
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ProductOffering
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductOfferingDelegate Declaration
#======================================================================
class ProductOfferingDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, productOfferingId ):
		try:	
			productOffering = ProductOffering.objects.filter(id=productOfferingId)
			return productOffering.first();
		except ProductOffering.DoesNotExist:
			raise ProcessingError("ProductOffering with id " + str(productOfferingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, productOffering):
		for model in serializers.deserialize("json", productOffering):
			model.save()
			return model;

	def create(self, productOffering):
		productOffering.save()
		return productOffering;

	def saveFromJson(self, productOffering):
		for model in serializers.deserialize("json", productOffering):
			model.save()
			return productOffering;
	
	def save(self, productOffering):
		productOffering.save()
		return productOffering;
	
	def delete(self, productOfferingId ):
		errMsg = "Failed to delete ProductOffering from db using id " + str(productOfferingId)
		
		try:
			productOffering = ProductOffering.objects.get(id=productOfferingId)
			productOffering.delete()
			return True
		except ProductOffering.DoesNotExist:
			raise ProcessingError("ProductOffering with id " + str(productOfferingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ProductOffering.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ProductOffering from db")
		except Exception:
			return None;
		
	def assignInstitution( self, productOfferingId, institutionId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FinancialInstitutionDelegate import FinancialInstitutionDelegate

		errMsg = "Failed to assign element " + str(institutionId) + " for Institution on ProductOffering"

		try:
			# get the ProductOffering from db
			productOffering = self.get( productOfferingId ).first()	
			
			# get the FinancialInstitution from db
			financialInstitution = FinancialInstitutionDelegate().get(institutionId).first();
			
			# assign the Institution		
			productOffering.institution = financialInstitution
			
			#save it
			productOffering.save()

			# reload and return the appropriate version					
			return self.get( productOfferingId );
		except ProductOffering.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductOffering with id " + str(productOfferingId) + " does not exist.")
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(institutionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInstitution( self, productOfferingId ):
		errMsg = "Failed to unassign element " + str(institutionId) + " for Institution on ProductOffering"

		try:
			# get the ProductOffering from db
			productOffering = self.get( productOfferingId ).first()	
			
			# assign to None for unassignment
			productOffering.financialInstitution = None			

			#save it
			productOffering.save()

			# reload and return the appropriate version					
			return self.get( productOfferingId );
		except ProductOffering.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductOffering with id " + str(productOfferingId) + " does not exist.")
		except Exception:
			return None;
		
	def addPricingPlans( self, productOfferingId, pricingPlansIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PricingPlanDelegate import PricingPlanDelegate

		errMsg = "Failed to add elements " + str(pricingPlansIds) + " for PricingPlans on ProductOffering"

		try:
			# get the ProductOffering
			productOffering = self.get( productOfferingId ).first()
				
			# split on a comma with no spaces
			idList = pricingPlansIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PricingPlan		
				pricingPlan = PricingPlanDelegate().get(id).first();	
				# add the PricingPlan
				productOffering.pricingPlans.add(pricingPlan)
				
			# save it		
			productOffering.save()
			
			# reload and return the appropriate version
			return self.get( productOfferingId );
		except ProductOffering.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductOffering with id " + str(productOfferingId) + " does not exist.")
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePricingPlans( self, productOfferingId, pricingPlansIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PricingPlanDelegate import PricingPlanDelegate

		errMsg = "Failed to remove elements " + str(pricingPlansIds) + " for PricingPlans on ProductOffering"

		try:
			# get the ProductOffering
			productOffering = self.get( productOfferingId ).first()
				
			# split on a comma with no spaces
			idList = pricingPlansIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PricingPlan		
				pricingPlan = PricingPlanDelegate().get(id).first();	
				# add the PricingPlan
				productOffering.pricingPlans.remove(pricingPlan)
				
			# save it		
			productOffering.save()
			
			# reload and return the appropriate version
			return self.get( productOfferingId );
		except ProductOffering.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductOffering with id " + str(productOfferingId) + " does not exist.")
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
