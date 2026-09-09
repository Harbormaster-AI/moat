from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.LoanTransaction import LoanTransaction
from fintechOnDjango.models.Loan import Loan
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LoanTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanTransactionDelegate Declaration
#======================================================================
class LoanTransactionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, loanTransactionId ):
		try:	
			loanTransaction = LoanTransaction.objects.filter(id=loanTransactionId)
			return loanTransaction.first();
		except LoanTransaction.DoesNotExist:
			raise ProcessingError("LoanTransaction with id " + str(loanTransactionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, loanTransaction):
		for model in serializers.deserialize("json", loanTransaction):
			model.save()
			return model;

	def create(self, loanTransaction):
		loanTransaction.save()
		return loanTransaction;

	def saveFromJson(self, loanTransaction):
		for model in serializers.deserialize("json", loanTransaction):
			model.save()
			return loanTransaction;
	
	def save(self, loanTransaction):
		loanTransaction.save()
		return loanTransaction;
	
	def delete(self, loanTransactionId ):
		errMsg = "Failed to delete LoanTransaction from db using id " + str(loanTransactionId)
		
		try:
			loanTransaction = LoanTransaction.objects.get(id=loanTransactionId)
			loanTransaction.delete()
			return True
		except LoanTransaction.DoesNotExist:
			raise ProcessingError("LoanTransaction with id " + str(loanTransactionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LoanTransaction.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LoanTransaction from db")
		except Exception:
			return None;
		
	def assignLoan( self, loanTransactionId, loanId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanDelegate import LoanDelegate

		errMsg = "Failed to assign element " + str(loanId) + " for Loan on LoanTransaction"

		try:
			# get the LoanTransaction from db
			loanTransaction = self.get( loanTransactionId ).first()	
			
			# get the Loan from db
			loan = LoanDelegate().get(loanId).first();
			
			# assign the Loan		
			loanTransaction.loan = loan
			
			#save it
			loanTransaction.save()

			# reload and return the appropriate version					
			return self.get( loanTransactionId );
		except LoanTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanTransaction with id " + str(loanTransactionId) + " does not exist.")
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLoan( self, loanTransactionId ):
		errMsg = "Failed to unassign element " + str(loanId) + " for Loan on LoanTransaction"

		try:
			# get the LoanTransaction from db
			loanTransaction = self.get( loanTransactionId ).first()	
			
			# assign to None for unassignment
			loanTransaction.loan = None			

			#save it
			loanTransaction.save()

			# reload and return the appropriate version					
			return self.get( loanTransactionId );
		except LoanTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanTransaction with id " + str(loanTransactionId) + " does not exist.")
		except Exception:
			return None;
		
