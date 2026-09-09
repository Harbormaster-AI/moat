from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.FXDeal import FXDeal
from fintechOnDjango.models.FXQuote import FXQuote
from fintechOnDjango.models.PaymentOrder import PaymentOrder
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model FXDeal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXDealDelegate Declaration
#======================================================================
class FXDealDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, fXDealId ):
		try:	
			fXDeal = FXDeal.objects.filter(id=fXDealId)
			return fXDeal.first();
		except FXDeal.DoesNotExist:
			raise ProcessingError("FXDeal with id " + str(fXDealId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, fXDeal):
		for model in serializers.deserialize("json", fXDeal):
			model.save()
			return model;

	def create(self, fXDeal):
		fXDeal.save()
		return fXDeal;

	def saveFromJson(self, fXDeal):
		for model in serializers.deserialize("json", fXDeal):
			model.save()
			return fXDeal;
	
	def save(self, fXDeal):
		fXDeal.save()
		return fXDeal;
	
	def delete(self, fXDealId ):
		errMsg = "Failed to delete FXDeal from db using id " + str(fXDealId)
		
		try:
			fXDeal = FXDeal.objects.get(id=fXDealId)
			fXDeal.delete()
			return True
		except FXDeal.DoesNotExist:
			raise ProcessingError("FXDeal with id " + str(fXDealId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = FXDeal.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all FXDeal from db")
		except Exception:
			return None;
		
	def assignQuote( self, fXDealId, quoteId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FXQuoteDelegate import FXQuoteDelegate

		errMsg = "Failed to assign element " + str(quoteId) + " for Quote on FXDeal"

		try:
			# get the FXDeal from db
			fXDeal = self.get( fXDealId ).first()	
			
			# get the FXQuote from db
			fXQuote = FXQuoteDelegate().get(quoteId).first();
			
			# assign the Quote		
			fXDeal.quote = fXQuote
			
			#save it
			fXDeal.save()

			# reload and return the appropriate version					
			return self.get( fXDealId );
		except FXDeal.DoesNotExist:
			raise ProcessingError(errMsg + " : FXDeal with id " + str(fXDealId) + " does not exist.")
		except FXQuote.DoesNotExist:
			raise ProcessingError(errMsg + " : FXQuote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignQuote( self, fXDealId ):
		errMsg = "Failed to unassign element " + str(quoteId) + " for Quote on FXDeal"

		try:
			# get the FXDeal from db
			fXDeal = self.get( fXDealId ).first()	
			
			# assign to None for unassignment
			fXDeal.fXQuote = None			

			#save it
			fXDeal.save()

			# reload and return the appropriate version					
			return self.get( fXDealId );
		except FXDeal.DoesNotExist:
			raise ProcessingError(errMsg + " : FXDeal with id " + str(fXDealId) + " does not exist.")
		except Exception:
			return None;
		
	def addPaymentOrders( self, fXDealId, paymentOrdersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentOrderDelegate import PaymentOrderDelegate

		errMsg = "Failed to add elements " + str(paymentOrdersIds) + " for PaymentOrders on FXDeal"

		try:
			# get the FXDeal
			fXDeal = self.get( fXDealId ).first()
				
			# split on a comma with no spaces
			idList = paymentOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentOrder		
				paymentOrder = PaymentOrderDelegate().get(id).first();	
				# add the PaymentOrder
				fXDeal.paymentOrders.add(paymentOrder)
				
			# save it		
			fXDeal.save()
			
			# reload and return the appropriate version
			return self.get( fXDealId );
		except FXDeal.DoesNotExist:
			raise ProcessingError(errMsg + " : FXDeal with id " + str(fXDealId) + " does not exist.")
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePaymentOrders( self, fXDealId, paymentOrdersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentOrderDelegate import PaymentOrderDelegate

		errMsg = "Failed to remove elements " + str(paymentOrdersIds) + " for PaymentOrders on FXDeal"

		try:
			# get the FXDeal
			fXDeal = self.get( fXDealId ).first()
				
			# split on a comma with no spaces
			idList = paymentOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentOrder		
				paymentOrder = PaymentOrderDelegate().get(id).first();	
				# add the PaymentOrder
				fXDeal.paymentOrders.remove(paymentOrder)
				
			# save it		
			fXDeal.save()
			
			# reload and return the appropriate version
			return self.get( fXDealId );
		except FXDeal.DoesNotExist:
			raise ProcessingError(errMsg + " : FXDeal with id " + str(fXDealId) + " does not exist.")
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
