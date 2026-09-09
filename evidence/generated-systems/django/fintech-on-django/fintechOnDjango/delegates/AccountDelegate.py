from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Account import Account
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.FinancialInstitution import FinancialInstitution
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.models.PaymentCard import PaymentCard
from fintechOnDjango.models.AccountStatement import AccountStatement
from fintechOnDjango.models.DirectDebitMandate import DirectDebitMandate
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Account
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccountDelegate Declaration
#======================================================================
class AccountDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, accountId ):
		try:	
			account = Account.objects.filter(id=accountId)
			return account.first();
		except Account.DoesNotExist:
			raise ProcessingError("Account with id " + str(accountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, account):
		for model in serializers.deserialize("json", account):
			model.save()
			return model;

	def create(self, account):
		account.save()
		return account;

	def saveFromJson(self, account):
		for model in serializers.deserialize("json", account):
			model.save()
			return account;
	
	def save(self, account):
		account.save()
		return account;
	
	def delete(self, accountId ):
		errMsg = "Failed to delete Account from db using id " + str(accountId)
		
		try:
			account = Account.objects.get(id=accountId)
			account.delete()
			return True
		except Account.DoesNotExist:
			raise ProcessingError("Account with id " + str(accountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Account.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Account from db")
		except Exception:
			return None;
		
	def assignCustomer( self, accountId, customerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			account.customer = customer
			
			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, accountId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# assign to None for unassignment
			account.customer = None			

			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInstitution( self, accountId, institutionId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FinancialInstitutionDelegate import FinancialInstitutionDelegate

		errMsg = "Failed to assign element " + str(institutionId) + " for Institution on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# get the FinancialInstitution from db
			financialInstitution = FinancialInstitutionDelegate().get(institutionId).first();
			
			# assign the Institution		
			account.institution = financialInstitution
			
			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(institutionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInstitution( self, accountId ):
		errMsg = "Failed to unassign element " + str(institutionId) + " for Institution on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# assign to None for unassignment
			account.financialInstitution = None			

			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
		
	def addTransactions( self, accountId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				account.transactions.add(transaction)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, accountId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				account.transactions.remove(transaction)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCards( self, accountId, cardsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentCardDelegate import PaymentCardDelegate

		errMsg = "Failed to add elements " + str(cardsIds) + " for Cards on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = cardsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentCard		
				paymentCard = PaymentCardDelegate().get(id).first();	
				# add the PaymentCard
				account.cards.add(paymentCard)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCards( self, accountId, cardsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentCardDelegate import PaymentCardDelegate

		errMsg = "Failed to remove elements " + str(cardsIds) + " for Cards on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = cardsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentCard		
				paymentCard = PaymentCardDelegate().get(id).first();	
				# add the PaymentCard
				account.cards.remove(paymentCard)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addStatements( self, accountId, statementsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountStatementDelegate import AccountStatementDelegate

		errMsg = "Failed to add elements " + str(statementsIds) + " for Statements on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = statementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AccountStatement		
				accountStatement = AccountStatementDelegate().get(id).first();	
				# add the AccountStatement
				account.statements.add(accountStatement)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except AccountStatement.DoesNotExist:
			raise ProcessingError(errMsg + " : AccountStatement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeStatements( self, accountId, statementsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountStatementDelegate import AccountStatementDelegate

		errMsg = "Failed to remove elements " + str(statementsIds) + " for Statements on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = statementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AccountStatement		
				accountStatement = AccountStatementDelegate().get(id).first();	
				# add the AccountStatement
				account.statements.remove(accountStatement)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except AccountStatement.DoesNotExist:
			raise ProcessingError(errMsg + " : AccountStatement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMandates( self, accountId, mandatesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DirectDebitMandateDelegate import DirectDebitMandateDelegate

		errMsg = "Failed to add elements " + str(mandatesIds) + " for Mandates on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = mandatesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DirectDebitMandate		
				directDebitMandate = DirectDebitMandateDelegate().get(id).first();	
				# add the DirectDebitMandate
				account.mandates.add(directDebitMandate)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError(errMsg + " : DirectDebitMandate does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMandates( self, accountId, mandatesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DirectDebitMandateDelegate import DirectDebitMandateDelegate

		errMsg = "Failed to remove elements " + str(mandatesIds) + " for Mandates on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = mandatesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DirectDebitMandate		
				directDebitMandate = DirectDebitMandateDelegate().get(id).first();	
				# add the DirectDebitMandate
				account.mandates.remove(directDebitMandate)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError(errMsg + " : DirectDebitMandate does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
