from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Merchant import Merchant
from fintechOnDjango.models.Terminal import Terminal
from fintechOnDjango.models.PaymentContract import PaymentContract
from fintechOnDjango.models.Payout import Payout
from fintechOnDjango.models.SettlementBatch import SettlementBatch
from fintechOnDjango.models.Dispute import Dispute
from fintechOnDjango.models.Invoice import Invoice
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Merchant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MerchantDelegate Declaration
#======================================================================
class MerchantDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, merchantId ):
		try:	
			merchant = Merchant.objects.filter(id=merchantId)
			return merchant.first();
		except Merchant.DoesNotExist:
			raise ProcessingError("Merchant with id " + str(merchantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, merchant):
		for model in serializers.deserialize("json", merchant):
			model.save()
			return model;

	def create(self, merchant):
		merchant.save()
		return merchant;

	def saveFromJson(self, merchant):
		for model in serializers.deserialize("json", merchant):
			model.save()
			return merchant;
	
	def save(self, merchant):
		merchant.save()
		return merchant;
	
	def delete(self, merchantId ):
		errMsg = "Failed to delete Merchant from db using id " + str(merchantId)
		
		try:
			merchant = Merchant.objects.get(id=merchantId)
			merchant.delete()
			return True
		except Merchant.DoesNotExist:
			raise ProcessingError("Merchant with id " + str(merchantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Merchant.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Merchant from db")
		except Exception:
			return None;
		
	def addTerminals( self, merchantId, terminalsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TerminalDelegate import TerminalDelegate

		errMsg = "Failed to add elements " + str(terminalsIds) + " for Terminals on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = terminalsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Terminal		
				terminal = TerminalDelegate().get(id).first();	
				# add the Terminal
				merchant.terminals.add(terminal)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Terminal.DoesNotExist:
			raise ProcessingError(errMsg + " : Terminal does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTerminals( self, merchantId, terminalsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TerminalDelegate import TerminalDelegate

		errMsg = "Failed to remove elements " + str(terminalsIds) + " for Terminals on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = terminalsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Terminal		
				terminal = TerminalDelegate().get(id).first();	
				# add the Terminal
				merchant.terminals.remove(terminal)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Terminal.DoesNotExist:
			raise ProcessingError(errMsg + " : Terminal does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPaymentContracts( self, merchantId, paymentContractsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentContractDelegate import PaymentContractDelegate

		errMsg = "Failed to add elements " + str(paymentContractsIds) + " for PaymentContracts on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = paymentContractsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentContract		
				paymentContract = PaymentContractDelegate().get(id).first();	
				# add the PaymentContract
				merchant.paymentContracts.add(paymentContract)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except PaymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentContract does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePaymentContracts( self, merchantId, paymentContractsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentContractDelegate import PaymentContractDelegate

		errMsg = "Failed to remove elements " + str(paymentContractsIds) + " for PaymentContracts on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = paymentContractsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentContract		
				paymentContract = PaymentContractDelegate().get(id).first();	
				# add the PaymentContract
				merchant.paymentContracts.remove(paymentContract)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except PaymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentContract does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPayouts( self, merchantId, payoutsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PayoutDelegate import PayoutDelegate

		errMsg = "Failed to add elements " + str(payoutsIds) + " for Payouts on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = payoutsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Payout		
				payout = PayoutDelegate().get(id).first();	
				# add the Payout
				merchant.payouts.add(payout)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayouts( self, merchantId, payoutsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PayoutDelegate import PayoutDelegate

		errMsg = "Failed to remove elements " + str(payoutsIds) + " for Payouts on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = payoutsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Payout		
				payout = PayoutDelegate().get(id).first();	
				# add the Payout
				merchant.payouts.remove(payout)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSettlements( self, merchantId, settlementsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.SettlementBatchDelegate import SettlementBatchDelegate

		errMsg = "Failed to add elements " + str(settlementsIds) + " for Settlements on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = settlementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SettlementBatch		
				settlementBatch = SettlementBatchDelegate().get(id).first();	
				# add the SettlementBatch
				merchant.settlements.add(settlementBatch)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSettlements( self, merchantId, settlementsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.SettlementBatchDelegate import SettlementBatchDelegate

		errMsg = "Failed to remove elements " + str(settlementsIds) + " for Settlements on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = settlementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SettlementBatch		
				settlementBatch = SettlementBatchDelegate().get(id).first();	
				# add the SettlementBatch
				merchant.settlements.remove(settlementBatch)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except SettlementBatch.DoesNotExist:
			raise ProcessingError(errMsg + " : SettlementBatch does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDisputes( self, merchantId, disputesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DisputeDelegate import DisputeDelegate

		errMsg = "Failed to add elements " + str(disputesIds) + " for Disputes on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = disputesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dispute		
				dispute = DisputeDelegate().get(id).first();	
				# add the Dispute
				merchant.disputes.add(dispute)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDisputes( self, merchantId, disputesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DisputeDelegate import DisputeDelegate

		errMsg = "Failed to remove elements " + str(disputesIds) + " for Disputes on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = disputesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dispute		
				dispute = DisputeDelegate().get(id).first();	
				# add the Dispute
				merchant.disputes.remove(dispute)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInvoices( self, merchantId, invoicesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

		errMsg = "Failed to add elements " + str(invoicesIds) + " for Invoices on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = invoicesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Invoice		
				invoice = InvoiceDelegate().get(id).first();	
				# add the Invoice
				merchant.invoices.add(invoice)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInvoices( self, merchantId, invoicesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

		errMsg = "Failed to remove elements " + str(invoicesIds) + " for Invoices on Merchant"

		try:
			# get the Merchant
			merchant = self.get( merchantId ).first()
				
			# split on a comma with no spaces
			idList = invoicesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Invoice		
				invoice = InvoiceDelegate().get(id).first();	
				# add the Invoice
				merchant.invoices.remove(invoice)
				
			# save it		
			merchant.save()
			
			# reload and return the appropriate version
			return self.get( merchantId );
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
