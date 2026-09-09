from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.AppliedFee import AppliedFee
from fintechOnDjango.models.PaymentOrder import PaymentOrder
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AppliedFee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AppliedFeeDelegate Declaration
#======================================================================
class AppliedFeeDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, appliedFeeId ):
		try:	
			appliedFee = AppliedFee.objects.filter(id=appliedFeeId)
			return appliedFee.first();
		except AppliedFee.DoesNotExist:
			raise ProcessingError("AppliedFee with id " + str(appliedFeeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, appliedFee):
		for model in serializers.deserialize("json", appliedFee):
			model.save()
			return model;

	def create(self, appliedFee):
		appliedFee.save()
		return appliedFee;

	def saveFromJson(self, appliedFee):
		for model in serializers.deserialize("json", appliedFee):
			model.save()
			return appliedFee;
	
	def save(self, appliedFee):
		appliedFee.save()
		return appliedFee;
	
	def delete(self, appliedFeeId ):
		errMsg = "Failed to delete AppliedFee from db using id " + str(appliedFeeId)
		
		try:
			appliedFee = AppliedFee.objects.get(id=appliedFeeId)
			appliedFee.delete()
			return True
		except AppliedFee.DoesNotExist:
			raise ProcessingError("AppliedFee with id " + str(appliedFeeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AppliedFee.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AppliedFee from db")
		except Exception:
			return None;
		
	def assignPaymentOrder( self, appliedFeeId, paymentOrderId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentOrderDelegate import PaymentOrderDelegate

		errMsg = "Failed to assign element " + str(paymentOrderId) + " for PaymentOrder on AppliedFee"

		try:
			# get the AppliedFee from db
			appliedFee = self.get( appliedFeeId ).first()	
			
			# get the PaymentOrder from db
			paymentOrder = PaymentOrderDelegate().get(paymentOrderId).first();
			
			# assign the PaymentOrder		
			appliedFee.paymentOrder = paymentOrder
			
			#save it
			appliedFee.save()

			# reload and return the appropriate version					
			return self.get( appliedFeeId );
		except AppliedFee.DoesNotExist:
			raise ProcessingError(errMsg + " : AppliedFee with id " + str(appliedFeeId) + " does not exist.")
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder with id " + str(paymentOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPaymentOrder( self, appliedFeeId ):
		errMsg = "Failed to unassign element " + str(paymentOrderId) + " for PaymentOrder on AppliedFee"

		try:
			# get the AppliedFee from db
			appliedFee = self.get( appliedFeeId ).first()	
			
			# assign to None for unassignment
			appliedFee.paymentOrder = None			

			#save it
			appliedFee.save()

			# reload and return the appropriate version					
			return self.get( appliedFeeId );
		except AppliedFee.DoesNotExist:
			raise ProcessingError(errMsg + " : AppliedFee with id " + str(appliedFeeId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTransaction( self, appliedFeeId, transactionId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to assign element " + str(transactionId) + " for Transaction on AppliedFee"

		try:
			# get the AppliedFee from db
			appliedFee = self.get( appliedFeeId ).first()	
			
			# get the Transaction from db
			transaction = TransactionDelegate().get(transactionId).first();
			
			# assign the Transaction		
			appliedFee.transaction = transaction
			
			#save it
			appliedFee.save()

			# reload and return the appropriate version					
			return self.get( appliedFeeId );
		except AppliedFee.DoesNotExist:
			raise ProcessingError(errMsg + " : AppliedFee with id " + str(appliedFeeId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTransaction( self, appliedFeeId ):
		errMsg = "Failed to unassign element " + str(transactionId) + " for Transaction on AppliedFee"

		try:
			# get the AppliedFee from db
			appliedFee = self.get( appliedFeeId ).first()	
			
			# assign to None for unassignment
			appliedFee.transaction = None			

			#save it
			appliedFee.save()

			# reload and return the appropriate version					
			return self.get( appliedFeeId );
		except AppliedFee.DoesNotExist:
			raise ProcessingError(errMsg + " : AppliedFee with id " + str(appliedFeeId) + " does not exist.")
		except Exception:
			return None;
		
