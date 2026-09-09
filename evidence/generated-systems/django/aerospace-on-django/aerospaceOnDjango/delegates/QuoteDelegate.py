from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.Quote import Quote
from aerospaceOnDjango.models.AircraftOrder import AircraftOrder
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Quote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteDelegate Declaration
#======================================================================
class QuoteDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, quoteId ):
		try:	
			quote = Quote.objects.filter(id=quoteId)
			return quote.first();
		except Quote.DoesNotExist:
			raise ProcessingError("Quote with id " + str(quoteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, quote):
		for model in serializers.deserialize("json", quote):
			model.save()
			return model;

	def create(self, quote):
		quote.save()
		return quote;

	def saveFromJson(self, quote):
		for model in serializers.deserialize("json", quote):
			model.save()
			return quote;
	
	def save(self, quote):
		quote.save()
		return quote;
	
	def delete(self, quoteId ):
		errMsg = "Failed to delete Quote from db using id " + str(quoteId)
		
		try:
			quote = Quote.objects.get(id=quoteId)
			quote.delete()
			return True
		except Quote.DoesNotExist:
			raise ProcessingError("Quote with id " + str(quoteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Quote.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Quote from db")
		except Exception:
			return None;
		
	def assignAircraftOrder( self, quoteId, aircraftOrderId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOrderDelegate import AircraftOrderDelegate

		errMsg = "Failed to assign element " + str(aircraftOrderId) + " for AircraftOrder on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# get the AircraftOrder from db
			aircraftOrder = AircraftOrderDelegate().get(aircraftOrderId).first();
			
			# assign the AircraftOrder		
			quote.aircraftOrder = aircraftOrder
			
			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except AircraftOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOrder with id " + str(aircraftOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAircraftOrder( self, quoteId ):
		errMsg = "Failed to unassign element " + str(aircraftOrderId) + " for AircraftOrder on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# assign to None for unassignment
			quote.aircraftOrder = None			

			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
		
