from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.CompetencyRating import CompetencyRating
from hrOnDjango.models.PerformanceReview import PerformanceReview
from hrOnDjango.models.Competency import Competency
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CompetencyRating
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompetencyRatingDelegate Declaration
#======================================================================
class CompetencyRatingDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, competencyRatingId ):
		try:	
			competencyRating = CompetencyRating.objects.filter(id=competencyRatingId)
			return competencyRating.first();
		except CompetencyRating.DoesNotExist:
			raise ProcessingError("CompetencyRating with id " + str(competencyRatingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, competencyRating):
		for model in serializers.deserialize("json", competencyRating):
			model.save()
			return model;

	def create(self, competencyRating):
		competencyRating.save()
		return competencyRating;

	def saveFromJson(self, competencyRating):
		for model in serializers.deserialize("json", competencyRating):
			model.save()
			return competencyRating;
	
	def save(self, competencyRating):
		competencyRating.save()
		return competencyRating;
	
	def delete(self, competencyRatingId ):
		errMsg = "Failed to delete CompetencyRating from db using id " + str(competencyRatingId)
		
		try:
			competencyRating = CompetencyRating.objects.get(id=competencyRatingId)
			competencyRating.delete()
			return True
		except CompetencyRating.DoesNotExist:
			raise ProcessingError("CompetencyRating with id " + str(competencyRatingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CompetencyRating.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CompetencyRating from db")
		except Exception:
			return None;
		
	def assignReview( self, competencyRatingId, reviewId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PerformanceReviewDelegate import PerformanceReviewDelegate

		errMsg = "Failed to assign element " + str(reviewId) + " for Review on CompetencyRating"

		try:
			# get the CompetencyRating from db
			competencyRating = self.get( competencyRatingId ).first()	
			
			# get the PerformanceReview from db
			performanceReview = PerformanceReviewDelegate().get(reviewId).first();
			
			# assign the Review		
			competencyRating.review = performanceReview
			
			#save it
			competencyRating.save()

			# reload and return the appropriate version					
			return self.get( competencyRatingId );
		except CompetencyRating.DoesNotExist:
			raise ProcessingError(errMsg + " : CompetencyRating with id " + str(competencyRatingId) + " does not exist.")
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview with id " + str(reviewId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignReview( self, competencyRatingId ):
		errMsg = "Failed to unassign element " + str(reviewId) + " for Review on CompetencyRating"

		try:
			# get the CompetencyRating from db
			competencyRating = self.get( competencyRatingId ).first()	
			
			# assign to None for unassignment
			competencyRating.performanceReview = None			

			#save it
			competencyRating.save()

			# reload and return the appropriate version					
			return self.get( competencyRatingId );
		except CompetencyRating.DoesNotExist:
			raise ProcessingError(errMsg + " : CompetencyRating with id " + str(competencyRatingId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCompetency( self, competencyRatingId, competencyId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompetencyDelegate import CompetencyDelegate

		errMsg = "Failed to assign element " + str(competencyId) + " for Competency on CompetencyRating"

		try:
			# get the CompetencyRating from db
			competencyRating = self.get( competencyRatingId ).first()	
			
			# get the Competency from db
			competency = CompetencyDelegate().get(competencyId).first();
			
			# assign the Competency		
			competencyRating.competency = competency
			
			#save it
			competencyRating.save()

			# reload and return the appropriate version					
			return self.get( competencyRatingId );
		except CompetencyRating.DoesNotExist:
			raise ProcessingError(errMsg + " : CompetencyRating with id " + str(competencyRatingId) + " does not exist.")
		except Competency.DoesNotExist:
			raise ProcessingError(errMsg + " : Competency with id " + str(competencyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCompetency( self, competencyRatingId ):
		errMsg = "Failed to unassign element " + str(competencyId) + " for Competency on CompetencyRating"

		try:
			# get the CompetencyRating from db
			competencyRating = self.get( competencyRatingId ).first()	
			
			# assign to None for unassignment
			competencyRating.competency = None			

			#save it
			competencyRating.save()

			# reload and return the appropriate version					
			return self.get( competencyRatingId );
		except CompetencyRating.DoesNotExist:
			raise ProcessingError(errMsg + " : CompetencyRating with id " + str(competencyRatingId) + " does not exist.")
		except Exception:
			return None;
		
