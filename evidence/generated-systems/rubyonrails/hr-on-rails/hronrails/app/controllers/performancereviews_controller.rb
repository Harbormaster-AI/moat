class PerformanceReviewsController < ApplicationController
  def index
    @performanceReviews = PerformanceReview.all
  end
 
  def show
    @performanceReview = PerformanceReview.find(params[:id])
  end
 
  def new
    @performanceReview = PerformanceReview.new
  end
 
  def edit
    @performanceReview = PerformanceReview.find(params[:id])
  end
 
  def create
    @performanceReview = PerformanceReview.new(performanceReview_params)
 
    if @performanceReview.save
      redirect_to performanceReviews_path
    else
      render 'new'
    end
  end
 
  def update
    @performanceReview = PerformanceReview.find(params[:id])
 
    if @performanceReview.update(performanceReview_params)
      redirect_to performanceReviews_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @performanceReview = PerformanceReview.find(params[:id])
    @performanceReview.destroy
    redirect_to performanceReviews_path
  end

 
  private
    def performanceReview_params
      params.require(:performanceReview).permit(:reviewNumber, :reviewDate, :reviewerComments, :Rating, :Status)
    end
end