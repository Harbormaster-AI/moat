class DispositionReviewsController < ApplicationController
  def index
    @dispositionReviews = DispositionReview.all
  end
 
  def show
    @dispositionReview = DispositionReview.find(params[:id])
  end
 
  def new
    @dispositionReview = DispositionReview.new
  end
 
  def edit
    @dispositionReview = DispositionReview.find(params[:id])
  end
 
  def create
    @dispositionReview = DispositionReview.new(dispositionReview_params)
 
    if @dispositionReview.save
      redirect_to dispositionReviews_path
    else
      render 'new'
    end
  end
 
  def update
    @dispositionReview = DispositionReview.find(params[:id])
 
    if @dispositionReview.update(dispositionReview_params)
      redirect_to dispositionReviews_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @dispositionReview = DispositionReview.find(params[:id])
    @dispositionReview.destroy
    redirect_to dispositionReviews_path
  end

 
  private
    def dispositionReview_params
      params.require(:dispositionReview).permit(:reviewDate, :reviewer, :notes, :Outcome)
    end
end