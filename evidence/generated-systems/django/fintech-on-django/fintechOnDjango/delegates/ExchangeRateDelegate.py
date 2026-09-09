from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.ExchangeRate import ExchangeRate
from fintechOnDjango.models.FXQuote import FXQuote
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ExchangeRate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExchangeRateDelegate Declaration
#======================================================================
class ExchangeRateDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, exchangeRateId ):
		try:	
			exchangeRate = ExchangeRate.objects.filter(id=exchangeRateId)
			return exchangeRate.first();
		except ExchangeRate.DoesNotExist:
			raise ProcessingError("ExchangeRate with id " + str(exchangeRateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, exchangeRate):
		for model in serializers.deserialize("json", exchangeRate):
			model.save()
			return model;

	def create(self, exchangeRate):
		exchangeRate.save()
		return exchangeRate;

	def saveFromJson(self, exchangeRate):
		for model in serializers.deserialize("json", exchangeRate):
			model.save()
			return exchangeRate;
	
	def save(self, exchangeRate):
		exchangeRate.save()
		return exchangeRate;
	
	def delete(self, exchangeRateId ):
		errMsg = "Failed to delete ExchangeRate from db using id " + str(exchangeRateId)
		
		try:
			exchangeRate = ExchangeRate.objects.get(id=exchangeRateId)
			exchangeRate.delete()
			return True
		except ExchangeRate.DoesNotExist:
			raise ProcessingError("ExchangeRate with id " + str(exchangeRateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ExchangeRate.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ExchangeRate from db")
		except Exception:
			return None;
		
	def addUsedByQuotes( self, exchangeRateId, usedByQuotesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FXQuoteDelegate import FXQuoteDelegate

		errMsg = "Failed to add elements " + str(usedByQuotesIds) + " for UsedByQuotes on ExchangeRate"

		try:
			# get the ExchangeRate
			exchangeRate = self.get( exchangeRateId ).first()
				
			# split on a comma with no spaces
			idList = usedByQuotesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FXQuote		
				fXQuote = FXQuoteDelegate().get(id).first();	
				# add the FXQuote
				exchangeRate.usedByQuotes.add(fXQuote)
				
			# save it		
			exchangeRate.save()
			
			# reload and return the appropriate version
			return self.get( exchangeRateId );
		except ExchangeRate.DoesNotExist:
			raise ProcessingError(errMsg + " : ExchangeRate with id " + str(exchangeRateId) + " does not exist.")
		except FXQuote.DoesNotExist:
			raise ProcessingError(errMsg + " : FXQuote does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeUsedByQuotes( self, exchangeRateId, usedByQuotesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FXQuoteDelegate import FXQuoteDelegate

		errMsg = "Failed to remove elements " + str(usedByQuotesIds) + " for UsedByQuotes on ExchangeRate"

		try:
			# get the ExchangeRate
			exchangeRate = self.get( exchangeRateId ).first()
				
			# split on a comma with no spaces
			idList = usedByQuotesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FXQuote		
				fXQuote = FXQuoteDelegate().get(id).first();	
				# add the FXQuote
				exchangeRate.usedByQuotes.remove(fXQuote)
				
			# save it		
			exchangeRate.save()
			
			# reload and return the appropriate version
			return self.get( exchangeRateId );
		except ExchangeRate.DoesNotExist:
			raise ProcessingError(errMsg + " : ExchangeRate with id " + str(exchangeRateId) + " does not exist.")
		except FXQuote.DoesNotExist:
			raise ProcessingError(errMsg + " : FXQuote does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
