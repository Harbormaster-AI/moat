from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Rate import Rate
from advertisingOnDjango.models.RateCard import RateCard
from advertisingOnDjango.models.AdSlot import AdSlot
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Rate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RateDelegate Declaration
#======================================================================
class RateDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, rateId ):
		try:	
			rate = Rate.objects.filter(id=rateId)
			return rate.first();
		except Rate.DoesNotExist:
			raise ProcessingError("Rate with id " + str(rateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, rate):
		for model in serializers.deserialize("json", rate):
			model.save()
			return model;

	def create(self, rate):
		rate.save()
		return rate;

	def saveFromJson(self, rate):
		for model in serializers.deserialize("json", rate):
			model.save()
			return rate;
	
	def save(self, rate):
		rate.save()
		return rate;
	
	def delete(self, rateId ):
		errMsg = "Failed to delete Rate from db using id " + str(rateId)
		
		try:
			rate = Rate.objects.get(id=rateId)
			rate.delete()
			return True
		except Rate.DoesNotExist:
			raise ProcessingError("Rate with id " + str(rateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Rate.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Rate from db")
		except Exception:
			return None;
		
	def assignRateCard( self, rateId, rateCardId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.RateCardDelegate import RateCardDelegate

		errMsg = "Failed to assign element " + str(rateCardId) + " for RateCard on Rate"

		try:
			# get the Rate from db
			rate = self.get( rateId ).first()	
			
			# get the RateCard from db
			rateCard = RateCardDelegate().get(rateCardId).first();
			
			# assign the RateCard		
			rate.rateCard = rateCard
			
			#save it
			rate.save()

			# reload and return the appropriate version					
			return self.get( rateId );
		except Rate.DoesNotExist:
			raise ProcessingError(errMsg + " : Rate with id " + str(rateId) + " does not exist.")
		except RateCard.DoesNotExist:
			raise ProcessingError(errMsg + " : RateCard with id " + str(rateCardId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRateCard( self, rateId ):
		errMsg = "Failed to unassign element " + str(rateCardId) + " for RateCard on Rate"

		try:
			# get the Rate from db
			rate = self.get( rateId ).first()	
			
			# assign to None for unassignment
			rate.rateCard = None			

			#save it
			rate.save()

			# reload and return the appropriate version					
			return self.get( rateId );
		except Rate.DoesNotExist:
			raise ProcessingError(errMsg + " : Rate with id " + str(rateId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAdSlot( self, rateId, adSlotId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdSlotDelegate import AdSlotDelegate

		errMsg = "Failed to assign element " + str(adSlotId) + " for AdSlot on Rate"

		try:
			# get the Rate from db
			rate = self.get( rateId ).first()	
			
			# get the AdSlot from db
			adSlot = AdSlotDelegate().get(adSlotId).first();
			
			# assign the AdSlot		
			rate.adSlot = adSlot
			
			#save it
			rate.save()

			# reload and return the appropriate version					
			return self.get( rateId );
		except Rate.DoesNotExist:
			raise ProcessingError(errMsg + " : Rate with id " + str(rateId) + " does not exist.")
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot with id " + str(adSlotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdSlot( self, rateId ):
		errMsg = "Failed to unassign element " + str(adSlotId) + " for AdSlot on Rate"

		try:
			# get the Rate from db
			rate = self.get( rateId ).first()	
			
			# assign to None for unassignment
			rate.adSlot = None			

			#save it
			rate.save()

			# reload and return the appropriate version					
			return self.get( rateId );
		except Rate.DoesNotExist:
			raise ProcessingError(errMsg + " : Rate with id " + str(rateId) + " does not exist.")
		except Exception:
			return None;
		
