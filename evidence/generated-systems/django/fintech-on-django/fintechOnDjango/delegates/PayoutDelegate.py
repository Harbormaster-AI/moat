from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Payout import Payout
from fintechOnDjango.models.Merchant import Merchant
from fintechOnDjango.models.SettlementBatch import SettlementBatch
from fintechOnDjango.models.Account import Account
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Payout
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayoutDelegate Declaration
#======================================================================
class PayoutDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, payoutId ):
		try:	
			payout = Payout.objects.filter(id=payoutId)
			return payout.first();
		except Payout.DoesNotExist:
			raise ProcessingError("Payout with id " + str(payoutId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, payout):
		for model in serializers.deserialize("json", payout):
			model.save()
			return model;

	def create(self, payout):
		payout.save()
		return payout;

	def saveFromJson(self, payout):
		for model in serializers.deserialize("json", payout):
			model.save()
			return payout;
	
	def save(self, payout):
		payout.save()
		return payout;
	
	def delete(self, payoutId ):
		errMsg = "Failed to delete Payout from db using id " + str(payoutId)
		
		try:
			payout = Payout.objects.get(id=payoutId)
			payout.delete()
			return True
		except Payout.DoesNotExist:
			raise ProcessingError("Payout with id " + str(payoutId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Payout.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Payout from db")
		except Exception:
			return None;
		
	def assignMerchant( self, payoutId, merchantId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Payout"

		try:
			# get the Payout from db
			payout = self.get( payoutId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			payout.merchant = merchant
			
			#save it
			payout.save()

			# reload and return the appropriate version					
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, payoutId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Payout"

		try:
			# get the Payout from db
			payout = self.get( payoutId ).first()	
			
			# assign to None for unassignment
			payout.merchant = None			

			#save it
			payout.save()

			# reload and return the appropriate version					
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSettlementBatch( self, payoutId, settlementBatchId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.SettlementBatchDelegate import SettlementBatchDelegate

		errMsg = "Failed to assign element " + str(settlementBatchId) + " for SettlementBatch on Payout"

		try:
			# get the Payout from db
			payout = self.get( payoutId ).first()	
			
			# get the SettlementBatch from db
			settlementBatch = SettlementBatchDelegate().get(settlementBatchId).first();
			
			# assign the SettlementBatch		
			payout.settlementBatch = settlementBatch
			
			#save it
			payout.save()

			# reload and return the appropriate version					
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch with id " + str(settlementBatchId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSettlementBatch( self, payoutId ):
		errMsg = "Failed to unassign element " + str(settlementBatchId) + " for SettlementBatch on Payout"

		try:
			# get the Payout from db
			payout = self.get( payoutId ).first()	
			
			# assign to None for unassignment
			payout.settlementBatch = None			

			#save it
			payout.save()

			# reload and return the appropriate version					
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDestinationAccount( self, payoutId, destinationAccountId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(destinationAccountId) + " for DestinationAccount on Payout"

		try:
			# get the Payout from db
			payout = self.get( payoutId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(destinationAccountId).first();
			
			# assign the DestinationAccount		
			payout.destinationAccount = account
			
			#save it
			payout.save()

			# reload and return the appropriate version					
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(destinationAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDestinationAccount( self, payoutId ):
		errMsg = "Failed to unassign element " + str(destinationAccountId) + " for DestinationAccount on Payout"

		try:
			# get the Payout from db
			payout = self.get( payoutId ).first()	
			
			# assign to None for unassignment
			payout.account = None			

			#save it
			payout.save()

			# reload and return the appropriate version					
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except Exception:
			return None;
		
