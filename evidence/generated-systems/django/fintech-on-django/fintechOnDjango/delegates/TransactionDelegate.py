from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.models.Account import Account
from fintechOnDjango.models.Wallet import Wallet
from fintechOnDjango.models.PaymentOrder import PaymentOrder
from fintechOnDjango.models.Merchant import Merchant
from fintechOnDjango.models.PaymentCard import PaymentCard
from fintechOnDjango.models.ComplianceAlert import ComplianceAlert
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Transaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransactionDelegate Declaration
#======================================================================
class TransactionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, transactionId ):
		try:	
			transaction = Transaction.objects.filter(id=transactionId)
			return transaction.first();
		except Transaction.DoesNotExist:
			raise ProcessingError("Transaction with id " + str(transactionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, transaction):
		for model in serializers.deserialize("json", transaction):
			model.save()
			return model;

	def create(self, transaction):
		transaction.save()
		return transaction;

	def saveFromJson(self, transaction):
		for model in serializers.deserialize("json", transaction):
			model.save()
			return transaction;
	
	def save(self, transaction):
		transaction.save()
		return transaction;
	
	def delete(self, transactionId ):
		errMsg = "Failed to delete Transaction from db using id " + str(transactionId)
		
		try:
			transaction = Transaction.objects.get(id=transactionId)
			transaction.delete()
			return True
		except Transaction.DoesNotExist:
			raise ProcessingError("Transaction with id " + str(transactionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Transaction.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Transaction from db")
		except Exception:
			return None;
		
	def assignAccount( self, transactionId, accountId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			transaction.account = account
			
			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, transactionId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# assign to None for unassignment
			transaction.account = None			

			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWallet( self, transactionId, walletId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.WalletDelegate import WalletDelegate

		errMsg = "Failed to assign element " + str(walletId) + " for Wallet on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# get the Wallet from db
			wallet = WalletDelegate().get(walletId).first();
			
			# assign the Wallet		
			transaction.wallet = wallet
			
			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Wallet.DoesNotExist:
			raise ProcessingError(errMsg + " : Wallet with id " + str(walletId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWallet( self, transactionId ):
		errMsg = "Failed to unassign element " + str(walletId) + " for Wallet on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# assign to None for unassignment
			transaction.wallet = None			

			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPaymentOrder( self, transactionId, paymentOrderId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentOrderDelegate import PaymentOrderDelegate

		errMsg = "Failed to assign element " + str(paymentOrderId) + " for PaymentOrder on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# get the PaymentOrder from db
			paymentOrder = PaymentOrderDelegate().get(paymentOrderId).first();
			
			# assign the PaymentOrder		
			transaction.paymentOrder = paymentOrder
			
			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPaymentOrder( self, transactionId ):
		errMsg = "Failed to unassign element " + str(paymentOrderId) + " for PaymentOrder on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# assign to None for unassignment
			transaction.paymentOrder = None			

			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignMerchant( self, transactionId, merchantId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			transaction.merchant = merchant
			
			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, transactionId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# assign to None for unassignment
			transaction.merchant = None			

			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCard( self, transactionId, cardId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentCardDelegate import PaymentCardDelegate

		errMsg = "Failed to assign element " + str(cardId) + " for Card on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# get the PaymentCard from db
			paymentCard = PaymentCardDelegate().get(cardId).first();
			
			# assign the Card		
			transaction.card = paymentCard
			
			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(cardId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCard( self, transactionId ):
		errMsg = "Failed to unassign element " + str(cardId) + " for Card on Transaction"

		try:
			# get the Transaction from db
			transaction = self.get( transactionId ).first()	
			
			# assign to None for unassignment
			transaction.paymentCard = None			

			#save it
			transaction.save()

			# reload and return the appropriate version					
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Exception:
			return None;
		
	def addRelatedTransactions( self, transactionId, relatedTransactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to add elements " + str(relatedTransactionsIds) + " for RelatedTransactions on Transaction"

		try:
			# get the Transaction
			transaction = self.get( transactionId ).first()
				
			# split on a comma with no spaces
			idList = relatedTransactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				transaction.relatedTransactions.add(transaction)
				
			# save it		
			transaction.save()
			
			# reload and return the appropriate version
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRelatedTransactions( self, transactionId, relatedTransactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to remove elements " + str(relatedTransactionsIds) + " for RelatedTransactions on Transaction"

		try:
			# get the Transaction
			transaction = self.get( transactionId ).first()
				
			# split on a comma with no spaces
			idList = relatedTransactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				transaction.relatedTransactions.remove(transaction)
				
			# save it		
			transaction.save()
			
			# reload and return the appropriate version
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAlerts( self, transactionId, alertsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ComplianceAlertDelegate import ComplianceAlertDelegate

		errMsg = "Failed to add elements " + str(alertsIds) + " for Alerts on Transaction"

		try:
			# get the Transaction
			transaction = self.get( transactionId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ComplianceAlert		
				complianceAlert = ComplianceAlertDelegate().get(id).first();	
				# add the ComplianceAlert
				transaction.alerts.add(complianceAlert)
				
			# save it		
			transaction.save()
			
			# reload and return the appropriate version
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceAlert does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAlerts( self, transactionId, alertsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ComplianceAlertDelegate import ComplianceAlertDelegate

		errMsg = "Failed to remove elements " + str(alertsIds) + " for Alerts on Transaction"

		try:
			# get the Transaction
			transaction = self.get( transactionId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ComplianceAlert		
				complianceAlert = ComplianceAlertDelegate().get(id).first();	
				# add the ComplianceAlert
				transaction.alerts.remove(complianceAlert)
				
			# save it		
			transaction.save()
			
			# reload and return the appropriate version
			return self.get( transactionId );
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceAlert does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
