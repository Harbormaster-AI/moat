from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.RepaymentSchedule import RepaymentSchedule
from fintechOnDjango.models.Loan import Loan
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model RepaymentSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RepaymentScheduleDelegate Declaration
#======================================================================
class RepaymentScheduleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, repaymentScheduleId ):
		try:	
			repaymentSchedule = RepaymentSchedule.objects.filter(id=repaymentScheduleId)
			return repaymentSchedule.first();
		except RepaymentSchedule.DoesNotExist:
			raise ProcessingError("RepaymentSchedule with id " + str(repaymentScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, repaymentSchedule):
		for model in serializers.deserialize("json", repaymentSchedule):
			model.save()
			return model;

	def create(self, repaymentSchedule):
		repaymentSchedule.save()
		return repaymentSchedule;

	def saveFromJson(self, repaymentSchedule):
		for model in serializers.deserialize("json", repaymentSchedule):
			model.save()
			return repaymentSchedule;
	
	def save(self, repaymentSchedule):
		repaymentSchedule.save()
		return repaymentSchedule;
	
	def delete(self, repaymentScheduleId ):
		errMsg = "Failed to delete RepaymentSchedule from db using id " + str(repaymentScheduleId)
		
		try:
			repaymentSchedule = RepaymentSchedule.objects.get(id=repaymentScheduleId)
			repaymentSchedule.delete()
			return True
		except RepaymentSchedule.DoesNotExist:
			raise ProcessingError("RepaymentSchedule with id " + str(repaymentScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = RepaymentSchedule.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all RepaymentSchedule from db")
		except Exception:
			return None;
		
	def assignLoan( self, repaymentScheduleId, loanId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanDelegate import LoanDelegate

		errMsg = "Failed to assign element " + str(loanId) + " for Loan on RepaymentSchedule"

		try:
			# get the RepaymentSchedule from db
			repaymentSchedule = self.get( repaymentScheduleId ).first()	
			
			# get the Loan from db
			loan = LoanDelegate().get(loanId).first();
			
			# assign the Loan		
			repaymentSchedule.loan = loan
			
			#save it
			repaymentSchedule.save()

			# reload and return the appropriate version					
			return self.get( repaymentScheduleId );
		except RepaymentSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RepaymentSchedule with id " + str(repaymentScheduleId) + " does not exist.")
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan with id " + str(loanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLoan( self, repaymentScheduleId ):
		errMsg = "Failed to unassign element " + str(loanId) + " for Loan on RepaymentSchedule"

		try:
			# get the RepaymentSchedule from db
			repaymentSchedule = self.get( repaymentScheduleId ).first()	
			
			# assign to None for unassignment
			repaymentSchedule.loan = None			

			#save it
			repaymentSchedule.save()

			# reload and return the appropriate version					
			return self.get( repaymentScheduleId );
		except RepaymentSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RepaymentSchedule with id " + str(repaymentScheduleId) + " does not exist.")
		except Exception:
			return None;
		
	def addPayments( self, repaymentScheduleId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to add elements " + str(paymentsIds) + " for Payments on RepaymentSchedule"

		try:
			# get the RepaymentSchedule
			repaymentSchedule = self.get( repaymentScheduleId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				repaymentSchedule.payments.add(transaction)
				
			# save it		
			repaymentSchedule.save()
			
			# reload and return the appropriate version
			return self.get( repaymentScheduleId );
		except RepaymentSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RepaymentSchedule with id " + str(repaymentScheduleId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayments( self, repaymentScheduleId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to remove elements " + str(paymentsIds) + " for Payments on RepaymentSchedule"

		try:
			# get the RepaymentSchedule
			repaymentSchedule = self.get( repaymentScheduleId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Transaction		
				transaction = TransactionDelegate().get(id).first();	
				# add the Transaction
				repaymentSchedule.payments.remove(transaction)
				
			# save it		
			repaymentSchedule.save()
			
			# reload and return the appropriate version
			return self.get( repaymentScheduleId );
		except RepaymentSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RepaymentSchedule with id " + str(repaymentScheduleId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
