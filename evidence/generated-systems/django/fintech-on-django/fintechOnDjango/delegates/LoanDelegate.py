from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Loan import Loan
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.RepaymentSchedule import RepaymentSchedule
from fintechOnDjango.models.Collateral import Collateral
from fintechOnDjango.models.LoanTransaction import LoanTransaction
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Loan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanDelegate Declaration
#======================================================================
class LoanDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, loanId ):
		try:	
			loan = Loan.objects.filter(id=loanId)
			return loan.first();
		except Loan.DoesNotExist:
			raise ProcessingError("Loan with id " + str(loanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, loan):
		for model in serializers.deserialize("json", loan):
			model.save()
			return model;

	def create(self, loan):
		loan.save()
		return loan;

	def saveFromJson(self, loan):
		for model in serializers.deserialize("json", loan):
			model.save()
			return loan;
	
	def save(self, loan):
		loan.save()
		return loan;
	
	def delete(self, loanId ):
		errMsg = "Failed to delete Loan from db using id " + str(loanId)
		
		try:
			loan = Loan.objects.get(id=loanId)
			loan.delete()
			return True
		except Loan.DoesNotExist:
			raise ProcessingError("Loan with id " + str(loanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Loan.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Loan from db")
		except Exception:
			return None;
		
	def assignCustomer( self, loanId, customerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Loan"

		try:
			# get the Loan from db
			loan = self.get( loanId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			loan.customer = customer
			
			#save it
			loan.save()

			# reload and return the appropriate version					
			return self.get( loanId );
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, loanId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Loan"

		try:
			# get the Loan from db
			loan = self.get( loanId ).first()	
			
			# assign to None for unassignment
			loan.customer = None			

			#save it
			loan.save()

			# reload and return the appropriate version					
			return self.get( loanId );
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except Exception:
			return None;
		
	def addSchedule( self, loanId, scheduleIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.RepaymentScheduleDelegate import RepaymentScheduleDelegate

		errMsg = "Failed to add elements " + str(scheduleIds) + " for Schedule on Loan"

		try:
			# get the Loan
			loan = self.get( loanId ).first()
				
			# split on a comma with no spaces
			idList = scheduleIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RepaymentSchedule		
				repaymentSchedule = RepaymentScheduleDelegate().get(id).first();	
				# add the RepaymentSchedule
				loan.schedule.add(repaymentSchedule)
				
			# save it		
			loan.save()
			
			# reload and return the appropriate version
			return self.get( loanId );
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except RepaymentSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RepaymentSchedule does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSchedule( self, loanId, scheduleIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.RepaymentScheduleDelegate import RepaymentScheduleDelegate

		errMsg = "Failed to remove elements " + str(scheduleIds) + " for Schedule on Loan"

		try:
			# get the Loan
			loan = self.get( loanId ).first()
				
			# split on a comma with no spaces
			idList = scheduleIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RepaymentSchedule		
				repaymentSchedule = RepaymentScheduleDelegate().get(id).first();	
				# add the RepaymentSchedule
				loan.schedule.remove(repaymentSchedule)
				
			# save it		
			loan.save()
			
			# reload and return the appropriate version
			return self.get( loanId );
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except RepaymentSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RepaymentSchedule does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCollateral( self, loanId, collateralIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CollateralDelegate import CollateralDelegate

		errMsg = "Failed to add elements " + str(collateralIds) + " for Collateral on Loan"

		try:
			# get the Loan
			loan = self.get( loanId ).first()
				
			# split on a comma with no spaces
			idList = collateralIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Collateral		
				collateral = CollateralDelegate().get(id).first();	
				# add the Collateral
				loan.collateral.add(collateral)
				
			# save it		
			loan.save()
			
			# reload and return the appropriate version
			return self.get( loanId );
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except Collateral.DoesNotExist:
			raise ProcessingError(errMsg + " : Collateral does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCollateral( self, loanId, collateralIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CollateralDelegate import CollateralDelegate

		errMsg = "Failed to remove elements " + str(collateralIds) + " for Collateral on Loan"

		try:
			# get the Loan
			loan = self.get( loanId ).first()
				
			# split on a comma with no spaces
			idList = collateralIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Collateral		
				collateral = CollateralDelegate().get(id).first();	
				# add the Collateral
				loan.collateral.remove(collateral)
				
			# save it		
			loan.save()
			
			# reload and return the appropriate version
			return self.get( loanId );
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except Collateral.DoesNotExist:
			raise ProcessingError(errMsg + " : Collateral does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTransactions( self, loanId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanTransactionDelegate import LoanTransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on Loan"

		try:
			# get the Loan
			loan = self.get( loanId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LoanTransaction		
				loanTransaction = LoanTransactionDelegate().get(id).first();	
				# add the LoanTransaction
				loan.transactions.add(loanTransaction)
				
			# save it		
			loan.save()
			
			# reload and return the appropriate version
			return self.get( loanId );
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except LoanTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanTransaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, loanId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanTransactionDelegate import LoanTransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on Loan"

		try:
			# get the Loan
			loan = self.get( loanId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LoanTransaction		
				loanTransaction = LoanTransactionDelegate().get(id).first();	
				# add the LoanTransaction
				loan.transactions.remove(loanTransaction)
				
			# save it		
			loan.save()
			
			# reload and return the appropriate version
			return self.get( loanId );
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except LoanTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanTransaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
