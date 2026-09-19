class RisksController < ApplicationController
  def index
    @risks = Risk.all
  end
 
  def show
    @risk = Risk.find(params[:id])
  end
 
  def new
    @risk = Risk.new
  end
 
  def edit
    @risk = Risk.find(params[:id])
  end
 
  def create
    @risk = Risk.new(risk_params)
 
    if @risk.save
      redirect_to risks_path
    else
      render 'new'
    end
  end
 
  def update
    @risk = Risk.find(params[:id])
 
    if @risk.update(risk_params)
      redirect_to risks_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @risk = Risk.find(params[:id])
    @risk.destroy
    redirect_to risks_path
  end

 
  private
    def risk_params
      params.require(:risk).permit(:name, :description, :inherentRiskScore, :residualRiskScore, :Category, :Impact, :Likelihood, :Status)
    end
end