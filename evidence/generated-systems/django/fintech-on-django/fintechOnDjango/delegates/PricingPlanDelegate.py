from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.PricingPlan import PricingPlan
from fintechOnDjango.models.ProductOffering import ProductOffering
from fintechOnDjango.models.FeeSchedule import FeeSchedule
from fintechOnDjango.models.UsageLimit import UsageLimit
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PricingPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PricingPlanDelegate Declaration
#======================================================================
class PricingPlanDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, pricingPlanId ):
		try:	
			pricingPlan = PricingPlan.objects.filter(id=pricingPlanId)
			return pricingPlan.first();
		except PricingPlan.DoesNotExist:
			raise ProcessingError("PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, pricingPlan):
		for model in serializers.deserialize("json", pricingPlan):
			model.save()
			return model;

	def create(self, pricingPlan):
		pricingPlan.save()
		return pricingPlan;

	def saveFromJson(self, pricingPlan):
		for model in serializers.deserialize("json", pricingPlan):
			model.save()
			return pricingPlan;
	
	def save(self, pricingPlan):
		pricingPlan.save()
		return pricingPlan;
	
	def delete(self, pricingPlanId ):
		errMsg = "Failed to delete PricingPlan from db using id " + str(pricingPlanId)
		
		try:
			pricingPlan = PricingPlan.objects.get(id=pricingPlanId)
			pricingPlan.delete()
			return True
		except PricingPlan.DoesNotExist:
			raise ProcessingError("PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PricingPlan.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PricingPlan from db")
		except Exception:
			return None;
		
	def assignProductOffering( self, pricingPlanId, productOfferingId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ProductOfferingDelegate import ProductOfferingDelegate

		errMsg = "Failed to assign element " + str(productOfferingId) + " for ProductOffering on PricingPlan"

		try:
			# get the PricingPlan from db
			pricingPlan = self.get( pricingPlanId ).first()	
			
			# get the ProductOffering from db
			productOffering = ProductOfferingDelegate().get(productOfferingId).first();
			
			# assign the ProductOffering		
			pricingPlan.productOffering = productOffering
			
			#save it
			pricingPlan.save()

			# reload and return the appropriate version					
			return self.get( pricingPlanId );
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except ProductOffering.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductOffering with id " + str(productOfferingId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProductOffering( self, pricingPlanId ):
		errMsg = "Failed to unassign element " + str(productOfferingId) + " for ProductOffering on PricingPlan"

		try:
			# get the PricingPlan from db
			pricingPlan = self.get( pricingPlanId ).first()	
			
			# assign to None for unassignment
			pricingPlan.productOffering = None			

			#save it
			pricingPlan.save()

			# reload and return the appropriate version					
			return self.get( pricingPlanId );
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except Exception:
			return None;
		
	def addFeeSchedules( self, pricingPlanId, feeSchedulesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FeeScheduleDelegate import FeeScheduleDelegate

		errMsg = "Failed to add elements " + str(feeSchedulesIds) + " for FeeSchedules on PricingPlan"

		try:
			# get the PricingPlan
			pricingPlan = self.get( pricingPlanId ).first()
				
			# split on a comma with no spaces
			idList = feeSchedulesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FeeSchedule		
				feeSchedule = FeeScheduleDelegate().get(id).first();	
				# add the FeeSchedule
				pricingPlan.feeSchedules.add(feeSchedule)
				
			# save it		
			pricingPlan.save()
			
			# reload and return the appropriate version
			return self.get( pricingPlanId );
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except FeeSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : FeeSchedule does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFeeSchedules( self, pricingPlanId, feeSchedulesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FeeScheduleDelegate import FeeScheduleDelegate

		errMsg = "Failed to remove elements " + str(feeSchedulesIds) + " for FeeSchedules on PricingPlan"

		try:
			# get the PricingPlan
			pricingPlan = self.get( pricingPlanId ).first()
				
			# split on a comma with no spaces
			idList = feeSchedulesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FeeSchedule		
				feeSchedule = FeeScheduleDelegate().get(id).first();	
				# add the FeeSchedule
				pricingPlan.feeSchedules.remove(feeSchedule)
				
			# save it		
			pricingPlan.save()
			
			# reload and return the appropriate version
			return self.get( pricingPlanId );
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except FeeSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : FeeSchedule does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLimits( self, pricingPlanId, limitsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.UsageLimitDelegate import UsageLimitDelegate

		errMsg = "Failed to add elements " + str(limitsIds) + " for Limits on PricingPlan"

		try:
			# get the PricingPlan
			pricingPlan = self.get( pricingPlanId ).first()
				
			# split on a comma with no spaces
			idList = limitsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the UsageLimit		
				usageLimit = UsageLimitDelegate().get(id).first();	
				# add the UsageLimit
				pricingPlan.limits.add(usageLimit)
				
			# save it		
			pricingPlan.save()
			
			# reload and return the appropriate version
			return self.get( pricingPlanId );
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except UsageLimit.DoesNotExist:
			raise ProcessingError(errMsg + " : UsageLimit does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLimits( self, pricingPlanId, limitsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.UsageLimitDelegate import UsageLimitDelegate

		errMsg = "Failed to remove elements " + str(limitsIds) + " for Limits on PricingPlan"

		try:
			# get the PricingPlan
			pricingPlan = self.get( pricingPlanId ).first()
				
			# split on a comma with no spaces
			idList = limitsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the UsageLimit		
				usageLimit = UsageLimitDelegate().get(id).first();	
				# add the UsageLimit
				pricingPlan.limits.remove(usageLimit)
				
			# save it		
			pricingPlan.save()
			
			# reload and return the appropriate version
			return self.get( pricingPlanId );
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except UsageLimit.DoesNotExist:
			raise ProcessingError(errMsg + " : UsageLimit does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
