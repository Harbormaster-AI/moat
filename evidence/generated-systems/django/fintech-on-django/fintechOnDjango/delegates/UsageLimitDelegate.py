from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.UsageLimit import UsageLimit
from fintechOnDjango.models.PricingPlan import PricingPlan
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model UsageLimit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UsageLimitDelegate Declaration
#======================================================================
class UsageLimitDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, usageLimitId ):
		try:	
			usageLimit = UsageLimit.objects.filter(id=usageLimitId)
			return usageLimit.first();
		except UsageLimit.DoesNotExist:
			raise ProcessingError("UsageLimit with id " + str(usageLimitId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, usageLimit):
		for model in serializers.deserialize("json", usageLimit):
			model.save()
			return model;

	def create(self, usageLimit):
		usageLimit.save()
		return usageLimit;

	def saveFromJson(self, usageLimit):
		for model in serializers.deserialize("json", usageLimit):
			model.save()
			return usageLimit;
	
	def save(self, usageLimit):
		usageLimit.save()
		return usageLimit;
	
	def delete(self, usageLimitId ):
		errMsg = "Failed to delete UsageLimit from db using id " + str(usageLimitId)
		
		try:
			usageLimit = UsageLimit.objects.get(id=usageLimitId)
			usageLimit.delete()
			return True
		except UsageLimit.DoesNotExist:
			raise ProcessingError("UsageLimit with id " + str(usageLimitId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = UsageLimit.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all UsageLimit from db")
		except Exception:
			return None;
		
	def assignPricingPlan( self, usageLimitId, pricingPlanId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PricingPlanDelegate import PricingPlanDelegate

		errMsg = "Failed to assign element " + str(pricingPlanId) + " for PricingPlan on UsageLimit"

		try:
			# get the UsageLimit from db
			usageLimit = self.get( usageLimitId ).first()	
			
			# get the PricingPlan from db
			pricingPlan = PricingPlanDelegate().get(pricingPlanId).first();
			
			# assign the PricingPlan		
			usageLimit.pricingPlan = pricingPlan
			
			#save it
			usageLimit.save()

			# reload and return the appropriate version					
			return self.get( usageLimitId );
		except UsageLimit.DoesNotExist:
			raise ProcessingError(errMsg + " : UsageLimit with id " + str(usageLimitId) + " does not exist.")
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPricingPlan( self, usageLimitId ):
		errMsg = "Failed to unassign element " + str(pricingPlanId) + " for PricingPlan on UsageLimit"

		try:
			# get the UsageLimit from db
			usageLimit = self.get( usageLimitId ).first()	
			
			# assign to None for unassignment
			usageLimit.pricingPlan = None			

			#save it
			usageLimit.save()

			# reload and return the appropriate version					
			return self.get( usageLimitId );
		except UsageLimit.DoesNotExist:
			raise ProcessingError(errMsg + " : UsageLimit with id " + str(usageLimitId) + " does not exist.")
		except Exception:
			return None;
		
