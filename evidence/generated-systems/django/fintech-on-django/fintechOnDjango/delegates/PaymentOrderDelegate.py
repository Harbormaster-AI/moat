from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.PaymentOrder import PaymentOrder
from fintechOnDjango.models.Account import Account
from fintechOnDjango.models.Beneficiary import Beneficiary
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.models.FXDeal import FXDeal
from fintechOnDjango.models.AppliedFee import AppliedFee
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PaymentOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentOrderDelegate Declaration
#======================================================================
class PaymentOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, paymentOrderId ):
		try:	
			paymentOrder = PaymentOrder.objects.filter(id=paymentOrderId)
			return paymentOrder.first();
		except PaymentOrder.DoesNotExist:
			raise ProcessingError("PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, paymentOrder):
		for model in serializers.deserialize("json", paymentOrder):
			model.save()
			return model;

	def create(self, paymentOrder):
		paymentOrder.save()
		return paymentOrder;

	def saveFromJson(self, paymentOrder):
		for model in serializers.deserialize("json", paymentOrder):
			model.save()
			return paymentOrder;
	
	def save(self, paymentOrder):
		paymentOrder.save()
		return paymentOrder;
	
	def delete(self, paymentOrderId ):
		errMsg = "Failed to delete PaymentOrder from db using id " + str(paymentOrderId)
		
		try:
			paymentOrder = PaymentOrder.objects.get(id=paymentOrderId)
			paymentOrder.delete()
			return True
		except PaymentOrder.DoesNotExist:
			raise ProcessingError("PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PaymentOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PaymentOrder from db")
		except Exception:
			return None;
		
	def assignSourceAccount( self, paymentOrderId, sourceAccountId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(sourceAccountId) + " for SourceAccount on PaymentOrder"

		try:
			# get the PaymentOrder from db
			paymentOrder = self.get( paymentOrderId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(sourceAccountId).first();
			
			# assign the SourceAccount		
			paymentOrder.sourceAccount = account
			
			#save it
			paymentOrder.save()

			# reload and return the appropriate version					
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(sourceAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSourceAccount( self, paymentOrderId ):
		errMsg = "Failed to unassign element " + str(sourceAccountId) + " for SourceAccount on PaymentOrder"

		try:
			# get the PaymentOrder from db
			paymentOrder = self.get( paymentOrderId ).first()	
			
			# assign to None for unassignment
			paymentOrder.account = None			

			#save it
			paymentOrder.save()

			# reload and return the appropriate version					
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDestinationAccount( self, paymentOrderId, destinationAccountId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(destinationAccountId) + " for DestinationAccount on PaymentOrder"

		try:
			# get the PaymentOrder from db
			paymentOrder = self.get( paymentOrderId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(destinationAccountId).first();
			
			# assign the DestinationAccount		
			paymentOrder.destinationAccount = account
			
			#save it
			paymentOrder.save()

			# reload and return the appropriate version					
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(destinationAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDestinationAccount( self, paymentOrderId ):
		errMsg = "Failed to unassign element " + str(destinationAccountId) + " for DestinationAccount on PaymentOrder"

		try:
			# get the PaymentOrder from db
			paymentOrder = self.get( paymentOrderId ).first()	
			
			# assign to None for unassignment
			paymentOrder.account = None			

			#save it
			paymentOrder.save()

			# reload and return the appropriate version					
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignBeneficiary( self, paymentOrderId, beneficiaryId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.BeneficiaryDelegate import BeneficiaryDelegate

		errMsg = "Failed to assign element " + str(beneficiaryId) + " for Beneficiary on PaymentOrder"

		try:
			# get the PaymentOrder from db
			paymentOrder = self.get( paymentOrderId ).first()	
			
			# get the Beneficiary from db
			beneficiary = BeneficiaryDelegate().get(beneficiaryId).first();
			
			# assign the Beneficiary		
			paymentOrder.beneficiary = beneficiary
			
			#save it
			paymentOrder.save()

			# reload and return the appropriate version					
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary with id " + str(beneficiaryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBeneficiary( self, paymentOrderId ):
		errMsg = "Failed to unassign element " + str(beneficiaryId) + " for Beneficiary on PaymentOrder"

		try:
			# get the PaymentOrder from db
			paymentOrder = self.get( paymentOrderId ).first()	
			
			# assign to None for unassignment
			paymentOrder.beneficiary = None			

			#save it
			paymentOrder.save()

			# reload and return the appropriate version					
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignFxDeal( self, paymentOrderId, fxDealId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FXDealDelegate import FXDealDelegate

		errMsg = "Failed to assign element " + str(fxDealId) + " for FxDeal on PaymentOrder"

		try:
			# get the PaymentOrder from db
			paymentOrder = self.get( paymentOrderId ).first()	
			
			# get the FXDeal from db
			fXDeal = FXDealDelegate().get(fxDealId).first();
			
			# assign the FxDeal		
			paymentOrder.fxDeal = fXDeal
			
			#save it
			paymentOrder.save()

			# reload and return the appropriate version					
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except FXDeal.DoesNotExist:
			raise ProcessingError(errMsg + " : FXDeal with id " + str(fxDealId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFxDeal( self, paymentOrderId ):
		errMsg = "Failed to unassign element " + str(fxDealId) + " for FxDeal on PaymentOrder"

		try:
			# get the PaymentOrder from db
			paymentOrder = self.get( paymentOrderId ).first()	
			
			# assign to None for unassignment
			paymentOrder.fXDeal = None			

			#save it
			paymentOrder.save()

			# reload and return the appropriate version					
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addTransactions( self, paymentOrderId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on PaymentOrder"

		try:
			# get the PaymentOrder
			paymentOrder = self.get( paymentOrderId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				paymentOrder.transactions.add(transaction)
				
			# save it		
			paymentOrder.save()
			
			# reload and return the appropriate version
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, paymentOrderId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on PaymentOrder"

		try:
			# get the PaymentOrder
			paymentOrder = self.get( paymentOrderId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				paymentOrder.transactions.remove(transaction)
				
			# save it		
			paymentOrder.save()
			
			# reload and return the appropriate version
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFees( self, paymentOrderId, feesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AppliedFeeDelegate import AppliedFeeDelegate

		errMsg = "Failed to add elements " + str(feesIds) + " for Fees on PaymentOrder"

		try:
			# get the PaymentOrder
			paymentOrder = self.get( paymentOrderId ).first()
				
			# split on a comma with no spaces
			idList = feesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AppliedFee		
				appliedFee = AppliedFeeDelegate().get(id).first();	
				# add the AppliedFee
				paymentOrder.fees.add(appliedFee)
				
			# save it		
			paymentOrder.save()
			
			# reload and return the appropriate version
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except AppliedFee.DoesNotExist:
			raise ProcessingError(errMsg + " : AppliedFee does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFees( self, paymentOrderId, feesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AppliedFeeDelegate import AppliedFeeDelegate

		errMsg = "Failed to remove elements " + str(feesIds) + " for Fees on PaymentOrder"

		try:
			# get the PaymentOrder
			paymentOrder = self.get( paymentOrderId ).first()
				
			# split on a comma with no spaces
			idList = feesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AppliedFee		
				appliedFee = AppliedFeeDelegate().get(id).first();	
				# add the AppliedFee
				paymentOrder.fees.remove(appliedFee)
				
			# save it		
			paymentOrder.save()
			
			# reload and return the appropriate version
			return self.get( paymentOrderId );
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except AppliedFee.DoesNotExist:
			raise ProcessingError(errMsg + " : AppliedFee does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
