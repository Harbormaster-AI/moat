class RecommendationScenariosController < ApplicationController
  def index
    @recommendationScenarios = RecommendationScenario.all
  end
 
  def show
    @recommendationScenario = RecommendationScenario.find(params[:id])
  end
 
  def new
    @recommendationScenario = RecommendationScenario.new
  end
 
  def edit
    @recommendationScenario = RecommendationScenario.find(params[:id])
  end
 
  def create
    @recommendationScenario = RecommendationScenario.new(recommendationScenario_params)
 
    if @recommendationScenario.save
      redirect_to recommendationScenarios_path
    else
      render 'new'
    end
  end
 
  def update
    @recommendationScenario = RecommendationScenario.find(params[:id])
 
    if @recommendationScenario.update(recommendationScenario_params)
      redirect_to recommendationScenarios_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @recommendationScenario = RecommendationScenario.find(params[:id])
    @recommendationScenario.destroy
    redirect_to recommendationScenarios_path
  end

 
  private
    def recommendationScenario_params
      params.require(:recommendationScenario).permit(:name, :objective, :RecommendationType)
    end
end