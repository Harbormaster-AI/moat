from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Dispute import Dispute
from fintechOnDjango.models.Transaction import Transaction
from fintechOnDjango.models.PaymentCard import PaymentCard
from fintechOnDjango.models.Merchant import Merchant
from fintechOnDjango.models.Chargeback import Chargeback
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Dispute
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DisputeDelegate Declaration
#======================================================================
class DisputeDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, disputeId ):
		try:	
			dispute = Dispute.objects.filter(id=disputeId)
			return dispute.first();
		except Dispute.DoesNotExist:
			raise ProcessingError("Dispute with id " + str(disputeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dispute):
		for model in serializers.deserialize("json", dispute):
			model.save()
			return model;

	def create(self, dispute):
		dispute.save()
		return dispute;

	def saveFromJson(self, dispute):
		for model in serializers.deserialize("json", dispute):
			model.save()
			return dispute;
	
	def save(self, dispute):
		dispute.save()
		return dispute;
	
	def delete(self, disputeId ):
		errMsg = "Failed to delete Dispute from db using id " + str(disputeId)
		
		try:
			dispute = Dispute.objects.get(id=disputeId)
			dispute.delete()
			return True
		except Dispute.DoesNotExist:
			raise ProcessingError("Dispute with id " + str(disputeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Dispute.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Dispute from db")
		except Exception:
			return None;
		
	def assignTransaction( self, disputeId, transactionId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TransactionDelegate import TransactionDelegate

		errMsg = "Failed to assign element " + str(transactionId) + " for Transaction on Dispute"

		try:
			# get the Dispute from db
			dispute = self.get( disputeId ).first()	
			
			# get the Transaction from db
			transaction = TransactionDelegate().get(transactionId).first();
			
			# assign the Transaction		
			dispute.transaction = transaction
			
			#save it
			dispute.save()

			# reload and return the appropriate version					
			return self.get( disputeId );
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute with id " + str(disputeId) + " does not exist.")
		except Transaction.DoesNotExist:
			raise ProcessingError(errMsg + " : Transaction with id " + str(transactionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTransaction( self, disputeId ):
		errMsg = "Failed to unassign element " + str(transactionId) + " for Transaction on Dispute"

		try:
			# get the Dispute from db
			dispute = self.get( disputeId ).first()	
			
			# assign to None for unassignment
			dispute.transaction = None			

			#save it
			dispute.save()

			# reload and return the appropriate version					
			return self.get( disputeId );
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute with id " + str(disputeId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCard( self, disputeId, cardId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentCardDelegate import PaymentCardDelegate

		errMsg = "Failed to assign element " + str(cardId) + " for Card on Dispute"

		try:
			# get the Dispute from db
			dispute = self.get( disputeId ).first()	
			
			# get the PaymentCard from db
			paymentCard = PaymentCardDelegate().get(cardId).first();
			
			# assign the Card		
			dispute.card = paymentCard
			
			#save it
			dispute.save()

			# reload and return the appropriate version					
			return self.get( disputeId );
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute with id " + str(disputeId) + " does not exist.")
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(cardId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCard( self, disputeId ):
		errMsg = "Failed to unassign element " + str(cardId) + " for Card on Dispute"

		try:
			# get the Dispute from db
			dispute = self.get( disputeId ).first()	
			
			# assign to None for unassignment
			dispute.paymentCard = None			

			#save it
			dispute.save()

			# reload and return the appropriate version					
			return self.get( disputeId );
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute with id " + str(disputeId) + " does not exist.")
		except Exception:
			return None;
		
	def assignMerchant( self, disputeId, merchantId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Dispute"

		try:
			# get the Dispute from db
			dispute = self.get( disputeId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			dispute.merchant = merchant
			
			#save it
			dispute.save()

			# reload and return the appropriate version					
			return self.get( disputeId );
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute with id " + str(disputeId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, disputeId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Dispute"

		try:
			# get the Dispute from db
			dispute = self.get( disputeId ).first()	
			
			# assign to None for unassignment
			dispute.merchant = None			

			#save it
			dispute.save()

			# reload and return the appropriate version					
			return self.get( disputeId );
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute with id " + str(disputeId) + " does not exist.")
		except Exception:
			return None;
		
	def addChargebacks( self, disputeId, chargebacksIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ChargebackDelegate import ChargebackDelegate

		errMsg = "Failed to add elements " + str(chargebacksIds) + " for Chargebacks on Dispute"

		try:
			# get the Dispute
			dispute = self.get( disputeId ).first()
				
			# split on a comma with no spaces
			idList = chargebacksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Chargeback		
				chargeback = ChargebackDelegate().get(id).first();	
				# add the Chargeback
				dispute.chargebacks.add(chargeback)
				
			# save it		
			dispute.save()
			
			# reload and return the appropriate version
			return self.get( disputeId );
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute with id " + str(disputeId) + " does not exist.")
		except Chargeback.DoesNotExist:
			raise ProcessingError(errMsg + " : Chargeback does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChargebacks( self, disputeId, chargebacksIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ChargebackDelegate import ChargebackDelegate

		errMsg = "Failed to remove elements " + str(chargebacksIds) + " for Chargebacks on Dispute"

		try:
			# get the Dispute
			dispute = self.get( disputeId ).first()
				
			# split on a comma with no spaces
			idList = chargebacksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Chargeback		
				chargeback = ChargebackDelegate().get(id).first();	
				# add the Chargeback
				dispute.chargebacks.remove(chargeback)
				
			# save it		
			dispute.save()
			
			# reload and return the appropriate version
			return self.get( disputeId );
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute with id " + str(disputeId) + " does not exist.")
		except Chargeback.DoesNotExist:
			raise ProcessingError(errMsg + " : Chargeback does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
