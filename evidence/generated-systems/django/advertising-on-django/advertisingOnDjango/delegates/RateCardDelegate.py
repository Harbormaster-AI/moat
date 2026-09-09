from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.RateCard import RateCard
from advertisingOnDjango.models.Publisher import Publisher
from advertisingOnDjango.models.Rate import Rate
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model RateCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RateCardDelegate Declaration
#======================================================================
class RateCardDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, rateCardId ):
		try:	
			rateCard = RateCard.objects.filter(id=rateCardId)
			return rateCard.first();
		except RateCard.DoesNotExist:
			raise ProcessingError("RateCard with id " + str(rateCardId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, rateCard):
		for model in serializers.deserialize("json", rateCard):
			model.save()
			return model;

	def create(self, rateCard):
		rateCard.save()
		return rateCard;

	def saveFromJson(self, rateCard):
		for model in serializers.deserialize("json", rateCard):
			model.save()
			return rateCard;
	
	def save(self, rateCard):
		rateCard.save()
		return rateCard;
	
	def delete(self, rateCardId ):
		errMsg = "Failed to delete RateCard from db using id " + str(rateCardId)
		
		try:
			rateCard = RateCard.objects.get(id=rateCardId)
			rateCard.delete()
			return True
		except RateCard.DoesNotExist:
			raise ProcessingError("RateCard with id " + str(rateCardId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = RateCard.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all RateCard from db")
		except Exception:
			return None;
		
	def assignPublisher( self, rateCardId, publisherId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PublisherDelegate import PublisherDelegate

		errMsg = "Failed to assign element " + str(publisherId) + " for Publisher on RateCard"

		try:
			# get the RateCard from db
			rateCard = self.get( rateCardId ).first()	
			
			# get the Publisher from db
			publisher = PublisherDelegate().get(publisherId).first();
			
			# assign the Publisher		
			rateCard.publisher = publisher
			
			#save it
			rateCard.save()

			# reload and return the appropriate version					
			return self.get( rateCardId );
		except RateCard.DoesNotExist:
			raise ProcessingError(errMsg + " : RateCard with id " + str(rateCardId) + " does not exist.")
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPublisher( self, rateCardId ):
		errMsg = "Failed to unassign element " + str(publisherId) + " for Publisher on RateCard"

		try:
			# get the RateCard from db
			rateCard = self.get( rateCardId ).first()	
			
			# assign to None for unassignment
			rateCard.publisher = None			

			#save it
			rateCard.save()

			# reload and return the appropriate version					
			return self.get( rateCardId );
		except RateCard.DoesNotExist:
			raise ProcessingError(errMsg + " : RateCard with id " + str(rateCardId) + " does not exist.")
		except Exception:
			return None;
		
	def addRates( self, rateCardId, ratesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.RateDelegate import RateDelegate

		errMsg = "Failed to add elements " + str(ratesIds) + " for Rates on RateCard"

		try:
			# get the RateCard
			rateCard = self.get( rateCardId ).first()
				
			# split on a comma with no spaces
			idList = ratesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Rate		
				rate = RateDelegate().get(id).first();	
				# add the Rate
				rateCard.rates.add(rate)
				
			# save it		
			rateCard.save()
			
			# reload and return the appropriate version
			return self.get( rateCardId );
		except RateCard.DoesNotExist:
			raise ProcessingError(errMsg + " : RateCard with id " + str(rateCardId) + " does not exist.")
		except Rate.DoesNotExist:
			raise ProcessingError(errMsg + " : Rate does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRates( self, rateCardId, ratesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.RateDelegate import RateDelegate

		errMsg = "Failed to remove elements " + str(ratesIds) + " for Rates on RateCard"

		try:
			# get the RateCard
			rateCard = self.get( rateCardId ).first()
				
			# split on a comma with no spaces
			idList = ratesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Rate		
				rate = RateDelegate().get(id).first();	
				# add the Rate
				rateCard.rates.remove(rate)
				
			# save it		
			rateCard.save()
			
			# reload and return the appropriate version
			return self.get( rateCardId );
		except RateCard.DoesNotExist:
			raise ProcessingError(errMsg + " : RateCard with id " + str(rateCardId) + " does not exist.")
		except Rate.DoesNotExist:
			raise ProcessingError(errMsg + " : Rate does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
