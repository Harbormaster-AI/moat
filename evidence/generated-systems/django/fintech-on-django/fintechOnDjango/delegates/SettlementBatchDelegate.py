from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.SettlementBatch import SettlementBatch
from fintechOnDjango.models.PaymentProcessor import PaymentProcessor
from fintechOnDjango.models.Merchant import Merchant
from fintechOnDjango.models.Payout import Payout
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SettlementBatch
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SettlementBatchDelegate Declaration
#======================================================================
class SettlementBatchDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, settlementBatchId ):
		try:	
			settlementBatch = SettlementBatch.objects.filter(id=settlementBatchId)
			return settlementBatch.first();
		except SettlementBatch.DoesNotExist:
			raise ProcessingError("SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, settlementBatch):
		for model in serializers.deserialize("json", settlementBatch):
			model.save()
			return model;

	def create(self, settlementBatch):
		settlementBatch.save()
		return settlementBatch;

	def saveFromJson(self, settlementBatch):
		for model in serializers.deserialize("json", settlementBatch):
			model.save()
			return settlementBatch;
	
	def save(self, settlementBatch):
		settlementBatch.save()
		return settlementBatch;
	
	def delete(self, settlementBatchId ):
		errMsg = "Failed to delete SettlementBatch from db using id " + str(settlementBatchId)
		
		try:
			settlementBatch = SettlementBatch.objects.get(id=settlementBatchId)
			settlementBatch.delete()
			return True
		except SettlementBatch.DoesNotExist:
			raise ProcessingError("SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SettlementBatch.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SettlementBatch from db")
		except Exception:
			return None;
		
	def assignProcessor( self, settlementBatchId, processorId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentProcessorDelegate import PaymentProcessorDelegate

		errMsg = "Failed to assign element " + str(processorId) + " for Processor on SettlementBatch"

		try:
			# get the SettlementBatch from db
			settlementBatch = self.get( settlementBatchId ).first()	
			
			# get the PaymentProcessor from db
			paymentProcessor = PaymentProcessorDelegate().get(processorId).first();
			
			# assign the Processor		
			settlementBatch.processor = paymentProcessor
			
			#save it
			settlementBatch.save()

			# reload and return the appropriate version					
			return self.get( settlementBatchId );
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except PaymentProcessor.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProcessor with id " + str(processorId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProcessor( self, settlementBatchId ):
		errMsg = "Failed to unassign element " + str(processorId) + " for Processor on SettlementBatch"

		try:
			# get the SettlementBatch from db
			settlementBatch = self.get( settlementBatchId ).first()	
			
			# assign to None for unassignment
			settlementBatch.paymentProcessor = None			

			#save it
			settlementBatch.save()

			# reload and return the appropriate version					
			return self.get( settlementBatchId );
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except Exception:
			return None;
		
	def assignMerchant( self, settlementBatchId, merchantId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on SettlementBatch"

		try:
			# get the SettlementBatch from db
			settlementBatch = self.get( settlementBatchId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			settlementBatch.merchant = merchant
			
			#save it
			settlementBatch.save()

			# reload and return the appropriate version					
			return self.get( settlementBatchId );
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, settlementBatchId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on SettlementBatch"

		try:
			# get the SettlementBatch from db
			settlementBatch = self.get( settlementBatchId ).first()	
			
			# assign to None for unassignment
			settlementBatch.merchant = None			

			#save it
			settlementBatch.save()

			# reload and return the appropriate version					
			return self.get( settlementBatchId );
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except Exception:
			return None;
		
	def addPayouts( self, settlementBatchId, payoutsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PayoutDelegate import PayoutDelegate

		errMsg = "Failed to add elements " + str(payoutsIds) + " for Payouts on SettlementBatch"

		try:
			# get the SettlementBatch
			settlementBatch = self.get( settlementBatchId ).first()
				
			# split on a comma with no spaces
			idList = payoutsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Payout		
				payout = PayoutDelegate().get(id).first();	
				# add the Payout
				settlementBatch.payouts.add(payout)
				
			# save it		
			settlementBatch.save()
			
			# reload and return the appropriate version
			return self.get( settlementBatchId );
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayouts( self, settlementBatchId, payoutsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PayoutDelegate import PayoutDelegate

		errMsg = "Failed to remove elements " + str(payoutsIds) + " for Payouts on SettlementBatch"

		try:
			# get the SettlementBatch
			settlementBatch = self.get( settlementBatchId ).first()
				
			# split on a comma with no spaces
			idList = payoutsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Payout		
				payout = PayoutDelegate().get(id).first();	
				# add the Payout
				settlementBatch.payouts.remove(payout)
				
			# save it		
			settlementBatch.save()
			
			# reload and return the appropriate version
			return self.get( settlementBatchId );
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTransactions( self, settlementBatchId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on SettlementBatch"

		try:
			# get the SettlementBatch
			settlementBatch = self.get( settlementBatchId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				settlementBatch.transactions.add(transaction)
				
			# save it		
			settlementBatch.save()
			
			# reload and return the appropriate version
			return self.get( settlementBatchId );
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, settlementBatchId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on SettlementBatch"

		try:
			# get the SettlementBatch
			settlementBatch = self.get( settlementBatchId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				settlementBatch.transactions.remove(transaction)
				
			# save it		
			settlementBatch.save()
			
			# reload and return the appropriate version
			return self.get( settlementBatchId );
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
