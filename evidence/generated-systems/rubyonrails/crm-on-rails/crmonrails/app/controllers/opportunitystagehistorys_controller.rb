class OpportunityStageHistorysController < ApplicationController
  def index
    @opportunityStageHistorys = OpportunityStageHistory.all
  end
 
  def show
    @opportunityStageHistory = OpportunityStageHistory.find(params[:id])
  end
 
  def new
    @opportunityStageHistory = OpportunityStageHistory.new
  end
 
  def edit
    @opportunityStageHistory = OpportunityStageHistory.find(params[:id])
  end
 
  def create
    @opportunityStageHistory = OpportunityStageHistory.new(opportunityStageHistory_params)
 
    if @opportunityStageHistory.save
      redirect_to opportunityStageHistorys_path
    else
      render 'new'
    end
  end
 
  def update
    @opportunityStageHistory = OpportunityStageHistory.find(params[:id])
 
    if @opportunityStageHistory.update(opportunityStageHistory_params)
      redirect_to opportunityStageHistorys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @opportunityStageHistory = OpportunityStageHistory.find(params[:id])
    @opportunityStageHistory.destroy
    redirect_to opportunityStageHistorys_path
  end

 
  private
    def opportunityStageHistory_params
      params.require(:opportunityStageHistory).permit(:changedAt, :comment, :FromStage, :ToStage)
    end
end