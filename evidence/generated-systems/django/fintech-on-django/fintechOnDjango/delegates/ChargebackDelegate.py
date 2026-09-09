from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Chargeback import Chargeback
from fintechOnDjango.models.Dispute import Dispute
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Chargeback
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChargebackDelegate Declaration
#======================================================================
class ChargebackDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, chargebackId ):
		try:	
			chargeback = Chargeback.objects.filter(id=chargebackId)
			return chargeback.first();
		except Chargeback.DoesNotExist:
			raise ProcessingError("Chargeback with id " + str(chargebackId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, chargeback):
		for model in serializers.deserialize("json", chargeback):
			model.save()
			return model;

	def create(self, chargeback):
		chargeback.save()
		return chargeback;

	def saveFromJson(self, chargeback):
		for model in serializers.deserialize("json", chargeback):
			model.save()
			return chargeback;
	
	def save(self, chargeback):
		chargeback.save()
		return chargeback;
	
	def delete(self, chargebackId ):
		errMsg = "Failed to delete Chargeback from db using id " + str(chargebackId)
		
		try:
			chargeback = Chargeback.objects.get(id=chargebackId)
			chargeback.delete()
			return True
		except Chargeback.DoesNotExist:
			raise ProcessingError("Chargeback with id " + str(chargebackId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Chargeback.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Chargeback from db")
		except Exception:
			return None;
		
	def assignDispute( self, chargebackId, disputeId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DisputeDelegate import DisputeDelegate

		errMsg = "Failed to assign element " + str(disputeId) + " for Dispute on Chargeback"

		try:
			# get the Chargeback from db
			chargeback = self.get( chargebackId ).first()	
			
			# get the Dispute from db
			dispute = DisputeDelegate().get(disputeId).first();
			
			# assign the Dispute		
			chargeback.dispute = dispute
			
			#save it
			chargeback.save()

			# reload and return the appropriate version					
			return self.get( chargebackId );
		except Chargeback.DoesNotExist:
			raise ProcessingError(errMsg + " : Chargeback with id " + str(chargebackId) + " does not exist.")
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute with id " + str(disputeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDispute( self, chargebackId ):
		errMsg = "Failed to unassign element " + str(disputeId) + " for Dispute on Chargeback"

		try:
			# get the Chargeback from db
			chargeback = self.get( chargebackId ).first()	
			
			# assign to None for unassignment
			chargeback.dispute = None			

			#save it
			chargeback.save()

			# reload and return the appropriate version					
			return self.get( chargebackId );
		except Chargeback.DoesNotExist:
			raise ProcessingError(errMsg + " : Chargeback with id " + str(chargebackId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTransaction( self, chargebackId, transactionId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to assign element " + str(transactionId) + " for Transaction on Chargeback"

		try:
			# get the Chargeback from db
			chargeback = self.get( chargebackId ).first()	
			
			# get the Transaction from db
			transaction = TransactionDelegate().get(transactionId).first();
			
			# assign the Transaction		
			chargeback.transaction = transaction
			
			#save it
			chargeback.save()

			# reload and return the appropriate version					
			return self.get( chargebackId );
		except Chargeback.DoesNotExist:
			raise ProcessingError(errMsg + " : Chargeback with id " + str(chargebackId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTransaction( self, chargebackId ):
		errMsg = "Failed to unassign element " + str(transactionId) + " for Transaction on Chargeback"

		try:
			# get the Chargeback from db
			chargeback = self.get( chargebackId ).first()	
			
			# assign to None for unassignment
			chargeback.transaction = None			

			#save it
			chargeback.save()

			# reload and return the appropriate version					
			return self.get( chargebackId );
		except Chargeback.DoesNotExist:
			raise ProcessingError(errMsg + " : Chargeback with id " + str(chargebackId) + " does not exist.")
		except Exception:
			return None;
		
