from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Candidate import Candidate
from hrOnDjango.models.JobApplication import JobApplication
from hrOnDjango.models.Interview import Interview
from hrOnDjango.models.Offer import Offer
from hrOnDjango.models.Document import Document
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Candidate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CandidateDelegate Declaration
#======================================================================
class CandidateDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, candidateId ):
		try:	
			candidate = Candidate.objects.filter(id=candidateId)
			return candidate.first();
		except Candidate.DoesNotExist:
			raise ProcessingError("Candidate with id " + str(candidateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, candidate):
		for model in serializers.deserialize("json", candidate):
			model.save()
			return model;

	def create(self, candidate):
		candidate.save()
		return candidate;

	def saveFromJson(self, candidate):
		for model in serializers.deserialize("json", candidate):
			model.save()
			return candidate;
	
	def save(self, candidate):
		candidate.save()
		return candidate;
	
	def delete(self, candidateId ):
		errMsg = "Failed to delete Candidate from db using id " + str(candidateId)
		
		try:
			candidate = Candidate.objects.get(id=candidateId)
			candidate.delete()
			return True
		except Candidate.DoesNotExist:
			raise ProcessingError("Candidate with id " + str(candidateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Candidate.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Candidate from db")
		except Exception:
			return None;
		
	def addApplications( self, candidateId, applicationsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobApplicationDelegate import JobApplicationDelegate

		errMsg = "Failed to add elements " + str(applicationsIds) + " for Applications on Candidate"

		try:
			# get the Candidate
			candidate = self.get( candidateId ).first()
				
			# split on a comma with no spaces
			idList = applicationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the JobApplication		
				jobApplication = JobApplicationDelegate().get(id).first();	
				# add the JobApplication
				candidate.applications.add(jobApplication)
				
			# save it		
			candidate.save()
			
			# reload and return the appropriate version
			return self.get( candidateId );
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except JobApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : JobApplication does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeApplications( self, candidateId, applicationsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobApplicationDelegate import JobApplicationDelegate

		errMsg = "Failed to remove elements " + str(applicationsIds) + " for Applications on Candidate"

		try:
			# get the Candidate
			candidate = self.get( candidateId ).first()
				
			# split on a comma with no spaces
			idList = applicationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the JobApplication		
				jobApplication = JobApplicationDelegate().get(id).first();	
				# add the JobApplication
				candidate.applications.remove(jobApplication)
				
			# save it		
			candidate.save()
			
			# reload and return the appropriate version
			return self.get( candidateId );
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except JobApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : JobApplication does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInterviews( self, candidateId, interviewsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.InterviewDelegate import InterviewDelegate

		errMsg = "Failed to add elements " + str(interviewsIds) + " for Interviews on Candidate"

		try:
			# get the Candidate
			candidate = self.get( candidateId ).first()
				
			# split on a comma with no spaces
			idList = interviewsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Interview		
				interview = InterviewDelegate().get(id).first();	
				# add the Interview
				candidate.interviews.add(interview)
				
			# save it		
			candidate.save()
			
			# reload and return the appropriate version
			return self.get( candidateId );
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInterviews( self, candidateId, interviewsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.InterviewDelegate import InterviewDelegate

		errMsg = "Failed to remove elements " + str(interviewsIds) + " for Interviews on Candidate"

		try:
			# get the Candidate
			candidate = self.get( candidateId ).first()
				
			# split on a comma with no spaces
			idList = interviewsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Interview		
				interview = InterviewDelegate().get(id).first();	
				# add the Interview
				candidate.interviews.remove(interview)
				
			# save it		
			candidate.save()
			
			# reload and return the appropriate version
			return self.get( candidateId );
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Interview.DoesNotExist:
			raise ProcessingError(errMsg + " : Interview does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOffers( self, candidateId, offersIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OfferDelegate import OfferDelegate

		errMsg = "Failed to add elements " + str(offersIds) + " for Offers on Candidate"

		try:
			# get the Candidate
			candidate = self.get( candidateId ).first()
				
			# split on a comma with no spaces
			idList = offersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Offer		
				offer = OfferDelegate().get(id).first();	
				# add the Offer
				candidate.offers.add(offer)
				
			# save it		
			candidate.save()
			
			# reload and return the appropriate version
			return self.get( candidateId );
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOffers( self, candidateId, offersIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OfferDelegate import OfferDelegate

		errMsg = "Failed to remove elements " + str(offersIds) + " for Offers on Candidate"

		try:
			# get the Candidate
			candidate = self.get( candidateId ).first()
				
			# split on a comma with no spaces
			idList = offersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Offer		
				offer = OfferDelegate().get(id).first();	
				# add the Offer
				candidate.offers.remove(offer)
				
			# save it		
			candidate.save()
			
			# reload and return the appropriate version
			return self.get( candidateId );
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Offer.DoesNotExist:
			raise ProcessingError(errMsg + " : Offer does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDocuments( self, candidateId, documentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DocumentDelegate import DocumentDelegate

		errMsg = "Failed to add elements " + str(documentsIds) + " for Documents on Candidate"

		try:
			# get the Candidate
			candidate = self.get( candidateId ).first()
				
			# split on a comma with no spaces
			idList = documentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Document		
				document = DocumentDelegate().get(id).first();	
				# add the Document
				candidate.documents.add(document)
				
			# save it		
			candidate.save()
			
			# reload and return the appropriate version
			return self.get( candidateId );
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDocuments( self, candidateId, documentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DocumentDelegate import DocumentDelegate

		errMsg = "Failed to remove elements " + str(documentsIds) + " for Documents on Candidate"

		try:
			# get the Candidate
			candidate = self.get( candidateId ).first()
				
			# split on a comma with no spaces
			idList = documentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Document		
				document = DocumentDelegate().get(id).first();	
				# add the Document
				candidate.documents.remove(document)
				
			# save it		
			candidate.save()
			
			# reload and return the appropriate version
			return self.get( candidateId );
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
