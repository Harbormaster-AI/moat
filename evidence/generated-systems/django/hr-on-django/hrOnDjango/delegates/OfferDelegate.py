from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Offer import Offer
from hrOnDjango.models.JobRequisition import JobRequisition
from hrOnDjango.models.Candidate import Candidate
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.EmploymentContract import EmploymentContract
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Offer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OfferDelegate Declaration
#======================================================================
class OfferDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, offerId ):
		try:	
			offer = Offer.objects.filter(id=offerId)
			return offer.first();
		except Offer.DoesNotExist:
			raise ProcessingError("Offer with id " + str(offerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, offer):
		for model in serializers.deserialize("json", offer):
			model.save()
			return model;

	def create(self, offer):
		offer.save()
		return offer;

	def saveFromJson(self, offer):
		for model in serializers.deserialize("json", offer):
			model.save()
			return offer;
	
	def save(self, offer):
		offer.save()
		return offer;
	
	def delete(self, offerId ):
		errMsg = "Failed to delete Offer from db using id " + str(offerId)
		
		try:
			offer = Offer.objects.get(id=offerId)
			offer.delete()
			return True
		except Offer.DoesNotExist:
			raise ProcessingError("Offer with id " + str(offerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Offer.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Offer from db")
		except Exception:
			return None;
		
	def assignRequisition( self, offerId, requisitionId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobRequisitionDelegate import JobRequisitionDelegate

		errMsg = "Failed to assign element " + str(requisitionId) + " for Requisition on Offer"

		try:
			# get the Offer from db
			offer = self.get( offerId ).first()	
			
			# get the JobRequisition from db
			jobRequisition = JobRequisitionDelegate().get(requisitionId).first();
			
			# assign the Requisition		
			offer.requisition = jobRequisition
			
			#save it
			offer.save()

			# reload and return the appropriate version					
			return self.get( offerId );
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer with id " + str(offerId) + " does not exist.")
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(requisitionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRequisition( self, offerId ):
		errMsg = "Failed to unassign element " + str(requisitionId) + " for Requisition on Offer"

		try:
			# get the Offer from db
			offer = self.get( offerId ).first()	
			
			# assign to None for unassignment
			offer.jobRequisition = None			

			#save it
			offer.save()

			# reload and return the appropriate version					
			return self.get( offerId );
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer with id " + str(offerId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCandidate( self, offerId, candidateId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CandidateDelegate import CandidateDelegate

		errMsg = "Failed to assign element " + str(candidateId) + " for Candidate on Offer"

		try:
			# get the Offer from db
			offer = self.get( offerId ).first()	
			
			# get the Candidate from db
			candidate = CandidateDelegate().get(candidateId).first();
			
			# assign the Candidate		
			offer.candidate = candidate
			
			#save it
			offer.save()

			# reload and return the appropriate version					
			return self.get( offerId );
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer with id " + str(offerId) + " does not exist.")
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCandidate( self, offerId ):
		errMsg = "Failed to unassign element " + str(candidateId) + " for Candidate on Offer"

		try:
			# get the Offer from db
			offer = self.get( offerId ).first()	
			
			# assign to None for unassignment
			offer.candidate = None			

			#save it
			offer.save()

			# reload and return the appropriate version					
			return self.get( offerId );
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer with id " + str(offerId) + " does not exist.")
		except Exception:
			return None;
		
	def assignApprovedBy( self, offerId, approvedById ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(approvedById) + " for ApprovedBy on Offer"

		try:
			# get the Offer from db
			offer = self.get( offerId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(approvedById).first();
			
			# assign the ApprovedBy		
			offer.approvedBy = employee
			
			#save it
			offer.save()

			# reload and return the appropriate version					
			return self.get( offerId );
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer with id " + str(offerId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(approvedById) + " does not exist.")
		except Exception:
			return None;
				
	def unassignApprovedBy( self, offerId ):
		errMsg = "Failed to unassign element " + str(approvedById) + " for ApprovedBy on Offer"

		try:
			# get the Offer from db
			offer = self.get( offerId ).first()	
			
			# assign to None for unassignment
			offer.employee = None			

			#save it
			offer.save()

			# reload and return the appropriate version					
			return self.get( offerId );
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer with id " + str(offerId) + " does not exist.")
		except Exception:
			return None;
		
	def assignContract( self, offerId, contractId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentContractDelegate import EmploymentContractDelegate

		errMsg = "Failed to assign element " + str(contractId) + " for Contract on Offer"

		try:
			# get the Offer from db
			offer = self.get( offerId ).first()	
			
			# get the EmploymentContract from db
			employmentContract = EmploymentContractDelegate().get(contractId).first();
			
			# assign the Contract		
			offer.contract = employmentContract
			
			#save it
			offer.save()

			# reload and return the appropriate version					
			return self.get( offerId );
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer with id " + str(offerId) + " does not exist.")
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract with id " + str(contractId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignContract( self, offerId ):
		errMsg = "Failed to unassign element " + str(contractId) + " for Contract on Offer"

		try:
			# get the Offer from db
			offer = self.get( offerId ).first()	
			
			# assign to None for unassignment
			offer.employmentContract = None			

			#save it
			offer.save()

			# reload and return the appropriate version					
			return self.get( offerId );
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer with id " + str(offerId) + " does not exist.")
		except Exception:
			return None;
		
