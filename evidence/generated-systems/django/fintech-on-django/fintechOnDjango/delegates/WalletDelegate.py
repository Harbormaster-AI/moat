from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Wallet import Wallet
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Wallet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WalletDelegate Declaration
#======================================================================
class WalletDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, walletId ):
		try:	
			wallet = Wallet.objects.filter(id=walletId)
			return wallet.first();
		except Wallet.DoesNotExist:
			raise ProcessingError("Wallet with id " + str(walletId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, wallet):
		for model in serializers.deserialize("json", wallet):
			model.save()
			return model;

	def create(self, wallet):
		wallet.save()
		return wallet;

	def saveFromJson(self, wallet):
		for model in serializers.deserialize("json", wallet):
			model.save()
			return wallet;
	
	def save(self, wallet):
		wallet.save()
		return wallet;
	
	def delete(self, walletId ):
		errMsg = "Failed to delete Wallet from db using id " + str(walletId)
		
		try:
			wallet = Wallet.objects.get(id=walletId)
			wallet.delete()
			return True
		except Wallet.DoesNotExist:
			raise ProcessingError("Wallet with id " + str(walletId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Wallet.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Wallet from db")
		except Exception:
			return None;
		
	def assignCustomer( self, walletId, customerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Wallet"

		try:
			# get the Wallet from db
			wallet = self.get( walletId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			wallet.customer = customer
			
			#save it
			wallet.save()

			# reload and return the appropriate version					
			return self.get( walletId );
		except Wallet.DoesNotExist:
			raise ProcessingError(errMsg + " : Wallet with id " + str(walletId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, walletId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Wallet"

		try:
			# get the Wallet from db
			wallet = self.get( walletId ).first()	
			
			# assign to None for unassignment
			wallet.customer = None			

			#save it
			wallet.save()

			# reload and return the appropriate version					
			return self.get( walletId );
		except Wallet.DoesNotExist:
			raise ProcessingError(errMsg + " : Wallet with id " + str(walletId) + " does not exist.")
		except Exception:
			return None;
		
	def addTransactions( self, walletId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on Wallet"

		try:
			# get the Wallet
			wallet = self.get( walletId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				wallet.transactions.add(transaction)
				
			# save it		
			wallet.save()
			
			# reload and return the appropriate version
			return self.get( walletId );
		except Wallet.DoesNotExist:
			raise ProcessingError(errMsg + " : Wallet with id " + str(walletId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, walletId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on Wallet"

		try:
			# get the Wallet
			wallet = self.get( walletId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				wallet.transactions.remove(transaction)
				
			# save it		
			wallet.save()
			
			# reload and return the appropriate version
			return self.get( walletId );
		except Wallet.DoesNotExist:
			raise ProcessingError(errMsg + " : Wallet with id " + str(walletId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
