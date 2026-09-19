class QualityRulesController < ApplicationController
  def index
    @qualityRules = QualityRule.all
  end
 
  def show
    @qualityRule = QualityRule.find(params[:id])
  end
 
  def new
    @qualityRule = QualityRule.new
  end
 
  def edit
    @qualityRule = QualityRule.find(params[:id])
  end
 
  def create
    @qualityRule = QualityRule.new(qualityRule_params)
 
    if @qualityRule.save
      redirect_to qualityRules_path
    else
      render 'new'
    end
  end
 
  def update
    @qualityRule = QualityRule.find(params[:id])
 
    if @qualityRule.update(qualityRule_params)
      redirect_to qualityRules_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @qualityRule = QualityRule.find(params[:id])
    @qualityRule.destroy
    redirect_to qualityRules_path
  end

 
  private
    def qualityRule_params
      params.require(:qualityRule).permit(:name, :threshold, :targetField, :Dimension, :Operator)
    end
end