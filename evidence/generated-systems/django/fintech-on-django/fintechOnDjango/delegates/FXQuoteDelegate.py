from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.FXQuote import FXQuote
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model FXQuote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXQuoteDelegate Declaration
#======================================================================
class FXQuoteDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, fXQuoteId ):
		try:	
			fXQuote = FXQuote.objects.filter(id=fXQuoteId)
			return fXQuote.first();
		except FXQuote.DoesNotExist:
			raise ProcessingError("FXQuote with id " + str(fXQuoteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, fXQuote):
		for model in serializers.deserialize("json", fXQuote):
			model.save()
			return model;

	def create(self, fXQuote):
		fXQuote.save()
		return fXQuote;

	def saveFromJson(self, fXQuote):
		for model in serializers.deserialize("json", fXQuote):
			model.save()
			return fXQuote;
	
	def save(self, fXQuote):
		fXQuote.save()
		return fXQuote;
	
	def delete(self, fXQuoteId ):
		errMsg = "Failed to delete FXQuote from db using id " + str(fXQuoteId)
		
		try:
			fXQuote = FXQuote.objects.get(id=fXQuoteId)
			fXQuote.delete()
			return True
		except FXQuote.DoesNotExist:
			raise ProcessingError("FXQuote with id " + str(fXQuoteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = FXQuote.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all FXQuote from db")
		except Exception:
			return None;
		
	def assignRequestedBy( self, fXQuoteId, requestedById ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(requestedById) + " for RequestedBy on FXQuote"

		try:
			# get the FXQuote from db
			fXQuote = self.get( fXQuoteId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(requestedById).first();
			
			# assign the RequestedBy		
			fXQuote.requestedBy = customer
			
			#save it
			fXQuote.save()

			# reload and return the appropriate version					
			return self.get( fXQuoteId );
		except FXQuote.DoesNotExist:
			raise ProcessingError(errMsg + " : FXQuote with id " + str(fXQuoteId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(requestedById) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRequestedBy( self, fXQuoteId ):
		errMsg = "Failed to unassign element " + str(requestedById) + " for RequestedBy on FXQuote"

		try:
			# get the FXQuote from db
			fXQuote = self.get( fXQuoteId ).first()	
			
			# assign to None for unassignment
			fXQuote.customer = None			

			#save it
			fXQuote.save()

			# reload and return the appropriate version					
			return self.get( fXQuoteId );
		except FXQuote.DoesNotExist:
			raise ProcessingError(errMsg + " : FXQuote with id " + str(fXQuoteId) + " does not exist.")
		except Exception:
			return None;
		
