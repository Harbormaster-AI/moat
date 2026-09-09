from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.FeeSchedule import FeeSchedule
from fintechOnDjango.models.PricingPlan import PricingPlan
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model FeeSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeeScheduleDelegate Declaration
#======================================================================
class FeeScheduleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, feeScheduleId ):
		try:	
			feeSchedule = FeeSchedule.objects.filter(id=feeScheduleId)
			return feeSchedule.first();
		except FeeSchedule.DoesNotExist:
			raise ProcessingError("FeeSchedule with id " + str(feeScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, feeSchedule):
		for model in serializers.deserialize("json", feeSchedule):
			model.save()
			return model;

	def create(self, feeSchedule):
		feeSchedule.save()
		return feeSchedule;

	def saveFromJson(self, feeSchedule):
		for model in serializers.deserialize("json", feeSchedule):
			model.save()
			return feeSchedule;
	
	def save(self, feeSchedule):
		feeSchedule.save()
		return feeSchedule;
	
	def delete(self, feeScheduleId ):
		errMsg = "Failed to delete FeeSchedule from db using id " + str(feeScheduleId)
		
		try:
			feeSchedule = FeeSchedule.objects.get(id=feeScheduleId)
			feeSchedule.delete()
			return True
		except FeeSchedule.DoesNotExist:
			raise ProcessingError("FeeSchedule with id " + str(feeScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = FeeSchedule.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all FeeSchedule from db")
		except Exception:
			return None;
		
	def assignPricingPlan( self, feeScheduleId, pricingPlanId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PricingPlanDelegate import PricingPlanDelegate

		errMsg = "Failed to assign element " + str(pricingPlanId) + " for PricingPlan on FeeSchedule"

		try:
			# get the FeeSchedule from db
			feeSchedule = self.get( feeScheduleId ).first()	
			
			# get the PricingPlan from db
			pricingPlan = PricingPlanDelegate().get(pricingPlanId).first();
			
			# assign the PricingPlan		
			feeSchedule.pricingPlan = pricingPlan
			
			#save it
			feeSchedule.save()

			# reload and return the appropriate version					
			return self.get( feeScheduleId );
		except FeeSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : FeeSchedule with id " + str(feeScheduleId) + " does not exist.")
		except PricingPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : PricingPlan with id " + str(pricingPlanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPricingPlan( self, feeScheduleId ):
		errMsg = "Failed to unassign element " + str(pricingPlanId) + " for PricingPlan on FeeSchedule"

		try:
			# get the FeeSchedule from db
			feeSchedule = self.get( feeScheduleId ).first()	
			
			# assign to None for unassignment
			feeSchedule.pricingPlan = None			

			#save it
			feeSchedule.save()

			# reload and return the appropriate version					
			return self.get( feeScheduleId );
		except FeeSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : FeeSchedule with id " + str(feeScheduleId) + " does not exist.")
		except Exception:
			return None;
		
