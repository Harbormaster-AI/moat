class CompetencyRatingsController < ApplicationController
  def index
    @competencyRatings = CompetencyRating.all
  end
 
  def show
    @competencyRating = CompetencyRating.find(params[:id])
  end
 
  def new
    @competencyRating = CompetencyRating.new
  end
 
  def edit
    @competencyRating = CompetencyRating.find(params[:id])
  end
 
  def create
    @competencyRating = CompetencyRating.new(competencyRating_params)
 
    if @competencyRating.save
      redirect_to competencyRatings_path
    else
      render 'new'
    end
  end
 
  def update
    @competencyRating = CompetencyRating.find(params[:id])
 
    if @competencyRating.update(competencyRating_params)
      redirect_to competencyRatings_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @competencyRating = CompetencyRating.find(params[:id])
    @competencyRating.destroy
    redirect_to competencyRatings_path
  end

 
  private
    def competencyRating_params
      params.require(:competencyRating).permit(:comment, :Rating)
    end
end