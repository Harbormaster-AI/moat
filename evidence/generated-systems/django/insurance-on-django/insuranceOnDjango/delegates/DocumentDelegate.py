from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Document import Document
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.Application import Application
from insuranceOnDjango.models.Customer import Customer
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Document
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DocumentDelegate Declaration
#======================================================================
class DocumentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, documentId ):
		try:	
			document = Document.objects.filter(id=documentId)
			return document.first();
		except Document.DoesNotExist:
			raise ProcessingError("Document with id " + str(documentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, document):
		for model in serializers.deserialize("json", document):
			model.save()
			return model;

	def create(self, document):
		document.save()
		return document;

	def saveFromJson(self, document):
		for model in serializers.deserialize("json", document):
			model.save()
			return document;
	
	def save(self, document):
		document.save()
		return document;
	
	def delete(self, documentId ):
		errMsg = "Failed to delete Document from db using id " + str(documentId)
		
		try:
			document = Document.objects.get(id=documentId)
			document.delete()
			return True
		except Document.DoesNotExist:
			raise ProcessingError("Document with id " + str(documentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Document.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Document from db")
		except Exception:
			return None;
		
	def assignPolicy( self, documentId, policyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			document.policy = policy
			
			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, documentId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# assign to None for unassignment
			document.policy = None			

			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignClaim( self, documentId, claimId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to assign element " + str(claimId) + " for Claim on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# get the Claim from db
			claim = ClaimDelegate().get(claimId).first();
			
			# assign the Claim		
			document.claim = claim
			
			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClaim( self, documentId ):
		errMsg = "Failed to unassign element " + str(claimId) + " for Claim on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# assign to None for unassignment
			document.claim = None			

			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignApplication( self, documentId, applicationId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ApplicationDelegate import ApplicationDelegate

		errMsg = "Failed to assign element " + str(applicationId) + " for Application on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# get the Application from db
			application = ApplicationDelegate().get(applicationId).first();
			
			# assign the Application		
			document.application = application
			
			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignApplication( self, documentId ):
		errMsg = "Failed to unassign element " + str(applicationId) + " for Application on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# assign to None for unassignment
			document.application = None			

			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCustomer( self, documentId, customerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			document.customer = customer
			
			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, documentId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# assign to None for unassignment
			document.customer = None			

			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Exception:
			return None;
		
