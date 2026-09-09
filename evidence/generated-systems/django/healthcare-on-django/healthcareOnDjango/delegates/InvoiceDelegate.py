from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Invoice import Invoice
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.Claim import Claim
from healthcareOnDjango.models.Payment import Payment
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Invoice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvoiceDelegate Declaration
#======================================================================
class InvoiceDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, invoiceId ):
		try:	
			invoice = Invoice.objects.filter(id=invoiceId)
			return invoice.first();
		except Invoice.DoesNotExist:
			raise ProcessingError("Invoice with id " + str(invoiceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, invoice):
		for model in serializers.deserialize("json", invoice):
			model.save()
			return model;

	def create(self, invoice):
		invoice.save()
		return invoice;

	def saveFromJson(self, invoice):
		for model in serializers.deserialize("json", invoice):
			model.save()
			return invoice;
	
	def save(self, invoice):
		invoice.save()
		return invoice;
	
	def delete(self, invoiceId ):
		errMsg = "Failed to delete Invoice from db using id " + str(invoiceId)
		
		try:
			invoice = Invoice.objects.get(id=invoiceId)
			invoice.delete()
			return True
		except Invoice.DoesNotExist:
			raise ProcessingError("Invoice with id " + str(invoiceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Invoice.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Invoice from db")
		except Exception:
			return None;
		
	def assignPatient( self, invoiceId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			invoice.patient = patient
			
			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, invoiceId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# assign to None for unassignment
			invoice.patient = None			

			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Exception:
			return None;
		
	def assignClaim( self, invoiceId, claimId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to assign element " + str(claimId) + " for Claim on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# get the Claim from db
			claim = ClaimDelegate().get(claimId).first();
			
			# assign the Claim		
			invoice.claim = claim
			
			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClaim( self, invoiceId ):
		errMsg = "Failed to unassign element " + str(claimId) + " for Claim on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# assign to None for unassignment
			invoice.claim = None			

			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Exception:
			return None;
		
	def addPayments( self, invoiceId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to add elements " + str(paymentsIds) + " for Payments on Invoice"

		try:
			# get the Invoice
			invoice = self.get( invoiceId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				invoice.payments.add(payment)
				
			# save it		
			invoice.save()
			
			# reload and return the appropriate version
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayments( self, invoiceId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to remove elements " + str(paymentsIds) + " for Payments on Invoice"

		try:
			# get the Invoice
			invoice = self.get( invoiceId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				invoice.payments.remove(payment)
				
			# save it		
			invoice.save()
			
			# reload and return the appropriate version
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
