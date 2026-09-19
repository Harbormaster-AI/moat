class UnderwritingDecisionsController < ApplicationController
  def index
    @underwritingDecisions = UnderwritingDecision.all
  end
 
  def show
    @underwritingDecision = UnderwritingDecision.find(params[:id])
  end
 
  def new
    @underwritingDecision = UnderwritingDecision.new
  end
 
  def edit
    @underwritingDecision = UnderwritingDecision.find(params[:id])
  end
 
  def create
    @underwritingDecision = UnderwritingDecision.new(underwritingDecision_params)
 
    if @underwritingDecision.save
      redirect_to underwritingDecisions_path
    else
      render 'new'
    end
  end
 
  def update
    @underwritingDecision = UnderwritingDecision.find(params[:id])
 
    if @underwritingDecision.update(underwritingDecision_params)
      redirect_to underwritingDecisions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @underwritingDecision = UnderwritingDecision.find(params[:id])
    @underwritingDecision.destroy
    redirect_to underwritingDecisions_path
  end

 
  private
    def underwritingDecision_params
      params.require(:underwritingDecision).permit(:notes, :decisionDate, :Decision)
    end
end