class CorrectiveActionsController < ApplicationController
  def index
    @correctiveActions = CorrectiveAction.all
  end
 
  def show
    @correctiveAction = CorrectiveAction.find(params[:id])
  end
 
  def new
    @correctiveAction = CorrectiveAction.new
  end
 
  def edit
    @correctiveAction = CorrectiveAction.find(params[:id])
  end
 
  def create
    @correctiveAction = CorrectiveAction.new(correctiveAction_params)
 
    if @correctiveAction.save
      redirect_to correctiveActions_path
    else
      render 'new'
    end
  end
 
  def update
    @correctiveAction = CorrectiveAction.find(params[:id])
 
    if @correctiveAction.update(correctiveAction_params)
      redirect_to correctiveActions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @correctiveAction = CorrectiveAction.find(params[:id])
    @correctiveAction.destroy
    redirect_to correctiveActions_path
  end

 
  private
    def correctiveAction_params
      params.require(:correctiveAction).permit(:actionTitle, :owner, :targetDate, :Status)
    end
end