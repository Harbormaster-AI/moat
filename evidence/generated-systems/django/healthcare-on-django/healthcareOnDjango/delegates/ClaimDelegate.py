from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Claim import Claim
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.Coverage import Coverage
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.Invoice import Invoice
from healthcareOnDjango.models.InsurancePayer import InsurancePayer
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Claim
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimDelegate Declaration
#======================================================================
class ClaimDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, claimId ):
		try:	
			claim = Claim.objects.filter(id=claimId)
			return claim.first();
		except Claim.DoesNotExist:
			raise ProcessingError("Claim with id " + str(claimId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, claim):
		for model in serializers.deserialize("json", claim):
			model.save()
			return model;

	def create(self, claim):
		claim.save()
		return claim;

	def saveFromJson(self, claim):
		for model in serializers.deserialize("json", claim):
			model.save()
			return claim;
	
	def save(self, claim):
		claim.save()
		return claim;
	
	def delete(self, claimId ):
		errMsg = "Failed to delete Claim from db using id " + str(claimId)
		
		try:
			claim = Claim.objects.get(id=claimId)
			claim.delete()
			return True
		except Claim.DoesNotExist:
			raise ProcessingError("Claim with id " + str(claimId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Claim.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Claim from db")
		except Exception:
			return None;
		
	def assignPatient( self, claimId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			claim.patient = patient
			
			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, claimId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# assign to None for unassignment
			claim.patient = None			

			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCoverage( self, claimId, coverageId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CoverageDelegate import CoverageDelegate

		errMsg = "Failed to assign element " + str(coverageId) + " for Coverage on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# get the Coverage from db
			coverage = CoverageDelegate().get(coverageId).first();
			
			# assign the Coverage		
			claim.coverage = coverage
			
			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCoverage( self, claimId ):
		errMsg = "Failed to unassign element " + str(coverageId) + " for Coverage on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# assign to None for unassignment
			claim.coverage = None			

			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEncounter( self, claimId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			claim.encounter = encounter
			
			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, claimId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# assign to None for unassignment
			claim.encounter = None			

			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPayer( self, claimId, payerId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InsurancePayerDelegate import InsurancePayerDelegate

		errMsg = "Failed to assign element " + str(payerId) + " for Payer on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# get the InsurancePayer from db
			insurancePayer = InsurancePayerDelegate().get(payerId).first();
			
			# assign the Payer		
			claim.payer = insurancePayer
			
			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except InsurancePayer.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePayer with id " + str(payerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPayer( self, claimId ):
		errMsg = "Failed to unassign element " + str(payerId) + " for Payer on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# assign to None for unassignment
			claim.insurancePayer = None			

			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
		
	def addInvoices( self, claimId, invoicesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

		errMsg = "Failed to add elements " + str(invoicesIds) + " for Invoices on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = invoicesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Invoice		
				invoice = InvoiceDelegate().get(id).first();	
				# add the Invoice
				claim.invoices.add(invoice)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInvoices( self, claimId, invoicesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

		errMsg = "Failed to remove elements " + str(invoicesIds) + " for Invoices on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = invoicesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Invoice		
				invoice = InvoiceDelegate().get(id).first();	
				# add the Invoice
				claim.invoices.remove(invoice)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
