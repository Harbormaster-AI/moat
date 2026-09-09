from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.PaymentContract import PaymentContract
from fintechOnDjango.models.Merchant import Merchant
from fintechOnDjango.models.PaymentProcessor import PaymentProcessor
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PaymentContract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentContractDelegate Declaration
#======================================================================
class PaymentContractDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, paymentContractId ):
		try:	
			paymentContract = PaymentContract.objects.filter(id=paymentContractId)
			return paymentContract.first();
		except PaymentContract.DoesNotExist:
			raise ProcessingError("PaymentContract with id " + str(paymentContractId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, paymentContract):
		for model in serializers.deserialize("json", paymentContract):
			model.save()
			return model;

	def create(self, paymentContract):
		paymentContract.save()
		return paymentContract;

	def saveFromJson(self, paymentContract):
		for model in serializers.deserialize("json", paymentContract):
			model.save()
			return paymentContract;
	
	def save(self, paymentContract):
		paymentContract.save()
		return paymentContract;
	
	def delete(self, paymentContractId ):
		errMsg = "Failed to delete PaymentContract from db using id " + str(paymentContractId)
		
		try:
			paymentContract = PaymentContract.objects.get(id=paymentContractId)
			paymentContract.delete()
			return True
		except PaymentContract.DoesNotExist:
			raise ProcessingError("PaymentContract with id " + str(paymentContractId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PaymentContract.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PaymentContract from db")
		except Exception:
			return None;
		
	def assignMerchant( self, paymentContractId, merchantId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on PaymentContract"

		try:
			# get the PaymentContract from db
			paymentContract = self.get( paymentContractId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			paymentContract.merchant = merchant
			
			#save it
			paymentContract.save()

			# reload and return the appropriate version					
			return self.get( paymentContractId );
		except PaymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentContract with id " + str(paymentContractId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, paymentContractId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on PaymentContract"

		try:
			# get the PaymentContract from db
			paymentContract = self.get( paymentContractId ).first()	
			
			# assign to None for unassignment
			paymentContract.merchant = None			

			#save it
			paymentContract.save()

			# reload and return the appropriate version					
			return self.get( paymentContractId );
		except PaymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentContract with id " + str(paymentContractId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAcquirer( self, paymentContractId, acquirerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentProcessorDelegate import PaymentProcessorDelegate

		errMsg = "Failed to assign element " + str(acquirerId) + " for Acquirer on PaymentContract"

		try:
			# get the PaymentContract from db
			paymentContract = self.get( paymentContractId ).first()	
			
			# get the PaymentProcessor from db
			paymentProcessor = PaymentProcessorDelegate().get(acquirerId).first();
			
			# assign the Acquirer		
			paymentContract.acquirer = paymentProcessor
			
			#save it
			paymentContract.save()

			# reload and return the appropriate version					
			return self.get( paymentContractId );
		except PaymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentContract with id " + str(paymentContractId) + " does not exist.")
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor with id " + str(acquirerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAcquirer( self, paymentContractId ):
		errMsg = "Failed to unassign element " + str(acquirerId) + " for Acquirer on PaymentContract"

		try:
			# get the PaymentContract from db
			paymentContract = self.get( paymentContractId ).first()	
			
			# assign to None for unassignment
			paymentContract.paymentProcessor = None			

			#save it
			paymentContract.save()

			# reload and return the appropriate version					
			return self.get( paymentContractId );
		except PaymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentContract with id " + str(paymentContractId) + " does not exist.")
		except Exception:
			return None;
		
